import { Routes, Route } from "react-router-dom"
import Signin from "./pages/Signin"
import Signup from "./pages/Signup"
import { Dispatch, SetStateAction, createContext, useState } from "react";
import Post from "./pages/Post";
import Feeds from "./pages/Feeds";
import LayoutFeed from "./components/LayoutFeed";
import Layout from "./components/Layout";

interface CredentialsContextType {
  payload: PayloadUser | null;
  setPayload: Dispatch<SetStateAction<any>>;
}

export const Playload = createContext<CredentialsContextType | null>(null);

function App() {
  const [payload, setPayload] = useState(null);

  return (
    <Playload.Provider value={{payload, setPayload}}>
      <Routes>
          <Route element={<Layout />}>
            <Route path="/" element={<></>} />
            <Route path="/auth/signin" element={<Signin />} />
            <Route path="/auth/signup" element={<Signup />} />
          </Route>

          <Route element={<LayoutFeed />}>
            <Route path="posts" element={<Feeds />} />
            <Route path="posts/:id" element={<Post />} />
          </Route>
      </Routes>
    </Playload.Provider>
    
  )
}

export default App
