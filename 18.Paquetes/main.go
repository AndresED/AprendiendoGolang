package main

import (
	"fmt"

	"github.com/AndresED/Curso/18.Paquetes/despedir"
	"github.com/AndresED/Curso/18.Paquetes/saludar"
)

func main() {
	nombre := "Andres Esquivel"
	saludar.Saludar(nombre)
	saludar.MeVes = "Este es un texto asignado desde el main"
	fmt.Println(saludar.MeVes)
	despedir.Despedirse("Andres")
}
