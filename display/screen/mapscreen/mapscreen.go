package mapscreen

import (
	displayEntity "github.com/Danice123/idk/display/entity"
	"github.com/Danice123/idk/display/tiledmap"
	"github.com/Danice123/idk/display/utils"
	"github.com/Danice123/idk/logic"
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
	ths.Player.Tick()
}

func (ths *MapScreen) Render(delta int64, window *pixelgl.Window) {
	ths.TiledMap.RenderBackground(window)

	tileRatio := window.Bounds().W() / float64(ths.TiledMap.TileSize*10)

	playerV := ths.TiledMap.MapSize().Scaled(0.5).Sub(ths.Player.Coord.Vector())
	playerScaled := playerV.Scaled(float64(ths.TiledMap.TileSize)).Sub(pixel.V(float64(ths.TiledMap.TileSize)/2, float64(ths.TiledMap.TileSize)/2))
	if ths.Player.Translation() != nil {
		playerScaled = playerScaled.Sub(ths.Player.Translation().Vector(ths.TiledMap.TileSize))
	}

	camera := pixel.IM.Moved(playerScaled).Scaled(pixel.Vec{}, tileRatio).Moved(window.Bounds().Center())

	for i := 0; i < ths.TiledMap.NumLayers(); i++ {
		layer := ths.TiledMap.RenderLayer(i)
		ths.EntityHandler.Render(layer, ths.TiledMap.TileSize, i)
		layer.Draw(window, camera)
	}
}

func (ths *MapScreen) HandleKey(key utils.KEY) {
	switch key {
	case utils.UP:
		ths.Player.Walk(logic.NORTH)
	case utils.DOWN:
		ths.Player.Walk(logic.SOUTH)
	case utils.LEFT:
		ths.Player.Walk(logic.WEST)
	case utils.RIGHT:
		ths.Player.Walk(logic.EAST)
	}
}
