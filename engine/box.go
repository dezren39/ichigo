package engine

import "image"

// Box describes an axis-aligned rectangular prism.
type Box struct {
	Min, Max Int3
}

// String returns a string representation of b like "(3,4,5)-(6,5,8)".
func (b Box) String() string {
	return b.Min.String() + "-" + b.Max.String()
}

// Empty reports whether the box contains no points.
func (b Box) Empty() bool {
	return b.Min.X >= b.Max.X || b.Min.Y >= b.Max.Y || b.Min.Z >= b.Max.Z
}

// Eq reports whether b and c contain the same set of points. All empty boxes
// are considered equal.
func (b Box) Eq(c Box) bool {
	return b == c || b.Empty() && c.Empty()
}

// Overlaps reports whether b and c have non-empty intersection.
func (b Box) Overlaps(c Box) bool {
	return !b.Empty() && !c.Empty() &&
		b.Min.X < c.Max.X && c.Min.X < b.Max.X &&
		b.Min.Y < c.Max.Y && c.Min.Y < b.Max.Y &&
		b.Min.Z < c.Max.Z && c.Min.Z < b.Max.Z
}

// Size returns b's width, height, and depth.
func (b Box) Size() Int3 {
	return b.Max.Sub(b.Min)
}

// Add offsets the box by vector p.
func (b Box) Add(p Int3) Box {
	return Box{
		Min: b.Min.Add(p),
		Max: b.Max.Add(p),
	}
}

// Sub offsets the box by (-p).
func (b Box) Sub(p Int3) Box {
	return Box{
		Min: b.Min.Sub(p),
		Max: b.Max.Sub(p),
	}
}

// Canon returns a copy of b that is well-formed.
func (b Box) Canon() Box {
	if b.Max.X < b.Min.X {
		b.Min.X, b.Max.X = b.Max.X, b.Min.X
	}
	if b.Max.Y < b.Min.Y {
		b.Min.Y, b.Max.Y = b.Max.Y, b.Min.Y
	}
	if b.Max.Z < b.Min.Z {
		b.Min.Z, b.Max.Z = b.Max.Z, b.Min.Z
	}
	return b
}

// Back returns an image.Rectangle representing the back of the box, using
// the given projection π.
func (b Box) Back(π IntProjection) image.Rectangle {
	b.Max.Z = b.Min.Z
	return image.Rectangle{
		Min: π.Project(b.Min),
		Max: π.Project(b.Max),
	}
}

// Front returns an image.Rectangle representing the front of the box, using
// the given projection π.
func (b Box) Front(π IntProjection) image.Rectangle {
	b.Min.Z = b.Max.Z
	return image.Rectangle{
		Min: π.Project(b.Min),
		Max: π.Project(b.Max),
	}
}

// XY returns the image.Rectangle representing the box if we forgot about Z.
func (b Box) XY() image.Rectangle {
	return image.Rectangle{
		Min: b.Min.XY(),
		Max: b.Max.XY(),
	}
}

// XZ returns the image.Rectangle representing the box if we forgot about Y.
func (b Box) XZ() image.Rectangle {
	return image.Rectangle{
		Min: b.Min.XZ(),
		Max: b.Max.XZ(),
	}
}
