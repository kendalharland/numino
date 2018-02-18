package numino

// GameState represents the state
type GameState struct {
	// blocks are blocks that have been placed on the grid.
	blocks [][]int
	// blockState tracks whether a block is dead or live.
	blockState [][]BlockState
}

// BlockState determines whether a block is dead or live.
type BlockState bool

const (
	// DeadBlock describes a block that cannot be modified.
	DeadBlock BlockState = true
	// LiveBlock describes a block that can be modified.
	LiveBlock BlockState = false

	// The maximum value a block can hold before it is marked as dead.
	maxLiveValue = 10
)

// NewGameState returns a GameState with the given number of rows and columns.
// All blocks are initially alive and empty.
func NewGameState(rows int, cols int) *GameState {
	g := &GameState{
		blocks:     make([][]int, rows),
		blockState: make([][]BlockState, rows),
	}
	for i := 0; i < rows; i++ {
		g.blocks[i] = make([]int, cols)
		g.blockState[i] = make([]BlockState, cols)
	}
	return g
}

func (gs GameState) RowCount() int {
	return len(gs.blocks)
}

func (gs GameState) ColCount() int {
	return len(gs.blocks[0])
}

// IsOver returns true iff this game is over.
//
// This game is over when the top-most row of any column contains a dead block.
func (gs *GameState) IsOver() bool {
	for i := 0; i < len(gs.blockState[0]); i++ {
		if gs.blockState[0][i] == DeadBlock {
			return true
		}
	}
	return false
}

func (gs *GameState) IsEmpty(row int, col int) bool {
	return gs.blocks[row][col] == 0
}

func (gs *GameState) IsDead(row int, col int) bool {
	return gs.blockState[row][col] == DeadBlock
}

// AddBlock adds the given block to this GameState.
//
// If the block overlaps a dead block, it is added to the row above its current
// row. If that row is above the top of the grid, nothing is done and IsDead()
// will return true.
//
// If the block overlaps a live block, its value is added to the live block's
// value. If the new value is outside the allowed bounds, the block becomes dead.
func (gs *GameState) AddBlock(block Block) {
	if block.Row >= gs.RowCount() {
		gs.blocks[block.Row-1][block.Col] = block.Value
		return
	}

	if gs.IsDead(block.Row, block.Col) {
		if block.Row == 0 {
			// Game over, Do nothing.
			return
		}
		gs.blocks[block.Row-1][block.Col] = block.Value
		return
	}

	// Merge live block values.
	gs.blocks[block.Row-1][block.Col] += block.Value
	if gs.blocks[block.Row-1][block.Col] > maxLiveValue {
		gs.blockState[block.Row-1][block.Col] = DeadBlock
	}
}
