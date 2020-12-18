package handler

import (
	"fmt"
	"net/http"
)

// SubIndexHandler ...
func SubIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You're in a subdirectory")
}
