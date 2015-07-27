package services

import(
	"net/http"
	"github.com/Orbittman/timeoff/dto"
	"github.com/Orbittman/timeoff/commands"
)

func PostGotchiHandler(w http.ResponseWriter, r *http.Request) {
	var request dto.NewGotchiRequest
	err := parsePost(r, &request)
	
	if err == nil {
		var command commands.NewGotchiCommand
		command.Execute(r, request)
	} else {
		http.Error(w, err.Error(), 400)
	}
}