package main

import "fmt"

func main() {
	//ARRAYS , ESTRUCTURA DE DATOS QUE PERMITE ALMANCENAR MAS DE UN VALOR
	//var nombres [3]string
	nombres := [3]string{"Estuardo", "Andres", "Esquivel"}
	/*nombres[0] = "Estuardo"
	nombres[1] = "Andres"
	nombres[2] = "Esquivel"*/
	//IMPRIMIENDO ARRAY
	fmt.Println(nombres)
	//IMPRIMIENDO ELEMENTO DE ARRAY
	fmt.Println(nombres[0])
	fmt.Println(nombres[1])
	fmt.Println(nombres[2])
	//IMPRIMIENDO TAMAÑO DE ARRAY
	size := len(nombres)
	fmt.Println(size)
	//IMPRIMIR POSICIONES CONTINUAS DE UN ARRAY[inicio:final_excluyente]
	fmt.Println(nombres[0:2])
	/*
	   CARACTERISTICAS DE ARRAYS EN GO
	   1. TODOS LOS CALORES DEBEN DE SER DEL MISMO TIPO
	   2. TAMAÑO FIJO
	*/
}
