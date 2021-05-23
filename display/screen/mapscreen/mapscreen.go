package mapscreen

import (
	"github.com/Danice123/idk/display/tiledmap"
	"github.com/faiface/pixel/pixelgl"
)

type MapScreen struct {
	TiledMap *tiledmap.OrthoMap
}

func (ths *MapScreen) Tick(delta int64) {

}

func (ths *MapScreen) Render(delta int64, window *pixelgl.Window) {
	for i := 0; i < ths.TiledMap.NumLayers(); i++ {
		ths.TiledMap.RenderLayer(delta, window, i)
	}
}
