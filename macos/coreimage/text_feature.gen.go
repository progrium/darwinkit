// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextFeature] class.
var TextFeatureClass = _TextFeatureClass{objc.GetClass("CITextFeature")}

type _TextFeatureClass struct {
	objc.Class
}

// An interface definition for the [TextFeature] class.
type ITextFeature interface {
	IFeature
	BottomRight() coregraphics.Point
	BottomLeft() coregraphics.Point
	SubFeatures() []objc.Object
	TopRight() coregraphics.Point
	TopLeft() coregraphics.Point
}

// Information about a region likely to contain text detected in a still or video image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citextfeature?language=objc
type TextFeature struct {
	Feature
}

func TextFeatureFrom(ptr unsafe.Pointer) TextFeature {
	return TextFeature{
		Feature: FeatureFrom(ptr),
	}
}

func (tc _TextFeatureClass) Alloc() TextFeature {
	rv := objc.Call[TextFeature](tc, objc.Sel("alloc"))
	return rv
}

func TextFeature_Alloc() TextFeature {
	return TextFeatureClass.Alloc()
}

func (tc _TextFeatureClass) New() TextFeature {
	rv := objc.Call[TextFeature](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextFeature() TextFeature {
	return TextFeatureClass.New()
}

func (t_ TextFeature) Init() TextFeature {
	rv := objc.Call[TextFeature](t_, objc.Sel("init"))
	return rv
}

// The lower-right corner of the detected text region, in image coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citextfeature/1437659-bottomright?language=objc
func (t_ TextFeature) BottomRight() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](t_, objc.Sel("bottomRight"))
	return rv
}

// The lower-left corner of the detected text region, in image coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citextfeature/1438004-bottomleft?language=objc
func (t_ TextFeature) BottomLeft() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](t_, objc.Sel("bottomLeft"))
	return rv
}

// An array containing additional features detected within the feature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citextfeature/1437810-subfeatures?language=objc
func (t_ TextFeature) SubFeatures() []objc.Object {
	rv := objc.Call[[]objc.Object](t_, objc.Sel("subFeatures"))
	return rv
}

// The upper-right corner of the detected text region, in image coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citextfeature/1438282-topright?language=objc
func (t_ TextFeature) TopRight() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](t_, objc.Sel("topRight"))
	return rv
}

// The upper-left corner of the detected text region, in image coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citextfeature/1438221-topleft?language=objc
func (t_ TextFeature) TopLeft() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](t_, objc.Sel("topLeft"))
	return rv
}
