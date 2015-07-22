package services

import(
	"net/http"
)

func gotchiHandler(w http.ResponseWriter, r *http.Request) {
	var request dto.NewGotchiRequest
	err := parsePost(r, &request)
}