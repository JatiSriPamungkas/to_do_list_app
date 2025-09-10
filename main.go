package main

import (
	"fmt"

	"to_do_list_app/src/task"
	"to_do_list_app/src/utils"
	"to_do_list_app/store"
)

func main() {
	store.LoadTasks()
	defer store.SaveTasks()

	for {
		utils.ClearScreen()
		
		var choice int
		utils.ShowMenu()
		fmt.Println()
		fmt.Print("Select Menu > ")
		fmt.Scanln(&choice)
		fmt.Println()
	
		switch choice {
		case 1:
			task.CreateTask()
		case 2:
			task.UpdateTask()
		case 3:
			task.ReadTask()
		case 4:
			task.DeleteTask()
		case 5:
			utils.TitleGenerator("Thanks for using this Application!")
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