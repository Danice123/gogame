package netbattle

import (
	"image/color"

	"github.com/Danice123/gogame/display/netbattle/field"
	"github.com/Danice123/gogame/display/netbattle/mettaur"
	"github.com/Danice123/gogame/display/screen"
	"github.com/Danice123/gogame/display/utils"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type NetBattleScreen struct {
	field  *field.BattleField
	Player *Player

	canvas *pixelgl.Canvas
	health *HealthDisplay

	screen.BaseScreen
}

func NewNetBattleScreen() *NetBattleScreen {
	screen := &NetBattleScreen{
		field: field.NewBattleField(),
	}
	screen.Player = NewPlayer(screen.field)
	mettaur.NewMettaur(screen.field)
	screen.health = NewHealthDisplay(&screen.Player.Health)
	return screen
}

func (ths *NetBattleScreen) ShouldRenderBehind() bool {
	return false
}

func (ths *NetBattleScreen) Tick(delta int64) {
	ths.field.AI(ths.Player.Coord())
	ths.field.Tick()
}

func (ths *NetBattleScreen) Render(delta int64, window *pixelgl.Window) {
	if ths.canvas == nil { // Or if window size is changed?
		ths.canvas = pixelgl.NewCanvas(pixel.R(0, 0, 240, 160))
	}

	ths.canvas.Clear(color.Black)
	ths.field.Render(ths.canvas)
	ths.health.Render(ths.canvas, 0, 145)

	scale := window.Bounds().Max.X / 240.0
	camera := pixel.IM.Moved(window.Bounds().Center()).Scaled(window.Bounds().Center(), scale)
	ths.canvas.Draw(window, camera)
}

func (ths *NetBattleScreen) HandleKey(pressed func(utils.KEY) bool) {
	ths.Player.HandleKey(pressed)
}
