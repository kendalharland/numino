package numino

import (
	"image/color"
	"log"
	"strconv"

	"github.com/faiface/pixel/pixelgl"
)

// ViewGame runs the numino game.
func ViewGame(win *pixelgl.Window, grid *Grid, done chan GoToCmd) {
	LoadSounds()
	bgMusic := LoopSound(BackgroundMusic)
	defer StopSound(bgMusic)

	game := NewGameState(grid.Rows, grid.Cols)

	// The starting speed of falling blocks. bigger == easier.
	startingTicksPerStep := 120.0
	fallingBlocks := NewFallingBlocks(startingTicksPerStep)

	var ticks float64
	var score float64
	nextSpeedup := startingTicksPerStep + 10000
	scoreRenderer := NewScoreRenderer(
		grid.ColumnToCell(grid.Cols-2),
		grid.RowToCell(0),
	)

	var slamTrailRenderTicks int
	var slamimgbuf *ImageBuffer

	for !win.Closed() {
		ticks++

		imgbuf := NewImageBuffer()

		if win.JustPressed(pixelgl.KeyQ) ||
			win.JustPressed(pixelgl.KeyEscape) {
			done <- GoToMenu
			return
		}

		if win.JustPressed(pixelgl.KeyS) {
			slamimgbuf = NewImageBuffer()
			PlaySound(SlamSound)
			blocksStart := fallingBlocks.Blocks()
			fallingBlocks.Slam(game)
			blocksEnd := fallingBlocks.Blocks()
			for i := range blocksStart {
				col := blocksStart[i].Col
				rowStart := blocksStart[i].Row
				rowEnd := blocksEnd[i].Row
				drawSlamTrail(col, rowStart, rowEnd, grid, slamimgbuf)
				slamTrailRenderTicks = 30
			}
		}

		if win.JustPressed(pixelgl.KeyA) {
			PlaySound(ShiftSound)
			fallingBlocks.ShiftLeft(game)
		}
		if win.JustPressed(pixelgl.KeyD) {
			PlaySound(ShiftSound)
			fallingBlocks.ShiftRight(game)
		}

		// Update sub systems.
		fallingBlocks.Update(ticks, game)
		if nextSpeedup <= ticks {
			fallingBlocks.Speedup()
			nextSpeedup = ticks + 10000
		}

		// Add landed blocks to the grid.
		var didBlockDie, didBlockMerge bool
		for _, block := range fallingBlocks.Blocks() {
			landingType, lrow, lcol := fallingBlocks.DescribeLanding(block, game)
			switch landingType {
			case Unlanded:
				continue
			case LandedOnLiveBlock:
				didBlockMerge = true
			}

			score++
			newBlock := Block{Row: lrow, Col: lcol, Value: block.Value}
			if err := game.AddBlock(newBlock); err != nil {
				log.Fatal(err)
			}
			if game.IsDead(newBlock.Row, newBlock.Col) {
				didBlockDie = true
			}
			fallingBlocks.Remove(block.Row, block.Col)
		}

		if didBlockMerge && didBlockDie {
			PlaySound(DieSound)
		} else if didBlockMerge {
			PlaySound(MergeSound)
		}

		// If all blocks have landed, generate a new wave of blocks.
		if fallingBlocks.Length() == 0 {
			fallingBlocks.Random(game.ColCount())
		}

		if game.IsOver() {
			println("GAME OVER!")
			return
		}

		// Render.
		win.Clear(ColorBg)
		drawGrid(game, grid, imgbuf)
		for _, block := range fallingBlocks.Blocks() {
			drawBlock(block, grid, ColorFallingBlock, imgbuf)
		}
		if slamTrailRenderTicks > 0 {
			slamTrailRenderTicks--
			slamimgbuf.Renderer().Render(win)
		}
		imgbuf.Renderer().Render(win)
		scoreRenderer.SetScore(score)
		scoreRenderer.Render(win)
		win.Update()
	}
}

// ViewMenu runs the main menu.
func ViewMenu(win *pixelgl.Window, grid *Grid, done chan GoToCmd) {
	const optNewGame = "New Game"
	const optCredits = "Credits"
	const optControls = "Controls"
	const optExit = "Exit"

	options := []string{
		optNewGame,
		optControls,
		optCredits,
		optExit,
	}

	selection := 0
	imgbuf := NewImageBuffer()

	for !win.Closed() {
		if win.JustPressed(pixelgl.KeyDown) ||
			win.JustPressed(pixelgl.KeyRight) ||
			win.JustPressed(pixelgl.KeyTab) {
			selection = (selection + 1) % len(options)
		}
		if win.JustPressed(pixelgl.KeyUp) ||
			win.JustPressed(pixelgl.KeyLeft) {
			selection--
			if selection < 0 {
				selection = len(options) - 1
			}
		}
		if win.JustPressed(pixelgl.KeyEnter) ||
			win.JustPressed(pixelgl.KeySpace) {
			switch options[selection] {
			case optNewGame:
				done <- GoToNewGame
				return
			case optExit:
				done <- GoToExit
				return
			case optControls:
				done <- GoToControls
				return
			}
		}

		for i, option := range options {
			var color color.RGBA
			if selection == i {
				color = ColorMenuOption
			} else {
				color = ColorBg
			}
			drawRect(imgbuf, grid.RowToPixel(i+1),
				grid.ColumnToPixel(1), 100, 50, color)
			imgbuf.Text(grid.RowToCell(i+1), grid.ColumnToCell(1),
				option)
		}

		win.Clear(ColorBg)
		imgbuf.Renderer().Render(win)
		win.Update()
	}
}

// ViewControls displays user controls.
func ViewControls(win *pixelgl.Window, grid *Grid, done chan GoToCmd) {
	controls := []struct{ Key, Desc string }{
		{"a", "shift left"},
		{"d", "shift right"},
		{"s", "slam blocks to bottom of screen"},
		{"q, Esc", "exit to main menu"},
	}

	for !win.Closed() {
		if win.JustPressed(pixelgl.KeyQ) ||
			win.JustPressed(pixelgl.KeyEscape) {
			done <- GoToMenu
			return
		}

		imgbuf := NewImageBuffer()

		for i, ctrl := range controls {
			imgbuf.Text(
				grid.RowToCell(i+1),
				grid.ColumnToPixel(1),
				ctrl.Key+": "+ctrl.Desc)
		}

		win.Clear(ColorBg)
		imgbuf.Renderer().Render(win)
		win.Update()
	}
}

func drawBlock(block Block, grid *Grid, color color.RGBA, buf *ImageBuffer) {
	col := grid.ColumnToPixel(block.Col)
	row := grid.RowToPixel(block.Row)
	drawSquare(row, col, grid.SquareSize, color, buf)
	col = grid.ColumnToCell(block.Col)
	row = grid.RowToCell(block.Row)
	buf.Text(row, col, strconv.Itoa(block.Value))
}

func drawGrid(game *GameState, grid *Grid, buf *ImageBuffer) {
	for row := 0; row < game.RowCount(); row++ {
		for col := 0; col < game.ColCount(); col++ {
			if !game.IsEmpty(row, col) {
				color := ColorLiveBlock
				if game.IsDead(row, col) {
					color = ColorDeadBlock
				}

				block := Block{
					Col:   col,
					Row:   row,
					Value: game.ValueAt(row, col),
				}
				drawBlock(block, grid, color, buf)
			}
		}
	}
}

func drawSlamTrail(col int, rowStart int, rowEnd int, grid *Grid, buf *ImageBuffer) {
	for row := rowStart; row < rowEnd; row++ {
		drawSquare(
			grid.RowToPixel(row),
			grid.ColumnToPixel(col),
			grid.SquareSize,
			ColorSlamTrail,
			buf)
	}
}

// CreateSquare returns a square at position x, y whose side length is size.
func drawSquare(y float64, x float64, size float64, color color.RGBA, buf *ImageBuffer) {
	buf.Color(color)
	buf.Vertex(x, y)
	buf.Vertex(x+size, y)
	buf.Vertex(x+size, y+size)
	buf.Vertex(x, y+size)
	buf.Polygon()
}

func drawRect(buf *ImageBuffer, y float64, x float64, w float64, h float64, color color.RGBA) {
	buf.Color(color)
	buf.Vertex(x, y)
	buf.Vertex(x+w, y)
	buf.Vertex(x+w, y+h)
	buf.Vertex(x, y+h)
	buf.Polygon()
}
