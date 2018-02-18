package numino

// GameState represents the state
type GameState struct {
	// cells are cells that have been placed on the grid.
	cells [][]int
	// cellState tracks whether a cell is dead or live.
	cellState [][]CellState
}

// CellState determines whether a cell is dead or live.
type CellState bool

const (
	// DeadCell describes a cell that cannot be modified.
	DeadCell CellState = true
	// LiveCell describes a cell that can be modified.
	LiveCell CellState = false

	// The maximum value a cell can hold before it is marked as dead.
	maxLiveValue = 10
)

// NewGameState returns a GameState with the given number of rows and columns.
// All cells are initially alive and empty.
func NewGameState(rows int, cols int) *GameState {
	g := &GameState{
		cells:     make([][]int, rows),
		cellState: make([][]CellState, rows),
	}
	for i := 0; i < rows; i++ {
		g.cells[i] = make([]int, cols)
		g.cellState[i] = make([]CellState, cols)
	}
	return g
}

func (gs GameState) RowCount() int {
	return len(gs.cells)
}

func (gs GameState) ColCount() int {
	return len(gs.cells[0])
}

// IsOver returns true iff this game is over.
//
// This game is over when the top-most row of any column contains a dead cell.
func (gs *GameState) IsOver() bool {
	for i := 0; i < len(gs.cellState[0]); i++ {
		if gs.cellState[0][i] == DeadCell {
			return true
		}
	}
	return false
}

func (gs *GameState) IsEmpty(row int, col int) bool {
	return gs.cells[row][col] == 0
}

func (gs *GameState) IsDead(row int, col int) bool {
	return gs.cellState[row][col] == DeadCell
}

// AddCell adds the given cell to this GameState.
//
// If the cell overlaps a dead cell, it is added to the row above its current
// row. If that row is above the top of the grid, nothing is done and IsDead()
// will return true.
//
// If the cell overlaps a live cell, its value is added to the live cell's
// value. If the new value is outside the allowed bounds, the cell becomes dead.
func (gs *GameState) AddCell(cell Cell) {
	if cell.Row >= gs.RowCount() {
		gs.cells[cell.Row-1][cell.Col] = cell.Value
		return
	}

	if gs.IsDead(cell.Row, cell.Col) {
		if cell.Row == 0 {
			// Game over, Do nothing.
			return
		}
		gs.cells[cell.Row-1][cell.Col] = cell.Value
		return
	}

	// Merge live cell values.
	gs.cells[cell.Row-1][cell.Col] += cell.Value
	if gs.cells[cell.Row-1][cell.Col] > maxLiveValue {
		gs.cellState[cell.Row-1][cell.Col] = DeadCell
	}
}
