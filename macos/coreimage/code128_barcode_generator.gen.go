// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a Code 128 barcode generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicode128barcodegenerator?language=objc
type PCode128BarcodeGenerator interface {
	// optional
	SetQuietSpace(value float64)
	HasSetQuietSpace() bool

	// optional
	QuietSpace() float64
	HasQuietSpace() bool

	// optional
	SetBarcodeHeight(value float64)
	HasSetBarcodeHeight() bool

	// optional
	BarcodeHeight() float64
	HasBarcodeHeight() bool

	// optional
	SetMessage(value []byte)
	HasSetMessage() bool

	// optional
	Message() []byte
	HasMessage() bool
}

// A concrete type wrapper for the [PCode128BarcodeGenerator] protocol.
type Code128BarcodeGeneratorWrapper struct {
	objc.Object
}

func (c_ Code128BarcodeGeneratorWrapper) HasSetQuietSpace() bool {
	return c_.RespondsToSelector(objc.Sel("setQuietSpace:"))
}

// The number of empty white pixels that should surround the barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicode128barcodegenerator/3228118-quietspace?language=objc
func (c_ Code128BarcodeGeneratorWrapper) SetQuietSpace(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setQuietSpace:"), value)
}

func (c_ Code128BarcodeGeneratorWrapper) HasQuietSpace() bool {
	return c_.RespondsToSelector(objc.Sel("quietSpace"))
}

// The number of empty white pixels that should surround the barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicode128barcodegenerator/3228118-quietspace?language=objc
func (c_ Code128BarcodeGeneratorWrapper) QuietSpace() float64 {
	rv := objc.Call[float64](c_, objc.Sel("quietSpace"))
	return rv
}

func (c_ Code128BarcodeGeneratorWrapper) HasSetBarcodeHeight() bool {
	return c_.RespondsToSelector(objc.Sel("setBarcodeHeight:"))
}

// The height, in pixels, of the generated barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicode128barcodegenerator/3228116-barcodeheight?language=objc
func (c_ Code128BarcodeGeneratorWrapper) SetBarcodeHeight(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setBarcodeHeight:"), value)
}

func (c_ Code128BarcodeGeneratorWrapper) HasBarcodeHeight() bool {
	return c_.RespondsToSelector(objc.Sel("barcodeHeight"))
}

// The height, in pixels, of the generated barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicode128barcodegenerator/3228116-barcodeheight?language=objc
func (c_ Code128BarcodeGeneratorWrapper) BarcodeHeight() float64 {
	rv := objc.Call[float64](c_, objc.Sel("barcodeHeight"))
	return rv
}

func (c_ Code128BarcodeGeneratorWrapper) HasSetMessage() bool {
	return c_.RespondsToSelector(objc.Sel("setMessage:"))
}

// The message to encode in the Code 128 barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicode128barcodegenerator/3228117-message?language=objc
func (c_ Code128BarcodeGeneratorWrapper) SetMessage(value []byte) {
	objc.Call[objc.Void](c_, objc.Sel("setMessage:"), value)
}

func (c_ Code128BarcodeGeneratorWrapper) HasMessage() bool {
	return c_.RespondsToSelector(objc.Sel("message"))
}

// The message to encode in the Code 128 barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicode128barcodegenerator/3228117-message?language=objc
func (c_ Code128BarcodeGeneratorWrapper) Message() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("message"))
	return rv
}
