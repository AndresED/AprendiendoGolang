package main

import "fmt"

// Persona struct que almacenara los datos de una persona
type Persona struct {
	nombre string
	edad   uint8
}

//  PARA ASIGNAR UN METODO A UN TIPO DE DATO BASTA CON COLOCAR ANTES DEL NOMBRE DE LA FUNCION
// EL TIPO DE DATO AL QUE QUIERES ASIGNARLO
func (p Persona) saludar() {
	fmt.Println("Hola soy una persona")
}
func (p *Persona) asignarEdad(i uint8) {
	if i >= 0 {
		p.edad = i
	} else {
		fmt.Println("No es valida la edad negativa")
	}
}

//Numero es un numero
//SE PUEDE ASOCIAR UN METODO A CUALQUIER TIPO DE DATO
type Numero int

func (n Numero) presentarse() {
	fmt.Println("Hola soy el", n)
}
func main() {
	//EN GO FUNCION QUE ESTA ASOCIADO A UN TIPO DE DATO
	var persona Persona
	//persona.nombre = "Andres Esquive"
	//persona.edad = 23
	persona.saludar()
	var numero Numero
	numero.presentarse()
	persona.asignarEdad(35)
	fmt.Println(persona)
}
