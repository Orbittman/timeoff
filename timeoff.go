package timeoff

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/Orbittman/timeoff/authentication"
)

// Run with: "/Applications/go_appengine 2/dev_appserver.py" timeoff

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
	ConfigureRoutes(r)
	http.Handle("/", r)
}

func main() {}

func MapToJSON(m interface{}) []byte {
	j, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return j
}
