package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"strconv"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/kharland/numino"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

const (
	// The number of rows to start the game with
	numRows = 15
	// The number of cols to start the game with
	numCols = 8
)

var (
	atlas = text.NewAtlas(basicfont.Face7x13, text.ASCII)
)

func run() {
	game := numino.NewGameState(numRows, numCols)
	random := rand.New(rand.NewSource(42))
	grid := &numino.Grid{
		Cols:       numCols,
		Rows:       numRows,
		SquareSize: 50,
	}
	fallingBlocks := numino.FallingBlocks{FramesPerStep: 120000}

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

		// Add landed blocks to the grid.
		landedBlocks := fallingBlocks.LandedBlocks(game)
		for _, block := range landedBlocks {
			game.AddBlock(block)
			fallingBlocks.Remove(block.Row, block.Col)
		}
		// If all blocks have landed, generate a new wave of blocks.
		if fallingBlocks.Length() == 0 {
			fallingBlocks.Random(random, game.ColCount())
		}

		if game.IsOver() {
			println("GAME OVER!")
			return
		}

		// Render.
		win.Clear(colornames.Aliceblue)
		renderGameState(game, grid).Render(win)
		renderBlocks(fallingBlocks.Blocks(), grid).Render(win)
		win.Update()
		time.Sleep(time.Second / 2)
	}
}

func main() {
	pixelgl.Run(run)
}

func renderBlocks(blocks []numino.Block, grid *numino.Grid) numino.Renderer {
	var renderers numino.MultiRenderer
	for _, block := range blocks {
		renderers = append(renderers, renderBlock(block, grid, colornames.Red))
	}

	return renderers
}

func renderBlock(block numino.Block, grid *numino.Grid, color color.RGBA) numino.Renderer {
	col := grid.ColumnToPixel(block.Col)
	row := grid.RowToPixel(block.Row)
	img := numino.CreateSquare(col, row, grid.SquareSize, numino.Red)

	col = grid.ColumnToCell(block.Col)
	row = grid.RowToCell(block.Row)
	txt := text.New(pixel.V(col, row), atlas)
	fmt.Fprintf(txt, strconv.Itoa(block.Value))

	return numino.MultiRenderer([]numino.Renderer{
		numino.NewImageRenderer(img),
		numino.NewTextRenderer(txt),
	})
}

func renderGameState(game *numino.GameState, grid *numino.Grid) numino.Renderer {
	var renderers numino.MultiRenderer
	for row := 0; row < game.RowCount(); row++ {
		for col := 0; col < game.ColCount(); col++ {
			if !game.IsEmpty(row, col) {
				var color pixel.RGBA
				if game.IsDead(row, col) {
					color = numino.Blue
				} else {
					color = numino.Green
				}

				img := numino.CreateSquare(
					grid.ColumnToPixel(col),
					grid.RowToPixel(row),
					grid.SquareSize, color)
				renderers = append(renderers, numino.NewImageRenderer(img))
			}
		}
	}

	return renderers
}
