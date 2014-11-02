package services

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	"timeoff/authentication"
	"timeoff/dto"
)

func Login(w http.ResponseWriter, r *http.Request) {
	loginRequest := dto.LoginRequest{}
	err := parsePost(r, &loginRequest)
	if err == nil {
		if loginRequest.UserName == "tim" && loginRequest.Password == "tim" {
			authentication.CreateAuthCookie(w, r, loginRequest)
		} else {
			http.Error(w, http.StatusText(401), 401)
		}
	} else {
		http.Error(w, err.Error(), 400)
	}
}

func GetHash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	value := vars["value"]

	output, _ := json.Marshal(dto.HashResponse{value})
	w.Write(output)
}
