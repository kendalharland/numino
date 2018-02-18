package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/kharland/numino"
	"golang.org/x/image/colornames"
)

const (
	// The number of rows to start the game with
	numRows = 15
	// The number of cols to start the game with
	numCols = 8
)

func run() {
	grid := &numino.Grid{
		Cols:       numCols,
		Rows:       numRows,
		SquareSize: 50,
	}

	game := numino.NewGameState(numRows, numCols)

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
