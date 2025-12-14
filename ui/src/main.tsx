import { createRoot } from 'react-dom/client'
import { RouterProvider } from "react-router/dom";
import './index.css'
import router from './router/router';



// const root = document.getElementById("root");

// ReactDOM.createRoot(root).render(
//   <RouterProvider router={router} />,
// );

createRoot(document.getElementById('root')!).render(
  <RouterProvider router={router} />,
)
