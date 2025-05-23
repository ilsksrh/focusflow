import React, { useState } from 'react';
import { updatePassword, uploadProfilePicture } from '../api/profile';
import { getCurrentUser } from '../api/auth';

export default function ProfilePanel({ user, onUpdated }) {
  const [newPassword, setNewPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [image, setImage] = useState(null);
  const [message, setMessage] = useState('');
  const [error, setError] = useState('');

  const validatePassword = (password) => {
    if (password.length < 6) {
      return 'Password must be at least 6 characters';
    }
    if (!/[A-Z]/.test(password)) {
      return 'Password must contain at least one uppercase letter';
    }
    if (!/\d/.test(password)) {
      return 'Password must contain at least one digit';
    }
    return '';
  };

  const handlePasswordChange = async () => {
    setMessage('');
    setError('');

    if (newPassword !== confirmPassword) {
      setError('Passwords do not match');
      return;
    }

    const validationError = validatePassword(newPassword);
    if (validationError) {
      setError(validationError);
      return;
    }

    try {
      await updatePassword(newPassword);
      setMessage('Password updated successfully');
      setNewPassword('');
      setConfirmPassword('');
    } catch (err) {
      setError(err.message || 'Failed to update password');
    }
  };

  const handlePictureUpload = async () => {
    setError('');
    setMessage('');

    if (!image) {
      setError('No image selected');
      return;
    }

    try {
      await uploadProfilePicture(image);
      setMessage('Profile picture updated successfully');
      if (onUpdated) {
        const updatedUser = await getCurrentUser();
        onUpdated(updatedUser);
      }
    } catch (err) {
      setError(err.message || 'Failed to upload photo');
    }
  };

  return (
    <div>
      <h4>Profile Info</h4>
      <img
        src={user.profile_picture ? `http://localhost:8080/${user.profile_picture}?t=${Date.now()}` : '/images/default-avatar.png'}
        alt="Avatar"
        className="rounded mb-2"
        width={80}
      />
      <p><strong>Username:</strong> {user.username}</p>
      <p><strong>Role:</strong> {user.role}</p>
      <p><strong>Blocked:</strong> {user.is_blocked ? 'Yes' : 'No'}</p>

      <hr />
      <h5>Change Password</h5>
      <input
        type="password"
        className="form-control mb-2"
        placeholder="New password"
        value={newPassword}
        onChange={(e) => setNewPassword(e.target.value)}
      />
      <input
        type="password"
        className="form-control mb-2"
        placeholder="Repeat new password"
        value={confirmPassword}
        onChange={(e) => setConfirmPassword(e.target.value)}
      />
      <button className="btn btn-sm btn-primary mb-3" onClick={handlePasswordChange}>Save Password</button>

      <h5>Upload Photo</h5>
      <input
        type="file"
        className="form-control mb-2"
        onChange={(e) => setImage(e.target.files[0])}
      />
      <button className="btn btn-sm btn-secondary" onClick={handlePictureUpload}>Upload Photo</button>

      {message && <p className="mt-3 text-success">{message}</p>}
      {error && <p className="mt-3 text-danger">{error}</p>}
    </div>
  );
}
