package main

import (
	"todolist-go/config"
	"todolist-go/controllers"
	"todolist-go/routes"
	"todolist-go/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	app := fiber.New()
	validator := validator.New()
	activityService := services.NewActivityService(db, validator)
	activityController := controllers.NewActivityController(activityService, validator)

	routes.SetupRoutes(app, activityController)

	app.Listen(":8080")
}
