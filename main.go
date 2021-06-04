package main

import (
	"path/filepath"
	"time"

	"github.com/Danice123/idk/display"
	displayEntity "github.com/Danice123/idk/display/entity"
	"github.com/Danice123/idk/display/screen/mapscreen"
	"github.com/Danice123/idk/display/texturepacker"
	"github.com/Danice123/idk/display/tiledmap"
	"github.com/Danice123/idk/logic/entity"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Game struct {
	display *display.Display
	player  *entity.Player
}

func (ths *Game) Start() {
	ss := texturepacker.NewSpriteSheet(filepath.Join("sheets", "Entity.json"))
	ths.player = entity.NewPlayer(ss)

	tmap := tiledmap.NewOrthoMap(filepath.Join("maps", "ortho.tmx"))
	mapscrn := &mapscreen.MapScreen{
		TiledMap:      tmap,
		EntityHandler: &displayEntity.EntityHandler{},
		Player:        ths.player,
	}
	mapscrn.EntityHandler.AddEntity(ths.player)
	ths.display.ChangeScreen(mapscrn)

	spf := 16666 * time.Microsecond
	go func() {
		for {
			tickTimer := time.Now()

			ths.display.Tick(0) // Update time?

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
