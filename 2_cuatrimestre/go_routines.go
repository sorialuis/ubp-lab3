package main

import "fmt"

func main() {
	//go fmt.Println("Hola mundo")
	//go mostrarMensaje(55)

	for i := 0; i < 100; i++ {
		go mostrarMensaje(i)
	}

}

func mostrarMensaje(i int) {
	msg := fmt.Sprintf("hola mundo %d", i)
	fmt.Println(msg)
}
