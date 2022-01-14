package mettaur

import (
	"path/filepath"

	"github.com/Danice123/gogame/display/netbattle/field"
	"github.com/Danice123/gogame/display/texturepacker"
	"github.com/Danice123/gogame/display/utils"
	"github.com/faiface/pixel/pixelgl"
)

type MettaurAttack struct {
	field   *field.BattleField
	sprites *texturepacker.SpriteSheet
	coord   utils.Coord
	dead    bool

	animationFrame int
}

func NewMettaurAttack(field *field.BattleField, start utils.Coord) *MettaurAttack {
	ma := &MettaurAttack{
		field:          field,
		sprites:        texturepacker.NewSpriteSheet(filepath.Join("resources", "sheets", "mettaur.json")),
		coord:          start,
		animationFrame: 1,
	}
	field.RegisterObject(ma)
	return ma
}

func (ths *MettaurAttack) Coord() utils.Coord {
	return ths.coord
}

func (ths *MettaurAttack) HighlightTile() bool {
	return !ths.dead
}

func (ths *MettaurAttack) AI(utils.Coord) {}

func (ths *MettaurAttack) Tick() {
	if ths.dead {
		return
	}

	animationLength := ths.sprites.FrameLength("mettaur-attack-effect")
	if ths.animationFrame == animationLength {
		ths.animationFrame = 1
		ths.coord.X--
	} else {
		ths.animationFrame++
	}

	hits := ths.field.HitReg(ths.coord)
	hitSomething := false
	for _, hit := range hits {
		if hit != ths {
			hit.Damage(10, "")
			hitSomething = true
		}
	}
	if hitSomething {
		ths.dead = true
	}
}

func (ths *MettaurAttack) Render(canvas *pixelgl.Canvas, x int, y int) {
	if ths.dead {
		return
	}

	ths.sprites.Clear()
	ths.sprites.DrawFrame("mettaur-attack-effect", ths.animationFrame-1, x, y)
	ths.sprites.Render(canvas)
}

func (ths *MettaurAttack) Damage(int, string) {}
