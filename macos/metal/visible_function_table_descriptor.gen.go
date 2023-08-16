// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [VisibleFunctionTableDescriptor] class.
var VisibleFunctionTableDescriptorClass = _VisibleFunctionTableDescriptorClass{objc.GetClass("MTLVisibleFunctionTableDescriptor")}

type _VisibleFunctionTableDescriptorClass struct {
	objc.Class
}

// An interface definition for the [VisibleFunctionTableDescriptor] class.
type IVisibleFunctionTableDescriptor interface {
	objc.IObject
	FunctionCount() uint
	SetFunctionCount(value uint)
}

// A specification of how to create a visible function table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvisiblefunctiontabledescriptor?language=objc
type VisibleFunctionTableDescriptor struct {
	objc.Object
}

func VisibleFunctionTableDescriptorFrom(ptr unsafe.Pointer) VisibleFunctionTableDescriptor {
	return VisibleFunctionTableDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (vc _VisibleFunctionTableDescriptorClass) Alloc() VisibleFunctionTableDescriptor {
	rv := objc.Call[VisibleFunctionTableDescriptor](vc, objc.Sel("alloc"))
	return rv
}

func VisibleFunctionTableDescriptor_Alloc() VisibleFunctionTableDescriptor {
	return VisibleFunctionTableDescriptorClass.Alloc()
}

func (vc _VisibleFunctionTableDescriptorClass) New() VisibleFunctionTableDescriptor {
	rv := objc.Call[VisibleFunctionTableDescriptor](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewVisibleFunctionTableDescriptor() VisibleFunctionTableDescriptor {
	return VisibleFunctionTableDescriptorClass.New()
}

func (v_ VisibleFunctionTableDescriptor) Init() VisibleFunctionTableDescriptor {
	rv := objc.Call[VisibleFunctionTableDescriptor](v_, objc.Sel("init"))
	return rv
}

// Creates a default visible function table descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvisiblefunctiontabledescriptor/3578282-visiblefunctiontabledescriptor?language=objc
func (vc _VisibleFunctionTableDescriptorClass) VisibleFunctionTableDescriptor() VisibleFunctionTableDescriptor {
	rv := objc.Call[VisibleFunctionTableDescriptor](vc, objc.Sel("visibleFunctionTableDescriptor"))
	return rv
}

// Creates a default visible function table descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvisiblefunctiontabledescriptor/3578282-visiblefunctiontabledescriptor?language=objc
func VisibleFunctionTableDescriptor_VisibleFunctionTableDescriptor() VisibleFunctionTableDescriptor {
	return VisibleFunctionTableDescriptorClass.VisibleFunctionTableDescriptor()
}

// The number of entries in the function table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvisiblefunctiontabledescriptor/3554059-functioncount?language=objc
func (v_ VisibleFunctionTableDescriptor) FunctionCount() uint {
	rv := objc.Call[uint](v_, objc.Sel("functionCount"))
	return rv
}

// The number of entries in the function table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvisiblefunctiontabledescriptor/3554059-functioncount?language=objc
func (v_ VisibleFunctionTableDescriptor) SetFunctionCount(value uint) {
	objc.Call[objc.Void](v_, objc.Sel("setFunctionCount:"), value)
}
