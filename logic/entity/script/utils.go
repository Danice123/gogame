package script

import (
	"strings"

	"github.com/Danice123/idk/logic"
)

func GetDirection(dir string) logic.Direction {
	switch strings.ToLower(dir) {
	case "n":
		fallthrough
	case "north":
		return logic.NORTH
	case "s":
		fallthrough
	case "south":
		return logic.SOUTH
	case "e":
		fallthrough
	case "east":
		return logic.EAST
	case "w":
		fallthrough
	case "west":
		return logic.WEST
	default:
		return ""
	}
}
