#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
HOST_DIR="$HOME/.local/share/chrome-to-terminal"
MANIFEST_DIR="$HOME/.config/google-chrome/NativeMessagingHosts"
mkdir -p "$HOST_DIR" "$MANIFEST_DIR"
cp "$ROOT/linux/run-anywhere-host" "$HOST_DIR/chrome-to-terminal-host"
chmod +x "$HOST_DIR/chrome-to-terminal-host"
cat > "$MANIFEST_DIR/com.digitalgenerationz.runanywhere.json" <<JSON
{
  "name":"com.digitalgenerationz.runanywhere",
  "description":"Chrome-to-Terminal native terminal connector",
  "path":"$HOST_DIR/chrome-to-terminal-host",
  "type":"stdio",
  "allowed_origins":["chrome-extension://pmkcmghjfmigbokmlpleiohnidpikpi/"]
}
JSON
echo "Chrome-to-Terminal native connector installed for Google Chrome."
