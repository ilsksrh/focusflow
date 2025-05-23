import React from 'react';

export default function Sidebar({ viewMode, setViewMode, onProfileClick, setShowProfile, user }) {
  return (
    <div className="p-3 d-flex flex-column h-100">

      <div className="d-flex flex-column align-items-center mb-3">
        <img
          src={user?.profile_picture ? `http://localhost:8080/${user.profile_picture}` : '/images/default-avatar.png'}
          alt="Avatar"
          className="rounded-circle"
          width={60}
          height={60}
        />
        <p className="mt-2 fw-bold">Profile</p>
        <button
          className="btn btn-sm btn-outline-secondary w-100"
          onClick={() => {
            setShowProfile(true);
          }}
        >
          View Profile
        </button>
      </div>


      <h5>View Mode</h5>
      <ul className="nav flex-column">
        <li className="nav-item">
          <button
            className={`btn btn-sm mb-2 w-100 ${viewMode === 'calendar' ? 'btn-primary' : 'btn-outline-primary'}`}
            onClick={() => {
              setViewMode('calendar');
              setShowProfile(false);
            }}
          >
            📅 Calendar
          </button>
        </li>
        <li className="nav-item">
          <button
            className={`btn btn-sm mb-2 w-100 ${viewMode === 'kanban' ? 'btn-primary' : 'btn-outline-primary'}`}
            onClick={() => {
              setViewMode('kanban');
              setShowProfile(false);
            }}
          >
            📋 Kanban
          </button>
        </li>
        <li className="nav-item">
          <button
            className={`btn btn-sm mb-2 w-100 ${viewMode === 'list' ? 'btn-primary' : 'btn-outline-primary'}`}
            onClick={() => {
              setViewMode('list');
              setShowProfile(false);
            }}
          >
            📃 List
          </button>
        </li>

        {user?.role === 'admin' && (
          <li className="nav-item">
            <button
              className={`btn btn-sm mb-2 w-100 ${viewMode === 'admin' ? 'btn-primary' : 'btn-outline-primary'}`}
              onClick={() => {
                setViewMode('admin');
                setShowProfile(false);
              }}
            >
              👥 Manage Users
            </button>
          </li>
        )}

        <li className="nav-item mt-auto">
          <button
            className="btn btn-sm btn-outline-danger w-100"
            onClick={() => {
              if (window.confirm('Are you sure you want to log out?')) {
                fetch('http://localhost:8080/logout', {
                  method: 'POST',
                  credentials: 'include',
                }).then(() => {
                  window.location.href = '/';
                });
              }
            }}
          >
            🔒 Logout
          </button>
        </li>
      </ul>
    </div>
  );
}
