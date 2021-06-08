package mapscreen

import (
	"image/color"

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

	mapCanvas *pixelgl.Canvas
}

func (ths *MapScreen) ShouldRenderBehind() bool {
	return false
}

func (ths *MapScreen) Tick(delta int64) {
	ths.Player.Tick()
}

func (ths *MapScreen) Render(delta int64, window *pixelgl.Window) {
	tileRatio := window.Bounds().W() / float64(ths.TiledMap.TileSize*10)

	if ths.mapCanvas == nil {
		ths.mapCanvas = pixelgl.NewCanvas(pixel.R(0, 0, ths.TiledMap.MapSize().X*float64(ths.TiledMap.TileSize)*tileRatio, ths.TiledMap.MapSize().Y*float64(ths.TiledMap.TileSize)*tileRatio))
	}
	ths.mapCanvas.Clear(color.White)

	for i := 0; i < ths.TiledMap.NumLayers(); i++ {
		ths.TiledMap.RenderLayer(ths.mapCanvas, i, tileRatio)
		ths.EntityHandler.Render(ths.mapCanvas, ths.TiledMap.TileSize, tileRatio, i)
	}

	playerV := ths.TiledMap.MapSize().Scaled(0.5).Sub(ths.Player.Coord.Vector())
	playerScaled := playerV.Scaled(float64(ths.TiledMap.TileSize) * tileRatio).Sub(pixel.V(float64(ths.TiledMap.TileSize)*tileRatio/2, float64(ths.TiledMap.TileSize)*tileRatio/2))
	if ths.Player.Translation() != nil {
		playerScaled = playerScaled.Sub(ths.Player.Translation().Vector(float64(ths.TiledMap.TileSize) * tileRatio))
	}

	camera := pixel.IM.Moved(playerScaled).Moved(window.Bounds().Center())
	ths.TiledMap.RenderBackground(window)
	ths.mapCanvas.Draw(window, camera)
}

func (ths *MapScreen) isValidDestination(coord logic.Coord) bool {
	return !ths.TiledMap.IsTileAt(coord.X, coord.Y, coord.Layer)
}

func (ths *MapScreen) HandleKey(key utils.KEY) {
	switch key {
	case utils.UP:
		if ths.isValidDestination(ths.Player.Coord.Translate(logic.NORTH)) {
			ths.Player.Walk(logic.NORTH)
		} else {
			ths.Player.Facing = logic.NORTH
		}
	case utils.DOWN:
		if ths.isValidDestination(ths.Player.Coord.Translate(logic.SOUTH)) {
			ths.Player.Walk(logic.SOUTH)
		} else {
			ths.Player.Facing = logic.SOUTH
		}
	case utils.LEFT:
		if ths.isValidDestination(ths.Player.Coord.Translate(logic.WEST)) {
			ths.Player.Walk(logic.WEST)
		} else {
			ths.Player.Facing = logic.WEST
		}
	case utils.RIGHT:
		if ths.isValidDestination(ths.Player.Coord.Translate(logic.EAST)) {
			ths.Player.Walk(logic.EAST)
		} else {
			ths.Player.Facing = logic.EAST
		}
	}
}
