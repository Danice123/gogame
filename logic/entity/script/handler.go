package script

import (
	"github.com/Danice123/gogame/display/screen"
	"github.com/Danice123/gogame/logic"
	lua "github.com/yuin/gopher-lua"
)

type ScriptHandler struct {
	Screen screen.Screen
}

func (ths *ScriptHandler) MakeLoaderFunction() func(*lua.LState) int {
	return func(luaState *lua.LState) int {
		exports := map[string]lua.LGFunction{
			"Display": ths.Display,
		}
		mod := luaState.SetFuncs(luaState.NewTable(), exports)
		luaState.Push(mod)
		return 1
	}
}

type Entity interface {
	GetCoord() logic.Coord
	GetFacing() logic.Direction
	Face(logic.Direction)
	FaceTowards(logic.Coord)
	Walk(logic.Direction) chan bool
}

type EntityHandler struct {
	Entity Entity
}

func (ths *EntityHandler) MakeLoaderFunction() func(*lua.LState) int {
	return func(luaState *lua.LState) int {
		exports := map[string]lua.LGFunction{
			"GetCoord":    ths.GetCoord,
			"GetFacing":   ths.GetFacing,
			"Face":        ths.Face,
			"FaceTowards": ths.FaceTowards,
			"Walk":        ths.Walk,
		}
		mod := luaState.SetFuncs(luaState.NewTable(), exports)
		luaState.Push(mod)
		return 1
	}
}
