package handler

import (
	"fmt"
	"net/http"
)

// ParamsHandler ...
func ParamsHandler(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "Header %v: %v\n", name, h)
		}
	}

	for name, queries := range r.URL.Query() {
		for _, h := range queries {
			fmt.Fprintf(w, "Query %v: %v\n", name, h)
		}
	}
}
