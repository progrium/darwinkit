// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [IntersectionFunctionTableDescriptor] class.
var IntersectionFunctionTableDescriptorClass = _IntersectionFunctionTableDescriptorClass{objc.GetClass("MTLIntersectionFunctionTableDescriptor")}

type _IntersectionFunctionTableDescriptorClass struct {
	objc.Class
}

// An interface definition for the [IntersectionFunctionTableDescriptor] class.
type IIntersectionFunctionTableDescriptor interface {
	objc.IObject
	FunctionCount() uint
	SetFunctionCount(value uint)
}

// A specification of how to create an intersection function table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlintersectionfunctiontabledescriptor?language=objc
type IntersectionFunctionTableDescriptor struct {
	objc.Object
}

func IntersectionFunctionTableDescriptorFrom(ptr unsafe.Pointer) IntersectionFunctionTableDescriptor {
	return IntersectionFunctionTableDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ic _IntersectionFunctionTableDescriptorClass) Alloc() IntersectionFunctionTableDescriptor {
	rv := objc.Call[IntersectionFunctionTableDescriptor](ic, objc.Sel("alloc"))
	return rv
}

func IntersectionFunctionTableDescriptor_Alloc() IntersectionFunctionTableDescriptor {
	return IntersectionFunctionTableDescriptorClass.Alloc()
}

func (ic _IntersectionFunctionTableDescriptorClass) New() IntersectionFunctionTableDescriptor {
	rv := objc.Call[IntersectionFunctionTableDescriptor](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewIntersectionFunctionTableDescriptor() IntersectionFunctionTableDescriptor {
	return IntersectionFunctionTableDescriptorClass.New()
}

func (i_ IntersectionFunctionTableDescriptor) Init() IntersectionFunctionTableDescriptor {
	rv := objc.Call[IntersectionFunctionTableDescriptor](i_, objc.Sel("init"))
	return rv
}

// Creates an intersection function table descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlintersectionfunctiontabledescriptor/3667486-intersectionfunctiontabledescrip?language=objc
func (ic _IntersectionFunctionTableDescriptorClass) IntersectionFunctionTableDescriptor() IntersectionFunctionTableDescriptor {
	rv := objc.Call[IntersectionFunctionTableDescriptor](ic, objc.Sel("intersectionFunctionTableDescriptor"))
	return rv
}

// Creates an intersection function table descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlintersectionfunctiontabledescriptor/3667486-intersectionfunctiontabledescrip?language=objc
func IntersectionFunctionTableDescriptor_IntersectionFunctionTableDescriptor() IntersectionFunctionTableDescriptor {
	return IntersectionFunctionTableDescriptorClass.IntersectionFunctionTableDescriptor()
}

// The number of entries in the intersection function table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlintersectionfunctiontabledescriptor/3566557-functioncount?language=objc
func (i_ IntersectionFunctionTableDescriptor) FunctionCount() uint {
	rv := objc.Call[uint](i_, objc.Sel("functionCount"))
	return rv
}

// The number of entries in the intersection function table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlintersectionfunctiontabledescriptor/3566557-functioncount?language=objc
func (i_ IntersectionFunctionTableDescriptor) SetFunctionCount(value uint) {
	objc.Call[objc.Void](i_, objc.Sel("setFunctionCount:"), value)
}
