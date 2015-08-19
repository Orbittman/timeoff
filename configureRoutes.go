package timeoff

import(
	"github.com/gorilla/mux"
	"github.com/gorilla/context"
	"net/http"
	"github.com/Orbittman/timeoff/services"
	"github.com/Orbittman/timeoff/authentication"
)

func ConfigureRoutes(r *mux.Router){
	getRoutes := r.Methods("GET").Subrouter()
	getRoutes.HandleFunc("/", secureHandler(services.RootHandler))
	getRoutes.HandleFunc("/hash/{value}", services.GetHash)
	
	postRoutes := r.Methods("POST").Subrouter()
	postRoutes.HandleFunc("/gotchi", secureHandler(services.PostGotchiHandler))
	postRoutes.HandleFunc("/login", services.Login)
	postRoutes.HandleFunc("/register", services.Register)
	
	postRoutes.HandleFunc("/food", secureHandler(services.PostFoodHandler))
}

func apiKeyHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("X-ApiKey")
		if header == "" {
			http.Error(w, "No API Key", 401)
			return
		}
		
		context.Set(r, "X-ApiKey", header)	
		next(w, r)
	}
}

func secureHandler(next http.HandlerFunc) http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth_token")
		if cookie == nil || err != nil {
			http.Error(w, "No auth cookie", 401)
			return
		}
		
		_, err = authentication.GetSession(r, cookie.Value)
		if err != nil {
			http.Error(w, "Invalid cookie", 401)
			return
		}
		
		next(w, r)
	}
	
	return apiKeyHandler(handler)
}