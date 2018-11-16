package main

import (
	"fmt"

	"github.com/sorialuis/ubp-lab3/1_cuatrimestre/estructuras"
)

/*func main() {
	alumn := &estructuras.Alumnos{
		Nombre:   "Pedro",
		Legajo:   63107,
		Promedio: 7.83,
	}
	fmt.Println("Viejo ", alumn.Nombre)
	modificarAlumno(alumn)
	fmt.Println("Nuevo", alumn.Nombre)
}
func modificarAlumno(al *estructuras.Alumnos) {
	al.Nombre = "Jose"
	al.Promedio = 8.02
}*/

func main() {
	alumn := &estructuras.Alumnos{
		Nombre:   "Pedro",
		Legajo:   63107,
		Promedio: 7.83,
	}
	fmt.Println("Viejo ", alumn.Nombre)
	alumn.Modificar()
	fmt.Println("Nuevo", alumn.Nombre)
}
