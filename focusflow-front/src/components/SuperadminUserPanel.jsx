import React, { useEffect, useState } from 'react';
import {
  getTeams,
  createTeam,
  deleteTeam,
  assignUserToTeam,
  removeUserFromTeam,
} from '../api/superadmin';
import { getAllUsers } from '../api/admin';

export default function SuperadminUserPanel() {
  const [teams, setTeams] = useState([]);
  const [users, setUsers] = useState([]);
  const [newTeamName, setNewTeamName] = useState('');
  const [assignUserId, setAssignUserId] = useState('');
  const [assignTeamId, setAssignTeamId] = useState('');

  useEffect(() => {
    loadData();
  }, []);

  const loadData = async () => {
    const [teamList, userList] = await Promise.all([getTeams(), getAllUsers()]);
    setTeams(teamList);
    setUsers(userList);
  };

  const handleCreateTeam = async () => {
    if (!newTeamName) return;
    await createTeam({ name: newTeamName });
    setNewTeamName('');
    loadData();
  };

  const handleDeleteTeam = async (id) => {
    if (window.confirm('Удалить команду?')) {
      await deleteTeam(id);
      loadData();
    }
  };

  const handleAssign = async () => {
    if (!assignUserId || !assignTeamId) return;
    await assignUserToTeam({
      user_id: parseInt(assignUserId),
      team_id: parseInt(assignTeamId),
    });
    setAssignUserId('');
    setAssignTeamId('');
    loadData();
  };

  const handleRemoveUser = async (userId, username) => {
    if (window.confirm(`Удалить ${username} из команды?`)) {
      await removeUserFromTeam(userId);
      loadData();
    }
  };

  return (
    <div>
      <h3>Управление командами</h3>

      {/* Создание команды */}
      <div className="mb-4">
        <label className="form-label">Создать новую команду:</label>
        <div className="d-flex gap-2">
          <input
            className="form-control"
            value={newTeamName}
            onChange={(e) => setNewTeamName(e.target.value)}
            placeholder="Название команды"
          />
          <button className="btn btn-success" onClick={handleCreateTeam}>
            Создать
          </button>
        </div>
      </div>

      {/* Назначение пользователя */}
      <div className="mb-4">
        <label className="form-label">Назначить пользователя в команду:</label>
        <div className="d-flex gap-2">
          <select
            className="form-select"
            value={assignUserId}
            onChange={(e) => setAssignUserId(e.target.value)}
          >
            <option value="">Выберите пользователя</option>
            {users.map((u) => (
              <option key={u.id} value={u.id}>
                {u.username}
              </option>
            ))}
          </select>

          <select
            className="form-select"
            value={assignTeamId}
            onChange={(e) => setAssignTeamId(e.target.value)}
          >
            <option value="">Выберите команду</option>
            {teams.map((t) => (
              <option key={t.id} value={t.id}>
                {t.name}
              </option>
            ))}
          </select>

          <button className="btn btn-primary" onClick={handleAssign}>
            Назначить
          </button>
        </div>
      </div>

      {/* Список команд */}
      <h5>Существующие команды:</h5>
      {teams.map((team) => (
        <div key={team.id} className="card p-3 mb-3">
          <div className="d-flex justify-content-between">
            <strong>{team.name}</strong>
            <button
              onClick={() => handleDeleteTeam(team.id)}
              className="btn btn-sm btn-outline-danger"
            >
              Удалить
            </button>
          </div>

          <div className="mt-2">
            <small>Участники:</small>
            <ul className="mb-0">
              {team.Users?.length > 0 ? (
                team.Users.map((u) => (
                  <li
                    key={u.id}
                    className="d-flex justify-content-between align-items-center"
                  >
                    <span>{u.username}</span>
                    <button
                      className="btn btn-sm btn-outline-danger"
                      onClick={() => handleRemoveUser(u.id, u.username)}
                    >
                      Удалить
                    </button>
                  </li>
                ))
              ) : (
                <li className="text-muted">Нет участников</li>
              )}
            </ul>
          </div>
        </div>
      ))}
    </div>
  );
}
