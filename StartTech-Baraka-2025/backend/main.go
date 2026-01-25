package main

import (
    "fmt"
    "net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Backend is Healthy")
}

func main() {
    http.HandleFunc("/health", healthCheck)
    fmt.Println("Server starting on port 8080...")
    http.ListenAndServe(":8080", nil)
}