package main

import (
	"net/http"

	"github.com/JuanLopezAranzazu/go-restapi-crud/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r)
}
