package main

import "fmt"

//Persona es un tipo de dato struct
type Persona struct {
	Nombre string
	Edad   uint8
}

//Persona2 es un tipo de dato struct
type Persona2 struct {
	Nombre string
	Edad   uint8
	Emails []string
}

func main() {
	/*var persona1 Persona
	persona1.Nombre = "Andres Esquivel"
	persona1.Edad = 26
	fmt.Println(persona1.Nombre, persona1.Edad)
	*/
	/*
		  persona2 := Persona{}
			persona2.Nombre = "Andres Esquivel"
			persona2.Edad = 26

			fmt.Println(persona2)
			fmt.Println(persona2.Nombre, persona2.Edad)
	*/
	emails := []string{"a@b.com", "z@b.com"}
	persona3 := Persona2{
		Nombre: "Andres Esquivel",
		Edad:   26,
		//Emails: []string{"a@b.com", "z@b.com"},
		Emails: emails,
	}
	fmt.Println(persona3)
	fmt.Println(persona3.Nombre, persona3.Edad)
}
