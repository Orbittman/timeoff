package queries

import (
	"github.com/Orbittman/timeoff/models"
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
