// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [QRCodeFeature] class.
var QRCodeFeatureClass = _QRCodeFeatureClass{objc.GetClass("CIQRCodeFeature")}

type _QRCodeFeatureClass struct {
	objc.Class
}

// An interface definition for the [QRCodeFeature] class.
type IQRCodeFeature interface {
	IFeature
	BottomRight() coregraphics.Point
	BottomLeft() coregraphics.Point
	TopRight() coregraphics.Point
	SymbolDescriptor() QRCodeDescriptor
	TopLeft() coregraphics.Point
	MessageString() string
}

// Information about a Quick Response code (a kind of 2D barcode) detected in a still or video image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciqrcodefeature?language=objc
type QRCodeFeature struct {
	Feature
}

func QRCodeFeatureFrom(ptr unsafe.Pointer) QRCodeFeature {
	return QRCodeFeature{
		Feature: FeatureFrom(ptr),
	}
}

func (qc _QRCodeFeatureClass) Alloc() QRCodeFeature {
	rv := objc.Call[QRCodeFeature](qc, objc.Sel("alloc"))
	return rv
}

func QRCodeFeature_Alloc() QRCodeFeature {
	return QRCodeFeatureClass.Alloc()
}

func (qc _QRCodeFeatureClass) New() QRCodeFeature {
	rv := objc.Call[QRCodeFeature](qc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewQRCodeFeature() QRCodeFeature {
	return QRCodeFeatureClass.New()
}

func (q_ QRCodeFeature) Init() QRCodeFeature {
	rv := objc.Call[QRCodeFeature](q_, objc.Sel("init"))
	return rv
}

// The lower-right corner of the detected barcode, in image coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciqrcodefeature/1438245-bottomright?language=objc
func (q_ QRCodeFeature) BottomRight() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](q_, objc.Sel("bottomRight"))
	return rv
}

// The lower-left corner of the detected barcode, in image coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciqrcodefeature/1437985-bottomleft?language=objc
func (q_ QRCodeFeature) BottomLeft() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](q_, objc.Sel("bottomLeft"))
	return rv
}

// The upper-right corner of the detected barcode, in image coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciqrcodefeature/1437896-topright?language=objc
func (q_ QRCodeFeature) TopRight() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](q_, objc.Sel("topRight"))
	return rv
}

// An abstract representation of a QR Code symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciqrcodefeature/2875553-symboldescriptor?language=objc
func (q_ QRCodeFeature) SymbolDescriptor() QRCodeDescriptor {
	rv := objc.Call[QRCodeDescriptor](q_, objc.Sel("symbolDescriptor"))
	return rv
}

// The upper-left corner of the detected barcode, in image coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciqrcodefeature/1437780-topleft?language=objc
func (q_ QRCodeFeature) TopLeft() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](q_, objc.Sel("topLeft"))
	return rv
}

// The string decoded from the detected barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciqrcodefeature/1438035-messagestring?language=objc
func (q_ QRCodeFeature) MessageString() string {
	rv := objc.Call[string](q_, objc.Sel("messageString"))
	return rv
}
