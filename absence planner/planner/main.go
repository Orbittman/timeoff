package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()
	getRoutes := r.Methods("GET", "POST").Subrouter()
	getRoutes.HandleFunc("/products", ProductsHandler)
	getRoutes.HandleFunc("/", handler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
