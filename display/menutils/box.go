package menutils

import "github.com/faiface/pixel/pixelgl"

type VerticalBox struct {
	Contents []Content
	Spacing  float64
	AlignTop bool
}

func (ths *VerticalBox) Width() float64 {
	var width float64
	for _, content := range ths.Contents {
		w := content.Width()
		if w > width {
			width = w
		}
	}
	return width
}

func (ths *VerticalBox) Height() float64 {
	var height float64
	for _, content := range ths.Contents {
		height += content.Height()
	}
	height += (float64(len(ths.Contents)) - 1) * ths.Spacing

	return height
}

func (ths *VerticalBox) Render(canvas *pixelgl.Canvas, x float64, y float64) {
	if ths.AlignTop {
		yOffset := y + ths.Height()
		for _, content := range ths.Contents {
			content.Render(canvas, x, yOffset-content.Height())
			yOffset -= content.Height() + ths.Spacing
		}
	} else {
		yOffset := y
		for _, content := range ths.Contents {
			content.Render(canvas, x, yOffset)
			yOffset += content.Height() + ths.Spacing
		}
	}
}
