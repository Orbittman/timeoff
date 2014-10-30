package timeoff

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"timeoff/authentication"
	"timeoff/dto"
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
	request := dto.LoginRequest{}
	err := parsePost(r, &request)

	if err == nil {
		if request.UserName == "tim" && request.Password == "tim" {
			authentication.CreateAuthCookie(w, r)
		} else {
			http.Error(w, http.StatusText(403), 403)
		}
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	request := TestRequest{}

	err := parsePost(r, &request)
	if err == nil {
		tr := TestResponse{Message: request.Message, Success: request.Flag}
		w.Write(MapToJSON(tr))
	} else {
		panic(request)
	}
}

type Request interface{}

func parsePost(r *http.Request, generic Request) error {
	body, err := ioutil.ReadAll(r.Body)

	if err == nil {
		json.Unmarshal(body, &generic)
	}

	return err
}

func init() {
	r := mux.NewRouter()
	getRoutes := r.Methods("GET", "POST").Subrouter()
	getRoutes.HandleFunc("/login", login)
	getRoutes.HandleFunc("/post", post)
	getRoutes.HandleFunc("/", handler)
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

type TestResponse struct {
	Message string
	Success bool
}

type TestRequest struct {
	Message string
	Value   int
	Flag    bool
}
