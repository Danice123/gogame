package entity

import (
	"os"
	"path/filepath"

	"github.com/Danice123/idk/display/texturepacker"
	"github.com/Danice123/idk/logic"
	"gopkg.in/yaml.v3"
)

type EntityData struct {
	Name   string
	Coord  logic.Coord
	Facing logic.Direction
	Script string
}

type EntityMapData struct {
	Entities []EntityData
}

func LoadEntityMapData(mapName string) *EntityMapData {
	mapData := &EntityMapData{}
	if data, err := os.ReadFile(filepath.Join("maps", mapName+".entity.yaml")); err != nil {
		println(err.Error())
	} else {
		if err := yaml.Unmarshal(data, mapData); err != nil {
			println(err.Error())
		}
	}
	return mapData
}

func (ths *EntityMapData) Build(spritesheet *texturepacker.SpriteSheet) []Entity {
	entities := []Entity{}
	for _, data := range ths.Entities {
		entities = append(entities, &Base{
			EntityName: data.Name,
			Coord:      data.Coord,
			Facing:     data.Facing,
			script:     data.Script,

			Spritesheet: spritesheet,
		})
	}
	return entities
}
