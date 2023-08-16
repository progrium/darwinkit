// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CachedImageRep] class.
var CachedImageRepClass = _CachedImageRepClass{objc.GetClass("NSCachedImageRep")}

type _CachedImageRepClass struct {
	objc.Class
}

// An interface definition for the [CachedImageRep] class.
type ICachedImageRep interface {
	IImageRep
}

// An object that stores image data in a form that can be readily transferred to the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscachedimagerep?language=objc
type CachedImageRep struct {
	ImageRep
}

func CachedImageRepFrom(ptr unsafe.Pointer) CachedImageRep {
	return CachedImageRep{
		ImageRep: ImageRepFrom(ptr),
	}
}

func (cc _CachedImageRepClass) Alloc() CachedImageRep {
	rv := objc.Call[CachedImageRep](cc, objc.Sel("alloc"))
	return rv
}

func CachedImageRep_Alloc() CachedImageRep {
	return CachedImageRepClass.Alloc()
}

func (cc _CachedImageRepClass) New() CachedImageRep {
	rv := objc.Call[CachedImageRep](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCachedImageRep() CachedImageRep {
	return CachedImageRepClass.New()
}

func (c_ CachedImageRep) Init() CachedImageRep {
	rv := objc.Call[CachedImageRep](c_, objc.Sel("init"))
	return rv
}
