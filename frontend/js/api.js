/* api.js */
async function apiRequest(url,opts={}){
    let res;
    try{
        res=await fetch(url,{...opts,headers:{'Content-Type':'application/json',...opts.headers}});
    }catch(e){throw new Error(friendlyError(e.message||'NetworkError'));}
    if(res.status===204)return null;
    let body=null;
    try{body=await res.json();}catch(_){}
    if(!res.ok){
        const raw=body?.message||body?.error||body?.detail||`HTTP ${res.status}`;
        throw new Error(friendlyError(raw));
    }
    return body;
}
function makeResource(group,getWord){
    const b=`${API_BASE}/${group}`;
    return{
        list:    ()     =>apiRequest(`${b}/list`),
        getById: id     =>apiRequest(`${b}/get${getWord}${id}`),
        create:  data   =>apiRequest(`${b}/create`,{method:'POST',body:JSON.stringify(data)}),
        update:  (id,d) =>apiRequest(`${b}/update${id}`,{method:'PUT',body:JSON.stringify(d)}),
        delete:  id     =>apiRequest(`${b}/delete${id}`,{method:'DELETE'}),
    };
}
const api={auth:{}};
CATEGORY_MAP.forEach(cat=>{api[cat.key]=makeResource(cat.group,cat.getWord);});
api.auth.login=(username,password)=>
    apiRequest(`${API_BASE}/auth/login`,{method:'POST',body:JSON.stringify({username,password})});
api.users={
    list:  ()   =>apiRequest(`${API_BASE}/users/list`),
    create:data =>apiRequest(`${API_BASE}/users/create`,{method:'POST',body:JSON.stringify(data)}),
    delete:id   =>apiRequest(`${API_BASE}/users/delete${id}`,{method:'DELETE'}),
};
