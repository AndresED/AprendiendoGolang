package saludar

import "fmt"

// Meves es para utilizar desde otro paquete
var MeVes string

// noMeves No se puede utilizar desde otro paquete
var noMeves string

//Saludar => FUNCIONES QUE INICIAN CON MAYUSCULA INDICAN QUE ESTAS SON FUNCIONES EXPORTADAS , ES DECIR
// QUE PUEDEN SER USADAS EN OTRAS CLASES
func Saludar(nombre string) {
	fmt.Println("Hola", nombre)
}

func noVisible() {
	fmt.Println("No me ven desde otro paquete")
}
