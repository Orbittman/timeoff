package services

import (
	"net/http"

	"timeoff/authentication"
	"timeoff/dto"
)

func Login(w http.ResponseWriter, r *http.Request) {
	request := dto.LoginRequest{}
	err := parsePost(r, &request)
	if err == nil {
		if request.UserName == "tim" && request.Password == "tim" {
			authentication.CreateAuthCookie(w, r)
		} else {
			http.Error(w, http.StatusText(401), 401)
		}
	} else {
		http.Error(w, err.Error(), 400)
	}
}
