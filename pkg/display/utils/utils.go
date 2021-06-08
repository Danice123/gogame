package utils

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"os"
	"path/filepath"

	"github.com/Danice123/idk/pkg/data"
	"github.com/faiface/pixel"
)

var pathCache map[string]string

func init() {
	pathCache = map[string]string{}
}

func FindPath(containing string) string {
	if val, ok := pathCache[containing]; ok {
		return val
	}
	workingDirectory, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	lastDir := workingDirectory
	for {
		currentPath := fmt.Sprintf("%s/%s", lastDir, containing)
		fi, err := os.Stat(currentPath)
		if err == nil {
			switch mode := fi.Mode(); {
			case mode.IsDir():
				pathCache[containing] = currentPath
				return currentPath
			}
		}
		newDir := filepath.Dir(lastDir)
		if newDir == "/" || newDir == lastDir {
			pathCache[containing] = ""
			return ""
		}
		lastDir = newDir
	}
}

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
