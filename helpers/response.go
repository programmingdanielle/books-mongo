package helpers

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Err    string `json:"error,omitempty"`
	Errors Errors `json:"errors,omitempty"`
}

type Errors []error

// Respond responds back to the client.
func Respond(w http.ResponseWriter, code int, body []byte) {
	w.WriteHeader(code)
	if code == http.StatusNoContent {
		w.Write(nil)
	} else {
		w.Write(body)
	}
}

func RespondWithError(w http.ResponseWriter, code int, err error) {
	RespondWithJSON(w, code, ErrorResponse{Err: err.Error()})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	body, _ := json.Marshal(payload)
	Respond(w, code, body)
}
