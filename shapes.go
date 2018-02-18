package numino

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

// createSquare returns a square image at position x, y whose side length
// is size.
func CreateSquare(x float64, y float64, size float64) *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(1, 0, 0)
	imd.Push(pixel.V(x, y))
	imd.Push(pixel.V(x+size, y))
	imd.Push(pixel.V(x+size, y+size))
	imd.Push(pixel.V(x, y+size))
	imd.Polygon(0)
	return imd
}
