package menutils

import "github.com/faiface/pixel/pixelgl"

type VerticalBox struct {
	Contents []Content
	Spacing  int
}

func (ths *VerticalBox) Width() int {
	width := 0
	for _, content := range ths.Contents {
		w := content.Width()
		if w > width {
			width = w
		}
	}
	return width
}

func (ths *VerticalBox) Height() int {
	height := 0
	for _, content := range ths.Contents {
		height += content.Height()
	}
	height += (len(ths.Contents) - 1) * ths.Spacing
	return height
}

func (ths *VerticalBox) Render(canvas pixelgl.Canvas, x int, y int) {
	yOffset := y
	for _, content := range ths.Contents {
		content.Render(canvas, x, yOffset)
		yOffset += content.Height() + ths.Spacing
	}
}
