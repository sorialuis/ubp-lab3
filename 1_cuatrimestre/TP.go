package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := 6
	m := crearBosque(n)
	fmt.Println("Bosque creado")
	cargarBosque(m)
	fmt.Println("Bosque cargado")
	ejecutarSimulacion(m)
	mostrarBosque(m)
}

func crearBosque(n int) [][]int {
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
		for j := range m[i] {
			m[i][j] = 0
		}
	}

	return m
}

func cargarBosque(m [][]int) {
	posicion1 := 0
	posicion2 := 0
	rand.Seed(time.Now().Unix())
	for i := 0; i < (len(m)*len(m))/2; i++ {
		posicion1 = rand.Intn(len(m))
		posicion2 = rand.Intn(len(m))
		m[posicion1][posicion2] = rand.Intn(2) + 1
	}

}

func ejecutarPaso(m [][]int) {

}

func ejecutarSimulacion(m [][]int) {

}

func mostrarBosque(m [][]int) {
	fmt.Println(m)
}
