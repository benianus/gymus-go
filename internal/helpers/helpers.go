package helpers

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}

func ResponseWithJson(w http.ResponseWriter, code int, payload any) {
	response, err := json.Marshal(&payload)
	if err != nil {
		return
	}
	w.Header().Set("Contnet-type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ResponseWithError(w http.ResponseWriter, code int, msg string) {
	errorResponse := ErrorResponse{
		ErrorMessage: msg,
	}

	response, err := json.Marshal(&errorResponse)
	if err != nil {
		return
	}
	w.Header().Set("Contnet-type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

}
