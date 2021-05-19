package display

import (
	"github.com/faiface/pixel/pixelgl"
)

type Display struct {
	window *pixelgl.Window
}

func (ths Display) Start() {
	for !ths.window.Closed() {
		ths.window.Update()
	}
}

func New(cfg pixelgl.WindowConfig) Display {
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	return Display{
		window: win,
	}
}
