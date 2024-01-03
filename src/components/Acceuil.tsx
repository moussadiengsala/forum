import { Link } from "react-router-dom"
import { GitHubIcon, GoogleIcon, Logo } from "./Buttons"


function Acceuil() {
  return (
    <div className="grid grid-cols-2 w-full h-full bg-slate-900 text-white">
        <div className="flex justify-center items-center">
          <Logo style="w-1/4" />
        </div>
        <div className="flex flex-col justify-center">
            <div className="bg-red-500">
                <h1 className="text-4xl font-black">Happening now</h1>
            </div>
            <div className="bg-green-500 w-fit">
                <h2>Join today.</h2>
                <div className="flex space-y-2 flex-col">
                  <GitHubIcon text="Sign up with Google" href="/" />
                  <GoogleIcon text="Sign up with Github" href="/" />
                  <Link to="/" className=" bg-sky-500 h-10 rounded-full flex justify-center items-center">Create account</Link>
                  <i className="w-48 text-xs">By signing up, you agree to the Terms of Service and Privacy Policy, including Cookie Use.</i>
                </div>
                <div className="bg-red-500 mt-6">
                  <i className="font-bold text-xs">Already have an account?</i>
                  <Link to="/" className="border-2 h-10 rounded-full flex justify-center items-center">Create account</Link>
                </div>
            </div>
        </div>
    </div>
  )
}

export default Acceuil