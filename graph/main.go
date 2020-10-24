package main

import (
	"flag"
	"image"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var w = 640
var h = 480
var fullscreen bool = true
var scale = 2.0

func lerp(inMin, inMax, outMin, outMax, in float64) float64 {
	outputRange := float64(outMax - outMin)
	inputRange := float64(inMax - inMin)

	return ((outputRange / inputRange) * (in - float64(inMin))) + float64(outMin)
}

func curve(x float64) (y float64) {
	return x * x
}

func frame() *image.RGBA {
	m := image.NewRGBA(image.Rect(0, 0, w, h))

	for x := 0; x < w; x++ {
		xx := lerp(0, float64(w), -1.0, 1.0, float64(x))
		yy := curve(xx)
		y := lerp(-1.0, 1.0, 0, float64(h), yy)
		m.Set(x, int(y), color.White)
	}

	return m
}

func run() {
	cfg := pixelgl.WindowConfig{
		Bounds:      pixel.R(0, 0, float64(w)*scale, float64(h)*scale),
		VSync:       true,
		Undecorated: true,
	}
	if fullscreen {
		cfg.Monitor = pixelgl.PrimaryMonitor()
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	c := win.Bounds().Center()

	for !win.Closed() {
		if win.JustPressed(pixelgl.KeyEscape) || win.JustPressed(pixelgl.KeyQ) {
			return
		}

		win.Clear(color.Black)

		p := pixel.PictureDataFromImage(frame())
		pixel.NewSprite(p, p.Bounds()).
			Draw(win, pixel.IM.Moved(c).Scaled(c, scale))

		win.Update()
	}
}

func main() {
	flag.BoolVar(&fullscreen, "f", fullscreen, "fullscreen")
	flag.IntVar(&w, "w", w, "width")
	flag.IntVar(&h, "h", h, "height")
	flag.Float64Var(&scale, "s", scale, "scale")
	flag.Parse()

	pixelgl.Run(run)
}
