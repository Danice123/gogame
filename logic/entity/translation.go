package entity

import (
	"github.com/Danice123/idk/logic"
	"github.com/faiface/pixel"
)

type Translation struct {
	Direction logic.Direction
	Completed float64
	Signal    chan bool
}

func (ths *Translation) Vector(scaledTileSize float64) pixel.Vec {
	partial := ths.Completed * scaledTileSize

	switch ths.Direction {
	case logic.NORTH:
		return pixel.V(0, partial)
	case logic.SOUTH:
		return pixel.V(0, -partial)
	case logic.EAST:
		return pixel.V(partial, 0)
	case logic.WEST:
		return pixel.V(-partial, 0)
	default:
		panic("No direction given?")
	}
}
