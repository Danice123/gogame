package entity

import (
	"github.com/Danice123/idk/pkg/display/texturepacker"
	"github.com/Danice123/idk/pkg/logic/entity"
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
	if _, ok := ths.sheetCache[added.SpriteSheet().Name]; !ok {
		ths.sheetCache[added.SpriteSheet().Name] = added.SpriteSheet()
	}
}

func (ths *EntityHandler) Render(canvas *pixelgl.Canvas, tileSize int, tileRatio float64, layer int) {
	if ths.entities == nil {
		return
	}

	for _, ss := range ths.sheetCache {
		ss.Batch.Clear()
	}

	scaledTileSize := float64(tileSize) * tileRatio
	for _, e := range ths.entities {
		if e.GetCoord().Layer == layer {
			matrix := pixel.IM.Scaled(pixel.ZV, tileRatio).Moved(e.GetCoord().Vector().Scaled(scaledTileSize).Add(pixel.V(scaledTileSize/2, scaledTileSize/2+tileRatio*4)))
			if e.Translation() != nil {
				matrix = matrix.Moved(e.Translation().Vector(scaledTileSize))
			}
			e.Sprite().Draw(e.SpriteSheet().Batch, matrix)
		}
	}

	for _, ss := range ths.sheetCache {
		ss.Batch.Draw(canvas)
	}
}
