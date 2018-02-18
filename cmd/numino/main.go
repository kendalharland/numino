package main

import (
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/kharland/numino"
	"golang.org/x/image/colornames"
)

const (
	// The number of rows to start the game with
	numRows = 5
	// The number of cols to start the game with
	numCols = 8
)

func run() {
	game := numino.NewGameState(numRows, numCols)
	random := rand.New(rand.NewSource(42))
	grid := &numino.Grid{
		Cols:       numCols,
		Rows:       numRows,
		SquareSize: 50,
	}
	activeCells := numino.ActiveCells{FramesPerStep: 60}

	cfg := pixelgl.WindowConfig{
		Title:  "Numino",
		Bounds: pixel.R(0, 0, grid.PixelWidth(), grid.PixelHeight()),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	var imgs []*imdraw.IMDraw
	for !win.Closed() {
		// Update sub systems.
		activeCells.Update(10, game)
		landedCells := activeCells.GetLandedCells(game)
		// Do something with landed cells...
		if len(landedCells) == activeCells.Length() {
			activeCells.Clear()
			activeCells.Random(random, game.ColCount())
		}

		imgs = renderBlocks(activeCells.Cells(), grid)
		win.Clear(colornames.Aliceblue)
		for _, img := range imgs {
			img.Draw(win)
		}
		win.Update()
		time.Sleep(time.Second)
	}
}

func main() {
	pixelgl.Run(run)
}

func renderBlocks(blocks []numino.Cell, grid *numino.Grid) []*imdraw.IMDraw {
	var imgs []*imdraw.IMDraw
	for _, block := range blocks {
		col, _ := grid.ColumnToPixel(block.Col)
		row, _ := grid.RowToPixel(block.Row)
		img := numino.CreateSquare(col, row, grid.SquareSize)
		imgs = append(imgs, img)
	}
	return imgs
}
