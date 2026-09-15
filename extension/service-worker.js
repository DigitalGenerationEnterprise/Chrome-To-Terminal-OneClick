const HOST_NAME = 'com.digitalgenerationz.runanywhere';

chrome.runtime.onInstalled.addListener(async () => {
  chrome.contextMenus.create({ id: 'run-selected', title: 'Run in Terminal', contexts: ['selection'] });
  chrome.contextMenus.create({ id: 'run-selected-shell', title: 'Run in Terminal ▸', contexts: ['selection'] });
  for (const [id, title] of [['auto','Auto-detect'],['bash','Bash / Zsh'],['powershell','PowerShell'],['cmd','Command Prompt']]) {
    chrome.contextMenus.create({ id: `shell-${id}`, parentId: 'run-selected-shell', title, contexts: ['selection'] });
  }
});

chrome.contextMenus.onClicked.addListener(async (info, tab) => {
  if (!info.selectionText) return;
  let shell = typeof info.menuItemId === 'string' && info.menuItemId.startsWith('shell-') ? info.menuItemId.slice(6) : await getShell();
  runCommand(info.selectionText, shell, tab?.url || '');
});

chrome.runtime.onMessage.addListener((message, sender, sendResponse) => {
  if (message?.type !== 'run-command') return;
  (async () => sendResponse(await runCommand(message.command, message.shell || await getShell(), sender?.tab?.url || '')))();
  return true;
});

async function getShell() {
  const { shell = 'auto' } = await chrome.storage.sync.get({ shell: 'auto' });
  return shell;
}

async function runCommand(command, shell, sourceUrl) {
  if (!command?.trim()) return { ok: false, error: 'No command selected.' };
  try {
    return await chrome.runtime.sendNativeMessage(HOST_NAME, { action:'run', command:command.trim(), shell, sourceUrl });
  } catch (error) {
    const message = error?.message || String(error);
    return { ok:false, error:message, needsSetup:/host|native|not found|forbidden|connect/i.test(message) };
  }
}
