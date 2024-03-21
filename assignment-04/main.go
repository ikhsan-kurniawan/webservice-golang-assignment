package main

import (
	"mygram/database"
	"mygram/models"
	"mygram/router"
)

func main() {
	db, err := database.StartDB()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Comment{}, &models.Photo{}, &models.SocialMedia{})
	if err != nil {
		panic(err)
	}

	r := router.StartApp(db)
	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}