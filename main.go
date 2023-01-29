package main

import (
	"github.com/Shiroyasha19/task-5-vix-btpns-AdjiPrayoga/database"
	"github.com/Shiroyasha19/task-5-vix-btpns-AdjiPrayoga/models"
	"github.com/Shiroyasha19/task-5-vix-btpns-AdjiPrayoga/router"
)

func main() {
	db := database.SetupDB()
	db.AutoMigrate(&models.User{})

	r := router.SetupRoutes(db)
	r.Run()
}
