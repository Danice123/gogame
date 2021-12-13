package main

import (
	"time"

	"github.com/Danice123/gogame/display"
	"github.com/Danice123/gogame/display/netbattle"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Game struct {
	display *display.Display
}

func (ths *Game) Start() {
	spf := 16666 * time.Microsecond // 60fps
	lastFrameTime := time.Now()
	go func() {
		for {
			tickTimer := time.Now()

			ths.display.Tick(time.Since(lastFrameTime).Milliseconds())
			lastFrameTime = time.Now()

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

	nb := netbattle.NewNetBattleScreen()

	game.display.ChangeScreen(nb)
	game.Start()
	game.display.StartRenderLoop()
}

func main() {
	pixelgl.Run(run)
}
