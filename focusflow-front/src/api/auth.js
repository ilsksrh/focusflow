
export async function getCurrentUser() {
    const res = await fetch('http://localhost:8080/me', {
      credentials: 'include',
    });
    if (!res.ok) throw new Error('Not authenticated');
    return await res.json();
  }
  