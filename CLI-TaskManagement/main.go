package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID int
	Title string
	DueDate *time.Time
	Status string
}

type TaskStore struct {
	tasks []Task
}

func main () {
	store := TaskStore{}
	runCLI(&store)
}

func runCLI(store *TaskStore){
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter command: ")

		if !scanner.Scan() {
			break
		}

		command := scanner.Text()

		switch command {
		case "add":
			addTaskInteractive(store, scanner)
		case "list":
			store.listTasks()
		case "complete":
			fmt.Print("Enter task ID to complete: ")
			scanner.Scan()
			idStr := scanner.Text()

			id, err := strconv.Atoi(idStr)

			if err != nil {
				fmt.Println("Invalid ID")
				continue
			}

			store.completeTask(id)
		case "delete":
			fmt.Print("Enter task ID to delete: ")
			scanner.Scan()
			idStr := scanner.Text()

			id, err := strconv.Atoi(idStr)

			if err != nil {
				fmt.Println("Invalid ID")
				continue
			}

			store.deleteTask(id)
		case "exit":
			return
		}
	}
	
}

func (store *TaskStore) addTask(t *Task){
	store.tasks = append(store.tasks, *t)
}

func addTaskInteractive(store *TaskStore, scanner *bufio.Scanner){
	fmt.Print("Task title: ")
	scanner.Scan()
	title := scanner.Text()

	fmt.Print("Due date (YYYY-MM-DD): ")
	scanner.Scan()
	dueDateStr := scanner.Text()

	layout := "2006-01-02"
	parsed, err := time.Parse(layout, dueDateStr)

	if err != nil {
		fmt.Println("Invalid date format")
		return
	}

	task := Task{
		ID: len(store.tasks) + 1,
		Title: title,
		DueDate: &parsed,
		Status: "Pending",
	}

	store.addTask(&task)
	fmt.Println("Task added!")
}

func (store *TaskStore) listTasks(){
	fmt.Println("Tasks: ")
	fmt.Println(store.tasks)
}

func (store *TaskStore) completeTask(id int) error{
	for i := range store.tasks {
		if store.tasks[i].ID == id {
			store.tasks[i].Status = "Completed"
			return nil
		}
	}
	return fmt.Errorf("Task %d not found", id)
}

func (store *TaskStore) deleteTask(id int) error{
	for i, t := range store.tasks {
		if t.ID == id {
			store.tasks = append(store.tasks[:i], store.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Task %d not found", id)
}