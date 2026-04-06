/* catalog.js — с добавлением кнопок ПР2 и ПР3 */
let allPartsData=[],categoryFilter='all',searchText='';

async function loadCatalogData(){
    renderCatalog(false,true);
    const raw={};
    const loadErrors=[];
    for(const cat of CATEGORY_MAP){
        try{raw[cat.key]=await api[cat.key].list()||[];}
        catch(e){
            raw[cat.key]=[];
            loadErrors.push(`${cat.label}: ${e.message}`);
            console.error(`[catalog] ${cat.label}:`,e.message);
        }
    }
    if(loadErrors.length){
        const warn=document.createElement('div');
        warn.style.cssText='position:fixed;bottom:1rem;right:1rem;z-index:9999;'+
            'background:#7f1d1d;color:#fca5a5;border:1px solid #ef4444;'+
            'border-radius:.5rem;padding:.75rem 1rem;font-size:.8rem;max-width:360px;'+
            'box-shadow:0 4px 16px rgba(0,0,0,.5)';
        warn.innerHTML='<strong>⚠ Не удалось загрузить:</strong><br>'+loadErrors.join('<br>');
        document.body.appendChild(warn);
        setTimeout(()=>warn.remove(),12000);
    }
    const nameFor={};
    for(const cat of CATEGORY_MAP){
        nameFor[cat.idField]={};
        for(const item of raw[cat.key]){
            nameFor[cat.idField][item[cat.idField]]=cat.nameField
                ?item[cat.nameField]
                :`${cat.label} #${shortId(item[cat.idField])}`;
        }
    }
    const getInfo={
        powerPoints:   i=>[
            {l:'Двиг',v:nameFor['engine_id']?.[i.engine_id]            ||'?'},
            {l:'Инв', v:nameFor['inverter_id']?.[i.inverter_id]         ||'?'},
            {l:'КПП', v:nameFor['gearbox_id']?.[i.gearbox_id]           ||'?'},
        ],
        chassis:       i=>[
            {l:'Рама', v:nameFor['frame_id']?.[i.frame_id]              ||'?'},
            {l:'Подв', v:nameFor['suspension_id']?.[i.suspension_id]    ||'?'},
            {l:'Торм', v:nameFor['break_system_id']?.[i.break_system_id]||'?'},
        ],
        chargerSystems:i=>[
            {l:'Заряд',v:nameFor['charger_id']?.[i.charger_id]          ||'?'},
            {l:'Конн', v:nameFor['connector_id']?.[i.connector_id]      ||'?'},
        ],
        electronics:   i=>[
            {l:'Контр',v:nameFor['controller_id']?.[i.controller_id]    ||'?'},
            {l:'Сенс', v:nameFor['sensor_id']?.[i.sensor_id]            ||'?'},
            {l:'Пров', v:nameFor['wiring_id']?.[i.wiring_id]            ||'?'},
        ],
        bodies:        i=>[
            {l:'Карк',v:nameFor['carcass_id']?.[i.carcass_id]           ||'?'},
            {l:'Дв',  v:nameFor['doors_id']?.[i.doors_id]               ||'?'},
            {l:'Кр',  v:nameFor['wings_id']?.[i.wings_id]               ||'?'},
        ],
        emobiles:      i=>[
            {l:'СУ',     v:nameFor['power_point_id']?.[i.power_point_id]    ||'?'},
            {l:'Батарея',v:nameFor['battery_id']?.[i.battery_id]            ||'?'},
            {l:'Зарядка',v:nameFor['charger_system_id']?.[i.charger_system_id]||'?'},
            {l:'Шасси',  v:nameFor['chassis_id']?.[i.chassis_id]            ||'?'},
            {l:'Кузов',  v:nameFor['body_id']?.[i.body_id]                  ||'?'},
            {l:'Электро',v:nameFor['electronics_id']?.[i.electronics_id]    ||'?'},
        ],
    };
    const flat=[];
    for(const cat of CATEGORY_MAP){
        const resolver=getInfo[cat.key]||null;
        for(const item of raw[cat.key]){
            const name=cat.nameField
                ?(item[cat.nameField]||`${cat.label} #${shortId(item[cat.idField])}`)
                :`${cat.label} …${shortId(item[cat.idField])}`;
            const infoLines=resolver?resolver(item):null;
            const infoStr=resolver?infoLines.map(r=>`${r.l}: ${r.v}`).join(' · '):cat.infoFn(item);
            flat.push({id:item[cat.idField],name,category:cat.label,
                       categoryKey:cat.key,infoLines,infoStr,rawData:item});
        }
    }
    allPartsData=flat;
    renderCatalog(false,false);
}

function renderInfoCell(p){
    if(!p.infoLines)
        return`<td class="info-cell">${escapeHtml(p.infoStr)}</td>`;
    const lines=p.infoLines.map(r=>
        `<span class="info-line">
            <span class="info-label">${escapeHtml(r.l)}</span>
            <span class="info-value" title="${escapeHtml(r.v)}">${escapeHtml(r.v)}</span>
         </span>`
    ).join('');
    return`<td class="info-cell"><div class="info-lines">${lines}</div></td>`;
}

function _buildChipsHtml(){
    const allLabels=CATEGORY_MAP.map(c=>c.label);
    return['all',...allLabels].map(l=>
        `<span class="filter-chip ${categoryFilter===l?'active':''}" data-cat="${escapeHtml(l)}">
            ${l==='all'?'Все':escapeHtml(l)}</span>`).join('');
}

function _buildRowsHtml(isAdmin){
    const q=searchText.toLowerCase();
    const rows=allPartsData.filter(p=>{
        const mc=categoryFilter==='all'||p.category===categoryFilter;
        return mc&&(!q||p.name.toLowerCase().includes(q)||String(p.id).includes(q));
    });
    if(!rows.length)return`<tr><td colspan="${isAdmin?5:4}">
        <div class="empty-state"><i class="fas fa-inbox"></i> Нет данных</div></td></tr>`;
    return rows.map(p=>`<tr>
        <td><span class="text-mono">${escapeHtml(shortId(p.id))}</span></td>
        <td style="font-weight:500">${escapeHtml(p.name)}</td>
        <td><span class="badge badge-default">${escapeHtml(p.category)}</span></td>
        ${renderInfoCell(p)}
        ${isAdmin?`<td style="white-space:nowrap">
            <i class="fas fa-edit action-icon" style="color:#60a5fa"
               data-action="edit"   data-id="${escapeHtml(p.id)}"
               data-cat="${escapeHtml(p.categoryKey)}" title="Редактировать"></i>
            <i class="fas fa-trash-alt action-icon" style="color:#f87171"
               data-action="delete" data-id="${escapeHtml(p.id)}"
               data-cat="${escapeHtml(p.categoryKey)}"
               data-name="${escapeHtml(p.name)}" title="Удалить"></i>
        </td>`:''}`
    ).join('');
}

function renderCatalog(partial=true,loading=false){
    const app=document.getElementById('app');
    if(!app)return;
    const isAdmin=currentRole==='admin';
    if(partial&&app.querySelector('.table-wrap')){
        app.querySelector('tbody').innerHTML=loading
            ?`<tr><td colspan="${isAdmin?5:4}" style="text-align:center;padding:2rem;color:var(--muted)">
               <i class="fas fa-spinner fa-spin"></i> Загрузка…</td></tr>`
            :_buildRowsHtml(isAdmin);
        const fc=app.querySelector('#filterChips');
        if(fc)fc.innerHTML=_buildChipsHtml();
        _attachActions(app,isAdmin);
        return;
    }
    const roleBadge=isAdmin
        ?`<span class="badge badge-admin"><i class="fas fa-shield-alt"></i> Администратор</span>`
        :`<span class="badge badge-user"><i class="fas fa-user"></i> Пользователь</span>`;
    app.innerHTML=`
    <div class="container">
        <div class="page-header">
            <h1 class="page-title"><i class="fas fa-charging-station"></i> Каталог Emobile</h1>
            <div style="display:flex;align-items:center;gap:.6rem;flex-wrap:wrap">
                ${roleBadge}
                <button class="btn btn-secondary" id="logoutBtn" title="Выйти">
                    <i class="fas fa-sign-out-alt"></i></button>
            </div>
        </div>
        ${isAdmin?`<div class="toolbar">
            <button class="btn btn-primary"   id="createCarBtn">
                <i class="fas fa-car"></i> Создать автомобиль</button>
            <button class="btn btn-secondary" id="addPartBtn">
                <i class="fas fa-plus"></i> Добавить деталь</button>
            <span style="margin-left:auto;display:flex;gap:.4rem">
                <button class="btn btn-secondary" id="enumsBtn" title="ПР2: Справочники перечислений">
                    <i class="fas fa-list-ul"></i> Справочники</button>
                <button class="btn btn-secondary" id="paramsBtn" title="ПР3: Параметры изделий">
                    <i class="fas fa-sliders-h"></i> Параметры</button>
                <button class="btn btn-secondary" id="createUserBtn">
                    <i class="fas fa-user-plus"></i> Новый пользователь</button>
            </span>
        </div>`:''}`+
        `<div class="toolbar">
            <input type="text" id="searchInput" class="input-dark" style="max-width:320px"
                placeholder="Поиск по ID или названию…" value="${escapeHtml(searchText)}">
        </div>
        <div class="filter-chips" id="filterChips">${_buildChipsHtml()}</div>
        <div class="table-wrap">
            <table>
                <thead><tr>
                    <th>ID</th><th>Название</th><th>Категория</th><th>Информация</th>
                    ${isAdmin?'<th>Действия</th>':''}
                </tr></thead>
                <tbody>${loading
                    ?`<tr><td colspan="${isAdmin?5:4}" style="text-align:center;padding:2rem;color:var(--muted)">
                       <i class="fas fa-spinner fa-spin"></i> Загрузка…</td></tr>`
                    :_buildRowsHtml(isAdmin)}</tbody>
            </table>
        </div>
    </div>`;
    app.querySelector('#searchInput').addEventListener('input',e=>{
        searchText=e.target.value;renderCatalog(true);});
    app.querySelector('#filterChips').addEventListener('click',e=>{
        const c=e.target.closest('.filter-chip');if(!c)return;
        categoryFilter=c.dataset.cat;renderCatalog(true);});
    app.querySelector('#logoutBtn').addEventListener('click',()=>{
        currentRole=null;saveSession(null);
        allPartsData=[];categoryFilter='all';searchText='';
        showLogin();});
    if(isAdmin){
        app.querySelector('#createCarBtn').addEventListener('click',openWizard);
        app.querySelector('#addPartBtn').addEventListener('click',()=>openPartForm(null,null));
        app.querySelector('#createUserBtn').addEventListener('click',openCreateUserForm);
        // ── ПР2: Справочники перечислений ──
        app.querySelector('#enumsBtn').addEventListener('click',openEnumsManager);
        // ── ПР3: Параметры изделий ──
        app.querySelector('#paramsBtn').addEventListener('click',openParametersManager);
        _attachActions(app,isAdmin);
    }
}

function _attachActions(app,isAdmin){
    if(!isAdmin)return;
    app.querySelectorAll('[data-action="edit"]').forEach(el=>
        el.addEventListener('click',()=>openPartForm(el.dataset.id,el.dataset.cat)));
    app.querySelectorAll('[data-action="delete"]').forEach(el=>
        el.addEventListener('click',()=>confirmDelete(el.dataset.id,el.dataset.cat,el.dataset.name)));
}
