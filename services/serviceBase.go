package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/Orbittman/timeoff/dto"
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
