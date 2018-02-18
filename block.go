package numino

import (
	"math/rand"
	"time"
)

// Block represents an object that occupies a space on the game grid.
type Block struct {
	Col   int
	Row   int
	Value int
}

// LandingType describes a LandedEvent
type LandingType int

const (
	LandedOnDeadBlock LandingType = iota
	LandedOnLiveBlock
	LandedOnSpace
	Unlanded
)

// FallingBlocks manages the cells that currently under player-control.
type FallingBlocks struct {
	// The cells under player control.
	blocks  []Block
	counter counter
	random  *rand.Rand
}

// NewFallingBlocks returns a pointer to a new FallingBlocks.
func NewFallingBlocks(ticksPerStep float64) *FallingBlocks {
	return &FallingBlocks{
		counter: counter{Ticks: ticksPerStep},
		random:  rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
	}
}

// Blocks returns a copy of the cells in these FallingBlocks.
func (blocks FallingBlocks) Blocks() []Block {
	blocksCopy := make([]Block, len(blocks.blocks))
	copy(blocksCopy, blocks.blocks)
	return blocksCopy
}

// Update updates this FallingBlocks given the current ticks and gameState.
func (blocks *FallingBlocks) Update(ticks float64, game *GameState) {
	if blocks.counter.Update(ticks) {
		for i := range blocks.blocks {
			blocks.blocks[i].Row++
		}
	}
}

// Length ...
func (blocks FallingBlocks) Length() int {
	return len(blocks.blocks)
}

// DescribeLanding returns information about a block's landing.
//
// A block has landed iff:
// 1. It overlaps an inactive block.
// 2. It overlaps a dead block.
// 3. It is in the final row of the grid.
//
// If the block has not landed, the returned landing type is Unlanded and the
// position values should be ignored.
func (blocks FallingBlocks) DescribeLanding(block Block, game *GameState) (LandingType, int, int) {
	switch {
	case block.Row >= game.RowCount():
		return LandedOnSpace, block.Row - 1, block.Col
	case game.IsDead(block.Row, block.Col):
		return LandedOnDeadBlock, block.Row - 1, block.Col
	case !game.IsEmpty(block.Row, block.Col):
		return LandedOnLiveBlock, block.Row, block.Col

	default:
		return Unlanded, -1, -1
	}
}

// Clear clears all cells from this collection.
func (blocks *FallingBlocks) Clear() {
	blocks.blocks = []Block{}
}

// Add adds an ActiveCell to this collection.
func (blocks *FallingBlocks) Add(row int, col int, value int) {
	blocks.blocks = append(blocks.blocks, Block{
		Row:   row,
		Col:   col,
		Value: value,
	})
}

// Remove ...
func (blocks *FallingBlocks) Remove(row int, col int) {
	for i, block := range blocks.blocks {
		if block.Row == row && block.Col == col {
			blocks.blocks = append(blocks.blocks[:i], blocks.blocks[i+1:]...)
			return
		}
	}
}

// Random generates a new set of cells in the first row.
func (blocks *FallingBlocks) Random(count int) {
	for i := 0; i < count; i++ {
		if (blocks.random.Int() % 5) > 3 {
			value := blocks.random.Int()%10 - 3
			if value == 0 {
				value = 1
			}
			blocks.Add(0, i, value)
		}
	}
}

func (blocks *FallingBlocks) Speedup() {
	blocks.counter.Ticks *= .9
}

func (blocks *FallingBlocks) Slam(game *GameState) {
	for i := range blocks.blocks {
		for blocks.blocks[i].Row < game.RowCount() &&
			game.IsEmpty(blocks.blocks[i].Row, blocks.blocks[i].Col) {
			blocks.blocks[i].Row++
		}
	}
}

type direction int

const (
	left  direction = -1
	right           = 1
)

// ShiftLeft shifts these FallingBlocks to the left.
//
// A block can move to its left iff:
// 1. There is no dead block to its left.
// 2. There is no other falling block to its left.
// 3. It is not in the leftmost column.
func (blocks *FallingBlocks) ShiftLeft(game *GameState) {
	shifted := make(map[*Block]bool)
	for i := range blocks.blocks {
		blocks.shift(left, game, shifted, &blocks.blocks[i])
	}
}

// ShiftRight shifts these FallingBlocks to the right.
//
// A block can move to its right iff:
// 1. There is no dead block to its right.
// 2. There is no other falling block to its right.
// 3. It is not in the rightmost column.
func (blocks *FallingBlocks) ShiftRight(game *GameState) {
	shifted := make(map[*Block]bool)
	for i := range blocks.blocks {
		blocks.shift(right, game, shifted, &blocks.blocks[i])
	}
}

func (blocks *FallingBlocks) shift(dir direction, game *GameState, shifted map[*Block]bool, block *Block) {
	if val, ok := shifted[block]; ok || val {
		return
	}
	shifted[block] = true

	// If there's a neighbor blocking the current block's movement, shift that
	// nieghbor first to see if a space is created for this block to shift into.
	if neighb := blocks.neighbor(dir, *block); neighb != nil {
		blocks.shift(dir, game, shifted, neighb)
	}

	if blocks.neighbor(dir, *block) != nil ||
		(dir == left && block.Col <= 0) ||
		(dir == right && block.Col >= game.ColCount()-1) ||
		game.IsDead(block.Row, block.Col+int(dir)) {
		return
	}

	block.Col += int(dir)
}

func (blocks *FallingBlocks) neighbor(dir direction, block Block) *Block {
	for i, b := range blocks.blocks {
		if blocks.blocks[i] != block && b.Row == block.Row && b.Col == block.Col+int(dir) {
			return &blocks.blocks[i]
		}
	}
	return nil
}
