/* app.js — с сохранением сессии в localStorage */
let currentRole = null;

// ── Восстановление сессии при обновлении страницы ──────────────────────
(function restoreSession() {
    const saved = localStorage.getItem('mispris_role');
    if (saved === 'admin' || saved === 'user') {
        currentRole = saved;
    }
})();

function saveSession(role) {
    if (role) localStorage.setItem('mispris_role', role);
    else      localStorage.removeItem('mispris_role');
}

function showLogin(){
    const app = document.getElementById('app');
    app.innerHTML = `
        <div class="login-wrap">
            <div class="login-card">
                <h2><i class="fas fa-bolt"></i> MISPRIS</h2>
                <p class="text-muted" style="text-align:center;margin-bottom:1.5rem;font-size:0.9rem">
                    Управление сборкой электромобилей
                </p>
                <form id="lf">
                    <div class="form-group">
                        <label>Логин</label>
                        <input type="text" id="lu" class="input-dark" autocomplete="username" required>
                    </div>
                    <div class="form-group">
                        <label>Пароль</label>
                        <input type="password" id="lp" class="input-dark" autocomplete="current-password" required>
                    </div>
                    <button type="submit" class="btn btn-primary"
                        style="width:100%;justify-content:center;padding:0.65rem">
                        <i class="fas fa-sign-in-alt"></i> Войти
                    </button>
                    <div id="le" class="text-danger"
                        style="text-align:center;margin-top:0.75rem;font-size:0.88rem"></div>
                </form>
            </div>
        </div>`;

    document.getElementById('lf').addEventListener('submit', async e => {
        e.preventDefault();
        const username = document.getElementById('lu').value.trim();
        const password = document.getElementById('lp').value;
        document.getElementById('le').textContent = '';
        try {
            const data = await api.auth.login(username, password);
            currentRole = data.role;
            saveSession(currentRole);          // ← сохраняем роль
            categoryFilter = 'all';
            searchText = '';
            await loadCatalogData();
        } catch(_) {
            document.getElementById('le').textContent = 'Неверный логин или пароль';
        }
    });
}

document.addEventListener('DOMContentLoaded', async () => {
    if (currentRole) {
        // Роль восстановлена из localStorage — сразу загружаем каталог
        try {
            await loadCatalogData();
        } catch(_) {
            // Если API недоступен — показываем логин
            saveSession(null);
            currentRole = null;
            showLogin();
        }
    } else {
        showLogin();
    }
});
