package models

type Status int

const (
	NotStarted Status = iota
	InProgress
	Completed
)
