#!/bin/bash
set -e
ROOT="$(cd "$(dirname "$0")" && pwd)"
HOST_DIR="$HOME/Library/Application Support/Chrome-to-Terminal"
MANIFEST_DIR="$HOME/Library/Application Support/Google/Chrome/NativeMessagingHosts"
mkdir -p "$HOST_DIR" "$MANIFEST_DIR"
ARCH="$(uname -m)"
if [ "$ARCH" = "arm64" ]; then cp "$ROOT/macos/run-anywhere-host-arm64" "$HOST_DIR/run-anywhere-host"; else cp "$ROOT/macos/run-anywhere-host-intel" "$HOST_DIR/run-anywhere-host"; fi
chmod +x "$HOST_DIR/run-anywhere-host"
cat > "$MANIFEST_DIR/com.digitalgenerationz.runanywhere.json" <<JSON
{
  "name":"com.digitalgenerationz.runanywhere",
  "description":"Chrome-to-Terminal native terminal connector",
  "path":"$HOST_DIR/run-anywhere-host",
  "type":"stdio",
  "allowed_origins":["chrome-extension://pmkcmghjfmigbokmlpleiohnidpikpi/"]
}
JSON
echo "Chrome-to-Terminal native connector installed for Google Chrome."
