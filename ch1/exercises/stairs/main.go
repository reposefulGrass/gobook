// Stairs generates GIF animations of a set of stairs being generated.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"os"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	stairs(os.Stdout)
}

func stairs(out io.Writer) {
	const (
		dir_up     = 0
		dir_right  = 1
		dir_down   = 2
		dir_left   = 3
		stair_size = 20
		size       = 400 // image canvas covers [-size..+size]
		nframes    = 64  // number of animation frames
		delay      = 8   // delay between frames in 10ms units.
	)
	anim := gif.GIF{LoopCount: nframes}

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		x := stair_size
		y := 2 * size
		cur_dir := dir_up

		for j := 0; j < i; j++ {
			for k := 0; k < stair_size; k++ {
				if cur_dir == dir_up {
					y -= 1
				} else if cur_dir == dir_right {
					x += 1
				}
				img.SetColorIndex(x, y, blackIndex)
			}

			if cur_dir == dir_up {
				cur_dir = dir_right
			} else if cur_dir == dir_right {
				cur_dir = dir_up
			}
		}

		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
