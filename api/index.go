package handler

import (
	"fmt"
	"net/http"
)

// IndexHandler ...
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there")
}
