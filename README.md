# Chrome-To-Terminal-OneClick

**Run Anywhere — Web → Terminal**

A cross-platform Chrome extension that lets users select text on a webpage and run it in their local terminal with one click.

## What it does

Select a command on Stack Overflow, GitHub, Reddit, documentation, or another webpage, then right-click **Run in Terminal**. The extension also provides a popup for selecting a shell or pasting a command.

On Linux, Run Anywhere keeps a persistent terminal session so repeated commands can be run in the same terminal with previous commands, output and working directory visible. The persistent Linux session uses `tmux` when available. On systems without `tmux`, Run Anywhere falls back to keeping each terminal window open at a prompt.

## Supported platforms

- Linux — Bash/Zsh and common terminal emulators including Konsole
- macOS — Terminal.app on Apple Silicon and Intel
- Windows — PowerShell or Command Prompt

## Architecture

Chrome Extension → Chrome Native Messaging → local Run Anywhere connector → operating-system terminal.

The native connector is separate from the Chrome Web Store extension because Chrome does not allow an ordinary extension to directly launch arbitrary local applications.

## Repository layout

- `extension/` — Manifest V3 Chrome extension
- `connectors/` — native connector source, installers and binaries
- `packages/` — Linux Debian package
- `RELEASE-NOTES.md` — release notes

## Development

Load `extension/` as an unpacked extension from `chrome://extensions/` while developing. Install the native connector for your operating system before testing terminal execution.

### Linux persistent terminal

For the shared-session behavior, install `tmux` once with your distribution's normal package manager. Run Anywhere creates a dedicated `run-anywhere` tmux session automatically and opens a terminal attached to it. Closing the terminal does not destroy the session, so the next **Run in Terminal** can reconnect to the same history and working directory.

## Security

This software intentionally executes selected text with the user's normal operating-system permissions. Only execute commands you understand and trust. The extension does not require a remote execution server.

Developed by Digital Generationz.
