package tiledmap

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/lafriks/go-tiled"
)

type OrthoMap struct {
	tiledMap  *tiled.Map
	tiledSets map[uint32]*TiledSet
}

func NewOrthoMap(path string) *OrthoMap {
	loadedMap, err := tiled.LoadFromFile(path)
	if err != nil {
		panic(err)
	}

	m := &OrthoMap{
		tiledMap:  loadedMap,
		tiledSets: make(map[uint32]*TiledSet),
	}

	for _, tileset := range m.tiledMap.Tilesets {
		m.tiledSets[tileset.FirstGID] = NewTiledSet(tileset)
	}

	return m
}

func (ths *OrthoMap) NumLayers() int {
	return len(ths.tiledMap.Layers)
}

func (ths *OrthoMap) RenderBackground(delta int64, window *pixelgl.Window) {
	if ths.tiledMap.BackgroundColor != nil {
		window.Clear(ths.tiledMap.BackgroundColor)
	} else {
		window.Clear(color.White)
	}
}

func (ths *OrthoMap) RenderLayer(delta int64, window *pixelgl.Window, layer int) {
	tiledLayer := ths.tiledMap.Layers[layer]

	for _, tiledSet := range ths.tiledSets {
		tiledSet.batch.Clear()
	}

	tilePointer := 0
	for y := 0; y < ths.tiledMap.Height; y++ {
		for x := 0; x < ths.tiledMap.Width; x++ {
			if tiledLayer.Tiles[tilePointer].IsNil() {
				tilePointer++
				continue
			}
			location := pixel.V(float64(x*ths.tiledMap.TileWidth), float64(y*ths.tiledMap.TileHeight))
			centerV := pixel.V(float64(ths.tiledMap.Width*ths.tiledMap.TileWidth/2), float64(ths.tiledMap.Height*ths.tiledMap.TileHeight/2))

			matrix := pixel.IM.Moved(window.Bounds().Center().Sub(centerV))
			matrix = matrix.Moved(location)

			tileSprite := ths.tiledSets[tiledLayer.Tiles[tilePointer].Tileset.FirstGID].tileCache[tiledLayer.Tiles[tilePointer].ID]
			tileSprite.Draw(ths.tiledSets[tiledLayer.Tiles[tilePointer].Tileset.FirstGID].batch, matrix)
			tilePointer++
		}
	}

	for _, tiledSet := range ths.tiledSets {
		tiledSet.batch.Draw(window)
	}
}
