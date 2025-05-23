const BASE_URL = 'http://localhost:8080';

// Получить все задачи пользователя
export async function getAllTasks() {
  const res = await fetch(`${BASE_URL}/tasks`, {
    credentials: 'include',
  });
  return await res.json();
}

// Получить задачи по цели
export async function getTasksByGoalID(goalId) {
  const res = await fetch(`${BASE_URL}/goals/${goalId}/tasks`, {
    credentials: 'include',
  });
  return await res.json();
}

// Получить одну задачу по ID
export async function getTaskByID(taskId) {
  const res = await fetch(`${BASE_URL}/tasks/${taskId}`, {
    credentials: 'include',
  });
  return await res.json();
}

// Создать задачу
export async function createTask(data) {
  const res = await fetch(`${BASE_URL}/tasks`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
    body: JSON.stringify(data),
  });

  if (!res.ok) throw new Error('Ошибка при создании задачи');
  return await res.json();
}

// Обновить задачу
export async function updateTask(taskId, data) {
  const res = await fetch(`${BASE_URL}/tasks/${taskId}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
    body: JSON.stringify(data),
  });

  if (!res.ok) throw new Error('Ошибка при обновлении задачи');
  return await res.json();
}

// Удалить задачу
export async function deleteTask(taskId) {
  const res = await fetch(`http://localhost:8080/tasks/${taskId}`, {
    method: 'DELETE',
    credentials: 'include',
  });

  if (!res.ok) {
    const text = await res.text(); // для отладки
    console.error('Ответ сервера:', text);
    throw new Error('Ошибка при удалении задачи');
  }

  return await res.json();
}


// Пометить задачу как выполненную
export async function markTaskDone(taskId, isDone) {
  const res = await fetch(`${BASE_URL}/tasks/${taskId}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
    body: JSON.stringify({ is_done: isDone }),
  });

  if (!res.ok) throw new Error('Ошибка при обновлении статуса');
  return await res.json();
}

export async function getTeamTasks() {
  const res = await fetch('http://localhost:8080/tasks/team', {
    credentials: 'include',
  });
  if (!res.ok) throw new Error('Ошибка получения задач команды');
  return res.json();
}

export async function getAllTeamTasks() {
  const res = await fetch('http://localhost:8080/tasks/kanban/all_teams', {
    credentials: 'include',
  });
  if (!res.ok) throw new Error('Ошибка получения командных задач');
  return res.json();
}
