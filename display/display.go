package display

import (
	"time"

	"github.com/faiface/pixel/pixelgl"
)

type Display struct {
	window     *pixelgl.Window
	screen     *ScreenHandler
	frameTimer time.Time
}

func (ths Display) Start() {
	ths.frameTimer = time.Now()
	for !ths.window.Closed() {
		previousFrameTimer := ths.frameTimer
		ths.frameTimer = time.Now()
		ths.render(ths.frameTimer.Sub(previousFrameTimer).Milliseconds())
		ths.window.Update()
	}
}

func (ths Display) render(delta int64) {
	if ths.screen != nil {
		ths.screen.Tick(delta)
		ths.screen.Render(delta, ths.window)
	}
}

func NewDisplay(cfg pixelgl.WindowConfig) Display {
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	return Display{
		window: win,
	}
}
