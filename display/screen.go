package display

import (
	"image/color"

	"github.com/faiface/pixel/pixelgl"
)

type Screen interface {
	Tick(delta int64)
	Render(delta int64, window *pixelgl.Window)
}

type ScreenHandler struct {
	screen       Screen
	renderBehind bool
	child        *ScreenHandler
}

func (ths *ScreenHandler) Tick(delta int64) {
	ths.screen.Tick(delta)
	if ths.child != nil {
		ths.child.Tick(delta)
	}
}

func (ths *ScreenHandler) Render(delta int64, window *pixelgl.Window) {
	if !ths.renderBehind {
		window.Clear(color.White)
	}
	if ths.child == nil || ths.child.shouldRenderAbove() {
		ths.screen.Render(delta, window)
	}
}

func (ths *ScreenHandler) shouldRenderAbove() bool {
	if !ths.renderBehind || (ths.child != nil && !ths.child.shouldRenderAbove()) {
		return false
	}
	return true
}
