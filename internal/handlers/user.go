package handlers

import (
	"fmt"
	"net/http"

	"github.com/pithandev/fuzzy-octo-meme/internal/database"
	"github.com/pithandev/fuzzy-octo-meme/internal/models"
)

func CreateUser(user models.User) {
	db := database.GetDB()

	db.Create(&user)

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	result := db.Find(models.User{}, nil)

	fmt.Println(result)

	w.Write([]byte(result.Name()))
}
