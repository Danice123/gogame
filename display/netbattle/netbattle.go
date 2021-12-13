package netbattle

import (
	"image/color"

	"github.com/Danice123/gogame/display/screen"
	"github.com/Danice123/gogame/display/utils"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type NetBattleScreen struct {
	field *BattleField

	canvas *pixelgl.Canvas

	screen.BaseScreen
}

func NewNetBattleScreen() *NetBattleScreen {
	return &NetBattleScreen{
		field: NewBattleField(),
	}
}

func (ths *NetBattleScreen) ShouldRenderBehind() bool {
	return false
}

func (ths *NetBattleScreen) Tick(delta int64) {
	ths.field.Player.Tick(delta)
}

func (ths *NetBattleScreen) Render(delta int64, window *pixelgl.Window) {
	if ths.canvas == nil { // Or if window size is changed?
		ths.canvas = pixelgl.NewCanvas(pixel.R(0, 0, 240, 160))
	}

	ths.canvas.Clear(color.Black)
	ths.field.Render(ths.canvas)

	scale := window.Bounds().Max.X / 240.0
	camera := pixel.IM.Moved(window.Bounds().Center()).Scaled(window.Bounds().Center(), scale)
	ths.canvas.Draw(window, camera)
}

func (ths *NetBattleScreen) HandleKey(pressed func(utils.KEY) bool) {
	ths.field.Player.HandleKey(pressed)
}
