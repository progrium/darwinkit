// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [IntersectionFunctionDescriptor] class.
var IntersectionFunctionDescriptorClass = _IntersectionFunctionDescriptorClass{objc.GetClass("MTLIntersectionFunctionDescriptor")}

type _IntersectionFunctionDescriptorClass struct {
	objc.Class
}

// An interface definition for the [IntersectionFunctionDescriptor] class.
type IIntersectionFunctionDescriptor interface {
	IFunctionDescriptor
}

// A description of an intersection function that performs an intersection test. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlintersectionfunctiondescriptor?language=objc
type IntersectionFunctionDescriptor struct {
	FunctionDescriptor
}

func IntersectionFunctionDescriptorFrom(ptr unsafe.Pointer) IntersectionFunctionDescriptor {
	return IntersectionFunctionDescriptor{
		FunctionDescriptor: FunctionDescriptorFrom(ptr),
	}
}

func (ic _IntersectionFunctionDescriptorClass) Alloc() IntersectionFunctionDescriptor {
	rv := objc.Call[IntersectionFunctionDescriptor](ic, objc.Sel("alloc"))
	return rv
}

func IntersectionFunctionDescriptor_Alloc() IntersectionFunctionDescriptor {
	return IntersectionFunctionDescriptorClass.Alloc()
}

func (ic _IntersectionFunctionDescriptorClass) New() IntersectionFunctionDescriptor {
	rv := objc.Call[IntersectionFunctionDescriptor](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewIntersectionFunctionDescriptor() IntersectionFunctionDescriptor {
	return IntersectionFunctionDescriptorClass.New()
}

func (i_ IntersectionFunctionDescriptor) Init() IntersectionFunctionDescriptor {
	rv := objc.Call[IntersectionFunctionDescriptor](i_, objc.Sel("init"))
	return rv
}
