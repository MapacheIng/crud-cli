package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

func ListTask(tasks []Task) {
	// la funcion sirve para retornar la lista de tareas y en un formato
	if len(tasks) == 0 {
		fmt.Println("No hay tareas")
		return
	}

	for _, task := range tasks {
		status := " "
		if task.Complete {
			status = "✓"
		}
		fmt.Printf(" [%s] %d  %s \n", status, task.ID, task.Name)
	}
}

func AddTask(tasks []Task, name string) []Task {
	//esta funcion es para escribir una nueva tarea y rotorna un arreglo de la estructura
	newTask := Task{
		ID:       GetNextID(tasks),
		Name:     name,
		Complete: false,
	}
	return append(tasks, newTask)
}

func SaveTasks(file *os.File, tasks []Task) {
	// esta funcion es para guardar en json los datos (se utiliza con la funcion anterior)
	// este metodo lo que hacemos es convertir un arreglo en json
	bytes, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}
	// aqui es para ubicar el puntero
	_, err = file.Seek(0,0)
	if err != nil {
		panic(err)
	}
	// para borrar el json y dejarlo en blanco por ahora
	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}
	// creamos una instancia de escritura sobre el archivo. (utilizando bufio)
	write := bufio.NewWriter(file)
	// escribimos el json en el archivo
	_, err = write.Write(bytes)
	if err != nil {
		panic(err)
	}
	// nos aseguramos que los datos si se escriban
	err = write.Flush()
	if err != nil {
		panic(err)
	}

}

func GetNextID(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}
	return tasks[len(tasks) - 1].ID + 1
}