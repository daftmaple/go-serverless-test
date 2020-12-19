package main

import (
	"net/http"
	"os"

	api "github.com/daftmaple/serverless-test/api"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/api/", api.IndexHandler)
	http.HandleFunc("/api/document", api.DocumentHandler)
	http.HandleFunc("/api/params", api.ParamsHandler)
	http.HandleFunc("/api/post", api.PostHandler)

	port, exists := os.LookupEnv("APP_PORT")
	if !exists {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
