package main

import (
	"fmt"
	"log"
	"net/http"
)

func messageHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola este es un handler</h1>")
}
func messageHFCumstom(message string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, message)
		},
	)
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", messageHandlerFunc)
	mux.Handle("/saludar", messageHFCumstom("<h1>Hola amigos</h1>"))
	mux.Handle("/despedirse", messageHFCumstom("<h1>Chao amigos</h1>"))
	log.Println("Ejecutando servidor en http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))
}
