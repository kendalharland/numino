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
	random := rand.New(rand.NewSource(42))
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

	var imgs []*imdraw.IMDraw
	generate := true
	for !win.Closed() {
		if generate {
			generate = false
			randomActiveCells(random, grid, game)
		} else {
			generate = !game.ShiftActiveCellsDown()
		}
		imgs = renderBlocks(game.ActiveCells, grid)

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

func randomActiveCells(r *rand.Rand, grid *numino.Grid, game *numino.GameState) {
	for i, b := range randomBools(r, grid.Cols) {
		if b {
			game.AddActiveCell(0, i, 5)
		}
	}
}

func randomBools(r *rand.Rand, count int) []bool {
	var states []bool
	for i := 0; i < count; i++ {
		states = append(states, (r.Int()%5) > 3)
	}
	return states
}

func renderBlocks(blocks []numino.ActiveCell, grid *numino.Grid) []*imdraw.IMDraw {
	var imgs []*imdraw.IMDraw
	for _, block := range blocks {
		col, _ := grid.ColumnToPixel(block.Col)
		row, _ := grid.RowToPixel(block.Row)
		img := numino.CreateSquare(col, row, grid.SquareSize)
		imgs = append(imgs, img)
	}
	return imgs
}
