package models

type Task struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status Status `json:"completed"`
}
