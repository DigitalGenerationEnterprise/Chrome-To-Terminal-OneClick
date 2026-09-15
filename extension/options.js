const select=document.getElementById('shell');
chrome.storage.sync.get({shell:'auto'}).then(x=>select.value=x.shell);
select.onchange=()=>chrome.storage.sync.set({shell:select.value});
