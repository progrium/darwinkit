// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RenderPipelineReflection] class.
var RenderPipelineReflectionClass = _RenderPipelineReflectionClass{objc.GetClass("MTLRenderPipelineReflection")}

type _RenderPipelineReflectionClass struct {
	objc.Class
}

// An interface definition for the [RenderPipelineReflection] class.
type IRenderPipelineReflection interface {
	objc.IObject
}

// Information about the arguments of a graphics function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinereflection?language=objc
type RenderPipelineReflection struct {
	objc.Object
}

func RenderPipelineReflectionFrom(ptr unsafe.Pointer) RenderPipelineReflection {
	return RenderPipelineReflection{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RenderPipelineReflectionClass) Alloc() RenderPipelineReflection {
	rv := objc.Call[RenderPipelineReflection](rc, objc.Sel("alloc"))
	return rv
}

func RenderPipelineReflection_Alloc() RenderPipelineReflection {
	return RenderPipelineReflectionClass.Alloc()
}

func (rc _RenderPipelineReflectionClass) New() RenderPipelineReflection {
	rv := objc.Call[RenderPipelineReflection](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRenderPipelineReflection() RenderPipelineReflection {
	return RenderPipelineReflectionClass.New()
}

func (r_ RenderPipelineReflection) Init() RenderPipelineReflection {
	rv := objc.Call[RenderPipelineReflection](r_, objc.Sel("init"))
	return rv
}
