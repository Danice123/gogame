package utils

import (
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
)

type Coord struct {
	X int
	Y int
}

func LoadPicture(path string) *pixel.PictureData {
	if file, err := os.Open(path); err != nil {
		panic(err)
	} else {
		defer file.Close()
		if img, _, err := image.Decode(file); err != nil {
			panic(err)
		} else {
			img.ColorModel()
			return pixel.PictureDataFromImage(img)
		}
	}
}
