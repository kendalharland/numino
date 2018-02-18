package main

import (
	"fmt"
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

		// Add landed cells to the grid.
		landedCells := activeCells.GetLandedCells(game)
		for _, cell := range landedCells {
			game.AddCell(cell)
			activeCells.Remove(cell.Row, cell.Col)
		}
		if game.IsOver() {
			println("GAME OVER!")
			return
		}

		// Clear the active cells and generate new ones if all have landed.
		if len(landedCells) == activeCells.Length() {
			activeCells.Clear()
			activeCells.Random(random, game.ColCount())
		}

		win.Clear(colornames.Aliceblue)
		// Render game state.
		imgs = renderGameState(game, grid)
		for _, img := range imgs {
			img.Draw(win)
		}

		// Render active cells
		imgs = renderBlocks(activeCells.Cells(), grid)
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
		col := grid.ColumnToPixel(block.Col)
		row := grid.RowToPixel(block.Row)
		img := numino.CreateSquare(col, row, grid.SquareSize, numino.Red)
		imgs = append(imgs, img)
	}
	return imgs
}

func renderGameState(game *numino.GameState, grid *numino.Grid) []*imdraw.IMDraw {
	var imgs []*imdraw.IMDraw
	for row := 0; row < game.RowCount(); row++ {
		for col := 0; col < game.ColCount(); col++ {
			if !game.IsEmpty(row, col) {
				fmt.Println("Drawing: ", row, col)
				var color pixel.RGBA
				if game.IsDead(row, col) {
					color = numino.Blue
				} else {
					color = numino.Green
				}

				imgs = append(imgs, numino.CreateSquare(
					grid.ColumnToPixel(col),
					grid.RowToPixel(row),
					grid.SquareSize, color))
			}
		}
	}
	return imgs
}
