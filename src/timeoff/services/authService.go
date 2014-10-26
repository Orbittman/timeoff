package services

import (
	"net/http"
)

type LoginRequest struct {
	UserName string
	Password string
}

type LoginReply struct {
	Message string
	Success bool
}

type AuthService struct{}

func (h *AuthService) Login(r *http.Request, request *LoginRequest, reply *LoginReply) error {
	reply.Message = "Hello, " + request.UserName + "!"
	return nil
}
