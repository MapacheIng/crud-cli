package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	task "github.com/MapacheIng/crud-cli/tasks"
)

func main() {
	// aqui creamos la vairble del archivo json
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	// que siempre al final cierre el archivo y no gaste recursos
	defer file.Close()

	var tasks []task.Task

	// informacion del archivo json
	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	// la condicion esta respecto al peso del archivo, para saber si esta vacio o no.
	if info.Size() != 0 {
		//si el archivo no esta vacio entra aqui

		//aqui vamos a leer el archivo, el cual retorna un arreglo de bytes (texto)
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}
		// luego el texto lo convertimos en json, y lo guardas en la varible que ya teniamos (el slice de tareas)
		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}
	} else {
		tasks = []task.Task{}
	}

	//aqui ya estamos comprobando lo que entra por consola
	// al momento de utilizar esta funcion debe retornar 2 parametros
	// ["ruta del programa" parametro]
	if len(os.Args) < 2 {
		printUsage()
	}

	switch os.Args[1] {
	case "list":
		task.ListTask(tasks)

	case "add":
		fmt.Println("cual es tu tarea?: ")
		// leemos lo que entra por consola (la mejor forma de leer la terminal)
		reader := bufio.NewReader(os.Stdin)
		// aqui convertimos la cadea en un string
		name, _ := reader.ReadString('\n')
		// aqui borramos espacios adicionales (cadena en limpio)
		name = strings.TrimSpace(name)

		// aqui del paquete task la anadimos a la variable tasks
		tasks = task.AddTask(tasks, name)
		task.SaveTasks(file, tasks)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("debes de proporcionar un ID para eliminar")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("el id debe ser un numero")
			return
		}

		tasks = task.DeleteTask(tasks, id)
		task.SaveTasks(file, tasks)
		fmt.Println("Tarea eliminada")

	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("debes de proporcionar un ID para Completar")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("el id debe ser un numero")
			return
		}
		tasks = task.CompleteTask(tasks, id)
		task.SaveTasks(file, tasks)
		fmt.Println("la tarea fue completada")

	default:
		printUsage()
	}

}

func printUsage() {
	fmt.Println("Uso: go-clid-crud [list|add|complete|delete]")
}
