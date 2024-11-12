package controllers

import (
	"strconv"
	"todolist-go/models"
	"todolist-go/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ActivityController struct {
	Service  *services.ActivityService
	Validate *validator.Validate
}

func NewActivityController(service *services.ActivityService, validate *validator.Validate) *ActivityController {
	return &ActivityController{Service: service, Validate: validate}
}

func (controller *ActivityController) GetActivities(c *fiber.Ctx) error {
	activities, err := controller.Service.GetActivities()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(activities)
}

func (controller *ActivityController) CreateActivity(c *fiber.Ctx) error {
	var activity models.Activity
	if err := c.BodyParser(&activity); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := controller.Validate.Struct(activity); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := controller.Service.CreateActivity(activity)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Activity created successfully"})
}

func (controller *ActivityController) UpdateActivity(c *fiber.Ctx) error {
	var activity models.Activity
	if err := c.BodyParser(&activity); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := controller.Validate.Struct(activity); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := controller.Service.UpdateActivity(activity)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Activity updated successfully"})
}

func (controller *ActivityController) DeleteActivity(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = controller.Service.DeleteActivity(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Activity deleted successfully"})
}

func (controller *ActivityController) GetActivityById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	activity, err := controller.Service.GetActivityById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(activity)
}
