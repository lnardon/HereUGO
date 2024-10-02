// import { useEffect } from "react";

const DownloadFile: React.FC = () => {
  // useEffect(() => {
  //   fetch(
  //     `/api/getSharedFile?file=${new URLSearchParams(
  //       window.location.search
  //     ).get("file")}`
  //   ).then((response) => {
  //     if (response.ok)
  //       response.blob().then((blob) => {
  //         const url = window.URL.createObjectURL(new Blob([blob]));
  //         const link = document.createElement("a");
  //         link.href = url;
  //         link.setAttribute("download", "file.jpg");
  //         document.body.appendChild(link);
  //         link.click();
  //         link.parentNode?.removeChild(link);
  //       });
  //   });
  // }, []);

  function downloadFile() {
    fetch(
      `/api/getSharedFile?file=${new URLSearchParams(
        window.location.search
      ).get("file")}`
    ).then((response) => {
      if (response.ok)
        response.blob().then((blob) => {
          const url = window.URL.createObjectURL(new Blob([blob]));
          const link = document.createElement("a");
          link.href = url;
          link.setAttribute("download", "file.jpg");
          document.body.appendChild(link);
          link.click();
          link.parentNode?.removeChild(link);
        });
    });
  }

  return (
    <div className="w-full flex items-center flex-col justify-center gap-8 bg-gray-600 p-12 rounded-xl shadow-xl bg-opacity-20">
      <h1>HereUGO</h1>
      <div>
        <button onClick={downloadFile}>Download</button>
      </div>
    </div>
  );
};

export default DownloadFile;
