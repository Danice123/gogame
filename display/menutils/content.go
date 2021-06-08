package menutils

import "github.com/faiface/pixel/pixelgl"

type Content interface {
	Width() int
	Height() int
	Render(canvas pixelgl.Canvas, x, y int)
}

type ContentWithMargin struct {
	TopMargin    int
	BottomMargin int
	LeftMargin   int
	RightMargin  int
	Hidden       bool
	AlignRight   bool
	alignBottom  bool

	Content Content
}

func (ths *ContentWithMargin) Width() int {
	return ths.LeftMargin + ths.Content.Width() + ths.RightMargin
}

func (ths *ContentWithMargin) Height() int {
	return ths.TopMargin + ths.Content.Height() + ths.BottomMargin
}

func (ths *ContentWithMargin) Render(canvas pixelgl.Canvas, x, y int) {
	if !ths.Hidden {
		var rx int
		if ths.AlignRight {
			rx = x - ths.Width() + ths.RightMargin
		} else {
			rx = x + ths.LeftMargin
		}

		var ry int
		if ths.alignBottom {
			ry = y - ths.Height() + ths.BottomMargin
		} else {
			ry = y + ths.TopMargin
		}

		ths.Content.Render(canvas, rx, ry)
	}
}
