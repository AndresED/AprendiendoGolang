package main

import "fmt"

func main() {
	//SE ASEMEJA AL FOREACH
	//PERMITE  ITERAR SSOBRE ARRAYS , SLICES.MAPS , CANALES
	/**ARRAYS **/
	fmt.Println("ARRAYS")
	numeros := []string{"cero", "uno", "dos", "tres"}
	//VALORES Y POSICION
	for i, v := range numeros {
		fmt.Println(i, v)
	}
	//OMITIENDO POSICION
	fmt.Println("omitiendo posiciones")
	for _, v := range numeros {
		fmt.Println(v)
	}
	//OMITIENDO VALORES
	fmt.Println("omitiendo valores")
	for i := range numeros {
		fmt.Println(i)
	}
	/** MAPAS **/
	fmt.Println("MAPAS")
	nombres := map[string]string{"es": "España", "co": "Colombia", "pe": "Peru"}
	for i, v := range nombres {
		fmt.Println(i, v)
	}
	/*STRINGS*/
	fmt.Println("STRINGS")
	frase := "Hola Mundo"
	for i, v := range frase {
		fmt.Println(i, string(v))
	}

	//recorriendo un array de enteros declaro en el for
	for _, entero := range []int{15, 36, 24, 25} {
		fmt.Println(entero)
	}
	//recorriendo una cadena declarada en el for
	for _, cadena := range "curso go desde cero" {
		fmt.Println(string(cadena))
	}
}
