package main

import (
	"fmt"

	"github.com/SaItFish/GoN00B/sundries/todo_app/config"
	"github.com/SaItFish/GoN00B/sundries/todo_app/models"
	"github.com/SaItFish/GoN00B/sundries/todo_app/routes"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Todo{})

	r := routes.SetupRouter()

	r.Run()
}
