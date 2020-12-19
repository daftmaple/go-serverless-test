package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Payload ...
type Payload struct {
	ID    int    `json:"id" validate:"required"`
	Field string `json:"field" validate:"required"`
}

// PostHandler Increments id field of the JSON body
func PostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// Option is valid
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Invalid Content-Type: "+contentType, http.StatusUnsupportedMediaType)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payload := &Payload{}
	err = json.Unmarshal(body, payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set header
	w.Header().Set("Content-Type", "application/json")

	// Increment payload id
	payload.ID++

	err = json.NewEncoder(w).Encode(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func parseJSON() {

}
