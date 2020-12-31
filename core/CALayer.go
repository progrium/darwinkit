package core

import "github.com/progrium/macdriver/objc"

type CALayer struct {
	objc.Object
}

func (l CALayer) CornerRadius() float64 {
	return l.Get("cornerRadius").Float()
}

func (l CALayer) SetCornerRadius(r float64) {
	l.Set("cornerRadius:", r)
}

func (l CALayer) Contents() objc.Object {
	return l.Get("contents")
}

func (l CALayer) SetContents(o objc.Object) {
	l.Set("contents:", o)
}
