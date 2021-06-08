package menutils

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

type Text struct {
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

func (ths *Text) Width() int {
	return int(ths.text.Bounds().W())
}

func (ths *Text) Height() int {
	return int(ths.text.Bounds().H())
}

func (ths *Text) Render(canvas *pixelgl.Canvas, x int, y int) {
	ths.text.Draw(canvas, pixel.IM.Scaled(pixel.ZV, 10).Moved(pixel.V(float64(x), float64(y))))
}
