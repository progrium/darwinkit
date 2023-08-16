// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AccelerationStructureDescriptor] class.
var AccelerationStructureDescriptorClass = _AccelerationStructureDescriptorClass{objc.GetClass("MTLAccelerationStructureDescriptor")}

type _AccelerationStructureDescriptorClass struct {
	objc.Class
}

// An interface definition for the [AccelerationStructureDescriptor] class.
type IAccelerationStructureDescriptor interface {
	objc.IObject
	Usage() AccelerationStructureUsage
	SetUsage(value AccelerationStructureUsage)
}

// A base class for classes that define the configuration for a new acceleration structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuredescriptor?language=objc
type AccelerationStructureDescriptor struct {
	objc.Object
}

func AccelerationStructureDescriptorFrom(ptr unsafe.Pointer) AccelerationStructureDescriptor {
	return AccelerationStructureDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AccelerationStructureDescriptorClass) Alloc() AccelerationStructureDescriptor {
	rv := objc.Call[AccelerationStructureDescriptor](ac, objc.Sel("alloc"))
	return rv
}

func AccelerationStructureDescriptor_Alloc() AccelerationStructureDescriptor {
	return AccelerationStructureDescriptorClass.Alloc()
}

func (ac _AccelerationStructureDescriptorClass) New() AccelerationStructureDescriptor {
	rv := objc.Call[AccelerationStructureDescriptor](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAccelerationStructureDescriptor() AccelerationStructureDescriptor {
	return AccelerationStructureDescriptorClass.New()
}

func (a_ AccelerationStructureDescriptor) Init() AccelerationStructureDescriptor {
	rv := objc.Call[AccelerationStructureDescriptor](a_, objc.Sel("init"))
	return rv
}

// The options that describe how you intend to use the acceleration structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuredescriptor/3553866-usage?language=objc
func (a_ AccelerationStructureDescriptor) Usage() AccelerationStructureUsage {
	rv := objc.Call[AccelerationStructureUsage](a_, objc.Sel("usage"))
	return rv
}

// The options that describe how you intend to use the acceleration structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuredescriptor/3553866-usage?language=objc
func (a_ AccelerationStructureDescriptor) SetUsage(value AccelerationStructureUsage) {
	objc.Call[objc.Void](a_, objc.Sel("setUsage:"), value)
}
