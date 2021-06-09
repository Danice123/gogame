package menutils

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

type Text struct {
	Scale float64

	text *text.Text
}

func NewTextContent(fontAtlas *text.Atlas) *Text {
	return &Text{
		text: text.New(pixel.ZV, fontAtlas),
	}
}

func (ths *Text) SetText(content string) {
	ths.text.Color = color.Black
	fmt.Fprint(ths.text, content)
}

func (ths *Text) Width() float64 {
	return ths.text.Bounds().W() * ths.Scale
}

func (ths *Text) Height() float64 {
	return ths.text.Bounds().H() * ths.Scale
}

func (ths *Text) Render(canvas *pixelgl.Canvas, x float64, y float64) {
	if ths.Scale == 0 {
		ths.Scale = 1
	}

	ths.text.Draw(canvas, pixel.IM.Scaled(pixel.ZV, ths.Scale).Moved(pixel.V(x, y)))
}
