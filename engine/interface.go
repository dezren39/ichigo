package engine

import (
	"image"
	"io/fs"
	"reflect"

	"drjosh.dev/gurgle/geom"
	"github.com/hajimehoshi/ebiten/v2"
)

// Reflection types used for queries... Is there a better way?
var (
	// TypeOf(pointer to interface).Elem() is "idiomatic" -
	// see https://pkg.go.dev/reflect#example-TypeOf
	BoundingRecterType = reflect.TypeOf((*BoundingRecter)(nil)).Elem()
	BoundingBoxerType  = reflect.TypeOf((*BoundingBoxer)(nil)).Elem()
	ColliderType       = reflect.TypeOf((*Collider)(nil)).Elem()
	DisablerType       = reflect.TypeOf((*Disabler)(nil)).Elem()
	DrawerType         = reflect.TypeOf((*Drawer)(nil)).Elem()
	HiderType          = reflect.TypeOf((*Hider)(nil)).Elem()
	IdentifierType     = reflect.TypeOf((*Identifier)(nil)).Elem()
	LoaderType         = reflect.TypeOf((*Loader)(nil)).Elem()
	PrepperType        = reflect.TypeOf((*Prepper)(nil)).Elem()
	ScannerType        = reflect.TypeOf((*Scanner)(nil)).Elem()
	SaverType          = reflect.TypeOf((*Saver)(nil)).Elem()
	TransformerType    = reflect.TypeOf((*Transformer)(nil)).Elem()
	UpdaterType        = reflect.TypeOf((*Updater)(nil)).Elem()

	// Behaviours lists all the behaviours that can be queried with Game.Query.
	Behaviours = []reflect.Type{
		BoundingRecterType,
		BoundingBoxerType,
		ColliderType,
		DisablerType,
		DrawerType,
		HiderType,
		IdentifierType,
		LoaderType,
		PrepperType,
		ScannerType,
		SaverType,
		TransformerType,
		UpdaterType,
	}
)

// BoundingRecter components have a bounding rectangle.
type BoundingRecter interface {
	BoundingRect() image.Rectangle
}

// BoundingBoxer components have a bounding box.
type BoundingBoxer interface {
	BoundingBox() geom.Box
}

// Collider components have tangible form.
type Collider interface {
	CollidesWith(geom.Box) bool
}

// Disabler components can be disabled.
type Disabler interface {
	Disabled() bool
	Disable()
	Enable()
}

// Drawer components can draw themselves. Draw is called often. Each component
// must call Draw on any internal components not known to the engine (i.e. not
// passed to Game.Register or returned from Scan).
type Drawer interface {
	Draw(*ebiten.Image, *ebiten.DrawImageOptions)
	DrawAfter(Drawer) bool
	DrawBefore(Drawer) bool
}

// Hider components can be hidden.
type Hider interface {
	Hidden() bool
	Hide()
	Show()
}

// Identifier components have a sense of self. This makes it easier for
// components to find and interact with one another. Returning the empty string
// is treated as having no identifier.
type Identifier interface {
	Ident() string
}

// Loader components get the chance to load themselves. This happens
// before preparation.
type Loader interface {
	Load(fs.FS) error
}

// Prepper components can be prepared. It is called after the component
// database has been populated but before the game is run. The component can
// store the reference to game, if needed, and also query the component database.
type Prepper interface {
	Prepare(game *Game) error
}

// Scanner components can be scanned. It is called when the game tree is walked
// (such as when the game component database is constructed).
// Scan should return a slice containing all immediate subcomponents.
type Scanner interface {
	Scan() []interface{}
}

// Saver components can be saved to disk.
type Saver interface {
	Save() error
}

// Transformer components can provide draw options to apply to themselves and
// any child components. The opts passed to Draw of a component c will be the
// cumulative opts of all parents of c plus the value returned from c.Transform.
type Transformer interface {
	Transform() ebiten.DrawImageOptions
}

// Updater components can update themselves. Update is called repeatedly. Each
// component must call Update on any internal components not known to the engine
//  (i.e. not passed to Game.Register or returned from Scan).
type Updater interface {
	Update() error
}

// ZPositioner components opt into a simpler method of determining draw order
// than DrawAfter/DrawBefore. Drawer implementations should handle the
// ZPositioner case as though the component were a flat infinite plane given by
// z = ZPos().
type ZPositioner interface {
	ZPos() int
}
