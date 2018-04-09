package main

import "fmt"

func main() {
	f()
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%T\nr", r)
			fmt.Println("Recuperado en  f:", r)
		}
	}()
	g(5)
	fmt.Println("Retorna normalmente desde g")
}
func g(i int) {
	if i > 3 {
		fmt.Println("AaaaaaAAAAAAaaaHHH")
		panic("El numero no puede ser mayor a 3")
	}
	defer fmt.Println("Dfer en la funcion g")
	fmt.Println("Imprimiendo en g", i)
}
