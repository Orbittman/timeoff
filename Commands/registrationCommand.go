package commands

import (
	"appengine"
	"appengine/datastore"
	"net/http"

	"github.com/timeoff/dto"
	"github.com/timeoff/models"
)

type RegistrationCommand struct {
}

func (command RegistrationCommand) Execute(r *http.Request, registrationRequest dto.Requester) error {
	c := appengine.NewContext(r)

	bob := registrationRequest.(dto.RegistrationRequest)
	user := models.User{
		UserName:  bob.UserName,
		Password:  bob.Password,
		FirstName: bob.FirstName,
		LastName:  bob.LastName,
		Email:     bob.Email,
	}

	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "user", nil), &user)
	return err
}
