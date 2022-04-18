// Exercise 1.6
// Lissajous2 generates GIF animatinos of random lissajous figures
// in multiple colors
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

var palette = []color.Color{
	color.White,
	color.RGBA{0xFF, 0x99, 0xFF, 0xFF}, // pink
	color.RGBA{0x99, 0xCC, 0xFF, 0xFF}, // blue
	color.RGBA{0xFF, 0xFF, 0x99, 0xFF}, // yellow
}

const (
	blackIndex  = 0 // first color in palette
	pinkIndex   = 1 // next color in palette
	blueIndex   = 2 // and so on...
	yellowIndex = 3
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 400   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units.
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			if t < 0.33*cycles*2*math.Pi {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), pinkIndex)
			} else if t < 0.66*cycles*2*math.Pi {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blueIndex)
			} else {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), yellowIndex)
			}

		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
