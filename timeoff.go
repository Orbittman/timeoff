package timeoff

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// Run with: "/Applications/go_appengine 2/dev_appserver.py" timeoff

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
