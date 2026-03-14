package main

import (
	"time"
	"fmt"
)

func AddTask(title string) error {
	tasks, err := LoadTasks()

	if err != nil {
		return err
	}

	id := len(tasks) + 1

	task := Task{
		ID:        id,
		Title:     title,
		Status:    statusTodo,
		CreatedAt: time.Now(),
	}

	tasks = append(tasks, task)
	return SaveTasks(tasks)
}

func ListTasks() error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	for _, t := range tasks {
		fmt.Printf("\nTask %d\nTitle: %s\nStatus: %s\nCreated At: %s\n", t.ID, t.Title, t.Status, t.CreatedAt.Format(time.RFC1123)	)
	}
	return nil
}

func UpdateTask(id int, status string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	validStatus := map[string]bool{
		statusTodo:       true,
		statusInProgress: true,
		statusDone:       true,
	}

	if !validStatus[status] {
		return fmt.Errorf("invalid status")
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = status
			return SaveTasks(tasks)
		}
	}
	return fmt.Errorf("Task with ID %d is not found.", id)
}

func DeleteTask(id int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	index := -1

	for i, t := range tasks {
		if t.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("task with id %d not found", id)
	}

	tasks = append(tasks[:index], tasks[index+1:]...)

	return SaveTasks(tasks)
}