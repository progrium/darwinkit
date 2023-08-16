// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [EPSImageRep] class.
var EPSImageRepClass = _EPSImageRepClass{objc.GetClass("NSEPSImageRep")}

type _EPSImageRepClass struct {
	objc.Class
}

// An interface definition for the [EPSImageRep] class.
type IEPSImageRep interface {
	IImageRep
}

// An object that can render an image from encapsulated PostScript (EPS) code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsepsimagerep?language=objc
type EPSImageRep struct {
	ImageRep
}

func EPSImageRepFrom(ptr unsafe.Pointer) EPSImageRep {
	return EPSImageRep{
		ImageRep: ImageRepFrom(ptr),
	}
}

func (ec _EPSImageRepClass) Alloc() EPSImageRep {
	rv := objc.Call[EPSImageRep](ec, objc.Sel("alloc"))
	return rv
}

func EPSImageRep_Alloc() EPSImageRep {
	return EPSImageRepClass.Alloc()
}

func (ec _EPSImageRepClass) New() EPSImageRep {
	rv := objc.Call[EPSImageRep](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewEPSImageRep() EPSImageRep {
	return EPSImageRepClass.New()
}

func (e_ EPSImageRep) Init() EPSImageRep {
	rv := objc.Call[EPSImageRep](e_, objc.Sel("init"))
	return rv
}
