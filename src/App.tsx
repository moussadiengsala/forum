import { Routes, Route, Link } from "react-router-dom"
import Home from "./pages/Home"
import Signin from "./pages/Signin"
import Signup from "./pages/Signup"


function App() {
  return (
    <>
      <ul className="flex gap-4 text-red-500">
        <li><Link to="/">home</Link></li>
        <li><Link to="/auth/signin">signin</Link></li>
        <li><Link to="/auth/signup">signup</Link></li>
      </ul>
      <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/auth/signin" element={<Signin />} />
          <Route path="/auth/signup" element={<Signup />} />
      </Routes>
    </>
    
  )
}

export default App
