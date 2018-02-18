package numino

// Grid is a model for computing the boundaries and locations of game entities.
type Grid struct {
	Cols       int
	Rows       int
	SquareSize float64
}

func (g Grid) PixelWidth() float64 {
	return float64(g.Cols) * g.SquareSize
}

func (g Grid) PixelHeight() float64 {
	return float64(g.Rows) * g.SquareSize
}

func (g Grid) ColumnToPixel(n int) float64 {
	return float64(n) * g.SquareSize
}

func (g Grid) RowToPixel(n int) float64 {
	return float64(g.Rows-1-n) * g.SquareSize
}
