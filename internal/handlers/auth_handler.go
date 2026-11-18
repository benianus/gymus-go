package handlers

import (
	"encoding/json/v2"
	"log"
	"net/http"

	"github.com/benianus/gymus-go/internal/dtos"
	"github.com/benianus/gymus-go/internal/helpers"
	"github.com/benianus/gymus-go/internal/services"
)

type AuthHandler struct {
	*services.AuthService
}

func NewAuthHandler(s *services.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: s,
	}
}

func (h *AuthHandler) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	// decode the request body to go object
	var reqDto dtos.LoginRequestDto
	if err := json.UnmarshalRead(r.Body, &reqDto); err != nil {
		log.Println(err)
		helpers.ResponseWithError(w, http.StatusBadRequest, dtos.ErrorResponseDto[string]{
			Message: err.Error(),
		})
		return
	}
	defer r.Body.Close()

	// validate the request
	if err := h.Validate.Struct(reqDto); err != nil {
		log.Println(err)
		helpers.ResponseWithError(w, http.StatusBadRequest, dtos.ErrorResponseDto[string]{
			Message: err.Error(),
		})
		return
	}

	// login service logic
	authRes, err := h.Login(reqDto, r.Context())
	if err != nil {
		log.Println(err)
		helpers.ResponseWithError(w, http.StatusBadRequest, dtos.ErrorResponseDto[string]{
			Message: err.Error(),
		})
		return
	}

	// return json response to the client
	helpers.ResponseWithJson(w, http.StatusOK, authRes)
}

func (h *AuthHandler) HandlerRegister(w http.ResponseWriter, r *http.Request) {
	// decoding the requst to an object
	var reqDto dtos.RegisterRequestDto
	if err := json.UnmarshalRead(r.Body, &reqDto); err != nil {
		log.Println(err)
		helpers.ResponseWithError(
			w,
			http.StatusBadRequest,
			dtos.ErrorResponseDto[string]{
				Message: err.Error(),
			},
		)
		return
	}
	defer r.Body.Close()

	// validate the request
	if err := h.Validate.Struct(reqDto); err != nil {
		log.Println(err)
		helpers.ResponseWithError(w, http.StatusBadRequest, dtos.ErrorResponseDto[string]{
			Message: "password should not be empty or blank!",
		})
		return
	}

	// send it to the service layer to implement the logic
	authRes, err := h.Register(reqDto, r.Context())
	if err != nil {
		log.Println(err)
		helpers.ResponseWithError(w, http.StatusBadRequest, dtos.ErrorResponseDto[string]{
			Message: err.Error(),
		})
		return
	}

	// encode the result and return response to the client
	helpers.ResponseWithJson(w, http.StatusCreated, authRes)
}

func (h *AuthHandler) HandlerRefresh(w http.ResponseWriter, r *http.Request) {
	// deserialization
	var reqDto dtos.RefreshRequestDto
	if err := json.UnmarshalRead(r.Body, &reqDto); err != nil {
		log.Println(err)
		helpers.ResponseWithError(
			w,
			http.StatusBadRequest,
			dtos.ErrorResponseDto[string]{
				Message: err.Error(),
			},
		)
		return
	}
	defer r.Body.Close()

	// validation
	if err := h.Validate.Struct(reqDto); err != nil {
		log.Println(err)
		helpers.ResponseWithError(w, http.StatusBadRequest, dtos.ErrorResponseDto[string]{
			Message: err.Error(),
		})
		return
	}

	// logic
	refResDto, err := h.Refresh(reqDto, r.Context())
	if err != nil {
		log.Println(err)
		helpers.ResponseWithError(w, http.StatusInternalServerError, dtos.ErrorResponseDto[string]{
			Message: err.Error(),
		})
		return
	}

	//  response
	helpers.ResponseWithJson(w, http.StatusOK, refResDto)
}

func (h *AuthHandler) HandlerLogout(w http.ResponseWriter, r *http.Request) {
	// request decoding
	var reqDto dtos.LogoutRequestDto
	if err := json.UnmarshalRead(r.Body, &reqDto); err != nil {
		log.Println(err)
		helpers.ResponseWithError(
			w,
			http.StatusInternalServerError,
			dtos.ErrorResponseDto[string]{
				Message: err.Error(),
			},
		)
		return
	}
	defer r.Body.Close()

	// request validation
	if err := h.Validate.Struct(reqDto); err != nil {
		log.Println(err)
		helpers.ResponseWithError(
			w,
			http.StatusInternalServerError,
			dtos.ErrorResponseDto[string]{
				Message: err.Error(),
			},
		)
		return
	}

	// logic to the service layer
	_, err := h.Logout(reqDto, r.Context())
	if err != nil {
		log.Println(err)
		helpers.ResponseWithError(
			w,
			http.StatusInternalServerError,
			dtos.ErrorResponseDto[string]{
				Message: err.Error(),
			},
		)
		return
	}

	// json response
	helpers.ResponseWithJson(
		w,
		http.StatusNoContent,
		helpers.LOGOUT_MESSAGE,
	)
}
