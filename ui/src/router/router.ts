import { createBrowserRouter } from "react-router";
import IndexPage from "../pages/IndexPage";
import BaseLayout from "../layouts/BaseLayout";
import AboutPage from "../pages/AboutPage";

const router = createBrowserRouter([
  {
    path: "/",
    Component: BaseLayout,
    children: [
        { index: true, Component: IndexPage },
        { path: "about", Component: AboutPage },
    ]
  },
  {
    path: "/about",
    Component: IndexPage,
  },
]);

export default router;