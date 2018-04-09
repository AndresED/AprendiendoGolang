package animales

import "fmt"

// Perro es una struct que almacena los datos de perro
type Perro struct {
	Nombre string
}

// Comunicarse permite comunicarse con su nuevo dueño
func (p Perro) Comunicarse() {
	fmt.Println("Wofff")
}
