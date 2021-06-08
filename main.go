package main

import (
	"path/filepath"
	"time"

	"github.com/Danice123/idk/display"
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

	entitySprites *texturepacker.SpriteSheet
}

func (ths *Game) LoadMap(mapName string) {
	tmap := tiledmap.NewOrthoMap(filepath.Join("maps", mapName+".tmx"))
	entityData := entity.LoadEntityMapData(mapName)
	entities := entityData.Build(ths.entitySprites)

	mapscrn := mapscreen.New(tmap, entities, ths.player)
	ths.display.ChangeScreen(mapscrn)
}

func (ths *Game) Start() {
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
	ss := texturepacker.NewSpriteSheet(filepath.Join("sheets", "Entity.json"))

	game := Game{
		display: display.NewDisplay(pixelgl.WindowConfig{
			Title:  "Some Game",
			Bounds: pixel.R(0, 0, 1024, 768),
			VSync:  true,
		}),
		entitySprites: ss,
		player:        entity.NewPlayer(ss),
	}

	game.LoadMap("ortho")
	game.Start()
	game.display.StartRenderLoop()
}

func main() {
	pixelgl.Run(run)
}
