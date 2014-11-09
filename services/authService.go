package services

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/timeoff/authentication"
	"github.com/timeoff/commands"
	"github.com/timeoff/dto"
	"github.com/timeoff/queries"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var registrationRequest dto.RegistrationRequest
	err := parsePost(r, &registrationRequest)

	if err == nil {
		var registrationCommand commands.RegistrationCommand
		registrationCommand.Execute(r, registrationRequest)
		w.Write([]byte(http.StatusText(200)))
	} else {
		http.Error(w, err.Error(), 500)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest dto.LoginRequest
	err := parsePost(r, &loginRequest)
	if err == nil {
		var loginQuery queries.LoginQuery
		user := loginQuery.Execute()
		if loginRequest.UserName == user.UserName && loginRequest.Password == user.Password {
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
	value := dto.CreateHash(vars["value"])

	output, _ := json.Marshal(dto.HashResponse{value})
	w.Write(output)
}
