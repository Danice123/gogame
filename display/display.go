package display

import (
	"fmt"
	"time"

	"github.com/faiface/pixel/pixelgl"
)

type Display struct {
	window     *pixelgl.Window
	screen     *ScreenHandler
	frameTimer time.Time
}

func NewDisplay(cfg pixelgl.WindowConfig) *Display {
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	return &Display{
		window: win,
		screen: &ScreenHandler{},
	}
}

func (ths *Display) StartRenderLoop() {
	ths.frameTimer = time.Now()
	for !ths.window.Closed() {
		previousFrameTimer := ths.frameTimer
		ths.frameTimer = time.Now()

		ths.screen.Render(ths.frameTimer.Sub(previousFrameTimer).Milliseconds(), ths.window)

		ths.window.SetTitle(fmt.Sprintf("Game | Frame Time: %d us", time.Since(ths.frameTimer).Microseconds()))
		ths.window.Update()
	}
}

func (ths *Display) Tick(delta int64) {
	ths.screen.Tick(delta)
}

func (ths *Display) ChangeScreen(screen Screen) {
	ths.screen.screen = screen
}
