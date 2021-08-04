package engine

import (
	"encoding/gob"
	"image"
	"math"
)

const gravity = 0.5

func init() {
	gob.Register(Actor{})
}

// Thorson-style movement:
// https://maddythorson.medium.com/celeste-and-towerfall-physics-d24bd2ae0fc5

// Collider components have tangible form.
type Collider interface {
	CollidesWith(image.Rectangle) bool
}

// Actor handles basic movement.
type Actor struct {
	Position image.Point
	Size     image.Point

	game       *Game
	xRem, yRem float64
}

func (a *Actor) CollidesAt(p image.Point) bool {
	// TODO: more efficient test?
	hit := false
	Walk(a.game, func(c interface{}) bool {
		if coll, ok := c.(Collider); ok {
			if coll.CollidesWith(image.Rectangle{Min: p, Max: p.Add(a.Size)}) {
				hit = true
				return false
			}
		}
		return true
	})
	return hit
}

func (a *Actor) MoveX(dx float64, onCollide func()) {
	a.xRem += dx
	move := int(math.Round(a.xRem))
	if move == 0 {
		return
	}
	a.xRem -= float64(move)
	sign := sign(move)
	for move != 0 {
		if a.CollidesAt(a.Position.Add(image.Pt(sign, 0))) {
			if onCollide != nil {
				onCollide()
			}
			return
		}
		a.Position.X += sign
		move -= sign
	}
}

func (a *Actor) MoveY(dy float64, onCollide func()) {
	a.yRem += dy
	move := int(math.Round(a.yRem))
	if move == 0 {
		return
	}
	a.yRem -= float64(move)
	sign := sign(move)
	for move != 0 {
		if a.CollidesAt(a.Position.Add(image.Pt(0, sign))) {
			if onCollide != nil {
				onCollide()
			}
			return
		}
		a.Position.Y += sign
		move -= sign
	}
}

func (a *Actor) Build(g *Game) {
	a.game = g
}

func sign(m int) int {
	if m < 0 {
		return -1
	}
	return 1
}
