package entity

import (
	"math"

	"github.com/Danice123/idk/display/screen"
	"github.com/Danice123/idk/display/texturepacker"
	"github.com/Danice123/idk/logic"
	"github.com/Danice123/idk/logic/entity/script"
	"github.com/faiface/pixel"
	lua "github.com/yuin/gopher-lua"
)

var tps = 1.0 / 60.0

type Entity interface {
	Name() string
	SpriteSheet() *texturepacker.SpriteSheet
	Sprite() *pixel.Sprite
	GetCoord() logic.Coord
	Translation() *Translation
	Activate(screen screen.Screen, player *Player)
	Tick()
}

type Base struct {
	EntityName string
	Coord      logic.Coord
	Frame      int

	facing      logic.Direction
	translation *Translation
	script      string

	// Tentative
	Spritesheet *texturepacker.SpriteSheet
}

func (ths *Base) Name() string {
	return ths.EntityName
}

func (ths *Base) SpriteSheet() *texturepacker.SpriteSheet {
	return ths.Spritesheet
}

func (ths *Base) Sprite() *pixel.Sprite {
	return ths.Spritesheet.Sprites[ths.EntityName][string(ths.facing)][ths.Frame]
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
			if ths.Frame == len(ths.Spritesheet.Sprites[ths.EntityName][string(ths.facing)])-1 {
				ths.Frame = 0
			} else {
				ths.Frame++
			}
		}

		if ths.translation.Completed >= 1.0 {
			select {
			case ths.translation.Signal <- true:
			default:
			}
			ths.Coord = ths.Coord.Translate(ths.translation.Direction)
			ths.translation = nil
			ths.Frame = 0
		}
	}
}

func (ths *Base) GetFacing() logic.Direction {
	return ths.facing
}

func (ths *Base) Face(dir logic.Direction) {
	if ths.translation == nil {
		ths.facing = dir
	}
}

func (ths *Base) FaceTowards(coord logic.Coord) {
	if ths.translation == nil {
		dx := ths.Coord.X - coord.X
		dy := ths.Coord.Y - coord.Y

		if math.Abs(float64(dx)) > math.Abs(float64(dy)) {
			if dx > 0 {
				ths.facing = logic.WEST
			} else {
				ths.facing = logic.EAST
			}
		} else {
			if dy > 0 {
				ths.facing = logic.SOUTH
			} else {
				ths.facing = logic.NORTH
			}
		}
	}
}

func (ths *Base) Walk(dir logic.Direction) chan bool {
	if ths.translation == nil {
		ths.facing = dir
		ths.translation = &Translation{
			Direction: dir,
			Signal:    make(chan bool),
		}
		return ths.translation.Signal
	}
	return nil
}

func (ths *Base) Activate(screen screen.Screen, player *Player) {
	sh := &script.ScriptHandler{
		Screen: screen,
	}
	eh := &script.EntityHandler{
		Entity: ths,
	}
	ph := &script.EntityHandler{
		Entity: player,
	}

	go func() {
		player.Locked = true
		luaState := lua.NewState()
		defer luaState.Close()
		luaState.PreloadModule("game", sh.MakeLoaderFunction())
		luaState.PreloadModule("self", eh.MakeLoaderFunction())
		luaState.PreloadModule("player", ph.MakeLoaderFunction())
		if err := luaState.DoFile(ths.script); err != nil {
			panic(err)
		}
		player.Locked = false
	}()
}
