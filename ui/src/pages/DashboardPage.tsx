import { useAuth } from "../contexts/AuthContext";
import { useEffect } from "react";
import { useNavigate } from "react-router";

export default function DashboardPage() {
    const { user, loading } = useAuth();
    const navigate = useNavigate();

    useEffect(() => {
        // only redirect to /login when auth has finished and there's no user
        if (loading) return;
        if (!user) navigate('/login');
    }, [user, loading, navigate]);

    return (
        <>
            <div className="p-6">
                <h1 className="text-2xl font-bold">Dashboard</h1>
                <p className="mt-4">Welcome, {user?.name ?? user?.email ?? "guest"}!</p>
            </div>
        </>
    );
}