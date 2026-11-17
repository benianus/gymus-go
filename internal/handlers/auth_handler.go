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
		helpers.ResponseWithError(w, http.StatusBadRequest, &dtos.ErrorResponseDto[string]{
			Message: err.Error(),
		})
		return
	}
	defer r.Body.Close()

	// validate the request

	if reqDto.Username == "" || len(reqDto.Username) == 0 {
		helpers.ResponseWithError(w, http.StatusBadRequest, &dtos.ErrorResponseDto[string]{
			Message: "username should not be empty or blank!!",
		})
		return
	}

	if reqDto.Password == "" || len(reqDto.Password) == 0 {
		helpers.ResponseWithError(w, http.StatusBadRequest, &dtos.ErrorResponseDto[string]{
			Message: "password is required",
		})
		return
	}

	res, err := h.Login(reqDto)
	if err != nil {
		log.Println(err)
		helpers.ResponseWithError(w, http.StatusBadRequest, &dtos.ErrorResponseDto[string]{
			Message: err.Error(),
		})
		return
	}

	if err := json.MarshalWrite(w, &res); err != nil {
		log.Println(err)
		helpers.ResponseWithError(w, http.StatusBadRequest, &dtos.ErrorResponseDto[string]{
			Message: err.Error(),
		})
		return
	}
}

func (h *AuthHandler) HandlerRegister(w http.ResponseWriter, r *http.Request) {
	// decoding the requst to an object
	var reqDto dtos.RegisterRequestDto
	if err := json.UnmarshalRead(r.Body, &reqDto); err != nil {
		log.Println(err)
		helpers.ResponseWithError(w, http.StatusBadRequest, &dtos.ErrorResponseDto[string]{
			Message: err.Error(),
		})
		return
	}
	defer r.Body.Close()

	// validate the request
	if err := h.Validate.Struct(reqDto); err != nil {
		log.Println(err)
		helpers.ResponseWithError(w, http.StatusBadRequest, &dtos.ErrorResponseDto[string]{
			Message: "password should not be empty or blank!",
		})
		return
	}

	// send it to the service layer to implement the logic
	authRes, err := h.Register(reqDto, r.Context())
	if err != nil {
		log.Println(err)
		helpers.ResponseWithError(w, http.StatusBadRequest, &dtos.ErrorResponseDto[string]{
			Message: err.Error(),
		})
		return
	}

	// encode the result and return response to the client
	helpers.ResponseWithJson(w, http.StatusCreated, authRes)
}

func (h *AuthHandler) HandlerRefresh(w http.ResponseWriter, r *http.Request) {}

func (h *AuthHandler) HandlerLogout(w http.ResponseWriter, r *http.Request) {}
