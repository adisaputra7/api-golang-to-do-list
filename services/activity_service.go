package services

import (
	"database/sql"

	"todolist-go/models"

	"github.com/go-playground/validator/v10"
)

type ActivityService struct {
	DB       *sql.DB
	Validate *validator.Validate
}

func NewActivityService(db *sql.DB, validate *validator.Validate) *ActivityService {
	return &ActivityService{DB: db, Validate: validate}
}

func (service *ActivityService) GetActivities() ([]models.Activity, error) {
	rows, err := service.DB.Query("SELECT * FROM activities")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []models.Activity
	for rows.Next() {
		var activity models.Activity
		err := rows.Scan(&activity.ID, &activity.Title, &activity.Category, &activity.Description, &activity.ActivityDate, &activity.Status, &activity.CreatedAt, &activity.UpdatedAt)
		if err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}
	return activities, nil
}

func (service *ActivityService) CreateActivity(activity models.Activity) error {
	_, err := service.DB.Exec("INSERT INTO activities (title, category, description, activity_date, status) VALUES ($1, $2, $3, $4, $5)", activity.Title, activity.Category, activity.Description, activity.ActivityDate, activity.Status)
	if err != nil {
		return err
	}
	return nil
}

func (service *ActivityService) UpdateActivity(activity models.Activity) error {
	_, err := service.DB.Exec("UPDATE activities SET title=$1, category=$2, description=$3, activity_date=$4, status=$5 WHERE id=$6", activity.Title, activity.Category, activity.Description, activity.ActivityDate, activity.Status, activity.ID)
	if err != nil {
		return err
	}
	return nil
}

func (service *ActivityService) DeleteActivity(id int) error {
	_, err := service.DB.Exec("DELETE FROM activities WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (service *ActivityService) GetActivityById(id int) (models.Activity, error) {
	var activity models.Activity
	err := service.DB.QueryRow("SELECT * FROM activities WHERE id=$1", id).Scan(&activity.ID, &activity.Title, &activity.Category, &activity.Description, &activity.ActivityDate, &activity.Status, &activity.CreatedAt, &activity.UpdatedAt)
	if err != nil {
		return models.Activity{}, err
	}
	return activity, nil
}
