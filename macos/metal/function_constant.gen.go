// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FunctionConstant] class.
var FunctionConstantClass = _FunctionConstantClass{objc.GetClass("MTLFunctionConstant")}

type _FunctionConstantClass struct {
	objc.Class
}

// An interface definition for the [FunctionConstant] class.
type IFunctionConstant interface {
	objc.IObject
	Name() string
	Required() bool
	Type() DataType
	Index() uint
}

// A constant used to specialize the behavior of a shader. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionconstant?language=objc
type FunctionConstant struct {
	objc.Object
}

func FunctionConstantFrom(ptr unsafe.Pointer) FunctionConstant {
	return FunctionConstant{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FunctionConstantClass) Alloc() FunctionConstant {
	rv := objc.Call[FunctionConstant](fc, objc.Sel("alloc"))
	return rv
}

func FunctionConstant_Alloc() FunctionConstant {
	return FunctionConstantClass.Alloc()
}

func (fc _FunctionConstantClass) New() FunctionConstant {
	rv := objc.Call[FunctionConstant](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFunctionConstant() FunctionConstant {
	return FunctionConstantClass.New()
}

func (f_ FunctionConstant) Init() FunctionConstant {
	rv := objc.Call[FunctionConstant](f_, objc.Sel("init"))
	return rv
}

// The name of the function constant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionconstant/1639955-name?language=objc
func (f_ FunctionConstant) Name() string {
	rv := objc.Call[string](f_, objc.Sel("name"))
	return rv
}

// A Boolean value indicating whether the function constant must be provided to specialize the function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionconstant/1639988-required?language=objc
func (f_ FunctionConstant) Required() bool {
	rv := objc.Call[bool](f_, objc.Sel("required"))
	return rv
}

// The data type of the function constant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionconstant/1639950-type?language=objc
func (f_ FunctionConstant) Type() DataType {
	rv := objc.Call[DataType](f_, objc.Sel("type"))
	return rv
}

// The index of the function constant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionconstant/1639905-index?language=objc
func (f_ FunctionConstant) Index() uint {
	rv := objc.Call[uint](f_, objc.Sel("index"))
	return rv
}
