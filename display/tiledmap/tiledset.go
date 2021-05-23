package tiledmap

import (
	"image"

	"github.com/Danice123/idk/display/utils"
	"github.com/faiface/pixel"
	"github.com/lafriks/go-tiled"
)

type TiledSet struct {
	Batch     *pixel.Batch
	TileCache map[uint32]*pixel.Sprite
}

func NewTiledSet(tileset *tiled.Tileset) *TiledSet {
	tilesetSource := utils.LoadPicture(tileset.GetFileFullPath(tileset.Image.Source)) // TODO: Handle tilesets with multiple images

	tiledset := &TiledSet{
		Batch:     pixel.NewBatch(&pixel.TrianglesData{}, tilesetSource),
		TileCache: make(map[uint32]*pixel.Sprite),
	}

	for i := uint32(0); i < uint32(tileset.TileCount); i++ {
		tiledset.TileCache[i] = pixel.NewSprite(tilesetSource, transformRect(tilesetSource.Bounds(), tileset.GetTileRect(i)))
	}

	return tiledset
}

func transformRect(sourceBounds pixel.Rect, rect image.Rectangle) pixel.Rect {
	return pixel.R(float64(rect.Min.X), sourceBounds.H()-float64(rect.Max.Y), float64(rect.Max.X), sourceBounds.H()-float64(rect.Min.Y))
}
