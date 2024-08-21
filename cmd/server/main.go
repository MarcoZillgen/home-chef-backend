package main

import (
	"log"
	"net/http"

	"MarcoZillgen/homeChef/internal/api"
	"MarcoZillgen/homeChef/internal/database"
	"MarcoZillgen/homeChef/internal/storage"

	"github.com/gorilla/mux"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	repo := storage.NewRepository(db)
	handler := api.NewHandler(repo)

	r := mux.NewRouter()

	// routes
	r.HandleFunc("/storage", handler.GetItems).Methods("GET")
	r.HandleFunc("/storage", handler.CreateItem).Methods("POST")
	r.HandleFunc("/storage/{id}", handler.GetItemByID).Methods("GET")
	r.HandleFunc("/storage", handler.UpdateItem).Methods("PUT")
	r.HandleFunc("/storage/{id}", handler.DeleteItem).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
