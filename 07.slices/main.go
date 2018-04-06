package main

import "fmt"

func main() {
	//SLICES => REPRESENTAN UNA SECUENCIA DE TAMAÑO VARIABLE DE ELEMENTOS DEL MISMO TIPO, tienen un apuntador a un
	//array
	/*
			   Apuntador a un array
			   TAMAÑO
			   caapcidad

		     make=> (tipo de dato del slice|tamaño inicial|capacidad inicial "opcional")
	*/
	/*var nombres []string
	fmt.Printf("Su capacidad es %d y su capacidad es %d \n", len(nombres), cap(nombres))
	nombres = append(nombres, "Estuardo")
	fmt.Printf("Su capacidad es %d y su capacidad es %d \n", len(nombres), cap(nombres))
	nombres = append(nombres, "Andres")
	fmt.Printf("Su capacidad es %d y su capacidad es %d \n", len(nombres), cap(nombres))
	nombres = append(nombres, "Díaz")
	fmt.Printf("Su capacidad es %d y su capacidad es %d \n", len(nombres), cap(nombres))
	nombres = append(nombres, "Díaz")
	fmt.Printf("Su capacidad es %d y su capacidad es %d \n", len(nombres), cap(nombres))
	fmt.Println(nombres)
	fmt.Println("MODIFICANDO TERCER ELEMENTO")
	nombres[2] = "Esquivel"*/
	nombres := []string{"Estuard", "Andres", "Esquivel", "Díaz"}
	fmt.Println(nombres)
}
