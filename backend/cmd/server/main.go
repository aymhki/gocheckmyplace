package main

import (
	"backend/internal/db"
	"backend/internal/handlers"
	"backend/internal/middleware"
	"log"
	"net/http"
	"os"
)

func main() {
	database, err := db.ConnectDB(os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Printf("warning: db not reachable at startup: %v", err)
	} else {
		log.Println("connected to db successfully")
	}

	h := handlers.NewHandler(database)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", h.Health)
	mux.HandleFunc("/api/db-check", h.DBCheck)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Printf("listening at port %s", port)

	if err := http.ListenAndServe(":"+port, middleware.CORS(mux)); err != nil {
		log.Fatal(err)
	}
}
