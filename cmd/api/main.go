package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pithandev/fuzzy-octo-meme/internal/database"
)

func main() {
	// 1) connecto to db
	// 2) check if connection is successfully
	if err := database.Connect(); err != nil {
		log.Fatal("❌ Database connection failed:", err)
	}

	// 3) print successfull message

	fmt.Print("DB CONNECTED!")

	// 4) keep program running

	port := ":8080"
	fmt.Printf("🚀 Server starting on port http://localhost%s\n", port)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("✅ API is healthy"))
	})

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("❌ Server failed:", err)
	}

}
