// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var ShadowClass = _ShadowClass{objc.GetClass("NSShadow")}

type _ShadowClass struct {
	objc.Class
}

type IShadow interface {
	objc.IObject
	Set()
	ShadowOffset() foundation.Size
	SetShadowOffset(value foundation.Size)
	ShadowBlurRadius() float64
	SetShadowBlurRadius(value float64)
	ShadowColor() Color
	SetShadowColor(value IColor)
}

type Shadow struct {
	objc.Object
}

func MakeShadow(ptr unsafe.Pointer) Shadow {
	return Shadow{
		Object: objc.MakeObject(ptr),
	}
}

func (s_ Shadow) Init() Shadow {
	rv := objc.CallMethod[Shadow](s_, objc.GetSelector("init"))
	return rv
}

func Shadow_Init() Shadow {
	return ShadowClass.Alloc().Init()
}

func (sc _ShadowClass) Alloc() Shadow {
	rv := objc.CallMethod[Shadow](sc, objc.GetSelector("alloc"))
	return rv
}

func Shadow_Alloc() Shadow {
	return ShadowClass.Alloc()
}

func (sc _ShadowClass) New() Shadow {
	rv := objc.CallMethod[Shadow](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewShadow() Shadow {
	return ShadowClass.New()
}

func Shadow_New() Shadow {
	return ShadowClass.New()
}

func (s_ Shadow) Set() {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("set"))
}

func (s_ Shadow) ShadowOffset() foundation.Size {
	rv := objc.CallMethod[foundation.Size](s_, objc.GetSelector("shadowOffset"))
	return rv
}

func (s_ Shadow) SetShadowOffset(value foundation.Size) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setShadowOffset:"), value)
}

func (s_ Shadow) ShadowBlurRadius() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("shadowBlurRadius"))
	return rv
}

func (s_ Shadow) SetShadowBlurRadius(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setShadowBlurRadius:"), value)
}

func (s_ Shadow) ShadowColor() Color {
	rv := objc.CallMethod[Color](s_, objc.GetSelector("shadowColor"))
	return rv
}

func (s_ Shadow) SetShadowColor(value IColor) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setShadowColor:"), objc.ExtractPtr(value))
}
