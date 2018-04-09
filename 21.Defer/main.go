package main

import "fmt"

func main() {
	//NOS PERMITE EJECUTAR CIERTAS INSTRUCCIONES ANTES DE QUE TERMINE EL PROGRAMA
	//SE EJECUTEN EN ORDEN INVERSO A LA QUE SON CREADAS
	// EN EL EJEMPLO SE EJECUTARA CERRARSETDATOS Y LUYEGO CERRARBD
	fmt.Println("Conectando a nuestra Base de datos")
	defer cerrarBD()
	//VAS A EJECUTAR cerrarBD una vez que todas las instrucciones anteriores se ejecuten si por alguna razon no
	// se ejecuta correctamente lo anterior cerrarBD no se ejecuta
	fmt.Println("Consultando set de datos")
	defer cerrarSetDatos()
}
func cerrarBD() {
	fmt.Println("Cerrando Base de datos")
}
func cerrarSetDatos() {
	fmt.Println("Cerrar set de datos")
}
