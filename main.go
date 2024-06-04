package main

import (
        "net/http"
        "log"
)

func main(){
        http.HandleFunc("/", handleWebsiteFiles)
        http.HandleFunc("/api/uploadFile", handleUploadFile)
        http.HandleFunc("/api/getSharedFile", handleGetSharedFile)

        const PORT = ":8088"
        err := http.ListenAndServe(PORT, nil)
        if err != nil {
                log.Println("Error starting server")
                return
        }
}

func handleWebsiteFiles (w http.ResponseWriter, r *http.Request){
        http.FileServer(http.Dir("./frontend/dist"))
}

