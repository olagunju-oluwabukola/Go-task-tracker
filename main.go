package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct{
	ID int
	Description string
	Done bool
}

var tasks []Task

func main(){
	data, err :=os.ReadFile("tasks.json")
	if err == nil{
		json.Unmarshal(data, &tasks)
	}
	if len(os.Args) < 2{
		fmt.Println("Please provide a command: add, list, or delete")
		return
		}

		command:=os.Args[1]
		switch command {
		case "add":
			if len(os.Args) < 3{
				fmt.Println("Error: missing description")
				return
			}
			description := os.Args[2]
			maxId := 0
			for _, task := range tasks{
				if task.ID > maxId{
					maxId = task.ID
				}
			}
			newTask := Task{ID: maxId +1, Description: description, Done: false}
			tasks = append(tasks, newTask)
			saveTask()
			fmt.Println("Task added successfully")


		case "list" :
			fmt.Println("your Tasks:")
			for _, task := range tasks{
				status := "❌"
				if task.Done {
					status = "✔️"
				}
				fmt.Printf(" [%s] %d: %s\n", status, task.ID, task.Description)
			}

		case "delete":
			if len(os.Args) < 3{
			fmt.Println("Error: provide ID to delete")
			return
			}
			id,_ := strconv.Atoi(os.Args[2])
			index := -1
			for i, task := range tasks{
				if task.ID == id{
					index = i
					break
				}
			}
			if index == -1 {
				fmt.Println("Error : Task ID not found")
				return
			}

			tasks = append(tasks[:index], tasks[index+1:]... )
			saveTask()
			fmt.Println("Task deleted successfully")

				case "done" :
		if len(os.Args) < 3{
			fmt.Println("Error : provide the id to mark it as done")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		for i, task := range tasks{
			if task.ID == id{
				tasks[i].Done = true
				saveTask()
				fmt.Println("Task marked as done")
				return
			}
		}
		fmt.Print("Error: Task ID not found")
		}

	}

	func saveTask(){
		data, _:= json.Marshal(tasks)
		os.WriteFile("tasks.json", data, 0644)
	}
