package services

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/Orbittman/timeoff/authentication"
	"github.com/Orbittman/timeoff/commands"
	"github.com/Orbittman/timeoff/dto"
	"github.com/Orbittman/timeoff/queries"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var registrationRequest dto.RegistrationRequest
	err := parsePost(r, &registrationRequest)

	if err == nil {
		var registrationCommand commands.RegistrationCommand
		registrationCommand.Execute(r, registrationRequest)
		if registrationRequest.LoginAfterRegistration {
			login(w, r, dto.LoginRequest{UserName:registrationRequest.UserName, Password:registrationRequest.Password})
		}
	} else {
		http.Error(w, err.Error(), 400)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest dto.LoginRequest
	err := parsePost(r, &loginRequest)
	if err == nil {
		login(w, r, loginRequest)
	} else {
		http.Error(w, err.Error(), 400)
	}
}

func login(w http.ResponseWriter, r *http.Request, loginRequest dto.LoginRequest) {
	var loginQuery queries.LoginQuery
		_, err := loginQuery.Execute(r, loginRequest)
		
		if err == nil {
			authentication.CreateAuthCookie(w, r, loginRequest)
		} else {
			http.Error(w, http.StatusText(401), 401)
		}
}

func GetHash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	value := dto.CreateHash(vars["value"])
	
	output, _ := json.Marshal(dto.HashResponse{value})
	w.Write(output)
}
