package models

import(
	"time"
)

type User struct {
	UserName  string
	FirstName string
	LastName  string
	Password  string
	Email     string
	Registered time.Time
}

func (u *User) DataKey() string {
	return u.UserName
}
