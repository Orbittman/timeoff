package services

import (
	"fmt"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi there, welcome to the TimeOff app\n")
	
	s, err := getSession(r)
	
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprintf(w, "Hello %s", s.Name)
}