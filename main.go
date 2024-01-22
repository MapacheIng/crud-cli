package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/MapacheIng/crud-cli/tasks"
)

func main() {
	// aqui creamos la vairble del archivo json
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	// que siempre al final cierre el archivo y no gaste recursos
	defer file.Close()

	var tasks []tasks.Task

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
	}

	fmt.Println(tasks)



}
