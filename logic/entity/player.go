package entity

import (
	"github.com/Danice123/idk/display/texturepacker"
	"github.com/Danice123/idk/logic"
)

type Player struct {
	Base
}

func NewPlayer(spritesheet *texturepacker.SpriteSheet) *Player {
	player := &Player{
		Base: Base{
			Name: "Red",
			Coord: logic.Coord{
				X:     4,
				Y:     4,
				Layer: 0,
			},
			Facing:      logic.SOUTH,
			Spritesheet: spritesheet,
		},
	}

	return player
}

func (ths *Player) Walk(dir logic.Direction) {
	ths.Facing = dir
	ths.translation = &Translation{
		Direction: dir,
	}
}
