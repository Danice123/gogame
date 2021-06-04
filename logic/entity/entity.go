package entity

import (
	"github.com/Danice123/idk/display/texturepacker"
	"github.com/Danice123/idk/logic"
	"github.com/faiface/pixel"
)

type Entity interface {
	GetSpriteSheet() *texturepacker.SpriteSheet
	GetSprite() *pixel.Sprite
	GetCoord() logic.Coord
}

type Base struct {
	Name   string
	Coord  logic.Coord
	Facing logic.Direction
	Frame  int

	tickCount int

	// Tentative
	Spritesheet *texturepacker.SpriteSheet
}

func (ths *Base) GetSpriteSheet() *texturepacker.SpriteSheet {
	return ths.Spritesheet
}

func (ths *Base) GetSprite() *pixel.Sprite {
	return ths.Spritesheet.Sprites[ths.Name][string(ths.Facing)][ths.Frame]
}

func (ths *Base) GetCoord() logic.Coord {
	return ths.Coord
}

func (ths *Base) Tick() {
	ths.tickCount++
}
