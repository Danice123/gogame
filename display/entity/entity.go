package entity

import (
	"github.com/Danice123/idk/display/texturepacker"
	"github.com/Danice123/idk/logic/entity"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type EntityHandler struct {
	entities   []entity.Entity
	sheetCache map[string]*texturepacker.SpriteSheet
}

func (ths *EntityHandler) initialize() {
	if ths.entities == nil {
		ths.entities = []entity.Entity{}
	}
	if ths.sheetCache == nil {
		ths.sheetCache = make(map[string]*texturepacker.SpriteSheet)
	}
}

func (ths *EntityHandler) AddEntity(added entity.Entity) {
	ths.initialize()
	ths.entities = append(ths.entities, added)
	if _, ok := ths.sheetCache[added.GetSpriteSheet().Name]; !ok {
		ths.sheetCache[added.GetSpriteSheet().Name] = added.GetSpriteSheet()
	}
}

func (ths *EntityHandler) Render(canvas *pixelgl.Canvas, tileSize int, layer int) {
	if ths.entities == nil {
		return
	}

	for _, ss := range ths.sheetCache {
		ss.Batch.Clear()
	}

	for _, e := range ths.entities {
		if e.GetCoord().Layer == layer {
			matrix := pixel.IM.Moved(e.GetCoord().Vector().Scaled(float64(tileSize)).Sub(pixel.V(float64(tileSize/2), float64(tileSize/2))))
			e.GetSprite().Draw(e.GetSpriteSheet().Batch, matrix)
		}
	}

	for _, ss := range ths.sheetCache {
		ss.Batch.Draw(canvas)
	}
}
