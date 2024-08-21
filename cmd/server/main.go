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

	// debug logging
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s", r.Method, r.URL)
			next.ServeHTTP(w, r)
		})
	})

	// CORS
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	})

	// routes
	r.HandleFunc("/storage", handler.GetItems).Methods("GET")
	r.HandleFunc("/storage", handler.CreateItem).Methods("POST")
	r.HandleFunc("/storage/{id}", handler.GetItemByID).Methods("GET")
	r.HandleFunc("/storage", handler.UpdateItem).Methods("PUT")
	r.HandleFunc("/storage/{id}", handler.DeleteItem).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
