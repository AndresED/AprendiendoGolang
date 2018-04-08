package main

import "fmt"

func main() {
	/*****************************FOR CLASICO****************************************/
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
			//break
		}
		fmt.Println(i)
	}
	for i := 5; i > 0; i-- {
		fmt.Println(i)
	}
	//RELLENANDO UNA MATRIZ
	valor := 0
	matriz := [3][3]int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			valor++
			matriz[i][j] = valor
		}
	}
	//IMPRIMIENDO MATRIZ
	fmt.Println(matriz)
	//RECORRIENDO MATRIZ
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(matriz[i][j])
		}
		fmt.Println("  ")
	}
}
