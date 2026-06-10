package controllers

import (
	"encoding/json"

	"github.com/diogocorigo/tasks/backend/models"
)

func GetTasks() string {
	tasks := []models.Task{
		{ID: 1, Name: "Task 1", Status: models.NotStarted},
		{ID: 2, Name: "Task 2", Status: models.InProgress},
		{ID: 3, Name: "Task 3", Status: models.Completed},
	}

	result, _ := json.Marshal(tasks)
	return string(result)
}

func CreateTask(title string) string {
	task := models.Task{
		ID:     4,
		Name:   title,
		Status: models.NotStarted,
	}

	result, _ := json.Marshal(task)
	return string(result)
}
