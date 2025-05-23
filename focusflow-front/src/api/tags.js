const BASE_URL = 'http://localhost:8080/tags';

export async function getTags() {
  const res = await fetch(BASE_URL, {
    credentials: 'include',
  });
  return await res.json();
}

export async function createTag(tag) {
  const res = await fetch(BASE_URL, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: JSON.stringify(tag),
  });

  if (!res.ok) {
    const errText = await res.text();
    throw new Error(errText || 'Ошибка создания тега');
  }

  return await res.json();
}

export async function deleteTag(id) {
  const res = await fetch(`${BASE_URL}/${id}`, {
    method: 'DELETE',
    credentials: 'include',
  });

  if (!res.ok) {
    const errText = await res.text();
    throw new Error(errText || 'Ошибка удаления тега');
  }

  return true;
}

export async function updateTag(id, updatedTag) {
  const res = await fetch(`${BASE_URL}/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: JSON.stringify(updatedTag),
  });

  if (!res.ok) {
    const errText = await res.text();
    throw new Error(errText || 'Ошибка обновления тега');
  }

  return await res.json();
}


