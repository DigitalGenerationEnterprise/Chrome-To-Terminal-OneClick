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

type Request struct { Action string `json:"action"`; Command string `json:"command"`; Shell string `json:"shell"`; SourceURL string `json:"sourceUrl"` }
type Response struct { OK bool `json:"ok"`; Error string `json:"error,omitempty"`; Platform string `json:"platform,omitempty"`; Shell string `json:"shell,omitempty"` }

func readMessage(r io.Reader) ([]byte,error){var n uint32;if err:=binary.Read(r,binary.LittleEndian,&n);err!=nil{return nil,err};if n>64*1024*1024{return nil,fmt.Errorf("message too large")};b:=make([]byte,n);_,err:=io.ReadFull(r,b);return b,err}
func writeMessage(w io.Writer,obj any)error{d,err:=json.Marshal(obj);if err!=nil{return err};var b bytes.Buffer;binary.Write(&b,binary.LittleEndian,uint32(len(d)));b.Write(d);_,err=w.Write(b.Bytes());return err}
func exists(n string)bool{_,e:=exec.LookPath(n);return e==nil}
func quote(s string)string{return "'"+strings.ReplaceAll(s,"'","'\\''")+"'"}
func linux(c,s string)error{p:="/bin/bash";if s=="zsh"||(s=="auto"&&exists("zsh")){p="/bin/zsh"};script:=c+"; printf '\\n\\n[Run Anywhere] Command finished. Press Enter to close. '; read";terms:=[][]string{};if exists("konsole"){terms=append(terms,[]string{"konsole","--hold","-e",p,"-lc",script})};if exists("gnome-terminal"){terms=append(terms,[]string{"gnome-terminal","--",p,"-lc",script})};if exists("kgx"){terms=append(terms,[]string{"kgx","--",p,"-lc",script})};if exists("x-terminal-emulator"){terms=append(terms,[]string{"x-terminal-emulator","-e",p,"-lc",script})};for _,a:=range terms{if err:=exec.Command(a[0],a[1:]...).Start();err==nil{return nil}};return fmt.Errorf("no supported terminal emulator was found")}
func mac(c,s string)error{p:="/bin/zsh";if s=="bash"{p="/bin/bash"};script:=fmt.Sprintf("tell application \"Terminal\" to do script %s",`"`+strings.ReplaceAll(p+" -lc "+quote(c),`"`,`\"`)+`"`);return exec.Command("osascript","-e",script).Start()}
func windows(c,s string)error{if s=="cmd"{return exec.Command("cmd.exe","/K",c).Start()};p:="powershell.exe";if exists("pwsh.exe"){p="pwsh.exe"};return exec.Command(p,"-NoExit","-Command",c).Start()}
func run(r Request)Response{out:=Response{Platform:runtime.GOOS,Shell:r.Shell};if strings.TrimSpace(r.Command)==""{out.Error="No command selected";return out};var e error;switch runtime.GOOS{case "linux":e=linux(r.Command,r.Shell);case "darwin":e=mac(r.Command,r.Shell);case "windows":e=windows(r.Command,r.Shell);default:e=fmt.Errorf("unsupported operating system")};if e!=nil{out.Error=e.Error();return out};out.OK=true;return out}
func main(){for{d,e:=readMessage(os.Stdin);if e!=nil{return};var r Request;if e=json.Unmarshal(d,&r);e!=nil{writeMessage(os.Stdout,Response{Error:e.Error()});continue};if r.Action=="run"{writeMessage(os.Stdout,run(r))}else{writeMessage(os.Stdout,Response{Error:"Unknown action"})}}}
