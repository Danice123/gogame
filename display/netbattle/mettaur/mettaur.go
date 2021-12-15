package mettaur

import (
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

	aiTimer  uint64
	ignoreAI bool

	idle           string
	idleFrame      int
	animation      mettaurAnimation
	animationFrame int

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
	ths.sprites.Batch.Clear()
	ths.genericSprites.Batch.Clear()

	if ths.animation == MOVE_ANIMATION {
		ths.genericSprites.DrawFrame(string(ths.animation), ths.animationFrame-1, x, y+25)
		ths.genericSprites.Batch.Draw(canvas)
	} else {
		if ths.animation != NONE {
			ths.sprites.DrawFrame(string(ths.animation), ths.animationFrame-1, x, y)
		} else {
			ths.sprites.DrawFrame(ths.idle, ths.idleFrame-1, x, y)
		}

		ths.sprites.Batch.Draw(canvas)
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
		ths.idle = string(RAISE_ANIMATION)
		ths.idleFrame = ths.sprites.FrameLength(string(RAISE_ANIMATION)) - 1

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
