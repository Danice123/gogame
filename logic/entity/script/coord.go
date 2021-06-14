package script

import lua "github.com/yuin/gopher-lua"

func (ths *EntityHandler) GetCoord(l *lua.LState) int {
	coord := l.NewTable()
	coord.RawSetString("X", lua.LNumber(ths.Entity.GetCoord().X))
	coord.RawSetString("Y", lua.LNumber(ths.Entity.GetCoord().Y))
	coord.RawSetString("Layer", lua.LNumber(ths.Entity.GetCoord().Layer))
	l.Push(coord)
	return 1
}
