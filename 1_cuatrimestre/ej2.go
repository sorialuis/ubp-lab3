package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Es palindromo: ",
		esPalindromo("Anita lava la tina"))
}

func esPalindromo(cadena string) bool {
	cadena = preparar(cadena)
	n := len(cadena)
	for i, j := 0, n-1; i < n/2; i, j = i+1, j-1 {
		if cadena[i] != cadena[j] {
			return false
		}
	}

	return true
}

func preparar(cadena string) string {
	cadena = strings.ToLower(cadena)
	return borrar(cadena, " ")
}

func borrar(cadena string, patron string) string {
	return strings.Replace(cadena, patron, "", -1)
}
