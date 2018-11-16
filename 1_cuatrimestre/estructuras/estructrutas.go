
package estructuras

type Alumnos struct {
	Nombre   string
	Legajo   int64
	Promedio float32
}

func (al *Alumnos) Modificar() {
	al.Nombre = "Pepe"
}
