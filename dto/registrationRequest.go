package dto

type RegistrationRequest struct {
	RequestBase
	UserName  string
	Password  string
	FirstName string
	LastName  string
	Email     string
	LoginAfterRegistration bool
}

func (r RegistrationRequest) Key() string{
	return r.UserName + r.FirstName + r.LastName + r.Password + r.Email
}
