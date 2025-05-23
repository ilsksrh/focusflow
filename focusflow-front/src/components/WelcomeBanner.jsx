import React from 'react';

export default function WelcomeBanner({ user }) {
  if (!user) return null;

  return (
    <div className="mb-4">
      <h2 className="fw-bold">Welcome, {user.username} 👋</h2>
      <hr />
    </div>
  );
}
