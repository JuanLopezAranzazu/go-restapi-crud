package main

import (
	"net/http"

	"github.com/JuanLopezAranzazu/go-restapi-crud/db"
	"github.com/JuanLopezAranzazu/go-restapi-crud/routes"
	"github.com/gorilla/mux"
)

func main() {
	// conexion con la base de datos
	db.DBConnection()
	// manejar rutas
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r)
}
