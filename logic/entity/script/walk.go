package script

import lua "github.com/yuin/gopher-lua"

func (ths *EntityHandler) Walk(l *lua.LState) int {
	for i := 1; l.Get(i) != lua.LNil; i++ {
		dir := GetDirection(l.ToString(i))
		c := ths.Entity.Walk(dir)
		if c != nil {
			<-c
		}
	}
	return 0
}
