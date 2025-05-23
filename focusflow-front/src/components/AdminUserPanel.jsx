import React, { useEffect, useState } from 'react';
import { getAllUsers, deleteUser, blockUser, addUser, unblockUser } from '../api/admin';

export default function AdminUserPanel() {
  const [users, setUsers] = useState([]);
  const [newUser, setNewUser] = useState({ username: '', password: '', role: 'user' });
  const [errorMsg, setErrorMsg] = useState('');

  useEffect(() => {
    fetchUsers();
  }, []);

  const fetchUsers = async () => {
    const res = await getAllUsers();
    setUsers(res);
  };

  const handleAdd = async () => {
    try {
      await addUser(newUser);
      setNewUser({ username: '', password: '', role: 'user' });
      setErrorMsg('');
      fetchUsers();
    } catch (err) {
      setErrorMsg(err.message);
    }
  };

  const handleDelete = async (username) => {
    if (!username) return;
    await deleteUser(username);
    fetchUsers();
  };

  const handleBlock = async (username) => {
    await blockUser(username);
    fetchUsers();
  };

  const handleUnblock = async (username) => {
    await unblockUser(username);
    fetchUsers();
  };
  const activeUsers = users.filter(u => !u.is_blocked);
  const blockedUsers = users.filter(u => u.is_blocked);

  return (
    <div>
      <h4>Manage Users</h4>

      <div className="mb-3 d-flex gap-2 align-items-start">
        <input placeholder="Username" value={newUser.username} onChange={e => setNewUser({ ...newUser, username: e.target.value })} />
        <input placeholder="Password" value={newUser.password} onChange={e => setNewUser({ ...newUser, password: e.target.value })} />
        <select value={newUser.role} onChange={e => setNewUser({ ...newUser, role: e.target.value })}>
          <option value="user">User</option>
          <option value="admin">Admin</option>
          <option value="superadmin">Super Admin</option>
        </select>
        <div className="d-flex flex-column">
          <button onClick={handleAdd} className="btn btn-primary btn-sm">Add</button>
          {errorMsg && <span className="text-danger small mt-1" style={{ maxWidth: '200px' }}>{errorMsg}</span>}
        </div>
      </div>

      <h5>🟢 Active Users</h5>
      <table className="table table-bordered">
        <thead>
          <tr><th>Username</th><th>Role</th><th>Blocked</th><th>Actions</th></tr>
        </thead>
        <tbody>
          {activeUsers.map((u) => (
            <tr key={u.id}>
              <td>{u.username}</td>
              <td>{u.role}</td>
              <td>{u.is_blocked ? 'Yes' : 'No'}</td>
              <td>
                <button className="btn btn-sm btn-danger me-2" onClick={() => handleDelete(u.username)}>Delete</button>
                <button className="btn btn-sm btn-warning" onClick={() => handleBlock(u.username)}>Block</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>

      <h5 className="mt-4 text-danger">🔒 Blocked Users</h5>
      <table className="table table-bordered">
        <thead>
          <tr><th>Username</th><th>Role</th><th>Actions</th></tr>
        </thead>
        <tbody>
          {blockedUsers.map((u) => (
            <tr key={u.id}>
              <td>{u.username}</td>
              <td>{u.role}</td>
              <td>
                <button className="btn btn-sm btn-danger me-2" onClick={() => handleDelete(u.username)}>Delete</button>
                <button className="btn btn-sm btn-success" onClick={() => handleUnblock(u.username)}>Unblock</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
