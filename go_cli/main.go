package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cli_tracker <command> [arguments]")
		fmt.Println("  add <task title> - Add a new task")
		fmt.Println("  list - List all tasks")
		fmt.Println("  status <task id> <status> - Update task status (todo, in-progress, done)")
		fmt.Println("  delete <task id> - Delete a task")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide task title")
			return
		}

		err := AddTask(os.Args[2])
		if err != nil {
			fmt.Println(err)
		}
	
	case "list":
		err := ListTasks()
		if err != nil {
			fmt.Println(err)
		}
	
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Provide task id")
			return
		}

		id, _ := strconv.Atoi(os.Args[2])

		err := DeleteTask(id)
		if err != nil {
			fmt.Println(err)
		}

	case "status":
		if len(os.Args) < 4 {
			fmt.Println("Usage: status <id> <todo|in-progress|done>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		status := os.Args[3]

		err := UpdateTask(id, status)
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}