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

func (ths *OrthoMap) RenderBackground(window *pixelgl.Window) {
	if ths.tiledMap.BackgroundColor != nil {
		window.Clear(ths.tiledMap.BackgroundColor)
	} else {
		window.Clear(color.White)
	}
}

func (ths *OrthoMap) RenderLayer(layer int) *pixelgl.Canvas {
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
			matrix := pixel.IM.Moved(pixel.V(float64(x*ths.TileSize+ths.TileSize/2), float64(y*ths.TileSize+ths.TileSize/2)))

			tileSprite := ths.tiledSets[tiledLayer.Tiles[tilePointer].Tileset.FirstGID].tileCache[tiledLayer.Tiles[tilePointer].ID]
			tileSprite.Draw(ths.tiledSets[tiledLayer.Tiles[tilePointer].Tileset.FirstGID].batch, matrix)
			tilePointer++
		}
	}

	canvas := pixelgl.NewCanvas(pixel.R(0, 0, float64(ths.tiledMap.Width*ths.tiledMap.TileWidth), float64(ths.tiledMap.Height*ths.tiledMap.TileHeight)))

	for _, tiledSet := range ths.tiledSets {
		tiledSet.batch.Draw(canvas)
	}

	return canvas
}