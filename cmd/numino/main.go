package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/kharland/numino"
	"golang.org/x/image/colornames"
)

func run() {
	grid := &numino.Grid{
		Cols:       8,
		Rows:       10,
		SquareSize: 50,
	}

	cfg := pixelgl.WindowConfig{
		Title:  "Numino",
		Bounds: pixel.R(0, 0, grid.PixelWidth(), grid.PixelHeight()),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		col, _ := grid.ColumnToPixel(1)
		row, _ := grid.RowToPixel(2)
		img := numino.CreateSquare(col, row, grid.SquareSize)

		win.Clear(colornames.Aliceblue)
		img.Draw(win)
		win.Update()
		time.Sleep(time.Second)
	}
}

func main() {
	pixelgl.Run(run)
}
