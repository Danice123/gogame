package entity

import (
	"github.com/Danice123/gogame/display/texturepacker"
	"github.com/Danice123/gogame/logic"
)

type Player struct {
	Locked bool

	Base
}

func NewPlayer(spritesheet *texturepacker.SpriteSheet) *Player {
	player := &Player{
		Base: Base{
			EntityName: "Red",
			Coord: logic.Coord{
				X:     4,
				Y:     4,
				Layer: 1,
			},
			facing:      logic.SOUTH,
			Spritesheet: spritesheet,
		},
	}

	return player
}
