package mettaur

import (
	"image/color"
	"math/rand"
	"path/filepath"

	"github.com/Danice123/gogame/display/netbattle/field"
	"github.com/Danice123/gogame/display/texturepacker"
	"github.com/Danice123/gogame/display/utils"
	"github.com/faiface/pixel/pixelgl"
)

type MettaurAnimation string

const NONE MettaurAnimation = ""
const MOVE_ANIMATION MettaurAnimation = "generic-move-small"
const RAISE_ANIMATION MettaurAnimation = "mettaur-raise"
const ATTACK_ANIMATION MettaurAnimation = "mettaur-attack"
const WITHDRAW_ANIMATION MettaurAnimation = "mettaur-withdraw"

const BUSTER_HIT_EFFECT MettaurAnimation = "generic-effect-buster"
const BUSTER_GREEN_HIT_EFFECT MettaurAnimation = "generic-effect-buster-green"
const BUSTER_PURPLE_HIT_EFFECT MettaurAnimation = "generic-effect-buster-purple"

type Mettaur struct {
	field          *field.BattleField
	sprites        *texturepacker.SpriteSheet
	genericSprites *texturepacker.SpriteSheet
	coord          utils.Coord
	delay          *utils.DelayHandler

	health   int
	aiTimer  uint64
	ignoreAI bool
	dead     bool

	idle           string
	idleFrame      int
	animation      MettaurAnimation
	animationFrame int

	flash          bool
	effect         MettaurAnimation
	effectFrame    int
	effectXOffset  int
	effectYOffset  int
	exploding      bool
	explodingFrame int
}

func NewMettaur(field *field.BattleField) *Mettaur {
	m := &Mettaur{
		field:          field,
		sprites:        texturepacker.NewSpriteSheet(filepath.Join("resources", "sheets", "mettaur.json")),
		genericSprites: texturepacker.NewSpriteSheet(filepath.Join("resources", "sheets", "generic.json")),
		delay:          utils.NewDelayHandler(),
		coord: utils.Coord{
			X: 4,
			Y: 1,
		},
		idle:           "mettaur-idle",
		idleFrame:      1,
		animationFrame: 1,
		effectFrame:    1,
		health:         40,
	}
	field.RegisterObject(m)
	return m
}

func (ths *Mettaur) Coord() utils.Coord {
	return ths.coord
}

func (ths *Mettaur) HighlightTile() bool {
	return false
}

func (ths *Mettaur) Tick() {
	ths.delay.Tick()

	if ths.animation != NONE && !ths.exploding {
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

	if ths.effect != NONE {
		animationLength := ths.genericSprites.FrameLength(string(ths.effect))
		if ths.effectFrame == animationLength {
			ths.effect = NONE
			ths.effectFrame = 1
		} else {
			ths.effectFrame++
		}
	}

	if ths.exploding {
		ths.explodingFrame++
		ths.flash = ths.explodingFrame%4 > 1
	}
}

func (ths *Mettaur) Render(canvas *pixelgl.Canvas, x int, y int) {
	if ths.dead {
		return
	}

	x += 9
	y += 5
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

		if ths.exploding {
			if ths.explodingFrame < ths.genericSprites.FrameLength("generic-explosion") {
				ths.genericSprites.DrawFrame("generic-explosion", ths.explodingFrame, x+7, y-5)
			}
			if ths.explodingFrame-4 > 0 && ths.explodingFrame-4 < ths.genericSprites.FrameLength("generic-explosion") {
				ths.genericSprites.DrawFrame("generic-explosion", ths.explodingFrame-4, x+10, y-5)
			}
			if ths.explodingFrame-8 > 0 && ths.explodingFrame-8 < ths.genericSprites.FrameLength("generic-explosion") {
				ths.genericSprites.DrawFrame("generic-explosion", ths.explodingFrame-8, x+13, y-5)
			}

		} else {
			ths.genericSprites.DrawSpriteNumber("generic-health", ths.health, x+3, y+23)
		}

		if ths.effect != NONE {
			ths.genericSprites.DrawFrame(string(ths.effect), ths.effectFrame-1, x+ths.effectXOffset, y+25+ths.effectYOffset)
		}
		ths.genericSprites.Render(canvas)
	}
}

func (ths *Mettaur) AI(playerCoord utils.Coord) {
	ths.aiTimer++
	if ths.ignoreAI {
		return
	}
	if ths.aiTimer%60 == 0 {
		if ths.coord.Y > playerCoord.Y {
			ths.up()
		} else if ths.coord.Y < playerCoord.Y {
			ths.down()
		} else {
			ths.raise()
		}
	}
}

func (ths *Mettaur) Damage(amount int, hitEffect string) {
	if ths.health-amount < 0 {
		ths.health = 0
	} else {
		ths.health -= amount
	}

	ths.flash = true
	ths.effect = MettaurAnimation(hitEffect)
	ths.effectFrame = 1
	ths.effectXOffset = rand.Intn(20)
	ths.effectYOffset = rand.Intn(10) - 5
	ths.delay.AddDelayedAction(2, func() {
		ths.flash = false
		if ths.health == 0 {
			ths.death()
		}
	})
}

func (ths *Mettaur) death() {
	ths.explodingFrame = 1
	ths.exploding = true
	ths.ignoreAI = true
	ths.delay.CancelAll()
	ths.delay.AddDelayedAction(20, func() {
		ths.dead = true
	})
}

func (ths *Mettaur) up() {
	if ths.coord.Y > 0 {
		ths.animation = MOVE_ANIMATION
		ths.animationFrame = 1
		ths.delay.AddDelayedAction(4, func() {
			ths.coord.Y--
		})
	}
}

func (ths *Mettaur) down() {
	if ths.coord.Y < 2 {
		ths.animation = MOVE_ANIMATION
		ths.animationFrame = 1
		ths.delay.AddDelayedAction(4, func() {
			ths.coord.Y++
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
		NewMettaurAttack(ths.field, utils.Coord{
			X: ths.coord.X - 1,
			Y: ths.coord.Y,
		})
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
