
import Acceuil from '../components/Acceuil'
import { createPortal } from 'react-dom'
import { Link } from 'react-router-dom'
import { CheckBadgeIcon, XMarkIcon } from '@heroicons/react/24/outline'
import { useFormInput } from '../lib/formInput'
import { useQuery } from 'react-query'
import { fetcher } from '../utils/fetcher'

const initValue: RegisterForm = {firstname: "", lastname: "", email: "", username: "", bio: "", password: ""}

function Signup() {
 
  let [state, handleForm] = useFormInput<RegisterForm>(initValue)
  const { isLoading, isSuccess, error, refetch } = useQuery(['signup'], () => fetcher({data: state, endpoint: "/auth/signup", method: "POST"}), {enabled: false, retry: false})

  let handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    e.stopPropagation();
    refetch()
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
                  <form method='post' action="/auth/signin" onSubmit={handleSubmit} className="flex flex-col space-y-2">
                    <div className='flex justify-center items-center'>
                      <input type="text" name="firstname" className="input-form" onChange={handleForm} placeholder='firstname' required/>
                      <input type="text" name="lastname" className="input-form" onChange={handleForm} placeholder='lastname' required/>
                    </div>

                    <input type="email" name="email" onChange={handleForm} placeholder='abc@gmail.com' className="input-form" />
                    <input type="text" name="username" onChange={handleForm} placeholder='username' className="input-form" />

                    <textarea name="bio" id="" onChange={handleForm} cols={30} rows={10} placeholder='bio'>

                    </textarea>

                    <input type="password" name="password" className="h-10 input-form" onChange={handleForm} placeholder='password' required/>
                    <div className='text-xs'>
                      <span>{(error as Error)?.message}</span>
                    </div>
                    <input type="submit" value={isLoading ? "loading..." : "next"} className="h-10 bg-white rounded-full" disabled={isLoading} />
                  </form>
              
                  <div className="space-x-2 text-xs">
                    <i>Already have an account?</i>
                    <Link to="/auth/signin" className="text-sky-500">Sign in</Link>
                  </div>
                </div>
            </div>
          </div>
        </div>,
        document.body
      )}

      {isSuccess && createPortal(
        <div className='bg-green-500 p-8 rounded-lg absolute left-1/2 top-1/2 transform -translate-y-1/2 -translate-x-1/2'>
          <div>Signup successful!</div>
          <CheckBadgeIcon />
          <Link to="/auth/signin" className='h-10 w-full rounded-full flex justify-center items-center bg-white'>Sign in</Link>
        </div>
        ,document.body
      )}
    </>
  )
}

export default Signup

