export async function getAllUsers() {
    const res = await fetch('http://localhost:8080/admin/users', { credentials: 'include' });
    return await res.json();
  }
  
  export async function addUser(user) {
    const res = await fetch('http://localhost:8080/admin/add_user', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
      body: JSON.stringify({
        username: user.username,
        password_hash: user.password, 
        role: user.role,
      }),
    });
  
    if (!res.ok) {
      const errText = await res.text();
      throw new Error(errText || 'Ошибка добавления пользователя');
    }
  
    return res.json();
  }
  
  
  export async function deleteUser(username) {
    await fetch(`http://localhost:8080/admin/delete_user?username=${username}`, {
      method: 'DELETE',
      credentials: 'include',
    });
  }
  
  export async function blockUser(username) {
    await fetch(`http://localhost:8080/admin/block_user?username=${username}`, {
      method: 'POST',
      credentials: 'include',
    });
  }
  
  export async function unblockUser(username) {
    await fetch(`http://localhost:8080/admin/unblock_user?username=${username}`, {
      method: 'POST',
      credentials: 'include',
    });
  }
  