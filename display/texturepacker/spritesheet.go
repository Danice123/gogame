package texturepacker

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Danice123/gogame/display/utils"
	"github.com/faiface/pixel"
)

type Pivot struct {
	X float64
	Y float64
}

type Size struct {
	W int
	H int
}

type Bounds struct {
	utils.Coord
	Size
}

type Texture struct {
	Frame            Bounds
	Rotated          bool
	Trimmed          bool
	SpriteSourceSize Bounds
	SourceSize       Size
	Pivot            Pivot
	Frametime        int
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
	Name  string
	Batch *pixel.Batch

	data            *PackedTextures
	sprites         map[string][]*pixel.Sprite
	animationFrames map[string][]int
}

func breakDownDescriptor(descriptor string) (string, int) {
	if strings.ContainsRune(descriptor, '#') {
		s := strings.Split(descriptor, "#")
		frameId, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		return s[0], frameId
	} else {
		return descriptor, 1
	}
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
		data: data,
	}
	image := utils.LoadPicture(filepath.Join(filepath.Dir(path), data.Meta.Image))
	sheet.Batch = pixel.NewBatch(&pixel.TrianglesData{}, image)

	frameSize := map[string]int{}
	for descriptor := range data.Frames {
		name, frameId := breakDownDescriptor(descriptor)
		if size, ok := frameSize[name]; !ok {
			frameSize[name] = frameId
		} else if frameId > size {
			frameSize[name] = frameId
		}
	}

	sheet.sprites = map[string][]*pixel.Sprite{}
	sheet.animationFrames = map[string][]int{}
	for name, size := range frameSize {
		sheet.sprites[name] = make([]*pixel.Sprite, size)

		sheet.animationFrames[name] = []int{}
		if size == 1 {
			sheet.animationFrames[name] = append(sheet.animationFrames[name], 0)
			continue
		}
		for i := 0; i < size; i++ {
			spriteFrameTime := data.Frames[fmt.Sprintf("%s#%d", name, i+1)].Frametime
			if spriteFrameTime == 0 {
				spriteFrameTime = 1
			}
			for j := 0; j < spriteFrameTime; j++ {
				sheet.animationFrames[name] = append(sheet.animationFrames[name], i)
			}
		}
	}

	for descriptor, sprite := range data.Frames {
		name, frameId := breakDownDescriptor(descriptor)
		s := pixel.NewSprite(image, pixel.R(float64(sprite.Frame.X), float64(data.Meta.Size.H-sprite.Frame.Y), float64(sprite.Frame.X+sprite.Frame.W), float64(data.Meta.Size.H-sprite.Frame.Y-sprite.Frame.H)))
		sheet.sprites[name][frameId-1] = s
	}

	return sheet
}

func (ths *SpriteSheet) Draw(name string, x int, y int) {
	ths.DrawFrame(name, 0, x, y)
}

func (ths *SpriteSheet) DrawFrame(name string, frame int, x int, y int) {
	if sprite, ok := ths.sprites[name]; ok {
		frame = ths.animationFrames[name][frame]
		if frame >= len(sprite) {
			fmt.Fprintf(os.Stderr, "Bad sprite ref: %s, frame: %d", name, frame)
		}

		spriteData, validName := ths.data.Frames[fmt.Sprintf("%s#%d", name, frame+1)]
		if !validName && frame == 0 {
			spriteData = ths.data.Frames[name]
		}

		spriteBounds := sprite[frame].Frame().Norm()
		xOffset := spriteBounds.W()/2 - spriteData.Pivot.X*spriteBounds.W()
		yOffset := -spriteBounds.H()/2 + spriteData.Pivot.Y*spriteBounds.H()
		sprite[frame].Draw(ths.Batch, pixel.IM.Moved(pixel.V(xOffset, yOffset)).Moved(pixel.V(float64(x), float64(y))))
	} else {
		fmt.Fprintf(os.Stderr, "Bad sprite ref: %s", name)
	}
}

func (ths *SpriteSheet) FrameLength(name string) int {
	return len(ths.animationFrames[name])
}
