import { Outlet } from "react-router-dom"
import Navbar from "./Navbar"

function LayoutFeed() {
  return (
    <>
        <Navbar />
        <Outlet />
    </>
  )
}

export default LayoutFeed