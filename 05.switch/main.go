package main

import "fmt"

func main() {

	/*
		    var a uint8
				  a = 2
				  switch a {
					case 1:
						fmt.Println("LUNES")
					case 2:
						fmt.Println("MARTES")
					case 3:
						fmt.Println("MIERCOLES")
					case 4:
						fmt.Println("JUEVES")
					case 5:
						fmt.Println("VIERNES")
					case 6:
						fmt.Println("SABADO")
					case 7:
						fmt.Println("DOMINGO")
					default:
						fmt.Println("NO ES UN DÍA DE LA SEMANA")
					}*/
	switch a := 2; {
	case a == 1:
		fallthrough
	case a == 2:
		fallthrough
	case a == 3:
		fallthrough
	case a == 4:
		fallthrough
	case a == 5:
		fmt.Println("ESTAS ENTRE SEMANA")
	case a == 6:
		fallthrough
	case a == 7:
		fmt.Println("ESTAS EN FIN DE SEMANA - DOMINGO")
	default:
		fmt.Println("NO ES UN DÍA DE LA SEMANA")
	}

}
