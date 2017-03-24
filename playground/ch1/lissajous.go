package ch1

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"time"

	"github.com/lucasb-eyer/go-colorful"
)

// Lissajous draw a lissajous gif to out with cycles
func Lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001
		size    = 256
		nframes = 64
		delay   = 8
	)

	rand.Seed(time.Now().UTC().UnixNano())

	colors := colorful.FastHappyPalette(nframes)
	colorPalette := make([]color.Color, 64)
	colorPalette[0] = color.White
	for i := 1; i < nframes; i++ {
		r, g, b := colors[i].RGB255()
		colorPalette[i] = color.RGBA{r, g, b, 0xFF}
	}

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, colorPalette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(i))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
