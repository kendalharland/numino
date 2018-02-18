package main

import (
	"fmt"
	"math/rand"

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
	game := numino.NewGameState(numRows, numCols)
	random := rand.New(rand.NewSource(42))
	grid := &numino.Grid{
		Cols:       numCols,
		Rows:       numRows,
		SquareSize: 50,
	}
	fallingBlocks := numino.FallingBlocks{FramesPerStep: 10000}

	cfg := pixelgl.WindowConfig{
		Title:  "Numino",
		Bounds: pixel.R(0, 0, grid.PixelWidth(), grid.PixelHeight()),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		// Update sub systems.
		fallingBlocks.Update(1, game)

		// Add landed blocks to the grid.  If all blocks have landed, generate
		// a new wave of blocks.
		landedBlocks := fallingBlocks.LandedBlocks(game)
		for _, block := range landedBlocks {
			game.AddBlock(block)
			fallingBlocks.Remove(block.Row, block.Col)
		}
		if fallingBlocks.Length() == 0 {
			fallingBlocks.Random(random, game.ColCount())
		}

		if game.IsOver() {
			println("GAME OVER!")
			return
		}

		// Render.
		win.Clear(colornames.Aliceblue)
		renderGameState(game, grid).Draw(win)
		renderBlocks(fallingBlocks.Blocks(), grid).Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

func renderBlocks(blocks []numino.Block, grid *numino.Grid) numino.Drawer {
	var imgs []*imdraw.IMDraw
	for _, block := range blocks {
		col := grid.ColumnToPixel(block.Col)
		row := grid.RowToPixel(block.Row)
		img := numino.CreateSquare(col, row, grid.SquareSize, numino.Red)
		imgs = append(imgs, img)
	}
	return numino.ImagesDrawer(imgs)
}

func renderGameState(game *numino.GameState, grid *numino.Grid) numino.Drawer {
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
	return numino.ImagesDrawer(imgs)
}
