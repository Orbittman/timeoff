package commands

import (
	"net/http"

	"github.com/Orbittman/timeoff/dto"
)

type Command interface {
	Execute(r *http.Request, request dto.Requester)
}
