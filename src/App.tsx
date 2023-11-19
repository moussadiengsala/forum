import { useState } from 'react'
import './App.css'
import { Route, Routes } from 'react-router-dom'
import Home from './pages/Home'
import Posts from './pages/Posts'

function App() {
  const [count, setCount] = useState(0)

  return (
    <Routes>
      <Route path='/' element={<Home />}/>
      <Route path='posts' element={<Posts />}/>
    </Routes>
  )
}

export default App
