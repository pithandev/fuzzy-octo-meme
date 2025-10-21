package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pithandev/fuzzy-octo-meme/internal/database"
	"github.com/pithandev/fuzzy-octo-meme/internal/handlers"
)

func main() {
	// 1) connecto to db
	// 2) check if connection is successfully
	if err := database.Connect(); err != nil {
		log.Fatal("❌ Database connection failed:", err)
	}

	// 3) print successfull message

	fmt.Println("DB CONNECTED!")

	//setting up routes
	setupRoutes()

	port := ":8080"
	fmt.Printf("🚀 Server starting on port http://localhost%s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("❌ Server failed:", err)
	}

}

func setupRoutes() {

	//health endpoints
	http.HandleFunc("/health", handlers.HealthCheck)
	http.HandleFunc("/health/db", handlers.DatabaseHealthCheck)

}
