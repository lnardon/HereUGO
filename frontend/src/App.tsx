import { useState } from "react";
import "./App.css";

function App() {
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
    <div className="w-full font-bold text-3xl flex items-center justify-center">
      <h1>HereUGO</h1>
      <div>
        <input onChange={(e: any) => setFile(e.target.files[0])} type="file" />
        <button onClick={uploadFile}>Upload</button>
      </div>
    </div>
  );
}

export default App;
