import React from "react"
import ReactDOM from "react-dom/client"
import { createBrowserRouter, RouterProvider } from "react-router-dom"
import App from "./App"

//define your routes here
const router = createBrowserRouter([
    {
        path: "/",
        element: <App />
    }
])

const root = ReactDOM.createRoot(document.querySelector("#root")!)
root.render(
    <React.StrictMode>
        <RouterProvider router={router}></RouterProvider>
    </React.StrictMode>
)
