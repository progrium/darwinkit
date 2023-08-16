// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Feature] class.
var FeatureClass = _FeatureClass{objc.GetClass("CIFeature")}

type _FeatureClass struct {
	objc.Class
}

// An interface definition for the [Feature] class.
type IFeature interface {
	objc.IObject
	Bounds() coregraphics.Rect
	Type() string
}

// The abstract superclass for objects representing notable features detected in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifeature?language=objc
type Feature struct {
	objc.Object
}

func FeatureFrom(ptr unsafe.Pointer) Feature {
	return Feature{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FeatureClass) Alloc() Feature {
	rv := objc.Call[Feature](fc, objc.Sel("alloc"))
	return rv
}

func Feature_Alloc() Feature {
	return FeatureClass.Alloc()
}

func (fc _FeatureClass) New() Feature {
	rv := objc.Call[Feature](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFeature() Feature {
	return FeatureClass.New()
}

func (f_ Feature) Init() Feature {
	rv := objc.Call[Feature](f_, objc.Sel("init"))
	return rv
}

// The rectangle that holds discovered feature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifeature/1437782-bounds?language=objc
func (f_ Feature) Bounds() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](f_, objc.Sel("bounds"))
	return rv
}

// The type of feature that was discovered. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifeature/1438092-type?language=objc
func (f_ Feature) Type() string {
	rv := objc.Call[string](f_, objc.Sel("type"))
	return rv
}
