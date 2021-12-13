package netbattle

import (
	"path/filepath"

	"github.com/Danice123/gogame/display/texturepacker"
	"github.com/Danice123/gogame/display/utils"
	"github.com/faiface/pixel/pixelgl"
)

type playerAnimation string

const NONE playerAnimation = ""
const MOVE_ANIMATION playerAnimation = "rockman-move"
const BUSTER_ANIMATION playerAnimation = "rockman-buster"

const CHARGING_BLUE_ANIMATION = "rockman-buster-charging-blue"
const CHARGING_GREEN_ANIMATION = "rockman-buster-charging-green"
const CHARGING_PURPLE_ANIMATION = "rockman-buster-charging-purple"

type Player struct {
	sprites        *texturepacker.SpriteSheet
	animation      playerAnimation
	animationFrame int

	inputHandler utils.InputHandler

	chargingAnimation      playerAnimation
	chargingAnimationFrame int
	busterCharge           uint64
	isCharging             bool

	Coord utils.Coord
}

func NewPlayer() *Player {
	p := &Player{
		sprites: texturepacker.NewSpriteSheet(filepath.Join("resources", "sheets", "rockman.json")),
		Coord: utils.Coord{
			X: 1,
			Y: 1,
		},
		animationFrame:         1,
		chargingAnimationFrame: 1,
	}

	p.inputHandler = utils.InputHandler{
		PressHandlers: map[utils.KEY]func() uint64{
			utils.UP:      p.up,
			utils.DOWN:    p.down,
			utils.LEFT:    p.left,
			utils.RIGHT:   p.right,
			utils.DECLINE: p.charge,
		},
		ReleaseHandlers: map[utils.KEY]func() uint64{
			utils.DECLINE: p.shoot,
		},
	}

	return p
}

func (ths *Player) Tick(delta int64) {
	ths.inputHandler.Tick()
	ths.busterCharge++

	if ths.animation != NONE {
		if ths.animationFrame == len(ths.sprites.Sprites[string(ths.animation)]) {
			ths.animation = NONE
			ths.animationFrame = 1
		} else {
			ths.animationFrame++
		}
	}

	if ths.isCharging {
		if ths.busterCharge > 60 && ths.chargingAnimation == CHARGING_BLUE_ANIMATION {
			ths.chargingAnimation = CHARGING_GREEN_ANIMATION
			ths.chargingAnimationFrame = 1
		} else if ths.busterCharge > 120 && ths.chargingAnimation == CHARGING_GREEN_ANIMATION {
			ths.chargingAnimation = CHARGING_PURPLE_ANIMATION
			ths.chargingAnimationFrame = 1
		} else if ths.chargingAnimationFrame == len(ths.sprites.Sprites[string(ths.chargingAnimation)]) {
			ths.chargingAnimationFrame = 1
		} else {
			ths.chargingAnimationFrame++
		}
	}
}

func (ths *Player) HandleKey(pressed func(utils.KEY) bool) {
	ths.inputHandler.HandleKey(pressed)
}

func (ths *Player) Render(canvas *pixelgl.Canvas, x int, y int) {
	ths.sprites.Batch.Clear()

	sprite := "rockman-idle"
	if ths.animation != NONE {
		sprite = string(ths.animation)
	}
	ths.sprites.DrawFrame(sprite, ths.animationFrame-1, x, y)

	if ths.isCharging && ths.busterCharge > 20 {
		ths.sprites.DrawFrame(string(ths.chargingAnimation), ths.chargingAnimationFrame-1, x+20, y+20)
	}

	ths.sprites.Batch.Draw(canvas)
}

func (ths *Player) up() uint64 {
	if ths.Coord.Y > 0 {
		ths.Coord.Y--
		ths.animation = MOVE_ANIMATION
		ths.animationFrame = 1
		return 10
	}
	return 0
}

func (ths *Player) down() uint64 {
	if ths.Coord.Y < 2 {
		ths.Coord.Y++
		ths.animation = MOVE_ANIMATION
		ths.animationFrame = 1
		return 10
	}
	return 0
}

func (ths *Player) left() uint64 {
	if ths.Coord.X > 0 {
		ths.Coord.X--
		ths.animation = MOVE_ANIMATION
		ths.animationFrame = 1
		return 10
	}
	return 0
}

func (ths *Player) right() uint64 {
	if ths.Coord.X < 2 {
		ths.Coord.X++
		ths.animation = MOVE_ANIMATION
		ths.animationFrame = 1
		return 10
	}
	return 0
}

func (ths *Player) charge() uint64 {
	if !ths.isCharging {
		ths.busterCharge = 0
		ths.isCharging = true
		ths.chargingAnimation = CHARGING_BLUE_ANIMATION
		ths.chargingAnimationFrame = 1
	}
	return 0
}

func (ths *Player) shoot() uint64 {
	ths.isCharging = false
	ths.animation = BUSTER_ANIMATION
	ths.animationFrame = 1
	return 15
}
