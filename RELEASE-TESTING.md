# Release testing

## Package layout

Each connector archive must be self-contained. The installer must resolve files relative to the installer itself, not its parent directory.

### Linux

```text
Run-Anywhere-Linux-Connector-X.Y.Z/
├── install.sh
└── linux/
    └── run-anywhere-host
```

Run from the extracted connector directory with:

```bash
chmod +x install.sh
./install.sh
```

Do not require `sudo` for the per-user Chrome Native Messaging installation.

### macOS

```text
Run-Anywhere-macOS-Connector-X.Y.Z/
├── install.sh
└── macos/
    ├── run-anywhere-host-intel
    └── run-anywhere-host-arm64
```

The installer selects the correct binary from the package's own `macos/` directory.

### Windows

```text
Run-Anywhere-Windows-Connector-X.Y.Z/
├── install.ps1
└── windows/
    └── run-anywhere-host.exe
```

The installer selects the executable from the package's own `windows/` directory.

## Customer flow

1. Install **Run Anywhere** from the Chrome Web Store.
2. Install the small connector for the operating system once.
3. Restart Chrome if it was running during connector installation.
4. Select a command on a webpage.
5. Right-click → **Run in Terminal**.

No extension-ID copy/paste or JSON editing should be required for normal users.
