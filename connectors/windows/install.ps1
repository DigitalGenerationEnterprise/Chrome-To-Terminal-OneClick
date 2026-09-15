$ErrorActionPreference = "Stop"
$Root = Split-Path -Parent $PSScriptRoot
$HostDir = Join-Path $env:LOCALAPPDATA "Run Anywhere"
New-Item -ItemType Directory -Force -Path $HostDir | Out-Null
$Target = Join-Path $HostDir "run-anywhere-host.exe"
Copy-Item (Join-Path $Root "windows\run-anywhere-host.exe") $Target -Force
$Manifest = Join-Path $HostDir "com.digitalgenerationz.runanywhere.json"
$ManifestPath = $Target.Replace('\','\\')
$Json = @"
{
  "name":"com.digitalgenerationz.runanywhere",
  "description":"Chrome-to-Terminal native terminal connector",
  "path":"$ManifestPath",
  "type":"stdio",
  "allowed_origins":["chrome-extension://pmkcmghjfmigbokmlpleiohnidpikpi/"]
}
"@
$Json | Set-Content -Encoding UTF8 $Manifest
$Key = "HKCU:\Software\Google\Chrome\NativeMessagingHosts\com.digitalgenerationz.runanywhere"
New-Item -Path $Key -Force | Out-Null
Set-ItemProperty -Path $Key -Name "(default)" -Value $Manifest
Write-Host "Chrome-to-Terminal native connector installed for Google Chrome."
