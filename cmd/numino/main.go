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
	numRows = 15
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

	// game := numino.NewGameState(numRows, numCols)

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
		imgs = randomBlocks(random, grid)
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

func randomBlocks(r *rand.Rand, grid *numino.Grid) []*imdraw.IMDraw {
	var imgs []*imdraw.IMDraw
	for i, b := range randomBools(r, grid.Cols) {
		if b {
			col, _ := grid.ColumnToPixel(i)
			row, _ := grid.RowToPixel(0)
			img := numino.CreateSquare(col, row, grid.SquareSize)
			imgs = append(imgs, img)
		}
	}
	return imgs
}

func randomBools(r *rand.Rand, count int) []bool {
	var states []bool
	for i := 0; i < count; i++ {
		states = append(states, (r.Int()%5) > 3)
	}
	return states
}
