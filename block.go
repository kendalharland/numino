package numino

import "math/rand"

// Block represents an object that occupies a space on the game grid.
type Block struct {
	Col   int
	Row   int
	Value int
}

// FallingBlocks manages the cells that currently under player-control.
type FallingBlocks struct {
	// FramesPerStep is the number of frames that pass between block-movements.
	FramesPerStep int
	// The cells under player control.
	blocks []Block
}

// Blocks returns a copy of the cells in these FallingBlocks.
func (blocks FallingBlocks) Blocks() []Block {
	blocksCopy := make([]Block, len(blocks.blocks))
	copy(blocksCopy, blocks.blocks)
	return blocksCopy
}

// Update updates this FallingBlocks given the elapsed frameCount and gameState.
func (blocks *FallingBlocks) Update(frameCount int, game *GameState) {
	// TODO: Check if frameCount > fpStep
	for i := range blocks.blocks {
		blocks.blocks[i].Row++
	}
}

// Length ...
func (blocks FallingBlocks) Length() int {
	return len(blocks.blocks)
}

// LandedBlocks returns the list of active cells that are landed.
//
// An active block is landed iff:
// 1. It overlaps an inactive block.
// 2. It overlaps a dead block.
// 3. It is in the final row of the grid.
//
// cellState is a two-dimensional grid of CellState for looking up inactive
// cells.
func (blocks FallingBlocks) LandedBlocks(game *GameState) []Block {
	var landed []Block
	for i, block := range blocks.blocks {
		if block.Row >= game.RowCount() ||
			!game.IsEmpty(block.Row, block.Col) ||
			game.IsDead(block.Row, block.Col) {
			landed = append(landed, blocks.blocks[i])
		}
	}
	return landed
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
func (blocks *FallingBlocks) Random(r *rand.Rand, count int) {
	for i := 0; i < count; i++ {
		if (r.Int() % 5) > 3 {
			blocks.Add(0, i, 5)
		}
	}
}
