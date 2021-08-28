package engine

import (
	"encoding/gob"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Ensure Camera satisfies interfaces.
var _ interface {
	Identifier
	Prepper
} = &Camera{}

func init() {
	gob.Register(&Camera{})
}

// Camera models a camera that is viewing a scene. (Camera is a child of the
// scene it is viewing, for various reasons.) Changes to the fields take effect
// immediately.
type Camera struct {
	ID

	// Camera controls
	Centre image.Point // world coordinates
	Filter ebiten.Filter
	Zoom   float64 // unitless

	game *Game
}

// Prepare grabs a copy of game (needed for screen dimensions)
func (c *Camera) Prepare(game *Game) error {
	c.game = game
	return nil
}
