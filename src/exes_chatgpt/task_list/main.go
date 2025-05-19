package main

import (
	"fmt"
	"time"
)

type Task struct {
	Name       string
	Status     bool
	DateCreate time.Time
}

type TaskList struct {
	Tasks []Task
}

func (t *TaskList) AddTask(title string) {
	res := Task{
		Name:       title,
		Status:     false,
		DateCreate: time.Now(),
	}

	t.Tasks = append(t.Tasks, res)
}

func (t *TaskList) ListTasks() {

	for i, task := range t.Tasks {
		if task.Status {
			fmt.Printf("%d: %s | Готово    | %s\n", i, task.Name, task.DateCreate)
		} else {
			fmt.Printf("%d: %s | Не готово | %s\n", i, task.Name, task.DateCreate)
		}
	}
}

func (t *TaskList) CompleteTask(index uint16) {
	t.Tasks[index].Status = true
}

func main() {
	petr := TaskList{}
	petr.AddTask("Купить хлеб")
	petr.AddTask("Купить воду")
	petr.ListTasks()
	petr.CompleteTask(0)
	petr.ListTasks()

}
