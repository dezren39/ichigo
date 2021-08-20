package engine

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Collider components have tangible form.
type Collider interface {
	CollidesWith(image.Rectangle) bool
}

// Drawer components can draw themselves. Draw is called often.
// Each component is responsible for calling Draw on its child components
// (so that hiding the parent can hide the children, etc).
type Drawer interface {
	Draw(screen *ebiten.Image, opts ebiten.DrawImageOptions)
}

// DrawOrderer components can provide a number for determining draw order.
type DrawOrderer interface {
	DrawOrder() float64
}

// Identifier components have a sense of self. This makes it easier for
// components to find and interact with one another.
type Identifier interface {
	Ident() string
}

// Loader components get the chance to load themselves. This happens
// before preparation.
type Loader interface {
	Load() error
}

// ParallaxScaler components have a scaling factor. This is used for
// parallax layers in a scene, and can be thought of as 1/distance.
type ParallaxScaler interface {
	ParallaxFactor() float64
}

// Prepper components can be prepared. It is called after the component
// database has been populated but before the game is run. The component can
// store the reference to game, if needed, and also query the component database.
type Prepper interface {
	Prepare(game *Game)
}

// Scanner components can be scanned. It is called when the game tree is walked
// (such as when the game component database is constructed).
// Scan should return a slice containing all immediate subcomponents.
type Scanner interface {
	Scan() []interface{}
}

// Scener components are a scene (Scene or SceneRef).
type Scener interface {
	// Q: Why not make Scener a small interface with just Scene() ?
	// A: Everything in the engine would then need to type switch for Scener or SceneRef, i.e.
	// switch x := i.(type) {
	// case Drawer:
	//     i.Draw(screen, opts)
	// case Scener:
	//     i.Scene().Draw(screen, opts)
	// }
	// It seems cleaner to let the engine assert only for the interface it needs at that moment.

	Drawer
	DrawOrderer
	Identifier
	Prepper
	Scanner
	Updater

	Scene() *Scene
}

// Updater components can update themselves. Update is called repeatedly.
// Each component is responsible for calling Update on its child components
// (so that disabling the parent prevents updates to the children, etc).
type Updater interface {
	Update() error
}
