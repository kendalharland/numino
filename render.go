package numino

import (
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

// Drawer is an interface for objects that draw to a window.
type Drawer interface {
	Draw(*pixelgl.Window)
}

// ImagesDrawer draws images to a window.
type ImagesDrawer []*imdraw.IMDraw

// Draw implements the Drawer interface
func (drawer ImagesDrawer) Draw(win *pixelgl.Window) {
	for _, img := range drawer {
		img.Draw(win)
	}
}
