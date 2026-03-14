package main

import (
	"encoding/json"
	"os"
)

const fileName = "tasks.json"

func LoadTasks() ([]Task, error) {
	var tasks []Task

	file, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

func SaveTasks(tasks []Task) error {

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}