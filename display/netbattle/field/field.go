package field

import (
	"path/filepath"

	"github.com/Danice123/gogame/display/texturepacker"
	"github.com/Danice123/gogame/display/utils"
	"github.com/faiface/pixel/pixelgl"
)

type BattleFieldObject interface {
	Coord() utils.Coord
	HighlightTile() bool
	AI(utils.Coord)
	Tick()
	Render(canvas *pixelgl.Canvas, x int, y int)
	Damage(int, string)
}

type BattleField struct {
	tileSprites *texturepacker.SpriteSheet

	objects []BattleFieldObject
}

func NewBattleField() *BattleField {
	return &BattleField{
		tileSprites: texturepacker.NewSpriteSheet(filepath.Join("resources", "sheets", "battle_tiles.json")),
		objects:     []BattleFieldObject{},
	}
}

func (ths *BattleField) RegisterObject(object BattleFieldObject) {
	ths.objects = append(ths.objects, object)
}

func (ths *BattleField) AI(playerCoord utils.Coord) {
	for _, object := range ths.objects {
		object.AI(playerCoord)
	}
}

func (ths *BattleField) Tick() {
	for _, object := range ths.objects {
		object.Tick()
	}
}

func (ths *BattleField) Render(canvas *pixelgl.Canvas) {
	x := 0
	y := 64

	highlighted := map[utils.Coord]bool{}
	for _, object := range ths.objects {
		if object.HighlightTile() {
			highlighted[object.Coord()] = true
		}
	}

	for col := 0; col < 6; col++ {
		side := "red"
		if col > 2 {
			side = "blue"
		}

		if highlighted[utils.Coord{X: col, Y: 0}] {
			ths.tileSprites.Draw("highlighted", x+40*col, y)
		} else {
			ths.tileSprites.Draw(side+"-top", x+40*col, y)
		}
		if highlighted[utils.Coord{X: col, Y: 1}] {
			ths.tileSprites.Draw("highlighted", x+40*col, y-24)
		} else {
			ths.tileSprites.Draw(side+"-mid", x+40*col, y-24)
		}
		if highlighted[utils.Coord{X: col, Y: 2}] {
			ths.tileSprites.Draw("highlighted", x+40*col, y-48)
		} else {
			ths.tileSprites.Draw(side+"-bot", x+40*col, y-48)
		}
		ths.tileSprites.Draw("edge", x+40*col, y-72)
	}
	ths.tileSprites.Render(canvas)

	for _, object := range ths.objects {
		object.Render(canvas, 40*object.Coord().X, 64-24*object.Coord().Y)
	}
}

type Registered interface {
	Damage(int, string)
}

func (ths *BattleField) HitReg(loc utils.Coord) []Registered {
	hits := []Registered{}
	for _, object := range ths.objects {
		if object.Coord() == loc {
			hits = append(hits, object)
		}
	}
	return hits
}
