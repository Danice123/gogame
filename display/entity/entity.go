package entity

import (
	"github.com/Danice123/gogame/display/texturepacker"
	"github.com/Danice123/gogame/logic"
	"github.com/Danice123/gogame/logic/entity"
	"github.com/faiface/pixel/pixelgl"
)

type EntityHandler struct {
	entities   []entity.Entity
	sheetCache map[string]*texturepacker.SpriteSheet
}

func NewEntityHandler() *EntityHandler {
	return &EntityHandler{
		entities:   []entity.Entity{},
		sheetCache: make(map[string]*texturepacker.SpriteSheet),
	}
}

func (ths *EntityHandler) AddEntity(added entity.Entity) {
	ths.entities = append(ths.entities, added)
	if _, ok := ths.sheetCache[added.SpriteSheet().Name]; !ok {
		ths.sheetCache[added.SpriteSheet().Name] = added.SpriteSheet()
	}
}

func (ths *EntityHandler) EntityAtTile(coord logic.Coord) entity.Entity {
	for _, entity := range ths.entities {
		if coord == entity.GetCoord() {
			return entity
		}
		if entity.Translation() != nil && coord == entity.GetCoord().Translate(entity.Translation().Direction) {
			return entity
		}
	}
	return nil
}

func (ths *EntityHandler) Tick() {
	for _, entity := range ths.entities {
		entity.Tick()
	}
}

func (ths *EntityHandler) Render(canvas *pixelgl.Canvas, tileSize int, tileRatio float64, layer int) {
	// if ths.entities == nil {
	// 	return
	// }

	// for _, ss := range ths.sheetCache {
	// 	ss.Batch.Clear()
	// }

	// scaledTileSize := float64(tileSize) * tileRatio
	// for _, e := range ths.entities {
	// 	if e.GetCoord().Layer == layer {
	// 		matrix := pixel.IM.Scaled(pixel.ZV, tileRatio).Moved(e.GetCoord().Vector().Scaled(scaledTileSize).Add(pixel.V(scaledTileSize/2, scaledTileSize/2+tileRatio*4)))
	// 		if e.Translation() != nil {
	// 			matrix = matrix.Moved(e.Translation().Vector(scaledTileSize))
	// 		}
	// 		e.Sprite().Draw(e.SpriteSheet().Batch, matrix)
	// 	}
	// }

	// for _, ss := range ths.sheetCache {
	// 	ss.Batch.Draw(canvas)
	// }
}
