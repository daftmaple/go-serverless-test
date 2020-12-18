package main

import (
	"net/http"

	api "github.com/daftmaple/serverless-test/api"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/api/", api.IndexHandler)
	http.HandleFunc("/api/document", api.DocumentHandler)
	http.HandleFunc("/api/params", api.ParamsHandler)

	http.ListenAndServe(":8080", nil)
}
