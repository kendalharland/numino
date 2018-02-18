package numino

import "errors"

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

func (g Grid) ColumnToPixel(n int) (float64, error) {
	if n < 0 || n >= g.Cols {
		return -1, errors.New("invalid column: " + string(n))
	}
	return float64(n) * g.SquareSize, nil
}

func (g Grid) RowToPixel(n int) (float64, error) {
	if n < 0 || n >= g.Cols {
		return -1, errors.New("invalid row: " + string(n))
	}
	return float64(n) * g.SquareSize, nil
}
