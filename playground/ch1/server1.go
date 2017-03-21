package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/lucasb-eyer/go-colorful"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	defaultCycle := 5
	val := r.FormValue("cycles")
	if cycles, err := strconv.Atoi(val); err == nil {
		log.Println("lissajous with cycles=", cycles)
		lissajous(w, cycles)
	} else {
		lissajous(w, defaultCycle)
	}
}

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001
		size    = 100
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
