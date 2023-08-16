// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BarcodeDescriptor] class.
var BarcodeDescriptorClass = _BarcodeDescriptorClass{objc.GetClass("CIBarcodeDescriptor")}

type _BarcodeDescriptorClass struct {
	objc.Class
}

// An interface definition for the [BarcodeDescriptor] class.
type IBarcodeDescriptor interface {
	objc.IObject
}

// An abstract base class that represents a machine readable code's attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibarcodedescriptor?language=objc
type BarcodeDescriptor struct {
	objc.Object
}

func BarcodeDescriptorFrom(ptr unsafe.Pointer) BarcodeDescriptor {
	return BarcodeDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (bc _BarcodeDescriptorClass) Alloc() BarcodeDescriptor {
	rv := objc.Call[BarcodeDescriptor](bc, objc.Sel("alloc"))
	return rv
}

func BarcodeDescriptor_Alloc() BarcodeDescriptor {
	return BarcodeDescriptorClass.Alloc()
}

func (bc _BarcodeDescriptorClass) New() BarcodeDescriptor {
	rv := objc.Call[BarcodeDescriptor](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBarcodeDescriptor() BarcodeDescriptor {
	return BarcodeDescriptorClass.New()
}

func (b_ BarcodeDescriptor) Init() BarcodeDescriptor {
	rv := objc.Call[BarcodeDescriptor](b_, objc.Sel("init"))
	return rv
}
