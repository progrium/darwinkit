// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ComputePipelineReflection] class.
var ComputePipelineReflectionClass = _ComputePipelineReflectionClass{objc.GetClass("MTLComputePipelineReflection")}

type _ComputePipelineReflectionClass struct {
	objc.Class
}

// An interface definition for the [ComputePipelineReflection] class.
type IComputePipelineReflection interface {
	objc.IObject
}

// Information about the arguments of a compute function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinereflection?language=objc
type ComputePipelineReflection struct {
	objc.Object
}

func ComputePipelineReflectionFrom(ptr unsafe.Pointer) ComputePipelineReflection {
	return ComputePipelineReflection{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ComputePipelineReflectionClass) Alloc() ComputePipelineReflection {
	rv := objc.Call[ComputePipelineReflection](cc, objc.Sel("alloc"))
	return rv
}

func ComputePipelineReflection_Alloc() ComputePipelineReflection {
	return ComputePipelineReflectionClass.Alloc()
}

func (cc _ComputePipelineReflectionClass) New() ComputePipelineReflection {
	rv := objc.Call[ComputePipelineReflection](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewComputePipelineReflection() ComputePipelineReflection {
	return ComputePipelineReflectionClass.New()
}

func (c_ ComputePipelineReflection) Init() ComputePipelineReflection {
	rv := objc.Call[ComputePipelineReflection](c_, objc.Sel("init"))
	return rv
}
