package screen

import (
	"image/color"

	"github.com/Danice123/gogame/display/utils"
	"github.com/faiface/pixel/pixelgl"
)

type Screen interface {
	Child() Screen
	SetChild(screen Screen)
	ShouldRenderBehind() bool
	Tick(delta int64)
	Render(delta int64, window *pixelgl.Window)
	HandleKey(key utils.KEY)
}

type BaseScreen struct {
	child Screen
}

func (ths *BaseScreen) Child() Screen {
	return ths.child
}

func (ths *BaseScreen) SetChild(screen Screen) {
	ths.child = screen
}

type ScreenHandler struct {
	Screen Screen
}

func (ths *ScreenHandler) Tick(delta int64) {

	if ths.Screen != nil {
		s := ths.Screen
		for {
			s.Tick(delta)
			if s.Child() != nil {
				s = s.Child()
				continue
			}
			break
		}
	}
}

func (ths *ScreenHandler) Render(delta int64, window *pixelgl.Window) {
	if ths.Screen != nil {
		s := ths.Screen
		for {
			if !s.ShouldRenderBehind() {
				window.Clear(color.White)
			}
			s.Render(delta, window)
			if s.Child() != nil {
				s = s.Child()
				continue
			}
			break
		}
	}
}

func (ths *ScreenHandler) Input(key utils.KEY) {
	if ths.Screen != nil {
		s := ths.Screen
		for {
			if s.Child() != nil {
				s = s.Child()
				continue
			}
			s.HandleKey(key)
			break
		}
	}
}
