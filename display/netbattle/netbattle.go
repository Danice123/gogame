package netbattle

import (
	"image/color"

	"github.com/Danice123/gogame/display/netbattle/mettaur"
	"github.com/Danice123/gogame/display/netbattle/state"
	"github.com/Danice123/gogame/display/screen"
	"github.com/Danice123/gogame/display/utils"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type NetBattleScreen struct {
	field   *BattleField
	Player  *Player
	Mettaur *mettaur.Mettaur

	canvas *pixelgl.Canvas

	screen.BaseScreen
}

func NewNetBattleScreen() *NetBattleScreen {
	screen := &NetBattleScreen{
		field:   NewBattleField(),
		Mettaur: mettaur.NewMettaur(),
	}
	screen.Player = NewPlayer(screen.HitReg)
	return screen
}

func (ths *NetBattleScreen) ShouldRenderBehind() bool {
	return false
}

func (ths *NetBattleScreen) Tick(delta int64) {
	ths.Mettaur.AI(state.BoardState{
		PlayerCoord: ths.Player.Coord,
	})

	ths.Player.Tick(delta)
	ths.Mettaur.Tick(delta)
}

func (ths *NetBattleScreen) Render(delta int64, window *pixelgl.Window) {
	if ths.canvas == nil { // Or if window size is changed?
		ths.canvas = pixelgl.NewCanvas(pixel.R(0, 0, 240, 160))
	}

	ths.canvas.Clear(color.Black)
	ths.field.Render(ths.canvas)
	ths.Player.Render(ths.canvas, 40*ths.Player.Coord.X+3, 64-24*ths.Player.Coord.Y+5)
	ths.Mettaur.Render(ths.canvas, 40*ths.Mettaur.Coord.X+9, 64-24*ths.Mettaur.Coord.Y+5)

	scale := window.Bounds().Max.X / 240.0
	camera := pixel.IM.Moved(window.Bounds().Center()).Scaled(window.Bounds().Center(), scale)
	ths.canvas.Draw(window, camera)
}

func (ths *NetBattleScreen) HandleKey(pressed func(utils.KEY) bool) {
	ths.Player.HandleKey(pressed)
}

func (ths *NetBattleScreen) HitReg(loc utils.Coord) *mettaur.Mettaur {
	if ths.Mettaur.Coord == loc {
		return ths.Mettaur
	}
	return nil
}
