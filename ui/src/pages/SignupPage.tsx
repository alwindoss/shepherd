import { useState } from "react";
import { useNavigate } from "react-router";
import { useAuth } from "../contexts/AuthContext";

export default function SignupPage() {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string | null>(null);
  const navigate = useNavigate();
  const { setUser, fetchUser } = useAuth();

  async function submit(e: React.FormEvent) {
    e.preventDefault();
    setError(null);
    try {
      const res = await fetch("/api/signup", {
        method: "POST",
        headers: { "content-type": "application/json" },
        credentials: "include",
        body: JSON.stringify({ name, email, password }),
      });
      const json = await res.json();
      if (!res.ok) throw new Error(json.error || "signup failed");
      setUser(json.user || null);
      // refresh background auth state but don't block navigation
      fetchUser().catch(() => {});
      navigate("/dashboard");
    } catch (err: any) {
      setError(err.message);
    }
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-base-200">
      <div className="card w-full max-w-md shadow-xl bg-base-100">
        <div className="card-body">
          <h2 className="card-title">Create account</h2>
          <form onSubmit={submit} className="space-y-4">
            <div className="form-control">
              <label className="label"><span className="label-text">Full name</span></label>
              <input type="text" className="input input-bordered" value={name} onChange={e => setName(e.target.value)} required />
            </div>
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
              <button className="btn btn-primary">Create account</button>
            </div>
          </form>

          <div className="divider">OR</div>

          <div>
            <button className="btn btn-outline w-full" onClick={() => (window.location.href = "/auth/google/login")}>Sign up with Google</button>
          </div>

          <p className="text-sm text-center mt-4">Already have an account? <a className="link link-primary" href="/login">Sign in</a></p>
        </div>
      </div>
    </div>
  )
}
