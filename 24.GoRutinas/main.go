package main

import (
	"fmt"
	"time"
)

func main() {
	//EJECUTAR UNA ACTIVIDAD INDEPENDIENTE , MANTENIENDO UNA COMUNICACION ENTRE ELLOS
	var h string
	go mostrarNumeros() //CONVERTIMOS EN UNA GORUTINA AL ANTEPONER GO, EL CUAL NOS PERMITE EJECUTAR DE MANERA INDIPENDIENTE
	fmt.Print("Digita algo: ")
	fmt.Scan(&h)
	fmt.Println("Digitaste algo: ", h)
}
func mostrarNumeros() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}
