import { Outlet } from "react-router-dom"
import Home from "./Home"

function Layout() {
  return (
    <>
        <Home />
        <Outlet />
    </>
  )
}

export default Layout