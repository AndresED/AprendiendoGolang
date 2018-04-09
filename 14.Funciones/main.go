package main

import "fmt"

func main() {
	saludar()
	saludar2("Andres", 26)
	eresMayorEdad(16)
	fmt.Println(suma(12, 12))
	n := []int8{52, 63, 47, 5, 5, 3, 7, 6, 100, 54, -5, -35, 0}
	maximo, minimo := maxymin(n)
	fmt.Printf("Maximo %d \n", maximo)
	fmt.Printf("Minimo %d \n", minimo)
	saludarVarios(18, "Estuardo", "Andres", "Antonio", "Alcides")
}

//FUNCION SIN PARAMETROS
func saludar() {
	fmt.Println("Hola mundo")
}

//FUNCION CON PARAMETROS PARAMETROS
func saludar2(nombre string, edad uint8) {
	fmt.Printf("Hola %s, tu edad es %d \n", nombre, edad)
}

//FUNCIONES CON RETURN
func eresMayorEdad(edad uint8) {
	if edad >= 18 {
		return //EVIAOS QUYE SE SIGAN CONTINUAND QUE SE EJECUTEN LO QUE CONTINUA
	}
	fmt.Println("Eres menor de edad")
}
func suma(a int8, b int8) int8 {
	return a + b
}

//RETORNANDO MULTIPLES VALORES
func maxymin(num []int8) (int8, int8) {
	var max, min int8
	for _, v := range num {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min
}

//FUNCIONES VARIATICAS: CUANDO NO SE SABE EXACTAMENTE CUANTOS PARAMETROS SERAN ENVIADOS
//SOLO PUEDE ENVIARSE UN UNICO PARAMETRO VARIATICO Y DEBE DE ESTAR AL FINAL DE LOS PARAMETROS DE LA FUNCION
func saludarVarios(edad uint8, nombres ...string) {
	//NOMBRES ES UN SLICE
	fmt.Printf("%T \n", nombres)
	for _, v := range nombres {
		fmt.Printf("Hola %s tienes %d años de edad \n", v, edad)
	}
}
