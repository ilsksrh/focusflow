import React, { useState, useEffect } from 'react';
import { updatePassword, uploadProfilePicture } from '../api/profile';
import { getCurrentUser } from '../api/auth';
import { getMyTeam } from '../api/superadmin';


export default function ProfilePanel({ user, onUpdated }) {
  const [currentPassword, setCurrentPassword] = useState('');
  const [newPassword, setNewPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [image, setImage] = useState(null);
  const [message, setMessage] = useState('');
  const [error, setError] = useState('');
  const [teamMembers, setTeamMembers] = useState([]);
  const [teamName, setTeamName] = useState('');



  useEffect(() => {
    async function fetchTeam() {
      try {
        const team = await getMyTeam();
        setTeamMembers(team.Users || []);
        setTeamName(team.name || '');

      } catch (e) {
        console.error('Ошибка загрузки команды:', e);
      }
    }

    fetchTeam();
  }, []);

  const validatePassword = (password) => {
    if (password.length < 6) return 'Минимум 6 символов';
    if (!/[A-Z]/.test(password)) return 'Хотя бы одна заглавная буква';
    if (!/\d/.test(password)) return 'Хотя бы одна цифра';
    return '';
  };

  const handlePasswordChange = async () => {
    setMessage('');
    setError('');

    if (!currentPassword) {
      setError('Введите текущий пароль');
      return;
    }

    if (newPassword !== confirmPassword) {
      setError('Пароли не совпадают');
      return;
    }

    const validationError = validatePassword(newPassword);
    if (validationError) {
      setError(validationError);
      return;
    }

    try {
      await updatePassword(currentPassword, newPassword);
      setMessage('Пароль успешно обновлён');
      setCurrentPassword('');
      setNewPassword('');
      setConfirmPassword('');
    } catch (err) {
      setError(err.message || 'Ошибка обновления пароля');
    }
  };

  const handlePictureUpload = async () => {
    setError('');
    setMessage('');

    if (!image) {
      setError('Выберите изображение');
      return;
    }

    try {
      await uploadProfilePicture(image);
      setMessage('Фото обновлено');
      if (onUpdated) {
        const updatedUser = await getCurrentUser();
        onUpdated(updatedUser);
      }
    } catch (err) {
      setError(err.message || 'Ошибка загрузки фото');
    }
  };

  return (
    <div className="container mt-4">
      <div className="row">
        {/* Левая колонка: профиль */}
        <div className="col-md-6 mb-4">
          <div className="card shadow-sm p-4">
            <div className="d-flex align-items-center mb-4">
              <img
                src={user.profile_picture ? `http://localhost:8080/${user.profile_picture}?t=${Date.now()}` : '/images/default-avatar.png'}
                alt="Avatar"
                className="rounded-circle me-4 border"
                style={{ width: '100px', height: '100px', objectFit: 'cover' }}
              />
              <div>
                <h5 className="mb-1">{user.username}</h5>
                <p className="mb-0"><strong>Роль:</strong> {user.role}</p>
                <p className="mb-0"><strong>Блокировка:</strong> {user.is_blocked ? 'Да' : 'Нет'}</p>
              </div>
            </div>

            <hr />

            <h5 className="mb-3"><i className="fa fa-lock me-2"></i>Сменить пароль</h5>
            <input
              type="password"
              className="form-control mb-2"
              placeholder="Текущий пароль"
              value={currentPassword}
              onChange={(e) => setCurrentPassword(e.target.value)}
            />
            <input
              type="password"
              className="form-control mb-2"
              placeholder="Новый пароль"
              value={newPassword}
              onChange={(e) => setNewPassword(e.target.value)}
            />
            <input
              type="password"
              className="form-control mb-3"
              placeholder="Повторите новый пароль"
              value={confirmPassword}
              onChange={(e) => setConfirmPassword(e.target.value)}
            />
            <button className="btn btn-primary w-100 mb-4" onClick={handlePasswordChange}>
              {/* <i className="fa fa-save me-2"></i> */}
              Сохранить пароль
            </button>

            <h5 className="mb-3"><i className="fa fa-image me-2"></i>Фото профиля</h5>
            <input
              type="file"
              className="form-control mb-2"
              onChange={(e) => setImage(e.target.files[0])}
            />
            <button className="btn btn-outline-secondary w-100" onClick={handlePictureUpload}>
              {/* <i className="fa fa-upload me-2"></i> */}
              Загрузить фото
            </button>

            {message && <div className="alert alert-success mt-4 mb-0">{message}</div>}
            {error && <div className="alert alert-danger mt-4 mb-0">{error}</div>}
          </div>
        </div>

        {/* Правая колонка: заглушка / доп. секция */}
        <div className="col-md-6 mb-4">
          <div className="card shadow-sm p-4 h-100">
           
              <h5><i className="fa fa-users me-2"></i>Моя команда</h5>
                {teamName && <p className="fw-bold mb-3">Название: {teamName}</p>}
                {teamMembers.length > 0 ? (
                  <ul className="list-group">
                    {teamMembers.map((member) => (
                      <li key={member.id} className="list-group-item d-flex justify-content-between">
                        {member.username}
                        <span className="badge bg-secondary">{member.role}</span>
                      </li>
                    ))}
                  </ul>
                ) : (
                  <p className="text-muted">Вы не состоите в команде</p>
                )}

               <hr className="my-4" />
              <h5><i className="fa fa-history me-2"></i>История активности</h5>
            
            <p className="text-muted">Здесь можно отобразить:</p>
            <ul className="small text-muted">
              <li> Последние входы</li>
              <li> Изменения профиля</li>
              <li> Загрузки фото</li>
              <li>Попытки смены пароля</li>
              <li> IP-адреса или геолокация (если будет)</li>
            </ul>
            <p className="text-muted">...или любую другую важную статистику.</p>
          </div>
        </div>
      </div>
    </div>
  );
}
