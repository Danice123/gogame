package netbattle

import (
	"path/filepath"

	"github.com/Danice123/gogame/display/texturepacker"
	"github.com/faiface/pixel/pixelgl"
)

type HealthDisplay struct {
	sprites *texturepacker.SpriteSheet

	health *int
}

func NewHealthDisplay(health *int) *HealthDisplay {
	return &HealthDisplay{
		sprites: texturepacker.NewSpriteSheet(filepath.Join("resources", "sheets", "battleui.json")),
		health:  health,
	}
}

func (ths *HealthDisplay) Tick() {

}

func (ths *HealthDisplay) Render(canvas *pixelgl.Canvas, x int, y int) {
	ths.sprites.Clear()

	ths.sprites.DrawFrame("battleui-health", 12, x, y)
	ths.sprites.DrawSpriteNumberPadded("battleui-health", *ths.health, 4, 11, x+8, y)
	ths.sprites.DrawFrame("battleui-health", 13, x+40, y)

	ths.sprites.Render(canvas)
}
