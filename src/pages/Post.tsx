import { useParams } from "react-router-dom"


function Post() {
    let id = useParams()
    console.log(id)
    return (
      <div> single Post</div>
    )
}

export default Post