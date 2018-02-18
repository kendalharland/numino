package numino

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

// Renderer is an interface for objects that draw to a window.
type Renderer interface {
	Render(*pixelgl.Window)
}

// MultiRenderer executes multiple Renderers at once.
type MultiRenderer []Renderer

// Render implements the Renderer interface
func (drawer MultiRenderer) Render(win *pixelgl.Window) {
	for _, d := range drawer {
		d.Render(win)
	}
}

// ImageRenderer draws images to a window.
type ImageRenderer struct {
	img *imdraw.IMDraw
}

func NewImageRenderer(img *imdraw.IMDraw) *ImageRenderer {
	return &ImageRenderer{img}
}

// Render implements the Renderer interface
func (r ImageRenderer) Render(win *pixelgl.Window) {
	r.img.Draw(win)
}

// TextRenderer draws text to a window.
type TextRenderer struct {
	text *text.Text
}

func NewTextRenderer(text *text.Text) *TextRenderer {
	return &TextRenderer{text}
}

// Render implements the Renderer interface
func (r TextRenderer) Render(win *pixelgl.Window) {
	r.text.Draw(win, pixel.IM)
}
