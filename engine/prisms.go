package engine

import (
	"encoding/gob"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	_ interface {
		Prepper
		Transformer
	} = &PrismMap{}

	_ interface {
		Drawer
		Transformer
	} = &Prism{}
)

func init() {
	gob.Register(&PrismMap{})
	gob.Register(&Prism{})
}

type PrismMap struct {
	Map           map[Point3]*Prism
	DrawOrderBias image.Point // dot with (X,Y) = bias
	DrawOffset    image.Point // offset to apply to whole map
	DrawZStride   image.Point // draw offset for each unit in Z
	PrismSize     Point3      // cmul map key = world pos
	Sheet         Sheet
}

func (m *PrismMap) Prepare(*Game) error {
	for v, p := range m.Map {
		p.pos = v
		p.pm = m
	}
	return nil
}

func (m *PrismMap) Transform(pt Transform) (tf Transform) {
	tf.Opts.GeoM.Translate(cfloat(m.DrawOffset))
	return tf.Concat(pt)
}

type Prism struct {
	Cell int

	pos Point3
	pm  *PrismMap
}

func (p *Prism) Draw(screen *ebiten.Image, opts *ebiten.DrawImageOptions) {
	screen.DrawImage(p.pm.Sheet.SubImage(p.Cell), opts)
}

func (p *Prism) DrawOrder() (int, int) {
	return p.pos.Z * p.pm.PrismSize.Z,
		dot(p.pos.XY(), p.pm.DrawOrderBias)
}

func (p *Prism) Transform(pt Transform) (tf Transform) {
	tf.Opts.GeoM.Translate(cfloat(
		cmul(p.pos.XY(), p.pm.PrismSize.XY()).
			Add(p.pm.DrawZStride.Mul(p.pos.Z)),
	))
	return tf.Concat(pt)
}
