package script

import (
	"github.com/Danice123/idk/display/screen/chatbox"
	lua "github.com/yuin/gopher-lua"
)

func (ths *ScriptHandler) Display(l *lua.LState) int {
	text := l.ToString(1)

	chat := chatbox.New(text)
	ths.Screen.SetChild(chat)
	<-chat.Finished
	ths.Screen.SetChild(nil)
	return 0
}
