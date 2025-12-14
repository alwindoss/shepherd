import React, { createContext, useContext, useEffect, useState } from "react";
import { useNavigate } from "react-router";

type User = { email: string; name: string } | null;

type AuthContextType = {
  user: User;
  loading: boolean;
  fetchUser: () => Promise<void>;
  setUser: (u: User) => void;
  logout: () => Promise<void>;
};

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const useAuth = () => {
  const ctx = useContext(AuthContext);
  if (!ctx) throw new Error("useAuth must be used within AuthProvider");
  return ctx;
};

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [user, setUserState] = useState<User>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const navigate = useNavigate();

  async function fetchUser() {
    setLoading(true);
    try {
      const res = await fetch("/api/me", { credentials: "include" });
      if (!res.ok) {
        // If we already have a user (e.g., right after a successful login request), don't clobber it
        if (user === null) {
          setUserState(null);
        }
        setLoading(false);
        return;
      }
      const json = await res.json();
      setUserState(json.user || null);
    } catch (e) {
      setUserState(null);
    } finally {
      setLoading(false);
    }
  }

  async function logout() {
    try {
      await fetch("/api/logout", { method: "POST", credentials: "include" });
    } catch (e) {
      // ignore
    }
    setUserState(null);
    navigate("/");
  }

  useEffect(() => {
    fetchUser();
  }, []);

  return (
    <AuthContext.Provider value={{ user, loading, fetchUser, setUser: setUserState, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export default AuthContext;
