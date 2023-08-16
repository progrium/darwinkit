// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a barcode generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibarcodegenerator?language=objc
type PBarcodeGenerator interface {
	// optional
	SetBarcodeDescriptor(value BarcodeDescriptor)
	HasSetBarcodeDescriptor() bool

	// optional
	BarcodeDescriptor() IBarcodeDescriptor
	HasBarcodeDescriptor() bool
}

// A concrete type wrapper for the [PBarcodeGenerator] protocol.
type BarcodeGeneratorWrapper struct {
	objc.Object
}

func (b_ BarcodeGeneratorWrapper) HasSetBarcodeDescriptor() bool {
	return b_.RespondsToSelector(objc.Sel("setBarcodeDescriptor:"))
}

// The barcode descriptor to generate an image for. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibarcodegenerator/3228068-barcodedescriptor?language=objc
func (b_ BarcodeGeneratorWrapper) SetBarcodeDescriptor(value IBarcodeDescriptor) {
	objc.Call[objc.Void](b_, objc.Sel("setBarcodeDescriptor:"), objc.Ptr(value))
}

func (b_ BarcodeGeneratorWrapper) HasBarcodeDescriptor() bool {
	return b_.RespondsToSelector(objc.Sel("barcodeDescriptor"))
}

// The barcode descriptor to generate an image for. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibarcodegenerator/3228068-barcodedescriptor?language=objc
func (b_ BarcodeGeneratorWrapper) BarcodeDescriptor() BarcodeDescriptor {
	rv := objc.Call[BarcodeDescriptor](b_, objc.Sel("barcodeDescriptor"))
	return rv
}
