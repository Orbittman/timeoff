package services

import(
	"net/http"
	"github.com/Orbittman/timeoff/dto"
	"github.com/Orbittman/timeoff/commands"
)

func PostFoodHandler(w http.ResponseWriter, r *http.Request) {
	var request dto.UpsertFoodRequest
	err := parsePost(r, &request)
	
	if err == nil {
		var command commands.UpsertFoodCommand
		command.Execute(r, request)
	} else {
		http.Error(w, err.Error(), 400)
	}
}