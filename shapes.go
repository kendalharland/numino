package numino

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

// CreateSquare returns a square at position x, y whose side length is size.
func CreateSquare(x float64, y float64, size float64, color color.RGBA) *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = color
	imd.Push(pixel.V(x, y))
	imd.Push(pixel.V(x+size, y))
	imd.Push(pixel.V(x+size, y+size))
	imd.Push(pixel.V(x, y+size))
	imd.Polygon(0)
	return imd
}
