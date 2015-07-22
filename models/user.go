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
