package main

// EN GOLANG => SON UN CONTRATO QUE SE PUEDEN IMPLEMENTAR EN CUALQUIER TIPO DE DATO EN GO

import "github.com/AndresED/Curso/20.Interfaces/animales"

func main() {
	var p animales.Perro
	var g animales.Gato
	p.Nombre = "Firularis"
	g.Nombre = "Manchas"
	//AdoptarPerro(p)
	//AdoptarGato(g)
	adoptarMascota(p)
	adoptarMascota(g)
}

// AdoptarPerro nos permite adoptar un perro
/*func AdoptarPerro(p animales.Perro) {
	p.Comunicarse()
}*/

// AdoptarGato nos permite adoptar un gato
/*func AdoptarGato(g animales.Gato) {
	g.Comunicarse()
}*/

func adoptarMascota(m animales.Mascota) {
	m.Comunicarse()
}
