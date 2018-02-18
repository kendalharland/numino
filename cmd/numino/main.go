package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/kharland/numino"
)

// Game configuration options.
const (
	// The number of rows to start the game with
	numRows = 9
	// The number of cols to start the game with
	numCols = 6
)

func run() {
	grid := &numino.Grid{Cols: numCols, Rows: numRows, SquareSize: 50}
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "Numino",
		Bounds: pixel.R(0, 0, grid.PixelWidth(), grid.PixelHeight()),
	})
	if err != nil {
		panic(err)
	}

	// The various numino view.
	menuView := numino.ViewMenu
	gameView := numino.ViewGame
	controlsView := numino.ViewControls

	// The channel used by views to signal that they are done. A view should
	// signal the channel once it is done, with a value specifying the next
	// view to load.
	done := make(chan numino.GoToCmd)
	defer close(done)

	// Start off at the main menu.
	go menuView(win, grid, done)

	for evt := range done {
		switch evt {
		case numino.GoToExit:
			return
		case numino.GoToNewGame:
			go gameView(win, grid, done)
			break
		case numino.GoToMenu:
			go menuView(win, grid, done)
			break
		case numino.GoToControls:
			go controlsView(win, grid, done)
			break
		}
	}
}

func main() {
	pixelgl.Run(run)
}
