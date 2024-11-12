package tests

import (
	"net/http"
	"testing"
	"todolist-go/config"
	"todolist-go/controllers"
	"todolist-go/routes"
	"todolist-go/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// Setup test environment
func setupTestEnvironment() {
	var app *fiber.App
	db, _ := config.InitDB()
	validator := validator.New()
	activityService := services.NewActivityService(db, validator)
	activityController := controllers.NewActivityController(activityService, validator)

	app = fiber.New()
	routes.SetupRoutes(app, activityController)
}

// GET /activities test
func TestGetActivities(t *testing.T) {
	setupTestEnvironment()

	app := fiber.New()
	req, _ := http.NewRequest("GET", "/activities", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}
