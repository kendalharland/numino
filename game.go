package numino

// GameState represents the state
type GameState struct {
	// activeCells are the cells currently under player-control.
	ActiveCells []ActiveCell
	// cells are cells that have been placed on the grid.
	cells [][]int
	// cellState tracks whether a cell is dead or live.
	cellState [][]CellState
}

// ActiveCell is a cell under player-control.
type ActiveCell struct {
	Col   int
	Row   int
	Value int
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

func (gs *GameState) ClearActiveCells() {
	gs.ActiveCells = []ActiveCell{}
}

func (gs *GameState) AddActiveCell(row int, col int, value int) {
	gs.ActiveCells = append(gs.ActiveCells, ActiveCell{
		Row:   row,
		Col:   col,
		Value: value,
	})
}

func (gs *GameState) ShiftActiveCellsDown() bool {
	var cellsMoved bool
	for i, cell := range gs.ActiveCells {
		// Can't shift past the bottom of the grid.
		if cell.Row >= gs.RowCount()-1 {
			continue
		}
		// Can't shift past a dead cell.
		if gs.cellState[cell.Row+1][cell.Col] == DeadCell {
			continue
		}
		gs.ActiveCells[i].Row++
		cellsMoved = true
	}
	return cellsMoved
}
