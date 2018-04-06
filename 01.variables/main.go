/*LAS VARIABLES SON ESPACIO RESERVADOS DE MEMORIA QUE CONTIENEN UN VALOR Y QUE
ESE VALOR PUEDE SER OBTENIDO CUANDO ESTAN EN EJECUCION*/
package main

import "fmt"

func main() {
	//var nombre, apellido string;
	//nombre = "Andres"
	//apellido = "Esquivel"
	nombre, apellido := "Andres", "Esquivel"
	nombre, edad := "Estuardo", 26
	fmt.Println("Mi nombre es ", nombre, ", mi apellido es", apellido, "y mi edad es", edad)
}
