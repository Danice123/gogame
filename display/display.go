package display

import (
	"fmt"
	"time"

	"github.com/Danice123/gogame/display/screen"
	"github.com/Danice123/gogame/display/utils"
	"github.com/faiface/pixel/pixelgl"
)

type Display struct {
	window     *pixelgl.Window
	screen     *screen.ScreenHandler
	frameTimer time.Time

	lastActivation time.Time
}

func NewDisplay(cfg pixelgl.WindowConfig) *Display {
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	return &Display{
		window: win,
		screen: &screen.ScreenHandler{},
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

func (ths *Display) ChangeScreen(screen screen.Screen) {
	ths.screen.Screen = screen
}

func (ths *Display) Tick(delta int64) {
	ths.screen.Tick(delta)

	if time.Since(ths.lastActivation) > 500*time.Millisecond {
		if ths.window.Pressed(pixelgl.KeyZ) {
			ths.screen.Input(utils.ACTIVATE)
			ths.lastActivation = time.Now()
		} else if ths.window.Pressed(pixelgl.KeyX) {
			ths.screen.Input(utils.DECLINE)
			ths.lastActivation = time.Now()
		}
	}

	if ths.window.Pressed(pixelgl.KeyUp) {
		ths.screen.Input(utils.UP)
	} else if ths.window.Pressed(pixelgl.KeyDown) {
		ths.screen.Input(utils.DOWN)
	} else if ths.window.Pressed(pixelgl.KeyLeft) {
		ths.screen.Input(utils.LEFT)
	} else if ths.window.Pressed(pixelgl.KeyRight) {
		ths.screen.Input(utils.RIGHT)
	}
}
