package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/kharland/numino"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Numino",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	img := numino.CreateSquare(50, 50, 30)
	for !win.Closed() {
		win.Clear(colornames.Aliceblue)
		img.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
