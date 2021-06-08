package chatbox

import (
	"github.com/Danice123/idk/display/menutils"
	"github.com/Danice123/idk/display/screen"
	"github.com/Danice123/idk/display/utils"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

type ChatBox struct {
	Finished chan bool
	content  *menutils.Text

	screen.BaseScreen
}

func New(content string) *ChatBox {
	font := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	t := menutils.NewTextContent(font)
	t.SetText(content)

	return &ChatBox{
		content:  t,
		Finished: make(chan bool),
	}
}

func (ths *ChatBox) ShouldRenderBehind() bool {
	return true
}

func (ths *ChatBox) Tick(delta int64) {

}

func (ths *ChatBox) Render(delta int64, window *pixelgl.Window) {
	ths.content.Render(window.Canvas(), 0, 0)
}

func (ths *ChatBox) HandleKey(key utils.KEY) {
	switch key {
	case utils.ACTIVATE:
		fallthrough
	case utils.DECLINE:
		ths.Finished <- true
	}
}
