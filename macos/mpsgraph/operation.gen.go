// AUTO-GENERATED CODE, DO NOT MODIFY

package mpsgraph

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Operation] class.
var OperationClass = _OperationClass{objc.GetClass("MPSGraphOperation")}

type _OperationClass struct {
	objc.Class
}

// An interface definition for the [Operation] class.
type IOperation interface {
	objc.IObject
	Name() string
	ControlDependencies() []Operation
	InputTensors() []Tensor
	Graph() Graph
	OutputTensors() []Tensor
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphoperation?language=objc
type Operation struct {
	objc.Object
}

func OperationFrom(ptr unsafe.Pointer) Operation {
	return Operation{
		Object: objc.ObjectFrom(ptr),
	}
}

func (oc _OperationClass) Alloc() Operation {
	rv := objc.Call[Operation](oc, objc.Sel("alloc"))
	return rv
}

func Operation_Alloc() Operation {
	return OperationClass.Alloc()
}

func (oc _OperationClass) New() Operation {
	rv := objc.Call[Operation](oc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewOperation() Operation {
	return OperationClass.New()
}

func (o_ Operation) Init() Operation {
	rv := objc.Call[Operation](o_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphoperation/3564659-name?language=objc
func (o_ Operation) Name() string {
	rv := objc.Call[string](o_, objc.Sel("name"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphoperation/3564657-controldependencies?language=objc
func (o_ Operation) ControlDependencies() []Operation {
	rv := objc.Call[[]Operation](o_, objc.Sel("controlDependencies"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphoperation/3564658-inputtensors?language=objc
func (o_ Operation) InputTensors() []Tensor {
	rv := objc.Call[[]Tensor](o_, objc.Sel("inputTensors"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphoperation/3580488-graph?language=objc
func (o_ Operation) Graph() Graph {
	rv := objc.Call[Graph](o_, objc.Sel("graph"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphoperation/3564660-outputtensors?language=objc
func (o_ Operation) OutputTensors() []Tensor {
	rv := objc.Call[[]Tensor](o_, objc.Sel("outputTensors"))
	return rv
}
