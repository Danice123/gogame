package menutils

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/Danice123/idk/display/utils"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

type Text struct {
	Scale float64

	renderer  *text.Text
	timer     *utils.Timer
	content   string
	charIndex int
}

func NewTextContent(fontAtlas *text.Atlas) *Text {
	return &Text{
		renderer: text.New(pixel.ZV, fontAtlas),
	}
}

func (ths *Text) SetText(content string, progressive bool) {
	ths.renderer.Color = color.Black
	ths.content = content

	if progressive {
		ths.charIndex = 0
		ths.timer = &utils.Timer{
			Length: 50,
			Ring: func(reset func()) {
				ths.increment()
				reset()
			},
		}
	} else {
		ths.timer = nil
		fmt.Fprint(ths.renderer, content)
		ths.charIndex = len(content)
	}
}

func (ths *Text) increment() {
	if !ths.IsFinished() {
		fmt.Fprint(ths.renderer, string(ths.content[ths.charIndex]))
		ths.charIndex++
	}
}

func (ths *Text) IsFinished() bool {
	return !(ths.charIndex < len(ths.content))
}

func (ths *Text) Finish() {
	if !ths.IsFinished() {
		fmt.Fprint(ths.renderer, ths.content[ths.charIndex:])
		ths.charIndex = len(ths.content)
	}
}

func (ths *Text) Tick(delta int64) {
	if ths.timer != nil {
		ths.timer.Tick(delta)
	}
}

func (ths *Text) SetMaxWidth(maxWidth float64) {
	if ths.renderer.BoundsOf(ths.content).W()*ths.Scale > maxWidth {
		words := strings.Split(ths.content, " ")

		newContent := ""
		line := ""
		for _, word := range words {
			if ths.renderer.BoundsOf(line+word+" ").W()*ths.Scale > maxWidth {
				newContent += line + "\n"
				line = word + " "
			} else {
				line += word + " "
			}
		}
		newContent += line
		ths.content = newContent
	}
}

func (ths *Text) Width() float64 {
	b := ths.renderer.BoundsOf(ths.content)
	return b.Resized(b.Min, pixel.V(ths.Scale, ths.Scale)).W()
}

func (ths *Text) Height() float64 {
	b := ths.renderer.BoundsOf(ths.content)
	return b.Resized(b.Min, pixel.V(ths.Scale, ths.Scale)).H()
}

func (ths *Text) Render(canvas *pixelgl.Canvas, x float64, y float64) {
	if ths.Scale == 0 {
		ths.Scale = 1
	}

	ths.renderer.Draw(canvas, pixel.IM.Moved(pixel.V(x, y)).Scaled(pixel.V(x, y), ths.Scale))
}
