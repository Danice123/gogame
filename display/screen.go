package display

import (
	"image/color"

	"github.com/faiface/pixel/pixelgl"
)

type Screen interface {
	ShouldRenderBehind() bool
	Tick(delta int64)
	Render(delta int64, window *pixelgl.Window)
}

type ScreenHandler struct {
	screen Screen
	child  *ScreenHandler
}

func (ths *ScreenHandler) Tick(delta int64) {
	if ths.screen != nil {
		ths.screen.Tick(delta)

		if ths.child != nil {
			ths.child.Tick(delta)
		}
	}
}

func (ths *ScreenHandler) Render(delta int64, window *pixelgl.Window) {
	if ths.screen != nil {
		if !ths.screen.ShouldRenderBehind() {
			window.Clear(color.White)
		}
		if ths.child == nil || ths.child.shouldRenderAbove() {
			ths.screen.Render(delta, window)
		}
	}
}

func (ths *ScreenHandler) shouldRenderAbove() bool {
	if ths.screen == nil {
		return true
	}
	if !ths.screen.ShouldRenderBehind() || (ths.child != nil && !ths.child.shouldRenderAbove()) {
		return false
	}
	return true
}
