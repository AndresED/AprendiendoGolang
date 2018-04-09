package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)
	bodegaOrigen := []string{"php", "javascript", "html", "css", "Java", "c++", "Base de datos", "git"}
	bodegaDestino := []string{}
	//CANALES PERMITEN COMUNICARSE ENTRE SI A LAS GORUTINAS
	fotocopiadora := make(chan string)
	fotocopiado := make(chan string)
	go func() {
		for _, libro := range bodegaOrigen {
			fotocopiadora <- libro
		}
	}()

	//Se carga la fotocopiadora
	go func() {
		for {
			//SE ENTREGA EL CONTENIDO DE LA FOTOCOPIADORA A LA VARIABLE LIBRO
			libro := <-fotocopiadora
			fotocopiado <- libro
			//AGREGAR EL LIBRO A LA BODEGA DESTINO
			bodegaDestino = append(bodegaDestino, libro)
			if len(bodegaOrigen) == len(bodegaDestino) {
				//CERRAR CANAL FOTOCOPIADO
				close(fotocopiado)
			}
		}
	}()
	fmt.Println(" ** Listado de libros fotocopiados **")
	for {
		if libro, ok := <-fotocopiado; ok {
			fmt.Println(libro)
		} else {
			fmt.Println("Se cerro el canal")
			break
		}
	}
}
