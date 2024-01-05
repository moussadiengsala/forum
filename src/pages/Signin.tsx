import { createPortal } from "react-dom"
import Acceuil from "../components/Acceuil"
import { GitHubIcon, GoogleIcon } from "../components/Buttons"
import { Link } from "react-router-dom"
import { useState } from "react"
import { XMarkIcon } from '@heroicons/react/24/outline'


function Signin() {
  let [state, setState] = useState<LoginForm>({password: "", identifiers: ""})
  let [infoLogged, setInfoLogged] = useState("")

  let handleForm = (e: React.ChangeEvent<HTMLElement>) => {
    setState((prev) => ({
      ...prev,
      [(e.target as HTMLInputElement | HTMLTextAreaElement).name]: (e.target as HTMLInputElement | HTMLTextAreaElement).value
    }))
  }

  let handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    e.stopPropagation();

    try {
      let response = await fetch("http://localhost:8000/api/auth/signin", {
        method: "POST",
        body: JSON.stringify(state),
      });
  
      if (!response.ok) {
        const errorMessage = await response.json();
        throw new Error(errorMessage.message || "An error occurred during signup");
      }

      let data = await response.json();
      console.log(data)

    } catch (error) {
      setInfoLogged((error as Error).message); 
    }
  }

  return (
    <>
      <Acceuil />
      {createPortal(
        <div className="absolute w-full h-full left-0 top-0 bg-black/50 flex justify-center items-center">
          <div className="w-1/2 h-3/4 bg-red-500 rounded-lg flex justify-center items-center relative">
            <Link to="/" className="absolute top-4 left-4 w-8">
              <XMarkIcon />
            </Link>
            <div className="w-fit">
                <div className="flex space-y-2 flex-col">
                  <GitHubIcon text="Sign up with Google" href="/" />
                  <GoogleIcon text="Sign up with Github" href="/" />

                  <form action="POST" onSubmit={handleSubmit} className="flex flex-col space-y-2">
                    <input type="text" name="identifiers" className="input-form" onChange={handleForm} required/>
                    <input type="password" name="password" className="input-form" onChange={handleForm} required/>
                    <div>
                      {infoLogged}
                    </div>
                    <input type="submit" value="next" className="h-10 bg-white rounded-full" />
                  </form>

                  <div className="space-x-2 text-xs">
                    <i>Don't have an account?</i>
                    <Link to="/auth/signup" className="text-sky-500">Sign up</Link>
                  </div>
                </div>
            </div>
          </div>
        </div>,
        document.body
      )}
    </>
  )
}

export default Signin