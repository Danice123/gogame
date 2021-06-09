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

	textContent string
	content     menutils.Content
	font        *text.Atlas

	screen.BaseScreen
}

func New(content string) *ChatBox {
	return &ChatBox{
		textContent: content,
		font:        text.NewAtlas(basicfont.Face7x13, text.ASCII),
		Finished:    make(chan bool),
	}
}

func (ths *ChatBox) ShouldRenderBehind() bool {
	return true
}

func (ths *ChatBox) Tick(delta int64) {

}

func (ths *ChatBox) Render(delta int64, window *pixelgl.Window) {
	if ths.content == nil {
		t := menutils.NewTextContent(ths.font)
		t.Scale = 5
		t.SetText(ths.textContent)

		ths.content = &menutils.BorderBox{
			BorderSize: 10,
			Fill:       true,
			Content: &menutils.ContentWithMargin{
				AlignTop:  true,
				MinWidth:  window.Bounds().W() - 20,
				MinHeight: window.Bounds().H()*0.25 - 20,
				Content:   t,
			},
		}
	}

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
