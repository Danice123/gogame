package utils

import (
	"bytes"
	"image"
	_ "image/png"

	"github.com/Danice123/idk/pkg/data"
	"github.com/faiface/pixel"
)

func LoadPicture(path string) pixel.Picture {
	if bs, err := data.Asset(path); err != nil {
		panic(err)
	} else {
		if img, _, err := image.Decode(bytes.NewReader(bs)); err != nil {
			panic(err)
		} else {
			return pixel.PictureDataFromImage(img)
		}
	}
}
