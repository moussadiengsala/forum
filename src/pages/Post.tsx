import { useParams } from "react-router-dom"


function Post() {
    let id = useParams()
    console.log(id)
    return (
      <div className="col-span-3 bg-yellow-900"> single Post</div>
    )
}

export default Post