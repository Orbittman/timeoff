package commands

import (
	"net/http"

	"github.com/timeoff/dto"
)

type Command interface {
	Execute(r *http.Request, request dto.Requester)
}
