package logic

import (
	"github.com/faiface/pixel"
)

type Coord struct {
	X     int
	Y     int
	Layer int
}

func (ths Coord) Vector() pixel.Vec {
	return pixel.V(float64(ths.X), float64(ths.Y))
}

func (ths Coord) Translate(dir Direction) Coord {
	switch dir {
	case NORTH:
		return Coord{
			X:     ths.X,
			Y:     ths.Y + 1,
			Layer: ths.Layer,
		}
	case SOUTH:
		return Coord{
			X:     ths.X,
			Y:     ths.Y - 1,
			Layer: ths.Layer,
		}
	case EAST:
		return Coord{
			X:     ths.X + 1,
			Y:     ths.Y,
			Layer: ths.Layer,
		}
	case WEST:
		return Coord{
			X:     ths.X - 1,
			Y:     ths.Y,
			Layer: ths.Layer,
		}
	default:
		panic("No direction given?")
	}
}
