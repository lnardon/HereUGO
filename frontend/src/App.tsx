import { useState, useEffect } from "react";
import "./App.css";
import UploadFile from "./views/UploadFile";
import DownloadFile from "./views/DownloadFile";

function App() {
  const [view, setView] = useState("upload" as "upload" | "download");

  useEffect(() => {
    let urlParams = new URLSearchParams(window.location.search);
    if (urlParams.has("file")) {
      setView("download");
    }
  }, []);

  return (
    <div className="w-full font-bold text-3xl flex items-center justify-center">
      {view === "upload" ? <UploadFile /> : <DownloadFile />}
    </div>
  );
}

export default App;
