package migration

import (
	"fmt"

	"github.com/muhrobby/go-redirect/database"
	"github.com/muhrobby/go-redirect/model"
)

func AutoMigrate() {
	err := database.Connect().AutoMigrate(&model.Shortlink{})

	if err != nil {
		fmt.Println("failed to migrate")
	}

	fmt.Println("success to migrate")
}
