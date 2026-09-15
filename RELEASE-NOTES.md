# Chrome-to-Terminal 1.0.4

Maintenance release for the first Chrome Web Store candidate.

- Product branding standardized to **Chrome-to-Terminal**
- Persistent Linux terminal session uses `tmux` when installed
- Linux connector install path standardized to `~/.local/share/chrome-to-terminal`
- macOS connector install path standardized to `~/Library/Application Support/Chrome-to-Terminal`
- Windows connector installer corrected to use the packaged `windows` directory
- Chrome extension version updated to 1.0.4
- Internal Chrome Native Messaging host ID remains `com.digitalgenerationz.runanywhere` for compatibility

# Chrome-to-Terminal 1.0.3

- Persistent Linux terminal session using `tmux`
- Terminal remains available for subsequent right-click commands
- Chrome-to-Terminal product branding introduced

# Chrome-to-Terminal 1.0.0

Initial cross-platform release.

- Chrome Manifest V3 extension
- Right-click **Run in Terminal** on selected webpage text
- Popup workflow for selected/pasted commands
- Auto-detect, Bash/Zsh, PowerShell and Command Prompt
- Chrome Native Messaging connector
- Linux, macOS Intel/Apple Silicon, Windows x64 connector builds
- Linux Debian package
