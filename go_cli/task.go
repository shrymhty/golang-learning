package main

import "time"

type Task struct {
	ID        int
	Title     string
	Status    string
	CreatedAt time.Time
}

const (
	statusTodo = "todo"
	statusInProgress = "in_progress"
	statusDone = "done"
)