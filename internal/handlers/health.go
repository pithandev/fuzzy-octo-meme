package handlers

import (
	"net/http"

	"github.com/pithandev/fuzzy-octo-meme/internal/database"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "healthy", "service": "todo-api"}`))
}

func DatabaseHealthCheck(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	w.Header().Set("Content-Type", "application/json")

	// check db connection
	sqlDB, err := db.DB()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"status": "unhealthy", "database": "connection error"}`))
		return
	}

	if err := sqlDB.Ping(); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"status": "unhealthy", "database": "ping failed"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "healthy", "database": "connected"}`))

}
