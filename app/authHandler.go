package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/hexagonal-auth/dto"
	"example.com/hexagonal-auth/service"
)

type AuthHandler struct {
	service service.AuthService
}

func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		log.Println("Error while decoding login request: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		token, err := h.service.Login(loginRequest)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, err.Error())
		} else {
			fmt.Fprint(w, *token)
		}
	}
}

func (h AuthHandler) Verify(w http.ResponseWriter, r *http.Request) {

}

func (h AuthHandler) NotImplemented(w http.ResponseWriter, r *http.Request) {

}
