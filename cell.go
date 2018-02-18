package numino

import "math/rand"

// Cell represents a space on the game grid.
type Cell struct {
	Col   int
	Row   int
	Value int
}

// ActiveCells manages the cells that currently under player-control.
type ActiveCells struct {
	// FramesPerStep is the number of frames that pass between cell-movements.
	FramesPerStep int
	// The cells under player control.
	cells []Cell
}

// Cells returns a copy of the cells in these ActiveCells.
func (ac ActiveCells) Cells() []Cell {
	cellsCopy := make([]Cell, len(ac.cells))
	copy(cellsCopy, ac.cells)
	return cellsCopy
}

// Update updates this ActiveCells given the elapsed frameCount and gameState.
func (ac *ActiveCells) Update(frameCount int, game *GameState) {
	// TODO: Check if frameCount > fpStep
	for i := range ac.cells {
		ac.cells[i].Row++
	}
}

func (ac ActiveCells) Length() int {
	return len(ac.cells)
}

// GetLandedCells returns the list of active cells that are landed.
//
// An active cell is landed iff:
// 1. It overlaps an inactive cell.
// 2. It overlaps a dead cell.
// 3. It is in the final row of the grid.
//
// cellState is a two-dimensional grid of CellState for looking up inactive
// cells.
func (ac ActiveCells) GetLandedCells(game *GameState) []Cell {
	var landed []Cell
	for i, cell := range ac.cells {
		if cell.Row >= game.RowCount() ||
			!game.IsEmpty(cell.Row, cell.Col) ||
			game.IsDead(cell.Row, cell.Col) {
			landed = append(landed, ac.cells[i])
		}
	}
	return landed
}

// Clear clears all cells from this collection.
func (ac *ActiveCells) Clear() {
	ac.cells = []Cell{}
}

// Add adds an ActiveCell to this collection.
func (ac *ActiveCells) Add(row int, col int, value int) {
	ac.cells = append(ac.cells, Cell{
		Row:   row,
		Col:   col,
		Value: value,
	})
}

// Random generates a new set of cells in the first row.
func (ac *ActiveCells) Random(r *rand.Rand, count int) {
	for i := 0; i < count; i++ {
		if (r.Int() % 5) > 3 {
			ac.Add(0, i, 5)
		}
	}
}
