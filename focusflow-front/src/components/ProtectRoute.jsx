// src/components/ProtectedRoute.jsx
import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { getCurrentUser } from '../api/auth';

export default function ProtectedRoute({ children }) {
  const [checking, setChecking] = useState(true);
  const [allowed, setAllowed] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    async function checkAuth() {
      try {
        await getCurrentUser();
        setAllowed(true);
      } catch {
        navigate('/');
      } finally {
        setChecking(false);
      }
    }

    checkAuth();
  }, [navigate]);

  if (checking) return null; 
  return allowed ? children : null;
}
