package main

import (
	"fmt"
	"sync"
)

func main() {

	//creo un wait group
	var wg sync.WaitGroup

	//informo al wait group que voy a ejecutar 2 go routines
	wg.Add(2)

	go mostrarMensaje("Hola", &wg)
	go mostrarMensaje("Chau!", &wg)

	//hacemos que el wait group espere la finalizacion de las go routines
	wg.Wait()
}

func mostrarMensaje(msj string, wg *sync.WaitGroup) {
	fmt.Println(msj)

	wg.Done()
}
