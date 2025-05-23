export async function uploadProfilePicture(file) {
    const formData = new FormData();
    formData.append('profile_picture', file);
  
    const res = await fetch('http://localhost:8080/profile/upload_picture', {
      method: 'POST',
      credentials: 'include',
      body: formData,
    });
  
    if (!res.ok) {
      throw new Error('Ошибка загрузки фото');
    }
  
    return await res.json();
  }
  
  export async function updatePassword(newPassword) {
    const res = await fetch('http://localhost:8080/profile/change_password', {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
      body: JSON.stringify({ new_password: newPassword }),
    });
  
    if (!res.ok) {
      throw new Error('Ошибка смены пароля');
    }
  
    return await res.json();
  }
  