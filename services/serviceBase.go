package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/timeoff/dto"
)

func parsePost(r *http.Request, generic dto.Requester) error {
	body, err := ioutil.ReadAll(r.Body)

	if err == nil {
		json.Unmarshal(body, &generic)
		if !generic.ValidateChecksum() {
			err := errors.New("Invalid hash")

			return err
		}
	}

	return err
}
