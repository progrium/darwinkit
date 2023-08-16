// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [StitchedLibraryDescriptor] class.
var StitchedLibraryDescriptorClass = _StitchedLibraryDescriptorClass{objc.GetClass("MTLStitchedLibraryDescriptor")}

type _StitchedLibraryDescriptorClass struct {
	objc.Class
}

// An interface definition for the [StitchedLibraryDescriptor] class.
type IStitchedLibraryDescriptor interface {
	objc.IObject
	FunctionGraphs() []FunctionStitchingGraph
	SetFunctionGraphs(value []IFunctionStitchingGraph)
	Functions() []FunctionWrapper
	SetFunctions(value []PFunction)
}

// A description of a new library of procedurally generated functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstitchedlibrarydescriptor?language=objc
type StitchedLibraryDescriptor struct {
	objc.Object
}

func StitchedLibraryDescriptorFrom(ptr unsafe.Pointer) StitchedLibraryDescriptor {
	return StitchedLibraryDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _StitchedLibraryDescriptorClass) Alloc() StitchedLibraryDescriptor {
	rv := objc.Call[StitchedLibraryDescriptor](sc, objc.Sel("alloc"))
	return rv
}

func StitchedLibraryDescriptor_Alloc() StitchedLibraryDescriptor {
	return StitchedLibraryDescriptorClass.Alloc()
}

func (sc _StitchedLibraryDescriptorClass) New() StitchedLibraryDescriptor {
	rv := objc.Call[StitchedLibraryDescriptor](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewStitchedLibraryDescriptor() StitchedLibraryDescriptor {
	return StitchedLibraryDescriptorClass.New()
}

func (s_ StitchedLibraryDescriptor) Init() StitchedLibraryDescriptor {
	rv := objc.Call[StitchedLibraryDescriptor](s_, objc.Sel("init"))
	return rv
}

// The function graphs that define the new stitched library’s functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstitchedlibrarydescriptor/3750549-functiongraphs?language=objc
func (s_ StitchedLibraryDescriptor) FunctionGraphs() []FunctionStitchingGraph {
	rv := objc.Call[[]FunctionStitchingGraph](s_, objc.Sel("functionGraphs"))
	return rv
}

// The function graphs that define the new stitched library’s functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstitchedlibrarydescriptor/3750549-functiongraphs?language=objc
func (s_ StitchedLibraryDescriptor) SetFunctionGraphs(value []IFunctionStitchingGraph) {
	objc.Call[objc.Void](s_, objc.Sel("setFunctionGraphs:"), value)
}

// The list of functions for creating the stitched library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstitchedlibrarydescriptor/3750550-functions?language=objc
func (s_ StitchedLibraryDescriptor) Functions() []FunctionWrapper {
	rv := objc.Call[[]FunctionWrapper](s_, objc.Sel("functions"))
	return rv
}

// The list of functions for creating the stitched library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstitchedlibrarydescriptor/3750550-functions?language=objc
func (s_ StitchedLibraryDescriptor) SetFunctions(value []PFunction) {
	objc.Call[objc.Void](s_, objc.Sel("setFunctions:"), value)
}
