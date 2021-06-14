package script

import (
	"github.com/Danice123/idk/display/screen/chatbox"
	lua "github.com/yuin/gopher-lua"
)

func (ths *ScriptHandler) Display(l *lua.LState) int {
	chat := chatbox.New("I'm a big boy with big boy powers!")
	ths.Screen.SetChild(chat)
	<-chat.Finished
	ths.Screen.SetChild(nil)
	return 0
}
