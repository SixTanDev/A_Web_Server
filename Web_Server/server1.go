package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // Cada solicitud llama a hanlder (manejador)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// manejador hace eco del componente Path de la URL de solicitud r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
