package main

import (
	"net/http"

	api "github.com/daftmaple/serverless-test/api"
)

func main() {
	http.HandleFunc("/", api.IndexHandler)
	http.HandleFunc("/document", api.DocumentHandler)
	http.HandleFunc("/params", api.ParamsHandler)

	http.ListenAndServe(":8080", nil)
}
