const shell=document.getElementById('shell');const command=document.getElementById('command');const run=document.getElementById('run');const status=document.getElementById('status');
load();
document.getElementById('settings').onclick=()=>chrome.runtime.openOptionsPage();
shell.onchange=()=>chrome.storage.sync.set({shell:shell.value});
run.onclick=async()=>{status.className='status';status.textContent='Opening terminal…';const text=command.value.trim();if(!text){status.className='status error';status.textContent='Select or paste a command first.';return;}render(await send(text));};
async function load(){const saved=await chrome.storage.sync.get({shell:'auto'});shell.value=saved.shell;try{const [tab]=await chrome.tabs.query({active:true,currentWindow:true});if(tab?.id){const [{result}]=await chrome.scripting.executeScript({target:{tabId:tab.id},func:()=>window.getSelection()?.toString()||''});if(result)command.value=result;}}catch(_){} }
function send(text){return new Promise(resolve=>chrome.runtime.sendMessage({type:'run-command',command:text,shell:shell.value},resolve));}
function render(result){if(result?.ok){status.className='status ok';status.textContent='Terminal opened.';}else{status.className='status error';status.textContent=result?.needsSetup?'Native connector not installed. Open Settings & setup.':(result?.error||'Could not open terminal.');}}
