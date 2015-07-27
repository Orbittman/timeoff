package commands

import (
	"net/http"

	"github.com/Orbittman/timeoff/dto"
	"github.com/Orbittman/timeoff/models"
	"github.com/Orbittman/timeoff/data"
)

type NewGotchiCommand struct { }

func (command NewGotchiCommand) Execute (r *http.Request, request dto.Requester) error {
	gotchiRequest := request.(dto.NewGotchiRequest)
	gotchi := models.Gotchi {
		Name: gotchiRequest.Name,
	}

	err := data.Put(r, &gotchi)
	return err
}