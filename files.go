package main

import (
	"io"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func handleGetSharedFile(w http.ResponseWriter, r *http.Request){
	file, err := os.Open("shared_files/" + r.URL.Query().Get("file"))
	if err != nil {
		http.Error(w, "Error opening the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, "Error sending the file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=" + r.URL.Query().Get("file"))
}

func handleUploadFile(w http.ResponseWriter, r *http.Request){
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Error parsing multipart form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	file_id := uuid.New().String()
	dst, err := os.Create("shared_files/" + file_id)
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
