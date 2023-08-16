// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Shadow] class.
var ShadowClass = _ShadowClass{objc.GetClass("NSShadow")}

type _ShadowClass struct {
	objc.Class
}

// An interface definition for the [Shadow] class.
type IShadow interface {
	objc.IObject
	Set()
	ShadowBlurRadius() float64
	SetShadowBlurRadius(value float64)
	ShadowOffset() coregraphics.Size
	SetShadowOffset(value coregraphics.Size)
	ShadowColor() objc.Object
	SetShadowColor(value objc.IObject)
}

// An object you use to specify attributes to create and style a drop shadow during drawing operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsshadow?language=objc
type Shadow struct {
	objc.Object
}

func ShadowFrom(ptr unsafe.Pointer) Shadow {
	return Shadow{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ Shadow) Init() Shadow {
	rv := objc.Call[Shadow](s_, objc.Sel("init"))
	return rv
}

func (sc _ShadowClass) Alloc() Shadow {
	rv := objc.Call[Shadow](sc, objc.Sel("alloc"))
	return rv
}

func Shadow_Alloc() Shadow {
	return ShadowClass.Alloc()
}

func (sc _ShadowClass) New() Shadow {
	rv := objc.Call[Shadow](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewShadow() Shadow {
	return ShadowClass.New()
}

// Sets the shadow of subsequent drawing operations to the current shadow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsshadow/1429857-set?language=objc
func (s_ Shadow) Set() {
	objc.Call[objc.Void](s_, objc.Sel("set"))
}

// The blur radius of the shadow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsshadow/1429846-shadowblurradius?language=objc
func (s_ Shadow) ShadowBlurRadius() float64 {
	rv := objc.Call[float64](s_, objc.Sel("shadowBlurRadius"))
	return rv
}

// The blur radius of the shadow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsshadow/1429846-shadowblurradius?language=objc
func (s_ Shadow) SetShadowBlurRadius(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setShadowBlurRadius:"), value)
}

// The shadow’s relative position, which you specify with horizontal and vertical offset values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsshadow/1429851-shadowoffset?language=objc
func (s_ Shadow) ShadowOffset() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](s_, objc.Sel("shadowOffset"))
	return rv
}

// The shadow’s relative position, which you specify with horizontal and vertical offset values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsshadow/1429851-shadowoffset?language=objc
func (s_ Shadow) SetShadowOffset(value coregraphics.Size) {
	objc.Call[objc.Void](s_, objc.Sel("setShadowOffset:"), value)
}

// The color of the shadow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsshadow/1429855-shadowcolor?language=objc
func (s_ Shadow) ShadowColor() objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("shadowColor"))
	return rv
}

// The color of the shadow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsshadow/1429855-shadowcolor?language=objc
func (s_ Shadow) SetShadowColor(value objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setShadowColor:"), value)
}
