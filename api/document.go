package handler

import (
	"fmt"
	"net/http"
)

// Handler ...
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Document page</h1>")
}
