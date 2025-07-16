package main

import (
	"encoding/json"
	"fmt"
	"github.com/mergestat/timediff"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ./todo add "description"
func add_task(tasks []Task, args []string) []Task {
	newTask := Task{
		ID:          0,
		Description: args[0],
		Status:      0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append([]Task{newTask}, tasks...)
	for i := range tasks {
		tasks[i].ID = tasks[i].ID + 1
	}
	list_tasks([]Task{tasks[0]}, []string{})
	return tasks
}

// ./todo update 1 uncheck|check
func update_task(tasks []Task, args []string) []Task {
	if len(args) < 2 {
		fmt.Println("Error with update command: $ ./todo update id uncheck|check")
	}
	idStr := args[0]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Error with id conversion:", err)
	}

	var status int
	switch args[1] {
	case "check", "c", "1":
		status = 1
	case "uncheck", "u", "0":
		status = 0
	default:
		fmt.Println("Error updating task: $ ./todo update id uncheck|check")
	}

	flag := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			list_tasks([]Task{tasks[i]}, []string{})
			flag = true
			break
		}
	}
	if !flag {
		fmt.Println("Error updating id not in tasks")
		return tasks
	}

	return tasks
}

// ./todo delete id
func delete_task(tasks []Task, args []string) []Task {
	idStr := args[0]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(err)
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
	for i := range tasks {
		tasks[i].ID = i + 1
	}
	list_tasks(tasks, []string{})
	return tasks
}

// ./todo list
func list_tasks(tasks []Task, args []string) []Task {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)
	fmt.Fprintln(w, "ID\tDescription\tCreated At\tUpdated At\tStatus\t")
	for _, task := range tasks {
		var status rune
		if task.Status == 1 {
			status = '✓'
		} else {
			status = 'X'
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
			strconv.Itoa(task.ID),
			task.Description,
			timediff.TimeDiff(task.CreatedAt),
			timediff.TimeDiff(task.UpdatedAt),
			string(status),
		)
	}
	err := w.Flush()
	if err != nil {
		fmt.Println("Error flushing tabwriter:", err)
	}
	return tasks
}

func help() {
	fmt.Println("Error: no command found")
}

func main() {
	commands := map[string]func([]Task, []string) []Task{
		"add":    add_task,
		"update": update_task,
		"delete": delete_task,
		"list":   list_tasks,
	}

	cmd := os.Args[1]
	args := os.Args[2:]
	if cmd == "help" {
		help()
		return
	}

	file, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if handler, ok := commands[cmd]; ok {
		var tasks []Task
		if len(data) > 0 {
			err := json.Unmarshal(data, &tasks)
			if err != nil && data != nil {
				fmt.Println("Error with Unmarshal:", err)
				return
			}
		}
		tasks = handler(tasks, args)
		newData, err := json.Marshal(tasks)
		if err != nil {
			fmt.Println("Error with Marshal:", err)
			return
		}
		file.Truncate(0)
		file.Write(newData)
	} else {
		fmt.Println("Error: command not found")
	}
}
