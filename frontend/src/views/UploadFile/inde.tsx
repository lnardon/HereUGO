import { useState } from "react";

const UploadFile: React.FC = () => {
  const [file, setFile] = useState(null);

  function uploadFile() {
    const formData = new FormData();
    formData.append("file", file!);

    fetch("/api/uploadFile", {
      method: "POST",
      body: formData,
    })
      .then((response) => response.json())
      .then((result) => {
        console.log("Success:", result);
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  }

  return (
    <div className="w-full flex items-center flex-col justify-center gap-8 bg-gray-600 p-12 rounded-xl shadow-xl bg-opacity-20">
      <h1>HereUGO</h1>
      <div>
        <input onChange={(e: any) => setFile(e.target.files[0])} type="file" />
        <button onClick={uploadFile}>Upload</button>
      </div>
    </div>
  );
};

export default UploadFile;
