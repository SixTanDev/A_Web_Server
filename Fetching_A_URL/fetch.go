package main

import (
	"fmt"
	"io/ioutil" // Para utilizar la función ReadAll lee toda la entrada y no línea por línea
	"net/http"  // Para utilizar el método get.
	"os"        // Para utilizar la función exit.
)

func main() {

	for _, url := range os.Args[1:] { // igual que escribir range os.Args[1:len(os.Args)]

		resp, err := http.Get(url) // Obtenemos la URL ingresada en la línea de comandos.

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close() // Los recursos ¡NO CRECEN EN LOS ARBOLES!

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
