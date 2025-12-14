import React, { createContext, useContext, useEffect, useState } from "react";
import { useNavigate } from "react-router";

type User = { email: string; name: string } | null;

type AuthContextType = {
  user: User;
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
  const navigate = useNavigate();

  async function fetchUser() {
    try {
      const res = await fetch("/api/me");
      if (!res.ok) {
        setUserState(null);
        return;
      }
      const json = await res.json();
      setUserState(json.user || null);
    } catch (e) {
      setUserState(null);
    }
  }

  async function logout() {
    try {
      await fetch("/api/logout", { method: "POST" });
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
    <AuthContext.Provider value={{ user, fetchUser, setUser: setUserState, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export default AuthContext;
