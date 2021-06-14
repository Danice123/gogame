package chatbox

import (
	"github.com/Danice123/gogame/display/menutils"
	"github.com/Danice123/gogame/display/screen"
	"github.com/Danice123/gogame/display/utils"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

type ChatBox struct {
	Finished chan bool

	textContent *menutils.Text
	content     menutils.Content

	screen.BaseScreen
}

func New(content string) *ChatBox {
	font := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	t := menutils.NewTextContent(font)
	t.Scale = 5
	t.SetText(content, true)

	return &ChatBox{
		textContent: t,
		Finished:    make(chan bool),
	}
}

func (ths *ChatBox) ShouldRenderBehind() bool {
	return true
}

func (ths *ChatBox) Tick(delta int64) {
	ths.textContent.Tick(delta)
}

func (ths *ChatBox) Render(delta int64, window *pixelgl.Window) {
	if ths.content == nil {
		ths.textContent.SetMaxWidth(window.Bounds().W() - 20)
		ths.content = &menutils.BorderBox{
			BorderSize: 10,
			Fill:       true,
			Content: &menutils.ContentWithMargin{
				AlignTop:   true,
				LeftMargin: 4,
				MinWidth:   window.Bounds().W() - 20,
				MinHeight:  window.Bounds().H() * 0.22,
				Content:    ths.textContent,
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
		if ths.textContent.IsFinished() {
			ths.Finished <- true
		} else {
			ths.textContent.Finish()
		}
	}
}
