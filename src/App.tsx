import { Routes, Route } from "react-router-dom"
import Signin from "./pages/Signin"
import Signup from "./pages/Signup"
import { Dispatch, SetStateAction, createContext, useState } from "react";
import Post from "./pages/Post";
import Feeds from "./pages/Feeds";
import LayoutFeed from "./components/LayoutFeed";
import Layout from "./components/Layout";
import Tweet from "./pages/Tweet";

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
            <Route path="/home" element={<Feeds />} />
            <Route path="/explore" element={<></>} />
            <Route path="/list" element={<></>} />
            <Route path="/save" element={<></>} />
            <Route path="/communities" element={<></>} />
            <Route path="/:user/status/:id" element={<Post />} />
            <Route path="/:user" element={<></>} />
          </Route>
      </Routes>
    </Playload.Provider>
  )
}

export default App
