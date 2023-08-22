// AUTO-GENERATED CODE, DO NOT MODIFY

package mpsgraph

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CompilationDescriptor] class.
var CompilationDescriptorClass = _CompilationDescriptorClass{objc.GetClass("MPSGraphCompilationDescriptor")}

type _CompilationDescriptorClass struct {
	objc.Class
}

// An interface definition for the [CompilationDescriptor] class.
type ICompilationDescriptor interface {
	objc.IObject
	DisableTypeInference()
	OptimizationLevel() Optimization
	SetOptimizationLevel(value Optimization)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphcompilationdescriptor?language=objc
type CompilationDescriptor struct {
	objc.Object
}

func CompilationDescriptorFrom(ptr unsafe.Pointer) CompilationDescriptor {
	return CompilationDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CompilationDescriptorClass) Alloc() CompilationDescriptor {
	rv := objc.Call[CompilationDescriptor](cc, objc.Sel("alloc"))
	return rv
}

func CompilationDescriptor_Alloc() CompilationDescriptor {
	return CompilationDescriptorClass.Alloc()
}

func (cc _CompilationDescriptorClass) New() CompilationDescriptor {
	rv := objc.Call[CompilationDescriptor](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCompilationDescriptor() CompilationDescriptor {
	return CompilationDescriptorClass.New()
}

func (c_ CompilationDescriptor) Init() CompilationDescriptor {
	rv := objc.Call[CompilationDescriptor](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphcompilationdescriptor/3787577-disabletypeinference?language=objc
func (c_ CompilationDescriptor) DisableTypeInference() {
	objc.Call[objc.Void](c_, objc.Sel("disableTypeInference"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphcompilationdescriptor/3922624-optimizationlevel?language=objc
func (c_ CompilationDescriptor) OptimizationLevel() Optimization {
	rv := objc.Call[Optimization](c_, objc.Sel("optimizationLevel"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphcompilationdescriptor/3922624-optimizationlevel?language=objc
func (c_ CompilationDescriptor) SetOptimizationLevel(value Optimization) {
	objc.Call[objc.Void](c_, objc.Sel("setOptimizationLevel:"), value)
}
