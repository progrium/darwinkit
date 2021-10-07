package core

import (
	"github.com/progrium/macdriver/objc"
)

type CALayer struct {
	gen_CALayer
}

func (l CALayer) CornerRadius() float64 {
	return float64(l.gen_CALayer.CornerRadius())
}

func (l CALayer) SetCornerRadius(r float64) {
	l.SetCornerRadius_(CGFloat(r))
}

func (l CALayer) SetContents(o objc.Object) {
	l.SetContents_(objc.Object_fromRef(o))
}
