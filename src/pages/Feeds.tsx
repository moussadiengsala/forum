import { useContext, useEffect } from "react"
import { Playload } from "../App"
import CreattePost from "../components/CreattePost"

function Feeds() {
    let {payload, setPayload} = useContext(Playload) ?? { payload: "", setPayload: () => {} }

    useEffect(() => {
        console.log(payload)
    }, [])
    
    return (
        <main className="col-span-3 bg-yellow-900">
            <CreattePost />
            <div>
                posts
            </div>
        </main>
    )
}

export default Feeds