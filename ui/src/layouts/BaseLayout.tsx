import { Outlet } from "react-router";
import Navbar from "../components/Navbar";
import { AuthProvider } from "../contexts/AuthContext";

export default function BaseLayout() {
    return (
        <AuthProvider>
            <Navbar />
            <Outlet />
        </AuthProvider>
    )
}