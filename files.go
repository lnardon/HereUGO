package main

import (
        "net/http"
        "os"
        "io"
)

func handleGetSharedFile(w http.ResponseWriter, r *http.Request){
        w.WriteHeader(http.StatusOK)
}

func handleUploadFile(w http.ResponseWriter, r *http.Request){
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Error parsing multipart form", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	dst, err := os.Create("shared_files/" + handler.Filename)
	if err != nil {
		http.Error(w, "Error creating the file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("File uploaded successfully!"))
}
