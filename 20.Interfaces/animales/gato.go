package animales

import "fmt"

// Gato es una struct que almacena los datos de gato
type Gato struct {
	Nombre string
}

// Comunicarse permite comunicarse con su nuevo dueño
func (g Gato) Comunicarse() {
	fmt.Println("Miau")
}
