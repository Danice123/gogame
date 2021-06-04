package mapscreen

import (
	displayEntity "github.com/Danice123/idk/display/entity"
	"github.com/Danice123/idk/display/tiledmap"
	"github.com/Danice123/idk/logic/entity"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type MapScreen struct {
	TiledMap      *tiledmap.OrthoMap
	EntityHandler *displayEntity.EntityHandler
	Player        *entity.Player
}

func (ths *MapScreen) ShouldRenderBehind() bool {
	return false
}

func (ths *MapScreen) Tick(delta int64) {
}

func (ths *MapScreen) Render(delta int64, window *pixelgl.Window) {
	ths.TiledMap.RenderBackground(window)

	tileRatio := window.Bounds().W() / float64(ths.TiledMap.TileSize*10)
	playerV := ths.Player.Coord.Vector().Scaled(float64(ths.TiledMap.TileSize / 2)).Sub(pixel.V(float64(ths.TiledMap.TileSize/2), float64(ths.TiledMap.TileSize/2)))
	camera := pixel.IM.Moved(playerV).Scaled(pixel.Vec{}, tileRatio).Moved(window.Bounds().Center())

	for i := 0; i < ths.TiledMap.NumLayers(); i++ {
		layer := ths.TiledMap.RenderLayer(i)
		ths.EntityHandler.Render(layer, ths.TiledMap.TileSize, i)
		layer.Draw(window, camera)
	}
}
