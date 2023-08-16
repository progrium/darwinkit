// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FunctionConstantValues] class.
var FunctionConstantValuesClass = _FunctionConstantValuesClass{objc.GetClass("MTLFunctionConstantValues")}

type _FunctionConstantValuesClass struct {
	objc.Class
}

// An interface definition for the [FunctionConstantValues] class.
type IFunctionConstantValues interface {
	objc.IObject
	SetConstantValueTypeWithName(value unsafe.Pointer, type_ DataType, name string)
	SetConstantValuesTypeWithRange(values unsafe.Pointer, type_ DataType, range_ foundation.Range)
	Reset()
}

// A set of constant values used to specialize a graphics or compute function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionconstantvalues?language=objc
type FunctionConstantValues struct {
	objc.Object
}

func FunctionConstantValuesFrom(ptr unsafe.Pointer) FunctionConstantValues {
	return FunctionConstantValues{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FunctionConstantValuesClass) Alloc() FunctionConstantValues {
	rv := objc.Call[FunctionConstantValues](fc, objc.Sel("alloc"))
	return rv
}

func FunctionConstantValues_Alloc() FunctionConstantValues {
	return FunctionConstantValuesClass.Alloc()
}

func (fc _FunctionConstantValuesClass) New() FunctionConstantValues {
	rv := objc.Call[FunctionConstantValues](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFunctionConstantValues() FunctionConstantValues {
	return FunctionConstantValuesClass.New()
}

func (f_ FunctionConstantValues) Init() FunctionConstantValues {
	rv := objc.Call[FunctionConstantValues](f_, objc.Sel("init"))
	return rv
}

// Sets a value for a function constant with a specific name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionconstantvalues/1639530-setconstantvalue?language=objc
func (f_ FunctionConstantValues) SetConstantValueTypeWithName(value unsafe.Pointer, type_ DataType, name string) {
	objc.Call[objc.Void](f_, objc.Sel("setConstantValue:type:withName:"), value, type_, name)
}

// Sets values for a group of function constants within a specific index range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionconstantvalues/1639527-setconstantvalues?language=objc
func (f_ FunctionConstantValues) SetConstantValuesTypeWithRange(values unsafe.Pointer, type_ DataType, range_ foundation.Range) {
	objc.Call[objc.Void](f_, objc.Sel("setConstantValues:type:withRange:"), values, type_, range_)
}

// Deletes all previously set constant values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionconstantvalues/1639526-reset?language=objc
func (f_ FunctionConstantValues) Reset() {
	objc.Call[objc.Void](f_, objc.Sel("reset"))
}
