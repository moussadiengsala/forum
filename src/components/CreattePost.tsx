import { useState } from "react";
import { PhotoIcon } from "@heroicons/react/24/outline";
import { useFormInput } from "../lib/formInput";

type Post = {
    content: string;
    file: File | null
}

function CreattePost() {
  const [text, setText] = useFormInput<Post>({content: "", file: null})


  function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    e.stopPropagation(); 
  }

  return (
    <form onSubmit={handleSubmit} className="flex flex-col space-y-2 bg-yellow-400" encType="multipart/form-data">
      <textarea onChange={setText} cols={30} rows={5} placeholder="What's on your mind?"></textarea>
      <div className="flex">
        <label className="cursor-pointer">
          <PhotoIcon className="w-8" />
          <input onChange={setText} className="border-2 p-3 w-full hidden post-img" type="file" accept="image/*" name="file" />
        </label>

        {text.file && 
        <div className="">
        {/* Display a preview of the selected image */}
        <img src={URL.createObjectURL(text.file)} alt="Preview" className="w-10 h-10 object-cover" />
    </div>}
        <input type="submit" value="Post" className="h-10 bg-white rounded-full px-4 ml-auto" />
      </div>
    </form>
  );
}

export default CreattePost;
