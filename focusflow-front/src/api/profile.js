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
  
export async function updatePassword(oldPassword, newPassword) {
  const res = await fetch('http://localhost:8080/profile/change_password', {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: JSON.stringify({ old_password: oldPassword, new_password: newPassword }),
  });

  if (!res.ok) {
    const error = await res.json().catch(() => ({}));
    throw new Error(error?.error || 'Failed to update password');
  }

  return await res.json();
}
