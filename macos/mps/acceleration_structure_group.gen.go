// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AccelerationStructureGroup] class.
var AccelerationStructureGroupClass = _AccelerationStructureGroupClass{objc.GetClass("MPSAccelerationStructureGroup")}

type _AccelerationStructureGroupClass struct {
	objc.Class
}

// An interface definition for the [AccelerationStructureGroup] class.
type IAccelerationStructureGroup interface {
	objc.IObject
}

// A group of acceleration structures. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructuregroup?language=objc
type AccelerationStructureGroup struct {
	objc.Object
}

func AccelerationStructureGroupFrom(ptr unsafe.Pointer) AccelerationStructureGroup {
	return AccelerationStructureGroup{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AccelerationStructureGroupClass) Alloc() AccelerationStructureGroup {
	rv := objc.Call[AccelerationStructureGroup](ac, objc.Sel("alloc"))
	return rv
}

func AccelerationStructureGroup_Alloc() AccelerationStructureGroup {
	return AccelerationStructureGroupClass.Alloc()
}

func (ac _AccelerationStructureGroupClass) New() AccelerationStructureGroup {
	rv := objc.Call[AccelerationStructureGroup](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAccelerationStructureGroup() AccelerationStructureGroup {
	return AccelerationStructureGroupClass.New()
}

func (a_ AccelerationStructureGroup) Init() AccelerationStructureGroup {
	rv := objc.Call[AccelerationStructureGroup](a_, objc.Sel("init"))
	return rv
}
