package script

import (
	"github.com/Danice123/idk/display/screen"
	lua "github.com/yuin/gopher-lua"
)

type Player interface {
}

type ScriptHandler struct {
	Screen screen.Screen
	Player Player
}

func (ths *ScriptHandler) MakeLoaderFunction() func(*lua.LState) int {
	return func(luaState *lua.LState) int {
		exports := map[string]lua.LGFunction{
			"display": ths.Display,
		}
		mod := luaState.SetFuncs(luaState.NewTable(), exports)
		luaState.SetField(mod, "name", lua.LString("game"))
		luaState.Push(mod)
		return 1
	}
}
