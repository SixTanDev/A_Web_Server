package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Variables a nivel de paquete.
var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler) // Cada solicitud llama a hanlder (manejador)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// hace eco del componente Path de la URL solicitada
func handler(w http.ResponseWriter, request *http.Request) {

	// uso de las variables a nivel de paquete.
	mu.Lock()
	count++
	mu.Unlock()

	fmt.Fprintf(w, "URL.Path = %q\n", request.URL.Path)
}

// counter hace eco del número de llamadas hasta el momento.
func counter(w http.ResponseWriter, request *http.Request) {

	// Uso de las variables a nivel de paquete.
	mu.Lock()

	fmt.Fprintf(w, "Count %d\n", count)

	mu.Unlock()
}
