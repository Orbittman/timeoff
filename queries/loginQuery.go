package queries

import (
	"github.com/Orbittman/timeoff/models"
	"github.com/Orbittman/timeoff/dto"
	"net/http"
	"appengine"
	"appengine/datastore"
)

type LoginQuery struct {
}

func (l *LoginQuery) Execute(r *http.Request, loginRequest dto.LoginRequest) (models.User, error) {
	c := appengine.NewContext(r)

	query := datastore.NewQuery("user").Filter("UserName =", loginRequest.UserName).KeysOnly()
		
	var user models.User	
	t := query.Run(c)
	_, err := t.Next(&user)
	
	return user, err
}
