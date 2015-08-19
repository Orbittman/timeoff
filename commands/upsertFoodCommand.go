package commands

import (
	"net/http"

	"github.com/Orbittman/timeoff/dto"
	"github.com/Orbittman/timeoff/models"
	"github.com/Orbittman/timeoff/data"
)

type UpsertFoodCommand struct { }

func (command UpsertFoodCommand) Execute (r *http.Request, request dto.Requester) error {
	foodRequest := request.(dto.UpsertFoodRequest)
	food := models.Food {
		Name: foodRequest.Name,
	}

	err := data.Put(r, &food)
	
	return err
}