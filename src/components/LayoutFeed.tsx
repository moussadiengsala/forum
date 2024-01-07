import { Outlet } from "react-router-dom"
import Navbar from "./Navbar"
import AsideLeft from "./AsideLeft"
import AsideReight from "./AsideReight"

function LayoutFeed() {
  return (
    <>
        <div className="w-full grid grid-cols-6">
          <AsideLeft />
          <div className="col-span-5">
            <Navbar />
            <div className="grid grid-cols-4">
              <Outlet />
              <AsideReight />
            </div>
          </div>
        </div>
    </>
  )
}

export default LayoutFeed