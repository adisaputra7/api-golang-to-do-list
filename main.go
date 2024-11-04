package main // package declaration

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type Activity struct {
	ID           int    `json:"id"`
	Title        string `json:"title" validate:"required"`
	Category     string `json:"category" validate:"required,oneof=TASK EVENT"`
	Description  string `json:"description" validate:"required"`
	ActivityDate string `json:"activity_date" validate:"required"`
	Status       string `json:"status"`
	CretedAt     string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func initDB() (*sql.DB, error) {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// db_user := os.Getenv("API_DB")
	// initDB function
	dns := "user=postgres.orskzevwrnwccsmkiskg password=7#2PdZ2afYe-NWU host=aws-0-ap-southeast-1.pooler.supabase.com port=6543 dbname=postgres"
	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func main() {
	db, err := initDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	app := fiber.New()
	validate := validator.New()

	// GET /activities
	app.Get("/activities", func(c *fiber.Ctx) error {
		rows, err := db.Query("SELECT * FROM activities")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}
		defer rows.Close()

		var activities []Activity
		for rows.Next() {
			var activity Activity
			if err := rows.Scan(&activity.ID, &activity.Title, &activity.Category, &activity.Description, &activity.ActivityDate, &activity.Status, &activity.CretedAt, &activity.UpdatedAt); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
			}
			activities = append(activities, activity)
		}
		return c.JSON(activities)
	})

	app.Post("/activities", func(c *fiber.Ctx) error {
		var activity Activity
		if err := c.BodyParser(&activity); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}

		// validate activity
		if err = validate.Struct(&activity); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}

		sqlStatement := `INSERT INTO activities (title, category, description, activity_date, status) VALUES ($1, $2, $3, $4, $5) RETURNING id`
		err = db.QueryRow(sqlStatement, activity.Title, activity.Category, activity.Description, activity.ActivityDate, "NEW").Scan(&activity.ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}
		return c.Status(fiber.StatusCreated).JSON(activity)
	})

	app.Put("/activities/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var activity Activity
		if err := c.BodyParser(&activity); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}

		// validate activity
		if err = validate.Struct(&activity); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}

		sqlStatement := `UPDATE activities SET title=$1, category=$2, description=$3, activity_date=$4, status=$5, updated_at=NOW() WHERE id=$6 RETURNING id`
		err = db.QueryRow(sqlStatement, activity.Title, activity.Category, activity.Description, activity.ActivityDate, activity.Status, id).Scan(&activity.ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Activity updated successfully"})
	})

	// DELETE /activities/:id
	app.Delete("/activities/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		sqlStatement := `DELETE FROM activities WHERE id=$1`
		_, err := db.Exec(sqlStatement, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Activity deleted successfully"})
	})

	app.Listen(":8081")
}
