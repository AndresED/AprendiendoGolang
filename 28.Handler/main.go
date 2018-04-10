package main

import (
	"fmt"
	"log"
	"net/http"
)

type messageHandler struct {
	message string
}

func (m messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}

func main() {
	// ES UNA INTERFAZ QUE CONTIENE UN METODO QUE NOS PERMITE RECIBIR COMO PARAMETRO UN RESPONSE Y UN REQUEST
	//NOS PERMITEN CREAR LOS MANEJADORES PARA LAS RUTAS QUE DISEÑEMOS PARA NUESTRA APLICACION
	mux := http.NewServeMux()
	m1 := &messageHandler{message: "<h1>Hola desde un handler</h1>"}
	m2 := &messageHandler{message: "<h1>Perros</h1>"}
	m3 := &messageHandler{message: "<h1>Gatos</h1>"}
	mux.Handle("/saludar", m1)
	mux.Handle("/perros", m2)
	mux.Handle("/gatos", m3)
	log.Println("Ejecutando servidor en http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))
}
