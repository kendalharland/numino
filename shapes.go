package numino

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

var (
	Red   pixel.RGBA = pixel.RGB(1, 0, 0)
	Green            = pixel.RGB(0, 1, 0)
	Blue             = pixel.RGB(0, 0, 1)
)

// CreateSquare returns a square at position x, y whose side length is size.
func CreateSquare(x float64, y float64, size float64, color pixel.RGBA) *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = color
	imd.Push(pixel.V(x, y))
	imd.Push(pixel.V(x+size, y))
	imd.Push(pixel.V(x+size, y+size))
	imd.Push(pixel.V(x, y+size))
	imd.Polygon(0)
	return imd
}
