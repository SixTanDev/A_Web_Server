package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"sync"
)

// Variables a nivel de paquete.
var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler) // Cada solicitud llama a hanlder (manejador)
	/* inserción de la función lissajous*/
	http.HandleFunc("/Gif", func(w http.ResponseWriter, r *http.Request) { lissajous(w) })
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

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // Primer color en palette
	blackIndex = 1 // Siguiente color en palette
)

func lissajous(out io.Writer) {

	const (
		cycles  = 5      // número de revoluciones completas del oscilador x
		res     = 0.0001 // Resolusión angular
		size    = 500    // Tamaño de la imagen [-size..+size]
		nframes = 64     // Número de cuadros de animación.
		delay   = 8      // Velocidad entre fotogramas en unidades de 10 ms, Es decir:
		// velocidad con la que se mueve las dos funciones senoidales.
	)

	freq := rand.Float64() * 3.0 // Frecuencia relativa del oscilador y
	// si cambiamos el número creara diferentes imagenes.
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Diferenciia de fase

	for i := 0; i < nframes; i++ {

		rect := image.Rect(0, 0, (2*size)+1, (2*size)+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			img.SetColorIndex(size+int((x*size)+0.5), size+int((y*size)+0.5),
				blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Nota: Ignorar error encoding
}
