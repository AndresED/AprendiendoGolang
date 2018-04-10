package main

import (
	"html/template"
	"log"
	"net/http"
)

// Persona almacena los datos de una persona
type Persona struct {
	Nombre string
	Edad   uint8
}

func renderizar(w http.ResponseWriter, r *http.Request) {
	p := &Persona{"Andres", 26}
	t, err := template.ParseFiles("./views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, p)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", renderizar)
	log.Println("Ejecutando servidor en http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))
	log.Println("Ejecución terminada")
}
