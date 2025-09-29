package main

import (
	"log"
	"net/http"

	"github.com/JuanLopezAranzazu/go-restapi-crud/db"
	"github.com/JuanLopezAranzazu/go-restapi-crud/models"
	"github.com/JuanLopezAranzazu/go-restapi-crud/routes"
	"github.com/gorilla/mux"
)

func main() {
	// conexion con la base de datos
	db.DBConnection()

	// migraciones de las tablas
	if err := db.DB.AutoMigrate(models.Product{}); err != nil {
		log.Fatal("Error en migración de tablas: ", err)
	}

	// manejar rutas
	r := mux.NewRouter()
	// index
	r.HandleFunc("/", routes.HomeHandler)

	s := r.PathPrefix("/api/v1").Subrouter()

	// productos
	s.HandleFunc("/products", routes.GetProducts).Methods("GET")
	s.HandleFunc("/products/{id}", routes.GetProduct).Methods("GET")
	s.HandleFunc("/products", routes.CreateProduct).Methods("POST")
	s.HandleFunc("/products/{id}", routes.UpdateProduct).Methods("PUT")
	s.HandleFunc("/products/{id}", routes.DeleteProduct).Methods("DELETE")

	// iniciar servidor
	log.Println("Servidor iniciado en http://localhost:3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}
