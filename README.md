# Chrome-To-Terminal-OneClick

**Run Anywhere — Web → Terminal**

A cross-platform Chrome extension that lets users select text on a webpage and run it in their local terminal with one click.

## What it does

Select a command on Stack Overflow, GitHub, Reddit, documentation, or another webpage, then right-click **Run in Terminal**. The extension also provides a popup for selecting a shell or pasting a command.

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

## Security

This software intentionally executes selected text with the user's normal operating-system permissions. Only execute commands you understand and trust. The extension does not require a remote execution server.

Developed by Digital Generationz.
