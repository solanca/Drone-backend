package main

import (
	"log"
	"net/http"

	"drone-backend/internal/database"
	"drone-backend/internal/handlers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	err := database.Initialize()
	if err != nil {
		log.Fatal("Failed to initialize database: ", err)
	}

	handlers.InitDB(database.DB)
	
	r := mux.NewRouter()
	// r.HandleFunc("/ping", handlers.PingHandler).Methods(http.MethodGet)

	r.HandleFunc("/drones", handlers.CreateDroneHandler).Methods(http.MethodPost)
	r.HandleFunc("/drones", handlers.GetDronesHandler).Methods(http.MethodGet)
	r.HandleFunc("/drones/{id:[0-9]+}", handlers.GetDroneHandler).Methods(http.MethodGet)
	r.HandleFunc("/drones/{id:[0-9]+}", handlers.UpdateDroneHandler).Methods(http.MethodPut)
	r.HandleFunc("/drones/{id:[0-9]+}", handlers.DeleteDroneHandler).Methods(http.MethodDelete)
	r.HandleFunc("/drones/zone/{zone}", handlers.GetDroneByZoneHandler).Methods(http.MethodGet)

	r.HandleFunc("/attributes", handlers.CreateAttributeHandler).Methods(http.MethodPost)
	r.HandleFunc("/attributes", handlers.GetAttributesHandler).Methods(http.MethodGet)
	r.HandleFunc("/attributes/{id:[0-9]+}", handlers.GetAttributeHandler).Methods(http.MethodGet)
	r.HandleFunc("/attributes/{id:[0-9]+}", handlers.UpdateAttributeHandler).Methods(http.MethodPut)
	r.HandleFunc("/attributes/{id:[0-9]+}", handlers.DeleteAttributeHandler).Methods(http.MethodDelete)
	r.HandleFunc("/attributes/{name}", handlers.GetAttributeByNameHandler).Methods(http.MethodGet)

	r.HandleFunc("/policies", handlers.CreatePolicyHandler).Methods(http.MethodPost)
	r.HandleFunc("/policies", handlers.GetPoliciesHandler).Methods(http.MethodGet)
	r.HandleFunc("/policies/{id:[0-9]+}", handlers.GetPolicyHandler).Methods(http.MethodGet)
	r.HandleFunc("/policies/{id:[0-9]+}", handlers.UpdatePolicyHandler).Methods(http.MethodPut)
	r.HandleFunc("/policies/{id:[0-9]+}", handlers.DeletePolicyHandler).Methods(http.MethodDelete)

	r.HandleFunc("/access-request", handlers.AccessRequestHandler).Methods(http.MethodPost)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
		AllowCredentials: true,
	})
	// srv := server.NewServer()
	handler := c.Handler((r));
	log.Fatal(http.ListenAndServe(":8080", handler))
}