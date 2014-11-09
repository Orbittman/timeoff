package queries

import (
	"github.com/timeoff/models"
)

type LoginQuery struct {
}

func (l *LoginQuery) Execute() models.User {
	user := models.User{
		UserName: "tim",
		Password: "password",
	}

	return user
}
