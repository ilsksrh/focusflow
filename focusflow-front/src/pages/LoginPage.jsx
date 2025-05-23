import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';

export default function LoginPage() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');
  const navigate = useNavigate();

  const handleLogin = async (e) => {
    e.preventDefault();
    try {
      const res = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify({ username, password }),
      });
      const data = await res.json();
      if (!res.ok) throw new Error(data.error || 'Login failed');
      setMessage('Login successful! Redirecting...');
      setTimeout(() => navigate('/dashboard'), 1500);
    } catch (err) {
      setMessage(err.message);
    }
  };

  return (
    <section className="vh-100" style={{ backgroundColor: '#f3f1fa' }}>
      <div className="container h-100">
        <div className="row d-flex justify-content-center align-items-center h-100">
          <div className="col-lg-10 col-xl-9">
            <div className="card text-black" style={{ borderRadius: '20px' }}>
              <div className="card-body p-md-5">
                <div className="row justify-content-center">
                  
                  {/* ИЗОБРАЖЕНИЕ — слева */}
                  <div className="col-md-10 col-lg-6 col-xl-6 d-flex align-items-center order-1 order-lg-1">
                    <img
                      src="https://i.pinimg.com/736x/25/d0/af/25d0af05305e58730a448f00d7ec97d4.jpg  "
                      className="img-fluid"
                      alt="Login illustration"
                    />
                  </div>

                  {/* ФОРМА — справа */}
                  <div className="col-md-10 col-lg-6 col-xl-5 order-2 order-lg-2">

                  <h2
                      className="text-center fw-bold mb-4"
                      style={{ color: '#6f42c1', fontSize: '2rem' }}
                    >
                      Log in
                    </h2>

                    <form className="mx-1 mx-md-4" onSubmit={handleLogin}>
                      <div className="form-outline mb-2">
                        <label className="form-label" htmlFor="username">Username</label>
                        <input
                          type="text"
                          id="username"
                          className="form-control"
                          value={username}
                          onChange={(e) => setUsername(e.target.value)}
                          required
                        />
                      </div>

                      <div className="form-outline mb-3">
                        <label className="form-label" htmlFor="password">Password</label>
                        <input
                          type="password"
                          id="password"
                          className="form-control"
                          value={password}
                          onChange={(e) => setPassword(e.target.value)}
                          required
                        />
                      </div>

                      <div className="d-grid gap-2 d-md-flex justify-content-md-center mb-3">
                        <button
                          type="submit"
                          className="btn btn-lg px-5 text-white"
                          style={{ backgroundColor: '#6f42c1', border: 'none' }}
                        >
                          Log in
                        </button>
                        <button
                          type="button"
                          className="btn btn-outline-secondary btn-lg px-4"
                          onClick={() => navigate('/')}
                        >
                          Back
                        </button>
                      </div>

                      {message && (
                        <div className="alert alert-info text-center">{message}</div>
                      )}

                      <p className="text-center mt-3">
                        Don't have an account?{' '}
                        <a href="/register" style={{ color: '#6f42c1', fontWeight: 'bold' }}>
                          Sign up
                        </a>
                      </p>
                    </form>
                  </div>

                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
