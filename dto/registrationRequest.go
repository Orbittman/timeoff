package dto

type RegistrationRequest struct {
	RequestBase
	UserName  string
	Password  string
	FirstName string
	LastName  string
	Email     string
}
