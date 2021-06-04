package texturepacker

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/Danice123/idk/display/utils"
	"github.com/faiface/pixel"
)

type Size struct {
	W int
	H int
}

type Bounds struct {
	X int
	Y int
	W int
	H int
}

type Texture struct {
	Frame            Bounds
	Rotated          bool
	Trimmed          bool
	SpriteSourceSize Bounds
	SourceSize       Size
}

type Meta struct {
	App     string
	Version string
	Image   string
	Format  string
	Size    Size
	Scale   string
}

type PackedTextures struct {
	Frames map[string]Texture
	Meta   Meta
}

type SpriteSheet struct {
	Name    string
	Batch   *pixel.Batch
	Sprites map[string]map[string][]*pixel.Sprite
}

func NewSpriteSheet(path string) *SpriteSheet {
	jsonData, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}

	data := &PackedTextures{}
	err = json.Unmarshal(jsonData, data)
	if err != nil {
		panic(err.Error())
	}

	sheet := &SpriteSheet{
		Name: data.Meta.Image,
	}
	image := utils.LoadPicture(filepath.Join(filepath.Dir(path), data.Meta.Image))
	sheet.Batch = pixel.NewBatch(&pixel.TrianglesData{}, image)
	sheet.Sprites = make(map[string]map[string][]*pixel.Sprite)

	for descriptor, sprite := range data.Frames {
		splitDescriptor := strings.Split(descriptor, "_")
		spriteName := splitDescriptor[0]
		spriteState := splitDescriptor[1]

		s := pixel.NewSprite(image, pixel.R(float64(sprite.Frame.X), float64(sprite.Frame.Y), float64(sprite.Frame.X+sprite.Frame.W), float64(sprite.Frame.Y+sprite.Frame.H)))

		if sheet.Sprites[spriteName] == nil {
			sheet.Sprites[spriteName] = make(map[string][]*pixel.Sprite)
		}
		if sheet.Sprites[spriteName][spriteState] == nil {
			sheet.Sprites[spriteName][spriteState] = []*pixel.Sprite{}
		}
		sheet.Sprites[spriteName][spriteState] = append(sheet.Sprites[spriteName][spriteState], s)
	}

	return sheet
}
