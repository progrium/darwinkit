// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RenderPipelineFunctionsDescriptor] class.
var RenderPipelineFunctionsDescriptorClass = _RenderPipelineFunctionsDescriptorClass{objc.GetClass("MTLRenderPipelineFunctionsDescriptor")}

type _RenderPipelineFunctionsDescriptorClass struct {
	objc.Class
}

// An interface definition for the [RenderPipelineFunctionsDescriptor] class.
type IRenderPipelineFunctionsDescriptor interface {
	objc.IObject
	TileAdditionalBinaryFunctions() []FunctionWrapper
	SetTileAdditionalBinaryFunctions(value []PFunction)
	VertexAdditionalBinaryFunctions() []FunctionWrapper
	SetVertexAdditionalBinaryFunctions(value []PFunction)
	FragmentAdditionalBinaryFunctions() []FunctionWrapper
	SetFragmentAdditionalBinaryFunctions(value []PFunction)
}

// A collection of functions for updating a render pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinefunctionsdescriptor?language=objc
type RenderPipelineFunctionsDescriptor struct {
	objc.Object
}

func RenderPipelineFunctionsDescriptorFrom(ptr unsafe.Pointer) RenderPipelineFunctionsDescriptor {
	return RenderPipelineFunctionsDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RenderPipelineFunctionsDescriptorClass) Alloc() RenderPipelineFunctionsDescriptor {
	rv := objc.Call[RenderPipelineFunctionsDescriptor](rc, objc.Sel("alloc"))
	return rv
}

func RenderPipelineFunctionsDescriptor_Alloc() RenderPipelineFunctionsDescriptor {
	return RenderPipelineFunctionsDescriptorClass.Alloc()
}

func (rc _RenderPipelineFunctionsDescriptorClass) New() RenderPipelineFunctionsDescriptor {
	rv := objc.Call[RenderPipelineFunctionsDescriptor](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRenderPipelineFunctionsDescriptor() RenderPipelineFunctionsDescriptor {
	return RenderPipelineFunctionsDescriptorClass.New()
}

func (r_ RenderPipelineFunctionsDescriptor) Init() RenderPipelineFunctionsDescriptor {
	rv := objc.Call[RenderPipelineFunctionsDescriptor](r_, objc.Sel("init"))
	return rv
}

// The tile functions to add to the render pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinefunctionsdescriptor/3750579-tileadditionalbinaryfunctions?language=objc
func (r_ RenderPipelineFunctionsDescriptor) TileAdditionalBinaryFunctions() []FunctionWrapper {
	rv := objc.Call[[]FunctionWrapper](r_, objc.Sel("tileAdditionalBinaryFunctions"))
	return rv
}

// The tile functions to add to the render pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinefunctionsdescriptor/3750579-tileadditionalbinaryfunctions?language=objc
func (r_ RenderPipelineFunctionsDescriptor) SetTileAdditionalBinaryFunctions(value []PFunction) {
	objc.Call[objc.Void](r_, objc.Sel("setTileAdditionalBinaryFunctions:"), value)
}

// The vertex functions to add to the render pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinefunctionsdescriptor/3750580-vertexadditionalbinaryfunctions?language=objc
func (r_ RenderPipelineFunctionsDescriptor) VertexAdditionalBinaryFunctions() []FunctionWrapper {
	rv := objc.Call[[]FunctionWrapper](r_, objc.Sel("vertexAdditionalBinaryFunctions"))
	return rv
}

// The vertex functions to add to the render pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinefunctionsdescriptor/3750580-vertexadditionalbinaryfunctions?language=objc
func (r_ RenderPipelineFunctionsDescriptor) SetVertexAdditionalBinaryFunctions(value []PFunction) {
	objc.Call[objc.Void](r_, objc.Sel("setVertexAdditionalBinaryFunctions:"), value)
}

// The fragment functions to add to the render pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinefunctionsdescriptor/3750578-fragmentadditionalbinaryfunction?language=objc
func (r_ RenderPipelineFunctionsDescriptor) FragmentAdditionalBinaryFunctions() []FunctionWrapper {
	rv := objc.Call[[]FunctionWrapper](r_, objc.Sel("fragmentAdditionalBinaryFunctions"))
	return rv
}

// The fragment functions to add to the render pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinefunctionsdescriptor/3750578-fragmentadditionalbinaryfunction?language=objc
func (r_ RenderPipelineFunctionsDescriptor) SetFragmentAdditionalBinaryFunctions(value []PFunction) {
	objc.Call[objc.Void](r_, objc.Sel("setFragmentAdditionalBinaryFunctions:"), value)
}
