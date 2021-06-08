package entity

import (
	"github.com/Danice123/idk/pkg/display/texturepacker"
	"github.com/Danice123/idk/pkg/logic"
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
				Layer: 1,
			},
			Facing:      logic.SOUTH,
			Spritesheet: spritesheet,
		},
	}

	return player
}
