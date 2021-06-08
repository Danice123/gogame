package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"runtime/pprof"
	"time"

	"github.com/Danice123/idk/pkg/display"
	displayEntity "github.com/Danice123/idk/pkg/display/entity"
	"github.com/Danice123/idk/pkg/display/screen/mapscreen"
	"github.com/Danice123/idk/pkg/display/texturepacker"
	"github.com/Danice123/idk/pkg/display/tiledmap"
	"github.com/Danice123/idk/pkg/logic/entity"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")

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
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
	pixelgl.Run(run)
}
