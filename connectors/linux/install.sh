#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
HOST_DIR="$HOME/.local/share/run-anywhere"
MANIFEST_DIR="$HOME/.config/google-chrome/NativeMessagingHosts"
mkdir -p "$HOST_DIR" "$MANIFEST_DIR"
cp "$ROOT/linux/run-anywhere-host" "$HOST_DIR/run-anywhere-host"
chmod +x "$HOST_DIR/run-anywhere-host"
cat > "$MANIFEST_DIR/com.digitalgenerationz.runanywhere.json" <<JSON
{
  "name":"com.digitalgenerationz.runanywhere",
  "description":"Run Anywhere native terminal connector",
  "path":"$HOST_DIR/run-anywhere-host",
  "type":"stdio",
  "allowed_origins":["chrome-extension://plkdaddppeijjnnhpgmfpbdenhbhakhm/"]
}
JSON
echo "Run Anywhere native connector installed for Google Chrome."
