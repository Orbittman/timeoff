package dto

type LoginRequest struct {
	RequestBase
	UserName string
	Password string
}

func (r LoginRequest) Key() string{
	return r.UserName + r.Password
}