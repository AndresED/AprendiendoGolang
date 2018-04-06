package main

import "fmt"

func main() {
	if !(5 < 10 && 1 > 7) == false {
		fmt.Println("ESTO ESTA EN EL BLOQUE VERDADERO")
	} else {
		fmt.Println("ESTO ESTA EN EL BLOQUE FALSE")
	}
	fmt.Println("FIN DEL PROGRAMA")
}
