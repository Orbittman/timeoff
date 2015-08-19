package commands

import (
	"net/http"
	"time"

	"github.com/Orbittman/timeoff/dto"
	"github.com/Orbittman/timeoff/models"
	"github.com/Orbittman/timeoff/data"
)

type RegistrationCommand struct {
}

func (command RegistrationCommand) Execute(r *http.Request, registrationRequest dto.Requester) error {
	bob := registrationRequest.(dto.RegistrationRequest)
	user := models.User{
		UserName:  bob.UserName,
		Password:  bob.Password,
		FirstName: bob.FirstName,
		LastName:  bob.LastName,
		Email:     bob.Email,
		Registered:time.Now(),
	}

	err := data.Put(r, &user)
	
	return err
}
