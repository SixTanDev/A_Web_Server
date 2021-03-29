// Generar imagenes GIF Lissajoues.

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // Primer color en palette
	blackIndex = 1 // Siguiente color en palette
)

func main() {
	lissajous(os.Stdout)
}

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
