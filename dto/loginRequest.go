package dto

type LoginRequest struct {
	RequestBase
	UserName string
	Password string
}
