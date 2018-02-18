package numino

// GameState represents the state
type GameState struct {
	cells     [][]int
	cellState [][]CellState
}

// CellState determines whether a cell is dead or live.
type CellState bool

const (
	// DeadCell describes a cell that cannot be modified.
	DeadCell CellState = true
	// LiveCell describes a cell that can be modified.
	LiveCell CellState = false
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

// IsOver returns true iff this game is over.
//
// This game is over when the top-most row of any column contains a dead cell.
func (gs *GameState) IsOver() bool {
	for i := 0; i < len(gs.cells[0]); i++ {
		if gs.cellState[i][0] == DeadCell {
			return true
		}
	}
	return false
}

func (gs *GameState) SetCellValue(row int, col int, value int) {
	gs.cells[row][col] = value
}

func (gs *GameState) SetCellState(row int, col int, state CellState) {
	gs.cellState[row][col] = state
}
