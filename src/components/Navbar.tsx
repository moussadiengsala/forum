import { BellIcon, ChatBubbleOvalLeftEllipsisIcon, HomeIcon, HomeModernIcon } from "@heroicons/react/16/solid"
import { Link, NavLink } from "react-router-dom"

function Navbar() {
  return (
    <header className="flex items-center bg-green-500">
      <nav className="flex items-center space-x-8">
        <Link to="/posts" className="flex gap-2 justify-center items-center bg-red-700">
          <HomeIcon className="w-8" />
          <span>Home</span>
        </Link>
        <ul className="flex gap-4">
          <li><NavLink to="#" className={({isActive}) => isActive ? "bg-red-500":"bg-tranparent" }>For you</NavLink></li>
          <li><NavLink to="#" className={({isActive}) => isActive ? "bg-red-500":"bg-tranparent" }>Following</NavLink></li>
        </ul>
      </nav>

      <div className="flex ml-auto">
        <button>
          <ChatBubbleOvalLeftEllipsisIcon className="w-8"/>
        </button>
        <button>
          <BellIcon className="w-8"/>
        </button>
        <div>
          <span>My name</span>
          <Link to="#">
            P
          </Link>
        </div>
      </div>
    </header>
  )
}

export default Navbar