package handler

import (
	"fmt"
	"net/http"
)

// ParamsHandler ...
func ParamsHandler(w http.ResponseWriter, r *http.Request) {
	// If headers are needed
	// for name, headers := range r.Header {
	// 	for _, h := range headers {
	// 		fmt.Fprintf(w, "Header %v: %v\n", name, h)
	// 	}
	// }

	if len(r.URL.Query()) == 0 {
		fmt.Fprintf(w, "No query parameter passed")
	}

	for name, queries := range r.URL.Query() {
		for _, h := range queries {
			fmt.Fprintf(w, "Query %v: %v\n", name, h)
		}
	}
}
