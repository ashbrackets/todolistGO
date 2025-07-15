package main

import (
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func add_task(file *os.File, args []string) {

}

func update_task(file *os.File, args []string) {

}

func delete_task(file *os.File, args []string) {

}

func list_tasks(file *os.File, args []string) {

}

func help() {
	fmt.Println("Error: no command found")
}

func main() {
	commands := map[string]func(*os.File, []string){
		"add":    add_task,
		"update": update_task,
		"delete": delete_task,
		"list":   list_tasks,
	}

	cmd := os.Args[1]
	args := os.Args[2:]
	file, err := os.Create("data.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	if len(args) == 0 {
		help()
	}

	if handler, ok := commands[cmd]; ok {
		handler(file, args)
	} else {
		fmt.Println("Error: command not found")
	}
}
