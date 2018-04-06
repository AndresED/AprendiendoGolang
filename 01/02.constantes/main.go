/*LAS CONSTANTES SON UN TIPO DE VARIABLE QUE SOLO PUEDEN ALMACENAR UN UNICO VALOR , VALOR QUE NO PUEDE
SER CAMBIADO A LO LARGO DE LA EJECUCION DE UN PROGRAMA*/
package main

import "fmt"

func main() {
	const nombre = "Andres"
	/*
	   nombre = "Estuardo"   => incorrecto
	   No se puede cambiar el contenido de una constante o el compilador nos votara un error
	*/
	fmt.Println(nombre)
}
