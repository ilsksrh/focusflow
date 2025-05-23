const BASE_URL = 'http://localhost:8080';

    const fetchWithAuth = async (url, options = {}) => {
    const res = await fetch(url, {
        credentials: 'include',
        headers: { 'Content-Type': 'application/json' },
        ...options,
    });

    if (!res.ok) {
        const errorText = await res.text();
        throw new Error(errorText || 'Ошибка запроса');
    }

    return res.json();
    };

        export const getTeams = async () => {
        return await fetchWithAuth(`${BASE_URL}/superadmin/teams`);
        };

    // Создать команду
    export const createTeam = async (data) => {
    return await fetchWithAuth(`${BASE_URL}/superadmin/teams`, {
        method: 'POST',
        body: JSON.stringify(data),
    });
    };

    // Удалить команду
    export const deleteTeam = async (id) => {
    return await fetchWithAuth(`${BASE_URL}/superadmin/teams?id=${id}`, {
        method: 'DELETE',
    });
    };

    // Назначить пользователя в команду
    export const assignUserToTeam = async ({ user_id, team_id }) => {
    return await fetchWithAuth(`${BASE_URL}/superadmin/assign`, {
        method: 'POST',
        body: JSON.stringify({ user_id, team_id }),
    });
    };

    // Получить все теги, используемые задачами команды
    export const getTagsByTeam = async (teamId) => {
    return await fetchWithAuth(`${BASE_URL}/superadmin/team_tags?team_id=${teamId}`);
    };

    export const removeUserFromTeam = async (user_id) => {
    return await fetchWithAuth(`${BASE_URL}/superadmin/remove_from_team`, {
        method: 'POST',
        body: JSON.stringify({ user_id }),
    });
    };

    export async function getMyTeam() {
    const res = await fetch('http://localhost:8080/my_team', {
        credentials: 'include',
    });

    if (!res.ok) {
        throw new Error('Ошибка получения команды');
    }

    return await res.json();
    }

        
