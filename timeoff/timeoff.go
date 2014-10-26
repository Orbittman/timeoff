package timeoff

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"timeoff/authentication"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi there, welcome to the TimeOff app<br />")
	cookie, err := r.Cookie("auth_token")
	if cookie == nil || err != nil {
		fmt.Fprintf(w, "No cookie")
		return
	}

	s, err := authentication.GetSession(r, cookie.Value)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprintf(w, "Hello %s", s.Name)
}

func login(w http.ResponseWriter, r *http.Request) {
	authentication.CreateAuthCookie(w, r)
}

func init() {
	r := mux.NewRouter()
	getRoutes := r.Methods("GET", "POST").Subrouter()
	getRoutes.Handle("/login", handlers.HTTPMethodOverrideHandler(handler))
	getRoutes.HandleFunc("/", handler)
	http.Handle("/", r)
}

func main() {}
