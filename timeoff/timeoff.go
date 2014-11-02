package timeoff

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"timeoff/authentication"
	"timeoff/services"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi there, welcome to the TimeOff app\n")
	cookie, err := r.Cookie("auth_token")

	s, err := authentication.GetSession(r, cookie.Value)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprintf(w, "Hello %s", s.Name)
}

func init() {
	r := mux.NewRouter()
	getRoutes := r.Methods("GET").Subrouter()
	getRoutes.HandleFunc("/", secureHandler(rootHandler))
	getRoutes.HandleFunc("/hash/{value}", services.GetHash)

	postRoutes := r.Methods("POST").Subrouter()
	postRoutes.HandleFunc("/login", services.Login)
	http.Handle("/", r)
}

func secureHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth_token")
		if cookie == nil || err != nil {
			fmt.Fprintf(w, "No cookie")
			return
		}

		next(w, r)
	}
}

func main() {}

func MapToJSON(m interface{}) []byte {
	j, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return j
}
