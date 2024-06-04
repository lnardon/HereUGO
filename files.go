package main

import "net/http"

func handleGetSharedFile(w http.ResponseWriter, r *http.Response){
        w.WriteHeader(http.StatusOK)
}
