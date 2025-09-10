package store

import (
	"encoding/json"
	"fmt"
	"os"
	"to_do_list_app/src/models"
)

var tasks []models.Task

func SaveTasks() {
	file, err := os.Create("tasks.json")
	if err != nil {
		fmt.Println("Error saving tasks: ", err)
		return
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	encoder.Encode(tasks)
}

func LoadTasks() {
	file, err := os.Open("tasks.json")
	if err != nil {
		return
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.Decode(&tasks)
}

func TaskLength() int {
	return len(tasks)
}

func GetTasks() []models.Task {
    return tasks
}

func GetTaskByIndex(index int) models.Task {
	return tasks[index]
}

func CreateTask(taskTitle string) {
	newTask := models.Task {
		ID: len(tasks) + 1,
		Title: taskTitle,
		Done: false,
	}

	tasks = append(tasks, newTask)
}

func UpdateTask(index int, newTask models.Task) {
	tasks[index - 1] = newTask
}

func DeleteTask(index int) {
	tasks = append(tasks[:index - 1], tasks[index:]...)
}
