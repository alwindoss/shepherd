import { createBrowserRouter } from "react-router";
import IndexPage from "../pages/IndexPage";
import BaseLayout from "../layouts/BaseLayout";
import AboutPage from "../pages/AboutPage";
import LoginPage from "../pages/LoginPage";
import SignupPage from "../pages/SignupPage";
import DashboardPage from "../pages/DashboardPage";

const router = createBrowserRouter([
  {
    path: "/",
    Component: BaseLayout,
    children: [
        { index: true, Component: IndexPage },
        { path: "login", Component: LoginPage },
        { path: "signup", Component: SignupPage },
        { path: "dashboard", Component: DashboardPage },
        { path: "about", Component: AboutPage },
    ]
  },
  {
    path: "/about",
    Component: IndexPage,
  },
]);

export default router;