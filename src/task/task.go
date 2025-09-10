package task

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"to_do_list_app/src/models"
	"to_do_list_app/src/utils"
	"to_do_list_app/store"
)


func CreateTask() {
	utils.ClearScreen()

	utils.TitleGenerator("Create Task")
	fmt.Println()
	
	fmt.Print("Add your Task: ")
	reader := bufio.NewReader(os.Stdin)
	taskTitle, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error at Input", err)
		return
	}
	taskTitle = strings.TrimSpace(taskTitle)
	fmt.Println()
	
	store.CreateTask(taskTitle)
	
	fmt.Println("Success to add Task!")

	fmt.Println()
	fmt.Print("-- Press Enter to continue")
	fmt.Scanln()
}

func UpdateTask() {
	for {
		utils.ClearScreen()
		utils.TitleGenerator("Update Task")
		fmt.Println()

		if store.TaskLength() == 0 {
			fmt.Println("Empty task!")

			fmt.Println()
			fmt.Print("-- Press Enter to continue")
			fmt.Scanln()
			return
		}

		var choice int

		fmt.Println("1 | Update Task Name")
		fmt.Println("2 | Update Task Status")
		fmt.Println("3 | Back")
		
		fmt.Println()
		fmt.Print("Select Menu > ")
		fmt.Scanln(&choice)
		fmt.Println()

		switch choice {
		case 1:
			utils.ClearScreen()
			utils.TitleGenerator("Update Task Name")
			fmt.Println()

			var index int
			var taskTitle string

			utils.ShowList()
			
			fmt.Println()
			fmt.Print("ID: ")
			fmt.Scanln(&index)
			if index < 1 || index > store.TaskLength() {
				fmt.Println("ID can't found!")

				fmt.Println()
				fmt.Print("-- Press Enter to continue")
				fmt.Scanln()
				return
			}
			fmt.Print("New Task: ")
			reader := bufio.NewReader(os.Stdin)
			taskTitle, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error at Input", err)
				return
			}
			taskTitle = strings.TrimSpace(taskTitle)

			newTask := models.Task {
				ID: index,
				Title: taskTitle,
				Done: false,
			}

			fmt.Println()
			if index - 1 >= 0 && index - 1 < store.TaskLength() {
				store.UpdateTask(index, newTask)
				fmt.Println("Success to update Task Name!")
				} else {
					fmt.Println("Failed to update Task Name!")
				}

			fmt.Println()
			fmt.Print("-- Press Enter to continue")
			fmt.Scanln()

			utils.ClearScreen()
		case 2:
			utils.ClearScreen()
			utils.TitleGenerator("Update Task Status")
			fmt.Println()
			
			var index int
			var confirm string
			var done bool = false

			utils.ShowList()
			
			fmt.Println()
			fmt.Print("ID: ")
			fmt.Scanln(&index)
			if index < 1 || index > store.TaskLength() {
				fmt.Println("ID can't found!")
				
				fmt.Println()
				fmt.Print("-- Press Enter to continue")
				fmt.Scanln()
				return
			}
			fmt.Print("Are you sure to change Task Status? (y/n) ")
			fmt.Scanln(&confirm)
			
			if confirm == "y" {
				done = true
			}

			newTask := models.Task {
				ID: index,
				Title: store.GetTaskByIndex(index - 1).Title,
				Done: done,
			}

			fmt.Println()
			if index - 1 >= 0 && index - 1 < store.TaskLength() && confirm == "y" {
					store.UpdateTask(index, newTask)
					fmt.Println("Success to update Task Status!")
				} else {
					fmt.Println("Failed to update Task Status!")
				}

			fmt.Println()
			fmt.Print("-- Press Enter to continue")
			fmt.Scanln()

			utils.ClearScreen()
		case 3:
			return
		default:
			fmt.Println("Menu doesn't exist!")
			fmt.Println()
			fmt.Print("-- Press Enter to continue")
			fmt.Scanln()
			utils.ClearScreen()
		}
	}
}

func ReadTask() {
	utils.ClearScreen()
	utils.TitleGenerator("List Task")
	fmt.Println()

	if store.TaskLength() == 0 {
		fmt.Println("Empty task!")

		fmt.Println()
		fmt.Print("-- Press Enter to continue")
		fmt.Scanln()
		return
	}

	utils.ShowList()

	fmt.Println()
	fmt.Print("-- Press Enter to continue")
	fmt.Scanln()
}

func DeleteTask() {
	utils.ClearScreen()
	utils.TitleGenerator("Delete Task")
	fmt.Println()

	if store.TaskLength() == 0 {
		fmt.Println("Empty task!")

		fmt.Println()
		fmt.Print("-- Press Enter to continue")
		fmt.Scanln()
		return
	}
	
	utils.ShowList()

	var index int
	var confirm string

	fmt.Println()
	fmt.Print("ID: ")
	fmt.Scanln(&index)
	if index < 1 || index > store.TaskLength() {
		fmt.Println("ID can't found!")

		fmt.Println()
		fmt.Print("-- Press Enter to continue")
		fmt.Scanln()
		return
	}

	fmt.Print("Are you sure to change Task Status? (y/n) ")
	fmt.Scanln(&confirm)

	fmt.Println()
	if index - 1 >= 0 && index - 1 < store.TaskLength() && confirm == "y" {
		store.DeleteTask(index)
		fmt.Println("Success to delete Task!")
	} else {
		fmt.Println("Failed to delete Task!")
	}

	fmt.Println()
	fmt.Print("-- Press Enter to continue")
	fmt.Scanln()
}