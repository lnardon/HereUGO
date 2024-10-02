package main

import (
	"log"
	"net/http"
)

func main(){
    http.Handle("/", http.FileServer(http.Dir("./frontend/dist")))

    http.HandleFunc("/api/uploadFile", handleUploadFile)
    http.HandleFunc("/api/getSharedFile", handleGetSharedFile)

    const PORT = ":8080"
    err := http.ListenAndServe(PORT, nil)
    if err != nil {
        log.Println("Error starting server")
        return
    }
}
