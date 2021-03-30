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

	fmt.Fprintf(w, "%s %s %s\n", request.Method, request.URL, request.Proto)

	for k, v := range request.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", request.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", request.RemoteAddr)

	if err := request.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range request.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// counter hace eco del número de llamadas hasta el momento.
func counter(w http.ResponseWriter, request *http.Request) {

	// Uso de las variables a nivel de paquete.
	mu.Lock()

	fmt.Fprintf(w, "Count %d\n", count)

	mu.Unlock()
}
