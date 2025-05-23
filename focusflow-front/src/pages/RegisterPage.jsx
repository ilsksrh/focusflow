import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';

export default function RegisterPage() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [repeatPassword, setRepeatPassword] = useState('');
  const [message, setMessage] = useState('');
  const navigate = useNavigate();

  const handleRegister = async (e) => {
    e.preventDefault();
    if (password !== repeatPassword) {
      setMessage('Passwords do not match');
      return;
    }
    try {
      const res = await fetch('http://localhost:8080/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify({
          username,
          password: password,
          role: 'user',
        }),
      });
      const data = await res.json();
      if (!res.ok) throw new Error(data.error || 'Registration failed');
      setMessage('Registered successfully! Redirecting...');
      setTimeout(() => navigate('/login'), 1500);
    } catch (err) {
      setMessage(err.message);
    }
  };

  return (
    <section className="vh-100" style={{ backgroundColor: '#f8f2ff' }}>
      <div className="container h-100">
        <div className="row d-flex justify-content-center align-items-center h-100">
          <div className="col-lg-10 col-xl-9">
            <div className="card text-black" style={{ borderRadius: '20px' }}>
              <div className="card-body p-md-5">
                <div className="row justify-content-center">
                  <div className="col-md-10 col-lg-6 col-xl-5 order-2 order-lg-1">

                    <h2
                      className="text-center fw-bold mb-4"
                      style={{ color: '#6f42c1', fontSize: '2rem' }}
                    >
                      Sign Up
                    </h2>

                    <form className="mx-1 mx-md-4" onSubmit={handleRegister}>

                      <div className="form-outline mb-3">
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

                      <div className="form-outline mb-3">
                        <label className="form-label" htmlFor="repeatPassword">Repeat Password</label>
                        <input
                          type="password"
                          id="repeatPassword"
                          className="form-control"
                          value={repeatPassword}
                          onChange={(e) => setRepeatPassword(e.target.value)}
                          required
                        />
                      </div>

                      <div className="form-check d-flex justify-content-center mb-4">
                        <input
                          className="form-check-input me-2"
                          type="checkbox"
                          required
                          id="terms"
                        />
                        <label className="form-check-label" htmlFor="terms">
                          I agree to the <a href="#!">terms of service</a>
                        </label>
                      </div>

                      <div className="d-grid gap-2 d-md-flex justify-content-md-center mb-3">
                        <button
                          type="submit"
                          className="btn btn-lg px-5 text-white"
                          style={{ backgroundColor: '#6f42c1', border: 'none' }}
                        >
                          Register
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
                        Already have an account?{' '}
                        <a href="/login" style={{ color: '#6f42c1', fontWeight: 'bold' }}>
                          Log in
                        </a>
                      </p>
                    </form>
                  </div>

                  <div className="col-md-10 col-lg-6 col-xl-7 d-flex align-items-center order-1 order-lg-2">
                    <img
                      src="https://mdbcdn.b-cdn.net/img/Photos/new-templates/bootstrap-registration/draw1.webp"
                      className="img-fluid"
                      alt="Illustration"
                    />
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
