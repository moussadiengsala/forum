import { useContext, useEffect } from "react"
import { Playload } from "../App"

function Feeds() {
    let {payload, setPayload} = useContext(Playload) ?? { payload: "", setPayload: () => {} }

    useEffect(() => {
        console.log(payload)
    }, [])
    
    return (
        <div>
            posts
        </div>
    )
}

export default Feeds