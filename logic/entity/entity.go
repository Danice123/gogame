package entity

import (
	"github.com/Danice123/idk/display/texturepacker"
	"github.com/Danice123/idk/logic"
	"github.com/faiface/pixel"
)

var tps = 1.0 / 60.0

type Entity interface {
	SpriteSheet() *texturepacker.SpriteSheet
	Sprite() *pixel.Sprite
	GetCoord() logic.Coord
	Translation() *Translation
}

type Base struct {
	Name   string
	Coord  logic.Coord
	Facing logic.Direction
	Frame  int

	translation *Translation

	// Tentative
	Spritesheet *texturepacker.SpriteSheet
}

func (ths *Base) SpriteSheet() *texturepacker.SpriteSheet {
	return ths.Spritesheet
}

func (ths *Base) Sprite() *pixel.Sprite {
	return ths.Spritesheet.Sprites[ths.Name][string(ths.Facing)][ths.Frame]
}

func (ths *Base) GetCoord() logic.Coord {
	return ths.Coord
}

func (ths *Base) Translation() *Translation {
	return ths.translation
}

func (ths *Base) Tick() {
	if ths.translation != nil {
		ths.translation.Completed += 3 * tps

		if int(ths.translation.Completed*100)%25 == 0 {
			if ths.Frame == len(ths.Spritesheet.Sprites[ths.Name][string(ths.Facing)])-1 {
				ths.Frame = 0
			} else {
				ths.Frame++
			}
		}

		if ths.translation.Completed >= 1.0 {
			ths.Coord = ths.Coord.Translate(ths.translation.Direction)
			ths.translation = nil
			ths.Frame = 0
		}
	}
}

func (ths *Base) Walk(dir logic.Direction) {
	if ths.translation == nil {
		ths.Facing = dir
		ths.translation = &Translation{
			Direction: dir,
		}
	}
}
