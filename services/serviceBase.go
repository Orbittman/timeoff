package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/Orbittman/timeoff/dto"
	"github.com/Orbittman/timeoff/authentication"
)

func parsePost(r *http.Request, generic dto.Requester) error {
	body, err := ioutil.ReadAll(r.Body)

	if err == nil {
		json.Unmarshal(body, &generic)
		if !dto.ValidateChecksum(generic) {
			err := errors.New("Invalid hash #" + generic.Key() + "#")

			return err
		}
	}

	return err
}

func getSession(r *http.Request) (authentication.Session, error) {
	cookie, err := r.Cookie("auth_token")
	s, err := authentication.GetSession(r, cookie.Value)
	
	return s, err
}
