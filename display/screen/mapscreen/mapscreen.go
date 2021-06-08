package mapscreen

import (
	"image/color"

	displayEntity "github.com/Danice123/idk/display/entity"
	"github.com/Danice123/idk/display/screen"
	"github.com/Danice123/idk/display/tiledmap"
	"github.com/Danice123/idk/display/utils"
	"github.com/Danice123/idk/logic"
	"github.com/Danice123/idk/logic/entity"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type MapScreen struct {
	tMap          *tiledmap.OrthoMap
	entityHandler *displayEntity.EntityHandler
	player        *entity.Player
	mapCanvas     *pixelgl.Canvas

	screen.BaseScreen
}

func New(tMap *tiledmap.OrthoMap, entities []entity.Entity, player *entity.Player) *MapScreen {
	screen := &MapScreen{
		tMap:          tMap,
		entityHandler: displayEntity.NewEntityHandler(),
		player:        player,
	}
	screen.entityHandler.AddEntity(player)
	for _, entity := range entities {
		screen.entityHandler.AddEntity(entity)
	}
	return screen
}

func (ths *MapScreen) ShouldRenderBehind() bool {
	return false
}

func (ths *MapScreen) Tick(delta int64) {
	ths.player.Tick()
}

func (ths *MapScreen) Render(delta int64, window *pixelgl.Window) {
	tileRatio := window.Bounds().W() / float64(ths.tMap.TileSize*10)

	if ths.mapCanvas == nil { // Or if window size is changed?
		ths.mapCanvas = pixelgl.NewCanvas(pixel.R(0, 0, ths.tMap.MapSize().X*float64(ths.tMap.TileSize)*tileRatio, ths.tMap.MapSize().Y*float64(ths.tMap.TileSize)*tileRatio))
	}
	ths.mapCanvas.Clear(color.White)

	for i := 0; i < ths.tMap.NumLayers(); i++ {
		ths.tMap.RenderLayer(ths.mapCanvas, i, tileRatio)
		ths.entityHandler.Render(ths.mapCanvas, ths.tMap.TileSize, tileRatio, i)
	}

	playerV := ths.tMap.MapSize().Scaled(0.5).Sub(ths.player.Coord.Vector())
	playerScaled := playerV.Scaled(float64(ths.tMap.TileSize) * tileRatio).Sub(pixel.V(float64(ths.tMap.TileSize)*tileRatio/2, float64(ths.tMap.TileSize)*tileRatio/2))
	if ths.player.Translation() != nil {
		playerScaled = playerScaled.Sub(ths.player.Translation().Vector(float64(ths.tMap.TileSize) * tileRatio))
	}

	camera := pixel.IM.Moved(playerScaled).Moved(window.Bounds().Center())
	ths.tMap.RenderBackground(window)
	ths.mapCanvas.Draw(window, camera)
}

func (ths *MapScreen) isValidDestination(coord logic.Coord) bool {
	if ths.entityHandler.EntityAtTile(coord) != nil {
		return false
	}
	if ths.tMap.IsTileAt(coord.X, coord.Y, coord.Layer) {
		return false
	}
	return true
}

func (ths *MapScreen) HandleKey(key utils.KEY) {
	switch key {
	case utils.UP:
		if ths.isValidDestination(ths.player.Coord.Translate(logic.NORTH)) {
			ths.player.Walk(logic.NORTH)
		} else {
			ths.player.Face(logic.NORTH)
		}
	case utils.DOWN:
		if ths.isValidDestination(ths.player.Coord.Translate(logic.SOUTH)) {
			ths.player.Walk(logic.SOUTH)
		} else {
			ths.player.Face(logic.SOUTH)
		}
	case utils.LEFT:
		if ths.isValidDestination(ths.player.Coord.Translate(logic.WEST)) {
			ths.player.Walk(logic.WEST)
		} else {
			ths.player.Face(logic.WEST)
		}
	case utils.RIGHT:
		if ths.isValidDestination(ths.player.Coord.Translate(logic.EAST)) {
			ths.player.Walk(logic.EAST)
		} else {
			ths.player.Face(logic.EAST)
		}
	case utils.ACTIVATE:
		if entity := ths.entityHandler.EntityAtTile(ths.player.Coord.Translate(ths.player.Facing)); entity != nil {
			entity.Activate(ths, ths.player)
		}
	}
}
