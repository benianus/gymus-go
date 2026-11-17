package helpers

import (
	"encoding/json/v2"
	"log"
	"net/http"

	"github.com/benianus/gymus-go/internal/dtos"
)

func ResponseWithJson[T any](w http.ResponseWriter, code int, payload *T) {
	succesRes := dtos.Success[T]{
		Data: payload,
	}
	w.Header().Set("Contnet-type", "application/json")
	w.WriteHeader(code)
	if err := json.MarshalWrite(w, &succesRes); err != nil {
		log.Println(err)
		return
	}
}

func ResponseWithError[T any](w http.ResponseWriter, code int, msg *dtos.ErrorResponseDto[T]) {
	w.Header().Set("Contnet-type", "application/json")
	w.WriteHeader(code)
	errorResponse := dtos.Error[dtos.ErrorResponseDto[T]]{
		Error: msg,
	}

	if err := json.MarshalWrite(w, &errorResponse); err != nil {
		log.Println(err)
		return
	}
}
