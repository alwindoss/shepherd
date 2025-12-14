import { useState } from "react";
import { useNavigate } from "react-router";
import { useAuth } from "../contexts/AuthContext";

export default function LoginPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string | null>(null);
  const navigate = useNavigate();
  const { setUser, fetchUser } = useAuth();

  async function submit(e: React.FormEvent) {
    e.preventDefault();
    setError(null);
    try {
      const res = await fetch("/api/login", {
        method: "POST",
        headers: { "content-type": "application/json" },
        body: JSON.stringify({ email, password }),
      });
      const json = await res.json();
      if (!res.ok) throw new Error(json.error || "login failed");
      setUser(json.user || null);
      await fetchUser();
      navigate("/dashboard");
    } catch (err: any) {
      setError(err.message);
    }
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-base-200">
      <div className="card w-full max-w-md shadow-xl bg-base-100">
        <div className="card-body">
          <h2 className="card-title">Sign in</h2>
          <form onSubmit={submit} className="space-y-4">
            <div className="form-control">
              <label className="label"><span className="label-text">Email</span></label>
              <input type="email" className="input input-bordered" value={email} onChange={e => setEmail(e.target.value)} required />
            </div>
            <div className="form-control">
              <label className="label"><span className="label-text">Password</span></label>
              <input type="password" className="input input-bordered" value={password} onChange={e => setPassword(e.target.value)} required />
            </div>
            {error && <div className="text-sm text-error">{error}</div>}
            <div className="form-control mt-6">
              <button className="btn btn-primary">Sign in</button>
            </div>
          </form>

          <div className="divider">OR</div>

          <div>
            <button className="btn btn-outline w-full" onClick={() => (window.location.href = "/auth/google/login")}> 
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" className="inline mr-2 h-5 w-5"><path fill="#EA4335" d="M24 9.5c3.6 0 6.2 1.6 7.6 2.8l5.7-5.5C34.2 4 29.6 2.5 24 2.5 14.8 2.5 7.2 7.7 3.8 14.9l6.6 5.1C12.4 14.3 17.6 9.5 24 9.5z"></path><path fill="#34A853" d="M46.5 24.5c0-1.6-.2-2.8-.5-4.1H24v8h12.7c-.6 3.1-2.5 6-5.4 7.8l6.6 5.1C43.9 37.4 46.5 31.5 46.5 24.5z"></path><path fill="#4A90E2" d="M10.4 29.9c-1.1-3.1-1.1-6.7 0-9.8l-6.6-5.1C.7 17.1 0 20.6 0 24s.7 6.9 3.8 9.9l6.6-4z"></path><path fill="#FBBC05" d="M24 46.5c6.6 0 12.2-2.1 16.3-5.7l-6.6-5.1c-2 1.3-4.7 2.1-9.7 2.1-6.4 0-11.6-4.8-12.9-11.1l-6.7 5C7.2 40 14.8 46.5 24 46.5z"></path></svg>
              Continue with Google
            </button>
          </div>

          <p className="text-sm text-center mt-4">Don't have an account? <a className="link link-primary" href="/signup">Sign up</a></p>
        </div>
      </div>
    </div>
  )
}
