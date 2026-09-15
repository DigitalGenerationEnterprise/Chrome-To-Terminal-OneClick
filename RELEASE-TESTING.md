# Release testing

## v1.0.1

The GitHub Actions release workflow builds the native connector from `connectors/go`, packages the Chrome extension, and publishes Linux, Windows, and macOS connector downloads to the GitHub Release.

### Customer flow

1. Install **Run Anywhere** from the Chrome Web Store.
2. Install the small connector for the operating system once.
3. Restart Chrome if it was running during connector installation.
4. Select a command on a webpage.
5. Right-click → **Run in Terminal**.

No extension-ID copy/paste or JSON editing should be required for normal users.
