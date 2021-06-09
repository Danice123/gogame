package menutils

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type BorderBox struct {
	BorderSize float64
	Fill       bool

	Content Content

	draw *imdraw.IMDraw
}

func (ths *BorderBox) Width() float64 {
	return ths.BorderSize*2 + ths.Content.Width()
}

func (ths *BorderBox) Height() float64 {
	return ths.BorderSize*2 + ths.Content.Height()
}

func (ths *BorderBox) Render(canvas *pixelgl.Canvas, x, y float64) {
	if ths.draw == nil {
		ths.draw = imdraw.New(nil)
	}

	if ths.Fill {
		ths.draw.Color = color.White
		ths.draw.Push(pixel.V(x+ths.BorderSize/2, y+ths.BorderSize/2))
		ths.draw.Push(pixel.V(x+ths.Width()-ths.BorderSize/2, y+ths.Height()-ths.BorderSize/2))
		ths.draw.Rectangle(0)
	}

	ths.draw.Color = color.Black
	ths.draw.Push(pixel.V(x+ths.BorderSize/2, y+ths.BorderSize/2))
	ths.draw.Push(pixel.V(x+ths.Width()-ths.BorderSize/2, y+ths.Height()-ths.BorderSize/2))
	ths.draw.Rectangle(ths.BorderSize)
	ths.draw.Draw(canvas)

	ths.Content.Render(canvas, x+ths.BorderSize, y+ths.BorderSize)
}
