package task

import "fmt"

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

func ListTask(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No hay tareas")
		return
	}

	for _, task := range tasks {
		fmt.Printf("%s \n", task.Name)
	}

}