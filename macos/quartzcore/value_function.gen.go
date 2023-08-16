// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ValueFunction] class.
var ValueFunctionClass = _ValueFunctionClass{objc.GetClass("CAValueFunction")}

type _ValueFunctionClass struct {
	objc.Class
}

// An interface definition for the [ValueFunction] class.
type IValueFunction interface {
	objc.IObject
	Name() ValueFunctionName
}

// An object that provides a flexible method of defining animated transformations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cavaluefunction?language=objc
type ValueFunction struct {
	objc.Object
}

func ValueFunctionFrom(ptr unsafe.Pointer) ValueFunction {
	return ValueFunction{
		Object: objc.ObjectFrom(ptr),
	}
}

func (vc _ValueFunctionClass) FunctionWithName(name ValueFunctionName) ValueFunction {
	rv := objc.Call[ValueFunction](vc, objc.Sel("functionWithName:"), name)
	return rv
}

// Returns the value function object identified by the name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cavaluefunction/1522115-functionwithname?language=objc
func ValueFunction_FunctionWithName(name ValueFunctionName) ValueFunction {
	return ValueFunctionClass.FunctionWithName(name)
}

func (vc _ValueFunctionClass) Alloc() ValueFunction {
	rv := objc.Call[ValueFunction](vc, objc.Sel("alloc"))
	return rv
}

func ValueFunction_Alloc() ValueFunction {
	return ValueFunctionClass.Alloc()
}

func (vc _ValueFunctionClass) New() ValueFunction {
	rv := objc.Call[ValueFunction](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewValueFunction() ValueFunction {
	return ValueFunctionClass.New()
}

func (v_ ValueFunction) Init() ValueFunction {
	rv := objc.Call[ValueFunction](v_, objc.Sel("init"))
	return rv
}

// Returns the name of the value function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cavaluefunction/1521888-name?language=objc
func (v_ ValueFunction) Name() ValueFunctionName {
	rv := objc.Call[ValueFunctionName](v_, objc.Sel("name"))
	return rv
}
