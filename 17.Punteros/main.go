package main

import "fmt"

func main() {
	/*
	   PUNTERO => ES LA DIRECCION EN MEMORIA DE UNA VARIABLE DE UN TIPO DETERMINADO
	               El valor cero de un puntero es nil
	*/
	a := 3
	fmt.Println(&a)
	fmt.Println("Antes de duplicar,a=", a)
	duplicar(&a)
	fmt.Println("Despues de duplicar,a=", a)
}

func duplicar(x *int) {
	*x = *x * 2
	fmt.Println("Dentro de la funcion duplicar vale", *x)
}
