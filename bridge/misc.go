package bridge

import (
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
)

type Point struct {
	X float64
	Y float64
}

func (p *Point) NSPoint() core.NSPoint {
	return core.NSPoint{X: p.X, Y: p.Y}
}

type Size struct {
	W float64
	H float64
}

func (s *Size) NSSize() core.NSSize {
	return core.NSSize{Width: s.W, Height: s.H}
}

type Color struct {
	R float64
	G float64
	B float64
	A float64
}

func (c *Color) NSColor() cocoa.NSColor {
	return cocoa.NSColor_Init(c.R, c.G, c.B, c.A)
}
