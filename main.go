package main

import (
	"github.com/Danice123/idk/display"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	game := display.New(pixelgl.WindowConfig{
		Title:  "Some Game",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	})
	game.Start()
}

func main() {
	pixelgl.Run(run)
}
