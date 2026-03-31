/* forms.js */
function createModal(html,wide=false){
    const o=document.createElement('div');
    o.className='modal';
    o.innerHTML=`<div class="modal-content${wide?' modal-wide':''}">${html}</div>`;
    o.addEventListener('click',e=>{if(e.target===o)o.remove();});
    document.body.appendChild(o);return o;
}
function showError(msg){
    const m=createModal(`
        <div class="modal-title" style="color:#f87171">
            <i class="fas fa-exclamation-circle"></i> Ошибка</div>
        <p style="line-height:1.6;color:var(--text)">${escapeHtml(msg)}</p>
        <div class="modal-footer">
            <button class="btn btn-secondary" id="_ok">OK</button></div>`);
    m.querySelector('#_ok').onclick=()=>m.remove();
}
function setFormError(container,msg){
    const el=container.querySelector('.error-banner');if(!el)return;
    if(msg){el.querySelector('span').textContent=msg;el.classList.remove('hidden');}
    else el.classList.add('hidden');
}

/* ── Удаление ──────────────────────────────────── */
function confirmDelete(id,categoryKey,name){
    const m=createModal(`
        <div class="modal-title" style="color:#f87171">
            <i class="fas fa-trash-alt"></i> Удаление</div>
        <p>Вы уверены, что хотите удалить <strong>${escapeHtml(name)}</strong>?</p>
        <div class="error-banner hidden"><i class="fas fa-exclamation-circle"></i><span></span></div>
        <div class="modal-footer">
            <button class="btn btn-secondary" id="_c">Отмена</button>
            <button class="btn btn-danger"    id="_d">
                <i class="fas fa-trash-alt"></i> Удалить</button>
        </div>`);
    m.querySelector('#_c').onclick=()=>m.remove();
    m.querySelector('#_d').onclick=async()=>{
        const btn=m.querySelector('#_d');
        btn.disabled=true;btn.innerHTML='<i class="fas fa-spinner fa-spin"></i>';
        setFormError(m,'');
        try{
            await api[categoryKey].delete(id);
            await loadCatalogData();m.remove();
        }catch(e){
            setFormError(m,e.message);
            btn.disabled=false;
            btn.innerHTML='<i class="fas fa-trash-alt"></i> Удалить';
        }
    };
}

/* ── HTML поля ─────────────────────────────────── */
function buildFieldHtml(f,val){
    const v=escapeHtml(val??'');
    if(f.type==='textarea')
        return`<textarea id="_f_${f.id}" class="input-dark" rows="2" style="resize:none">${v}</textarea>`;
    if(f.type==='select'){
        const opts=(f.options||[]).map(o=>
            `<option value="${escapeHtml(o)}" ${val===o?'selected':''}>${escapeHtml(o)}</option>`
        ).join('');
        return`<select id="_f_${f.id}" class="input-dark">
                   <option value="">— выберите —</option>${opts}</select>`;
    }
    return`<input id="_f_${f.id}" type="${f.type}" class="input-dark"
               value="${v}" placeholder="${escapeHtml(f.label)}">`;
}

/* ── Форма листовой детали ─────────────────────── */
async function openPartForm(editId,categoryKey){
    if(!categoryKey){_openCategoryPicker();return;}
    const cat=getCatConfig(categoryKey);if(!cat)return;
    if(cat.composite){_openCompositeForm(cat,editId);return;}

    let existing={};
    if(editId){
        try{
            const list=await api[cat.key].list();
            existing=list.find(i=>String(i[cat.idField])===String(editId))||{};
        }catch(_){}
    }

    const fieldsHtml=cat.subFields.map(f=>{
        const val = f.id==='name'
            ? (existing[cat.nameField]??'')
            : (existing[f.dbField]??'');
        return`<div class="form-group">
                   <label>${escapeHtml(f.label)}</label>
                   ${buildFieldHtml(f,val)}
               </div>`;
    }).join('');

    const m=createModal(`
        <div class="modal-title">
            <i class="fas fa-${editId?'edit':'plus-circle'}"
               style="color:var(--accent)"></i>
            ${editId?'Редактировать':'Новая деталь'} — ${escapeHtml(cat.label)}</div>
        ${fieldsHtml}
        <div class="error-banner hidden">
            <i class="fas fa-exclamation-circle"></i><span></span></div>
        <div class="modal-footer">
            <button class="btn btn-secondary" id="_cc">Отмена</button>
            <button class="btn btn-primary"   id="_ss">
                <i class="fas fa-save"></i> Сохранить</button>
        </div>`);

    m.querySelector('#_cc').onclick=()=>m.remove();
    m.querySelector('#_ss').onclick=async()=>{
        /* payload строится по f.id — json:"name", json:"info", json:"engine_type" и т.д.
           Подтверждено: CreateBatteryInput.Name → json:"name" */
        const payload={};
        cat.subFields.forEach(f=>{
            const el=m.querySelector(`#_f_${f.id}`);
            payload[f.id]=el?el.value.trim():'';
        });

        console.debug('[save] category='+cat.key+' payload='+JSON.stringify(payload));

        if(!payload['name']){setFormError(m,'Заполните название');return;}

        const btn=m.querySelector('#_ss');
        btn.disabled=true;btn.innerHTML='<i class="fas fa-spinner fa-spin"></i>';
        setFormError(m,'');
        try{
            if(editId) await api[cat.key].update(editId,payload);
            else       await api[cat.key].create(payload);
            await loadCatalogData();m.remove();
        }catch(e){
            setFormError(m,e.message);
            btn.disabled=false;
            btn.innerHTML='<i class="fas fa-save"></i> Сохранить';
        }
    };
}

/* ── Выбор типа детали ─────────────────────────── */
function _openCategoryPicker(){
    const leaf=CATEGORY_MAP.filter(c=>!c.composite);
    const comp=CATEGORY_MAP.filter(c=> c.composite);
    const g=arr=>arr.map(c=>
        `<button class="btn btn-secondary _pc" data-key="${c.key}"
             style="text-align:left;justify-content:flex-start">
             ${escapeHtml(c.label)}</button>`).join('');
    const m=createModal(`
        <div class="modal-title">
            <i class="fas fa-list" style="color:var(--accent)"></i> Тип детали</div>
        <div class="section-label">Листовые</div>
        <div style="display:grid;grid-template-columns:1fr 1fr;gap:.4rem;margin-bottom:1rem">
            ${g(leaf)}</div>
        <div class="section-label">Составные</div>
        <div style="display:grid;grid-template-columns:1fr 1fr;gap:.4rem;margin-bottom:1rem">
            ${g(comp)}</div>
        <div class="modal-footer">
            <button class="btn btn-secondary" id="_cx">Отмена</button>
        </div>`,true);
    m.querySelector('#_cx').onclick=()=>m.remove();
    m.querySelectorAll('._pc').forEach(b=>
        b.addEventListener('click',()=>{m.remove();openPartForm(null,b.dataset.key);}));
}

/* ── Зависимости составных деталей ────────────── */
const COMPOSITE_DEPS={
    powerPoints:[
        {f:'engine_id',   c:'engines',     l:'Двигатель'},
        {f:'inverter_id', c:'inverters',   l:'Инвертор'},
        {f:'gearbox_id',  c:'gearboxes',   l:'КПП'},
    ],
    chassis:[
        {f:'frame_id',       c:'frames',       l:'Рама'},
        {f:'suspension_id',  c:'suspensions',  l:'Подвеска'},
        {f:'break_system_id',c:'breakSystems', l:'Тормозная система'},
    ],
    chargerSystems:[
        {f:'charger_id',   c:'chargers',   l:'Зарядное устройство'},
        {f:'connector_id', c:'connectors', l:'Коннектор'},
    ],
    electronics:[
        {f:'controller_id',c:'controllers',l:'Контроллер'},
        {f:'sensor_id',    c:'sensors',    l:'Датчик'},
        {f:'wiring_id',    c:'wirings',    l:'Проводка'},
    ],
    bodies:[
        {f:'carcass_id',c:'carcasses',l:'Каркас'},
        {f:'doors_id',  c:'doors',    l:'Двери'},
        {f:'wings_id',  c:'wings',    l:'Крылья'},
    ],
    emobiles:[
        {f:'name',              c:null,            l:'Название'},
        {f:'power_point_id',    c:'powerPoints',   l:'Силовая установка'},
        {f:'battery_id',        c:'batteries',     l:'Батарея'},
        {f:'charger_system_id', c:'chargerSystems',l:'Зарядная система'},
        {f:'chassis_id',        c:'chassis',       l:'Шасси'},
        {f:'body_id',           c:'bodies',        l:'Кузов'},
        {f:'electronics_id',    c:'electronics',   l:'Электроника'},
    ],
};

async function _openCompositeForm(cat,editId){
    const deps=COMPOSITE_DEPS[cat.key];if(!deps)return;
    const listsMap={};
    await Promise.all(deps.map(async dep=>{
        if(!dep.c)return;
        try{
            listsMap[dep.f]         =await api[dep.c].list();
            listsMap[`_c_${dep.f}`] =getCatConfig(dep.c);
        }catch(_){listsMap[dep.f]=[];}
    }));

    let existing={};
    if(editId){
        try{
            const list=await api[cat.key].list();
            existing=list.find(i=>String(i[cat.idField])===String(editId))||{};
        }catch(_){}
    }

    const fieldsHtml=deps.map(dep=>{
        if(!dep.c){
            const v=escapeHtml(existing[dep.f]||existing[cat.nameField]||'');
            return`<div class="form-group"><label>${escapeHtml(dep.l)}</label>
                   <input id="_f_${dep.f}" type="text" class="input-dark" value="${v}"></div>`;
        }
        const items=listsMap[dep.f]||[];
        const dc   =listsMap[`_c_${dep.f}`];
        const opts =items.map(i=>{
            const iId=i[dc.idField];
            const iN =dc.nameField?i[dc.nameField]:`${dc.label} #${shortId(iId)}`;
            return`<option value="${escapeHtml(iId)}"
                           ${String(existing[dep.f])===String(iId)?'selected':''}>
                       ${escapeHtml(iN)}</option>`;
        }).join('');
        const hint=items.length===0
            ?`<p style="font-size:.8rem;color:var(--muted);margin-top:.25rem">
               ⚠ Список пуст — сначала создайте эту деталь.</p>`:'';
        return`<div class="form-group"><label>${escapeHtml(dep.l)}</label>
               <select id="_f_${dep.f}" class="input-dark">
                   <option value="">— выберите —</option>${opts}
               </select>${hint}</div>`;
    }).join('');

    const m=createModal(`
        <div class="modal-title">
            <i class="fas fa-puzzle-piece" style="color:var(--accent)"></i>
            ${editId?'Редактировать':'Создать'} — ${escapeHtml(cat.label)}</div>
        ${fieldsHtml}
        <div class="error-banner hidden">
            <i class="fas fa-exclamation-circle"></i><span></span></div>
        <div class="modal-footer">
            <button class="btn btn-secondary" id="_cc">Отмена</button>
            <button class="btn btn-primary"   id="_ss">
                <i class="fas fa-save"></i> Сохранить</button>
        </div>`,true);

    m.querySelector('#_cc').onclick=()=>m.remove();
    m.querySelector('#_ss').onclick=async()=>{
        const payload={};
        deps.forEach(dep=>{
            const el=m.querySelector(`#_f_${dep.f}`);
            payload[dep.f]=el?el.value.trim():'';
        });

        console.debug('[save-composite] category='+cat.key+' payload='+JSON.stringify(payload));

        const miss=deps.find(d=>!payload[d.f]);
        if(miss){setFormError(m,`Заполните поле «${miss.l}»`);return;}
        const btn=m.querySelector('#_ss');
        btn.disabled=true;btn.innerHTML='<i class="fas fa-spinner fa-spin"></i>';
        setFormError(m,'');
        try{
            if(editId) await api[cat.key].update(editId,payload);
            else       await api[cat.key].create(payload);
            await loadCatalogData();m.remove();
        }catch(e){
            setFormError(m,e.message);
            btn.disabled=false;
            btn.innerHTML='<i class="fas fa-save"></i> Сохранить';
        }
    };
}

/* ── Пользователь ──────────────────────────────── */
function openCreateUserForm(){
    const m=createModal(`
        <div class="modal-title">
            <i class="fas fa-user-plus" style="color:var(--accent)"></i>
            Новый пользователь</div>
        <div class="form-group"><label>Логин</label>
            <input id="_cu" type="text" class="input-dark"
                placeholder="username" autocomplete="off"></div>
        <div class="form-group"><label>Пароль</label>
            <input id="_cp" type="password" class="input-dark"
                placeholder="password" autocomplete="new-password"></div>
        <div class="form-group"><label>Роль</label>
            <select id="_cr" class="input-dark">
                <option value="user">user — обычный пользователь</option>
                <option value="admin">admin — администратор</option>
            </select></div>
        <div class="error-banner hidden">
            <i class="fas fa-exclamation-circle"></i><span></span></div>
        <div class="modal-footer">
            <button class="btn btn-secondary" id="_ccc">Отмена</button>
            <button class="btn btn-primary"   id="_css">
                <i class="fas fa-save"></i> Создать</button>
        </div>`);
    m.querySelector('#_ccc').onclick=()=>m.remove();
    m.querySelector('#_css').onclick=async()=>{
        const username=m.querySelector('#_cu').value.trim();
        const password=m.querySelector('#_cp').value.trim();
        const role    =m.querySelector('#_cr').value;
        if(!username||!password){setFormError(m,'Заполните все поля');return;}
        const btn=m.querySelector('#_css');
        btn.disabled=true;btn.innerHTML='<i class="fas fa-spinner fa-spin"></i>';
        setFormError(m,'');
        try{
            await api.users.create({username,password,role});
            m.remove();
        }catch(e){
            setFormError(m,e.message);
            btn.disabled=false;
            btn.innerHTML='<i class="fas fa-save"></i> Создать';
        }
    };
}
