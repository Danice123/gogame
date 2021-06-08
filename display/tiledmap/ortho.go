package tiledmap

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/lafriks/go-tiled"
)

type OrthoMap struct {
	TileSize int

	tiledMap  *tiled.Map
	tiledSets map[uint32]*TiledSet
}

func NewOrthoMap(path string) *OrthoMap {
	loadedMap, err := tiled.LoadFromFile(path)
	if err != nil {
		panic(err)
	}
	if loadedMap.TileWidth != loadedMap.TileHeight {
		panic("Cannot handle rectangular tiles!")
	}

	m := &OrthoMap{
		TileSize:  loadedMap.TileWidth,
		tiledMap:  loadedMap,
		tiledSets: make(map[uint32]*TiledSet),
	}

	for _, tileset := range m.tiledMap.Tilesets {
		m.tiledSets[tileset.FirstGID] = NewTiledSet(tileset)
	}

	return m
}

func (ths *OrthoMap) MapSize() pixel.Vec {
	return pixel.Vec{
		X: float64(ths.tiledMap.Width),
		Y: float64(ths.tiledMap.Height),
	}
}

func (ths *OrthoMap) NumLayers() int {
	return len(ths.tiledMap.Layers)
}

func (ths *OrthoMap) IsTileAt(x, y, layer int) bool {
	tileId := y*ths.tiledMap.Width + x
	return !ths.tiledMap.Layers[layer].Tiles[tileId].IsNil()
}

func (ths *OrthoMap) RenderBackground(window *pixelgl.Window) {
	if ths.tiledMap.BackgroundColor != nil {
		window.Clear(ths.tiledMap.BackgroundColor)
	} else {
		window.Clear(color.White)
	}
}

func (ths *OrthoMap) RenderLayer(canvas *pixelgl.Canvas, layer int, scaleFactor float64) {
	tiledLayer := ths.tiledMap.Layers[layer]

	for _, tiledSet := range ths.tiledSets {
		tiledSet.batch.Clear()
	}

	tilePointer := 0
	tileSize := float64(ths.TileSize) * scaleFactor
	for y := 0; y < ths.tiledMap.Height; y++ {
		for x := 0; x < ths.tiledMap.Width; x++ {
			if tiledLayer.Tiles[tilePointer].IsNil() {
				tilePointer++
				continue
			}
			matrix := pixel.IM.Scaled(pixel.ZV, scaleFactor).Moved(pixel.V(float64(x)*tileSize+tileSize/2, float64(y)*tileSize+tileSize/2))

			tileSprite := ths.tiledSets[tiledLayer.Tiles[tilePointer].Tileset.FirstGID].tileCache[tiledLayer.Tiles[tilePointer].ID]
			tileSprite.Draw(ths.tiledSets[tiledLayer.Tiles[tilePointer].Tileset.FirstGID].batch, matrix)
			tilePointer++
		}
	}

	for _, tiledSet := range ths.tiledSets {
		tiledSet.batch.Draw(canvas)
	}
}
