package routes

import (
	"todolist-go/controllers"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes function initializes all routes for the application
func SetupRoutes(app *fiber.App, activityController *controllers.ActivityController) {
	app.Get("/activities", activityController.GetActivities)
	app.Post("/activities", activityController.CreateActivity)
	app.Put("/activities", activityController.UpdateActivity)
	app.Delete("/activities/:id", activityController.DeleteActivity)
}
