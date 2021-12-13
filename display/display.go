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

	debounce map[utils.KEY]bool

	lastActivation time.Time
}

func NewDisplay(cfg pixelgl.WindowConfig) *Display {
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	return &Display{
		window:   win,
		screen:   &screen.ScreenHandler{},
		debounce: make(map[utils.KEY]bool),
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
	var pressedFunc func(key utils.KEY) bool
	if ths.window.JoystickPresent(pixelgl.Joystick1) {
		pressedFunc = func(key utils.KEY) bool {
			var keyMap pixelgl.GamepadButton
			switch key {
			case utils.ACTIVATE:
				keyMap = pixelgl.ButtonB
			case utils.DECLINE:
				keyMap = pixelgl.ButtonA
			case utils.UP:
				keyMap = pixelgl.ButtonDpadUp
			case utils.DOWN:
				keyMap = pixelgl.ButtonDpadDown
			case utils.LEFT:
				keyMap = pixelgl.ButtonDpadLeft
			case utils.RIGHT:
				keyMap = pixelgl.ButtonDpadRight
			}
			return ths.window.JoystickPressed(pixelgl.Joystick1, keyMap)
		}
	} else {
		pressedFunc = func(key utils.KEY) bool {
			var keyMap pixelgl.Button
			switch key {
			case utils.ACTIVATE:
				keyMap = pixelgl.KeyZ
			case utils.DECLINE:
				keyMap = pixelgl.KeyX
			case utils.UP:
				keyMap = pixelgl.KeyUp
			case utils.DOWN:
				keyMap = pixelgl.KeyDown
			case utils.LEFT:
				keyMap = pixelgl.KeyLeft
			case utils.RIGHT:
				keyMap = pixelgl.KeyRight
			}
			return ths.window.Pressed(keyMap)
		}
	}
	ths.screen.Input(pressedFunc)
	ths.screen.Tick(delta)
}
