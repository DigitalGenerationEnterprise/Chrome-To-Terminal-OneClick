package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Request struct {
	Action    string `json:"action"`
	Command  string `json:"command"`
	Shell     string `json:"shell"`
	SourceURL string `json:"sourceUrl"`
}

type Response struct {
	OK       bool   `json:"ok"`
	Error    string `json:"error,omitempty"`
	Platform string `json:"platform,omitempty"`
	Shell    string `json:"shell,omitempty"`
}

func readMessage(r io.Reader) ([]byte, error) {
	var n uint32
	if err := binary.Read(r, binary.LittleEndian, &n); err != nil {
		return nil, err
	}
	if n > 64*1024*1024 {
		return nil, fmt.Errorf("message too large")
	}
	b := make([]byte, n)
	_, err := io.ReadFull(r, b)
	return b, err
}

func writeMessage(w io.Writer, obj any) error {
	d, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	var b bytes.Buffer
	if err := binary.Write(&b, binary.LittleEndian, uint32(len(d))); err != nil {
		return err
	}
	b.Write(d)
	_, err = w.Write(b.Bytes())
	return err
}

func exists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

func shellPath(shell string, fallback string) string {
	switch shell {
	case "bash":
		return "/bin/bash"
	case "zsh":
		return "/bin/zsh"
	case "powershell":
		return "powershell.exe"
	case "cmd":
		return "cmd.exe"
	default:
		return fallback
	}
}

func quoteSingle(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "'\\''") + "'"
}

func linux(c, s string) error {
	p := shellPath(s, "/bin/bash")
	if s == "auto" {
		if exists("zsh") {
			p = "/bin/zsh"
		}
	}

	// Keep one persistent Run Anywhere shell session so every right-click
	// command shares the same history, output and working directory.
	if exists("tmux") {
		return linuxTmux(c, p)
	}

	// Fallback for systems without tmux: keep the terminal open at a prompt,
	// but commands will use separate terminal windows instead of one shared session.
	script := c + "; exec " + quoteSingle(p)
	terms := [][]string{}
	if exists("konsole") {
		terms = append(terms, []string{"konsole", "-e", p, "-lc", script})
	}
	if exists("gnome-terminal") {
		terms = append(terms, []string{"gnome-terminal", "--", p, "-lc", script})
	}
	if exists("kgx") {
		terms = append(terms, []string{"kgx", "--", p, "-lc", script})
	}
	if exists("x-terminal-emulator") {
		terms = append(terms, []string{"x-terminal-emulator", "-e", p, "-lc", script})
	}
	for _, a := range terms {
		if err := exec.Command(a[0], a[1:]...).Start(); err == nil {
			return nil
		}
	}
	return fmt.Errorf("no supported terminal emulator was found")
}

func linuxTmux(command, shell string) error {
	const session = "run-anywhere"

	// Create the persistent shell session if it does not already exist.
	if err := exec.Command("tmux", "has-session", "-t", session).Run(); err != nil {
		if err := exec.Command("tmux", "new-session", "-d", "-s", session, shell).Run(); err != nil {
			return fmt.Errorf("could not create Run Anywhere terminal session: %w", err)
		}
	}

	// If no terminal client is currently attached, open one and attach to the
	// existing session. Closing the terminal leaves the tmux session alive.
	clients := exec.Command("tmux", "list-clients", "-t", session)
	if out, err := clients.Output(); err != nil || len(bytes.TrimSpace(out)) == 0 {
		if err := launchTerminalAttached(session); err != nil {
			return err
		}
	}

	// Send literally, then send Enter separately so commands beginning with '-'
	// are not interpreted as tmux options.
	if err := exec.Command("tmux", "send-keys", "-t", session, "-l", command).Run(); err != nil {
		return fmt.Errorf("could not send command to Run Anywhere terminal: %w", err)
	}
	if err := exec.Command("tmux", "send-keys", "-t", session, "Enter").Run(); err != nil {
		return fmt.Errorf("could not submit command to Run Anywhere terminal: %w", err)
	}
	return nil
}

func launchTerminalAttached(session string) error {
	attach := []string{"tmux", "attach-session", "-t", session}
	terms := [][]string{}
	if exists("konsole") {
		terms = append(terms, []string{"konsole", "-e"})
	}
	if exists("gnome-terminal") {
		terms = append(terms, []string{"gnome-terminal", "--"})
	}
	if exists("kgx") {
		terms = append(terms, []string{"kgx", "--"})
	}
	if exists("x-terminal-emulator") {
		terms = append(terms, []string{"x-terminal-emulator", "-e"})
	}
	for _, prefix := range terms {
		args := append(append([]string{}, prefix...), attach...)
		if err := exec.Command(args[0], args[1:]...).Start(); err == nil {
			return nil
		}
	}
	return fmt.Errorf("could not open a supported terminal for the persistent Run Anywhere session")
}

func mac(c, s string) error {
	p := "/bin/zsh"
	if s == "bash" {
		p = "/bin/bash"
	}
	script := fmt.Sprintf("tell application \"Terminal\" to do script %s", `"`+strings.ReplaceAll(p+" -lc "+quoteSingle(c)+"; exec "+p, `"`, `\"`)+`"`)
	return exec.Command("osascript", "-e", script).Start()
}

func windows(c, s string) error {
	if s == "cmd" {
		return exec.Command("cmd.exe", "/K", c).Start()
	}
	p := "powershell.exe"
	if exists("pwsh.exe") {
		p = "pwsh.exe"
	}
	return exec.Command(p, "-NoExit", "-Command", c).Start()
}

func run(r Request) Response {
	out := Response{Platform: runtime.GOOS, Shell: r.Shell}
	if strings.TrimSpace(r.Command) == "" {
		out.Error = "No command selected"
		return out
	}

	var err error
	switch runtime.GOOS {
	case "linux":
		err = linux(r.Command, r.Shell)
	case "darwin":
		err = mac(r.Command, r.Shell)
	case "windows":
		err = windows(r.Command, r.Shell)
	default:
		err = fmt.Errorf("unsupported operating system")
	}
	if err != nil {
		out.Error = err.Error()
		return out
	}
	out.OK = true
	return out
}

func main() {
	for {
		d, err := readMessage(os.Stdin)
		if err != nil {
			return
		}
		var r Request
		if err = json.Unmarshal(d, &r); err != nil {
			_ = writeMessage(os.Stdout, Response{Error: err.Error()})
			continue
		}
		if r.Action == "run" {
			_ = writeMessage(os.Stdout, run(r))
		} else {
			_ = writeMessage(os.Stdout, Response{Error: "Unknown action"})
		}
	}
}
