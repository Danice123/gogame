package main

import (
	"path/filepath"
	"time"

	"github.com/Danice123/idk/display"
	"github.com/Danice123/idk/display/screen/mapscreen"
	"github.com/Danice123/idk/display/tiledmap"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Game struct {
	display *display.Display
}

func (ths *Game) Start() {
	tmap := tiledmap.NewOrthoMap(filepath.Join("maps", "ortho.tmx"))
	mapscrn := &mapscreen.MapScreen{
		TiledMap: tmap,
	}
	ths.display.ChangeScreen(mapscrn)

	spf := 16666 * time.Microsecond
	go func() {
		for {
			tickTimer := time.Now()

			ths.display.Tick(0)

			sleepTime := spf - time.Since(tickTimer)
			time.Sleep(sleepTime)
		}
	}()
}

func run() {
	game := Game{
		display: display.NewDisplay(pixelgl.WindowConfig{
			Title:  "Some Game",
			Bounds: pixel.R(0, 0, 1024, 768),
			VSync:  true,
		}),
	}

	game.Start()
	game.display.StartRenderLoop()
}

func main() {
	pixelgl.Run(run)
}
