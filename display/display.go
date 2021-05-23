package display

import (
	"fmt"
	"time"

	"github.com/Danice123/idk/display/screen/mapscreen"
	"github.com/Danice123/idk/display/tiledmap"
	"github.com/faiface/pixel/pixelgl"
)

type Display struct {
	window     *pixelgl.Window
	screen     *ScreenHandler
	frameTimer time.Time
}

func (ths *Display) Start() {
	m := tiledmap.OrthoMap{
		Name: "ortho",
	}
	m.Init()

	ths.screen = &ScreenHandler{
		screen: &mapscreen.MapScreen{
			TiledMap: &m,
		},
		renderBehind: false,
	}

	ths.frameTimer = time.Now()
	for !ths.window.Closed() {
		previousFrameTimer := ths.frameTimer
		ths.frameTimer = time.Now()
		ths.render(ths.frameTimer.Sub(previousFrameTimer).Milliseconds())
		ths.window.SetTitle(fmt.Sprintf("Game | Frame Time: %d us", time.Since(ths.frameTimer).Microseconds()))
		ths.window.Update()
	}
}

func (ths *Display) render(delta int64) {
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
