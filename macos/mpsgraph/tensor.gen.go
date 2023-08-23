// AUTO-GENERATED CODE, DO NOT MODIFY

package mpsgraph

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/mps"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Tensor] class.
var TensorClass = _TensorClass{objc.GetClass("MPSGraphTensor")}

type _TensorClass struct {
	objc.Class
}

// An interface definition for the [Tensor] class.
type ITensor interface {
	objc.IObject
	Shape() *foundation.Array
	DataType() mps.DataType
	Operation() Operation
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensor?language=objc
type Tensor struct {
	objc.Object
}

func TensorFrom(ptr unsafe.Pointer) Tensor {
	return Tensor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (tc _TensorClass) Alloc() Tensor {
	rv := objc.Call[Tensor](tc, objc.Sel("alloc"))
	return rv
}

func Tensor_Alloc() Tensor {
	return TensorClass.Alloc()
}

func (tc _TensorClass) New() Tensor {
	rv := objc.Call[Tensor](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTensor() Tensor {
	return TensorClass.New()
}

func (t_ Tensor) Init() Tensor {
	rv := objc.Call[Tensor](t_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensor/3564666-shape?language=objc
func (t_ Tensor) Shape() *foundation.Array {
	rv := objc.Call[*foundation.Array](t_, objc.Sel("shape"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensor/3564663-datatype?language=objc
func (t_ Tensor) DataType() mps.DataType {
	rv := objc.Call[mps.DataType](t_, objc.Sel("dataType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensor/3564665-operation?language=objc
func (t_ Tensor) Operation() Operation {
	rv := objc.Call[Operation](t_, objc.Sel("operation"))
	return rv
}
