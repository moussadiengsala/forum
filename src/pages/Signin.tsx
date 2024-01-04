import { createPortal } from "react-dom"
import Acceuil from "../components/Acceuil"


function Signin() {
  return (
    <>
      <Acceuil />
      {createPortal(
        <div className="absolute w-full h-full left-0 top-0 bg-black/50 flex justify-center items-center">
          <div className="w-1/2 h-3/4 bg-red-500 rounded-lg"></div>
        </div>,
        document.body
      )}
    </>
  )
}

export default Signin