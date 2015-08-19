package queries

import (
	"github.com/Orbittman/timeoff/models"
	"github.com/Orbittman/timeoff/dto"
	"net/http"
	"appengine"
	"appengine/datastore"
)

type FoodQuery struct {
}

func (l *FoodQuery) Execute(r *http.Request, foodRequest dto.FoodRequest) ([]models.Food, error) {
	c := appengine.NewContext(r)

	query := datastore.NewQuery("food").Filter("Level <=", foodRequest.Level).KeysOnly()
		
	var food []models.Food
	_, err := query.GetAll(c, &food)

	return food, err
}
