package netbattle

import (
	"path/filepath"

	"github.com/Danice123/gogame/display/texturepacker"
	"github.com/faiface/pixel/pixelgl"
)

type BattleField struct {
	tileSprites *texturepacker.SpriteSheet
}

func NewBattleField() *BattleField {
	return &BattleField{
		tileSprites: texturepacker.NewSpriteSheet(filepath.Join("resources", "sheets", "battle_tiles.json")),
	}
}

func (ths *BattleField) Render(canvas *pixelgl.Canvas) {
	x := 0
	y := 64

	for col := 0; col < 6; col++ {
		side := "red"
		if col > 2 {
			side = "blue"
		}

		ths.tileSprites.Draw(side+"-top", x+40*col, y)
		ths.tileSprites.Draw(side+"-mid", x+40*col, y-24)
		ths.tileSprites.Draw(side+"-bot", x+40*col, y-48)
		ths.tileSprites.Draw("edge", x+40*col, y-72)
	}
	ths.tileSprites.Render(canvas)
}
