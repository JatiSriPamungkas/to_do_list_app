package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"text/tabwriter"
	"to_do_list_app/store"
)

func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func TitleGenerator(title string) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, len(title) + 14, '-', 0)
	h := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)

	fmt.Fprintln(w, "\t") 
	w.Flush()
	fmt.Fprintln(h, "|\t", title,"\t|")
	h.Flush()
	fmt.Fprintln(w, "\t") 
	w.Flush()
}

func ShowMenu() {
	TitleGenerator("Management Task CLI")
	fmt.Println()
	fmt.Println("1 | Create Task")
	fmt.Println("2 | Update Task")
	fmt.Println("3 | Read Task")
	fmt.Println("4 | Delete Task")
	fmt.Println("5 | Exit")
}

func ShowList() {
	for index, task := range store.GetTasks() {
		status := "Pending"
		if task.Done {
			status = "Done"
		}
		fmt.Printf("%d. %s [%s]\n", index + 1, task.Title, status)
	}
}