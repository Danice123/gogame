package script

import (
	"github.com/Danice123/gogame/logic"
	lua "github.com/yuin/gopher-lua"
)

func (ths *EntityHandler) GetFacing(l *lua.LState) int {
	l.Push(lua.LString(ths.Entity.GetFacing()))
	return 1
}

func (ths *EntityHandler) Face(l *lua.LState) int {
	dir := GetDirection(l.ToString(1))
	ths.Entity.Face(dir)
	return 0
}

func (ths *EntityHandler) FaceTowards(l *lua.LState) int {
	table := l.ToTable(1)
	x := lua.LVAsNumber(table.RawGetString("X"))
	y := lua.LVAsNumber(table.RawGetString("Y"))
	layer := lua.LVAsNumber(table.RawGetString("Layer"))

	coord := logic.Coord{
		X:     int(x),
		Y:     int(y),
		Layer: int(layer),
	}

	ths.Entity.FaceTowards(coord)
	return 0
}
