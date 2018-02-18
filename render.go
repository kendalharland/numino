package numino

import (
	"fmt"
	"image/color"

	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"

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
	r.text.DrawColorMask(win, pixel.IM, colornames.Black)
	win.SetColorMask(colornames.Aliceblue)
}

type ImageBuffer struct {
	img   *imdraw.IMDraw
	txt   []*text.Text
	atlas *text.Atlas
}

func NewImageBuffer() *ImageBuffer {
	return &ImageBuffer{
		img:   imdraw.New(nil),
		txt:   []*text.Text{},
		atlas: text.NewAtlas(basicfont.Face7x13, text.ASCII),
	}
}

func (buf ImageBuffer) Renderer() Renderer {
	var txtRs []Renderer
	for i := range buf.txt {
		txtRs = append(txtRs, NewTextRenderer(buf.txt[i]))
	}

	rs := []Renderer{NewImageRenderer(buf.img)}
	return MultiRenderer(append(rs, txtRs...))
}

func (buf *ImageBuffer) Color(c color.RGBA) {
	buf.img.Color = c
}

func (buf *ImageBuffer) Vertex(x float64, y float64) {
	buf.img.Push(pixel.V(x, y))
}

func (buf *ImageBuffer) Polygon() {
	buf.img.Polygon(0)
}

func (buf *ImageBuffer) Text(row float64, col float64, msg string) {
	txt := text.New(pixel.V(col, row), buf.atlas)
	fmt.Fprintf(txt, msg)
	buf.txt = append(buf.txt, txt)
}

type ScoreRenderer struct {
	txt *text.Text
}

func NewScoreRenderer(x float64, y float64) *ScoreRenderer {
	return &ScoreRenderer{
		txt: text.New(
			pixel.V(x, y),
			text.NewAtlas(basicfont.Face7x13, text.ASCII),
		),
	}
}

func (s *ScoreRenderer) SetScore(score float64) {
	s.txt.Clear()
	fmt.Fprintf(s.txt, "Score: %v", score)
}

func (s *ScoreRenderer) Render(win *pixelgl.Window) {
	s.txt.DrawColorMask(win, pixel.IM, colornames.Black)
	win.SetColorMask(colornames.Aliceblue)
}
