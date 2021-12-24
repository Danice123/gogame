package netbattle

import (
	"path/filepath"

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
	sprites        *texturepacker.SpriteSheet
	delay          *utils.DelayHandler
	animation      playerAnimation
	animationFrame int

	inputHandler utils.InputHandler

	chargingAnimation      playerAnimation
	chargingAnimationFrame int
	busterCharge           uint64
	isCharging             bool

	hitreg func(utils.Coord) *mettaur.Mettaur

	Coord utils.Coord
}

func NewPlayer(hitreg func(utils.Coord) *mettaur.Mettaur) *Player {
	p := &Player{
		sprites: texturepacker.NewSpriteSheet(filepath.Join("resources", "sheets", "rockman.json")),
		delay:   utils.NewDelayHandler(),
		Coord: utils.Coord{
			X: 1,
			Y: 1,
		},
		animationFrame:         1,
		chargingAnimationFrame: 1,
		hitreg:                 hitreg,
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
	if ths.Coord.Y > 0 {
		ths.animation = MOVE_ANIMATION
		ths.animationFrame = 1
		ths.delay.AddDelayedAction(4, func() {
			ths.Coord.Y--
		})
		return 10
	}
	return 0
}

func (ths *Player) down() uint64 {
	if ths.Coord.Y < 2 {
		ths.animation = MOVE_ANIMATION
		ths.animationFrame = 1
		ths.delay.AddDelayedAction(4, func() {
			ths.Coord.Y++
		})
		return 10
	}
	return 0
}

func (ths *Player) left() uint64 {
	if ths.Coord.X > 0 {
		ths.animation = MOVE_ANIMATION
		ths.animationFrame = 1
		ths.delay.AddDelayedAction(4, func() {
			ths.Coord.X--
		})
		return 10
	}
	return 0
}

func (ths *Player) right() uint64 {
	if ths.Coord.X < 2 {
		ths.animation = MOVE_ANIMATION
		ths.animationFrame = 1
		ths.delay.AddDelayedAction(4, func() {
			ths.Coord.X++
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

	for i := ths.Coord.X; i < 6; i++ {
		hit := ths.hitreg(utils.Coord{
			X: i,
			Y: ths.Coord.Y,
		})

		if hit != nil {
			if ths.busterCharge > 120 {
				hit.Damage(10, mettaur.BUSTER_PURPLE_HIT_EFFECT)
			} else if ths.busterCharge > 60 {
				hit.Damage(5, mettaur.BUSTER_GREEN_HIT_EFFECT)
			} else {
				hit.Damage(1, mettaur.BUSTER_HIT_EFFECT)
			}
			break
		}
	}
	ths.busterCharge = 0
	return 15
}
