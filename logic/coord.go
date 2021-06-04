package logic

import "github.com/faiface/pixel"

type Coord struct {
	X     int
	Y     int
	Layer int
}

func (ths Coord) Vector() pixel.Vec {
	return pixel.V(float64(ths.X), float64(ths.Y))
}
