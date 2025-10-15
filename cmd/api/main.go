package api

import (
	"fmt"

	"github.com/pithandev/fuzzy-octo-meme/internal/database"
)

func main() {
	// 1) connecto to db
	// 2) check if connection is successfully
	if err := database.Connect(); err != nil {
		fmt.Print("Error trying to connect with database!")
	}

	// 3) print successfull message

	fmt.Print("DB CONNECTED!")

	// 4) keep program running

}
