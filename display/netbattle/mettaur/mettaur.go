package mettaur

import (
	"image/color"
	"path/filepath"

	"github.com/Danice123/gogame/display/netbattle/state"
	"github.com/Danice123/gogame/display/texturepacker"
	"github.com/Danice123/gogame/display/utils"
	"github.com/faiface/pixel/pixelgl"
)

type mettaurAnimation string

const NONE mettaurAnimation = ""
const MOVE_ANIMATION mettaurAnimation = "generic-move-small"
const RAISE_ANIMATION mettaurAnimation = "mettaur-raise"
const ATTACK_ANIMATION mettaurAnimation = "mettaur-attack"
const WITHDRAW_ANIMATION mettaurAnimation = "mettaur-withdraw"

type Mettaur struct {
	sprites        *texturepacker.SpriteSheet
	genericSprites *texturepacker.SpriteSheet
	delay          *utils.DelayHandler

	health   int
	aiTimer  uint64
	ignoreAI bool

	idle           string
	idleFrame      int
	animation      mettaurAnimation
	animationFrame int
	flash          bool

	Coord utils.Coord
}

func NewMettaur() *Mettaur {
	return &Mettaur{
		sprites:        texturepacker.NewSpriteSheet(filepath.Join("resources", "sheets", "mettaur.json")),
		genericSprites: texturepacker.NewSpriteSheet(filepath.Join("resources", "sheets", "generic.json")),
		delay:          utils.NewDelayHandler(),
		Coord: utils.Coord{
			X: 4,
			Y: 1,
		},
		idle:           "mettaur-idle",
		idleFrame:      1,
		animationFrame: 1,
		health:         40,
	}
}

func (ths *Mettaur) Tick(delta int64) {
	ths.delay.Tick()

	if ths.animation != NONE {
		animationLength := ths.sprites.FrameLength(string(ths.animation))
		if ths.animation == MOVE_ANIMATION {
			animationLength = ths.genericSprites.FrameLength(string(ths.animation))
		}
		if ths.animationFrame == animationLength {
			ths.animation = NONE
			ths.animationFrame = 1
		} else {
			ths.animationFrame++
		}
	}
}

func (ths *Mettaur) Render(canvas *pixelgl.Canvas, x int, y int) {
	ths.sprites.Clear()
	ths.genericSprites.Clear()

	var damageMask color.Color
	if ths.flash {
		damageMask = color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 0,
		}
	}

	if ths.animation == MOVE_ANIMATION {
		ths.genericSprites.DrawFrame(string(ths.animation), ths.animationFrame-1, x+12, y+12)
		ths.genericSprites.Render(canvas)
	} else {
		if ths.animation != NONE {
			ths.sprites.DrawFrameWithMask(string(ths.animation), ths.animationFrame-1, x, y, damageMask)
		} else {
			ths.sprites.DrawFrameWithMask(ths.idle, ths.idleFrame-1, x, y, damageMask)
		}

		ths.sprites.Render(canvas)

		ths.genericSprites.DrawSpriteNumber("generic-health", ths.health, x+3, y+23)
		ths.genericSprites.Render(canvas)
	}
}

func (ths *Mettaur) AI(state state.BoardState) {
	ths.aiTimer++
	if ths.ignoreAI {
		return
	}
	if ths.aiTimer%60 == 0 {
		if ths.Coord.Y > state.PlayerCoord.Y {
			ths.up()
		} else if ths.Coord.Y < state.PlayerCoord.Y {
			ths.down()
		} else {
			ths.raise()
		}
	}
}

func (ths *Mettaur) Damage(amount int) {
	if ths.health-amount < 0 {
		ths.health = 0
	} else {
		ths.health -= amount
	}

	ths.flash = true
	ths.delay.AddDelayedAction(2, func() {
		ths.flash = false
	})
}

func (ths *Mettaur) up() {
	if ths.Coord.Y > 0 {
		ths.animation = MOVE_ANIMATION
		ths.animationFrame = 1
		ths.delay.AddDelayedAction(4, func() {
			ths.Coord.Y--
		})
	}
}

func (ths *Mettaur) down() {
	if ths.Coord.Y < 2 {
		ths.animation = MOVE_ANIMATION
		ths.animationFrame = 1
		ths.delay.AddDelayedAction(4, func() {
			ths.Coord.Y++
		})
	}
}

func (ths *Mettaur) raise() {
	ths.ignoreAI = true
	ths.animation = RAISE_ANIMATION
	ths.animationFrame = 1
	ths.delay.AddDelayedAction(ths.sprites.FrameLength(string(RAISE_ANIMATION))-1, func() {
		ths.idle = string(ATTACK_ANIMATION)
		ths.idleFrame = 1

		ths.delay.AddDelayedAction(10, func() {
			ths.attack()
		})
	})
}

func (ths *Mettaur) attack() {
	ths.animation = ATTACK_ANIMATION
	ths.animationFrame = 1
	ths.delay.AddDelayedAction(ths.sprites.FrameLength(string(ATTACK_ANIMATION))-1, func() {
		ths.idle = string(ATTACK_ANIMATION)
		ths.idleFrame = ths.sprites.FrameLength(string(ATTACK_ANIMATION)) - 1
		ths.withdraw()
	})
}

func (ths *Mettaur) withdraw() {
	ths.animation = WITHDRAW_ANIMATION
	ths.animationFrame = 1
	ths.delay.AddDelayedAction(ths.sprites.FrameLength(string(WITHDRAW_ANIMATION))-1, func() {
		ths.idle = "mettaur-idle"
		ths.idleFrame = 1
		ths.ignoreAI = false
	})
}
