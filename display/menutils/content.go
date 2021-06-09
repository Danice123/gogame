package menutils

import "github.com/faiface/pixel/pixelgl"

type Content interface {
	Width() float64
	Height() float64
	Render(canvas *pixelgl.Canvas, x, y float64)
}

type ContentWithMargin struct {
	TopMargin    float64
	BottomMargin float64
	LeftMargin   float64
	RightMargin  float64

	MinWidth  float64
	MinHeight float64

	AlignRight bool
	AlignTop   bool

	Content Content
}

func (ths *ContentWithMargin) Width() float64 {
	width := ths.LeftMargin + ths.Content.Width() + ths.RightMargin

	if ths.MinWidth > width {
		return ths.MinWidth
	} else {
		return width
	}
}

func (ths *ContentWithMargin) Height() float64 {
	height := ths.TopMargin + ths.Content.Height() + ths.BottomMargin

	if ths.MinHeight > height {
		return ths.MinHeight
	} else {
		return height
	}
}

func (ths *ContentWithMargin) Render(canvas *pixelgl.Canvas, x, y float64) {
	var rx float64
	if ths.AlignRight {
		rx = x - ths.Width() + ths.RightMargin
	} else {
		rx = x + ths.LeftMargin
	}

	var ry float64
	if ths.AlignTop {
		ry = y + ths.Height() - ths.TopMargin - ths.Content.Height()
	} else {
		ry = y + ths.BottomMargin
	}

	ths.Content.Render(canvas, rx, ry)
}
