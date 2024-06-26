// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/coreimage"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [CIImageRep] class.
var CIImageRepClass = _CIImageRepClass{objc.GetClass("NSCIImageRep")}

type _CIImageRepClass struct {
	objc.Class
}

// An interface definition for the [CIImageRep] class.
type ICIImageRep interface {
	IImageRep
	CIImage() coreimage.Image
}

// An object that can render an image from a Core Image object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsciimagerep?language=objc
type CIImageRep struct {
	ImageRep
}

func CIImageRepFrom(ptr unsafe.Pointer) CIImageRep {
	return CIImageRep{
		ImageRep: ImageRepFrom(ptr),
	}
}

func (cc _CIImageRepClass) ImageRepWithCIImage(image coreimage.IImage) CIImageRep {
	rv := objc.Call[CIImageRep](cc, objc.Sel("imageRepWithCIImage:"), image)
	return rv
}

// Creates and returns a representation of an image initialized to the specified Core Image instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsciimagerep/1550736-imagerepwithciimage?language=objc
func CIImageRep_ImageRepWithCIImage(image coreimage.IImage) CIImageRep {
	return CIImageRepClass.ImageRepWithCIImage(image)
}

func (c_ CIImageRep) InitWithCIImage(image coreimage.IImage) CIImageRep {
	rv := objc.Call[CIImageRep](c_, objc.Sel("initWithCIImage:"), image)
	return rv
}

// Returns a representation of an image initialized to the specified Core Image instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsciimagerep/1528642-initwithciimage?language=objc
func NewCIImageRepWithCIImage(image coreimage.IImage) CIImageRep {
	instance := CIImageRepClass.Alloc().InitWithCIImage(image)
	instance.Autorelease()
	return instance
}

func (cc _CIImageRepClass) Alloc() CIImageRep {
	rv := objc.Call[CIImageRep](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CIImageRepClass) New() CIImageRep {
	rv := objc.Call[CIImageRep](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCIImageRep() CIImageRep {
	return CIImageRepClass.New()
}

func (c_ CIImageRep) Init() CIImageRep {
	rv := objc.Call[CIImageRep](c_, objc.Sel("init"))
	return rv
}

// The Core Image instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsciimagerep/1525696-ciimage?language=objc
func (c_ CIImageRep) CIImage() coreimage.Image {
	rv := objc.Call[coreimage.Image](c_, objc.Sel("CIImage"))
	return rv
}
