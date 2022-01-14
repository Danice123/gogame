package netbattle

import (
	"path/filepath"

	"github.com/Danice123/gogame/display/netbattle/field"
	"github.com/Danice123/gogame/display/netbattle/mettaur"
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
	field          *field.BattleField
	sprites        *texturepacker.SpriteSheet
	coord          utils.Coord
	delay          *utils.DelayHandler
	animation      playerAnimation
	animationFrame int

	inputHandler utils.InputHandler

	chargingAnimation      playerAnimation
	chargingAnimationFrame int
	busterCharge           uint64
	isCharging             bool

	Health int
}

func NewPlayer(field *field.BattleField) *Player {
	p := &Player{
		field:   field,
		sprites: texturepacker.NewSpriteSheet(filepath.Join("resources", "sheets", "rockman.json")),
		delay:   utils.NewDelayHandler(),
		coord: utils.Coord{
			X: 1,
			Y: 1,
		},
		Health:                 100,
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
	field.RegisterObject(p)
	return p
}

func (ths *Player) Coord() utils.Coord {
	return ths.coord
}

func (ths *Player) HighlightTile() bool {
	return false
}

func (ths *Player) AI(utils.Coord) {}

func (ths *Player) Tick() {
	ths.inputHandler.Tick()
	ths.delay.Tick()
	ths.busterCharge++

	if ths.animation != NONE {
		if ths.animationFrame == ths.sprites.FrameLength(string(ths.animation)) {
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
		} else if ths.chargingAnimationFrame == ths.sprites.FrameLength(string(ths.chargingAnimation)) {
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
	x += 3
	y += 5
	ths.sprites.Clear()

	sprite := "rockman-idle"
	if ths.animation != NONE {
		sprite = string(ths.animation)
	}
	ths.sprites.DrawFrame(sprite, ths.animationFrame-1, x, y)

	if ths.isCharging && ths.busterCharge > 20 {
		ths.sprites.DrawFrame(string(ths.chargingAnimation), ths.chargingAnimationFrame-1, x+20, y+20)
	}

	ths.sprites.Render(canvas)
}

func (ths *Player) up() uint64 {
	if ths.coord.Y > 0 {
		ths.animation = MOVE_ANIMATION
		ths.animationFrame = 1
		ths.delay.AddDelayedAction(4, func() {
			ths.coord.Y--
		})
		return 10
	}
	return 0
}

func (ths *Player) down() uint64 {
	if ths.coord.Y < 2 {
		ths.animation = MOVE_ANIMATION
		ths.animationFrame = 1
		ths.delay.AddDelayedAction(4, func() {
			ths.coord.Y++
		})
		return 10
	}
	return 0
}

func (ths *Player) left() uint64 {
	if ths.coord.X > 0 {
		ths.animation = MOVE_ANIMATION
		ths.animationFrame = 1
		ths.delay.AddDelayedAction(4, func() {
			ths.coord.X--
		})
		return 10
	}
	return 0
}

func (ths *Player) right() uint64 {
	if ths.coord.X < 2 {
		ths.animation = MOVE_ANIMATION
		ths.animationFrame = 1
		ths.delay.AddDelayedAction(4, func() {
			ths.coord.X++
		})
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

	for i := ths.coord.X + 1; i < 6; i++ {
		hits := ths.field.HitReg(utils.Coord{
			X: i,
			Y: ths.coord.Y,
		})

		for _, hit := range hits {
			if ths.busterCharge > 120 {
				hit.Damage(10, string(mettaur.BUSTER_PURPLE_HIT_EFFECT))
			} else if ths.busterCharge > 60 {
				hit.Damage(5, string(mettaur.BUSTER_GREEN_HIT_EFFECT))
			} else {
				hit.Damage(1, string(mettaur.BUSTER_HIT_EFFECT))
			}
		}
		if len(hits) > 0 {
			break
		}
	}
	ths.busterCharge = 0
	return 15
}

func (ths *Player) Damage(amount int, hitEffect string) {
	if ths.Health-amount < 0 {
		ths.Health = 0
	} else {
		ths.Health -= amount
	}
}
