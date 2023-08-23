// AUTO-GENERATED CODE, DO NOT MODIFY

package mpsgraph

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/mps"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [VariableOp] class.
var VariableOpClass = _VariableOpClass{objc.GetClass("MPSGraphVariableOp")}

type _VariableOpClass struct {
	objc.Class
}

// An interface definition for the [VariableOp] class.
type IVariableOp interface {
	IOperation
	Shape() *foundation.Array
	DataType() mps.DataType
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphvariableop?language=objc
type VariableOp struct {
	Operation
}

func VariableOpFrom(ptr unsafe.Pointer) VariableOp {
	return VariableOp{
		Operation: OperationFrom(ptr),
	}
}

func (vc _VariableOpClass) Alloc() VariableOp {
	rv := objc.Call[VariableOp](vc, objc.Sel("alloc"))
	return rv
}

func VariableOp_Alloc() VariableOp {
	return VariableOpClass.Alloc()
}

func (vc _VariableOpClass) New() VariableOp {
	rv := objc.Call[VariableOp](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewVariableOp() VariableOp {
	return VariableOpClass.New()
}

func (v_ VariableOp) Init() VariableOp {
	rv := objc.Call[VariableOp](v_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphvariableop/3564695-shape?language=objc
func (v_ VariableOp) Shape() *foundation.Array {
	rv := objc.Call[*foundation.Array](v_, objc.Sel("shape"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphvariableop/3564693-datatype?language=objc
func (v_ VariableOp) DataType() mps.DataType {
	rv := objc.Call[mps.DataType](v_, objc.Sel("dataType"))
	return rv
}
