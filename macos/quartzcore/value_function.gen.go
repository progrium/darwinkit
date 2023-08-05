// AUTO-GENERATED CODE, DO NOT MODIFY
package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var ValueFunctionClass = _ValueFunctionClass{objc.GetClass("CAValueFunction")}

type _ValueFunctionClass struct {
	objc.Class
}

type IValueFunction interface {
	objc.IObject
	Name() ValueFunctionName
}

type ValueFunction struct {
	objc.Object
}

func MakeValueFunction(ptr unsafe.Pointer) ValueFunction {
	return ValueFunction{
		Object: objc.MakeObject(ptr),
	}
}

func (vc _ValueFunctionClass) FunctionWithName(name ValueFunctionName) ValueFunction {
	rv := objc.CallMethod[ValueFunction](vc, objc.GetSelector("functionWithName:"), name)
	return rv
}

func ValueFunction_FunctionWithName(name ValueFunctionName) ValueFunction {
	return ValueFunctionClass.FunctionWithName(name)
}

func (vc _ValueFunctionClass) Alloc() ValueFunction {
	rv := objc.CallMethod[ValueFunction](vc, objc.GetSelector("alloc"))
	return rv
}

func ValueFunction_Alloc() ValueFunction {
	return ValueFunctionClass.Alloc()
}

func (vc _ValueFunctionClass) New() ValueFunction {
	rv := objc.CallMethod[ValueFunction](vc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewValueFunction() ValueFunction {
	return ValueFunctionClass.New()
}

func ValueFunction_New() ValueFunction {
	return ValueFunctionClass.New()
}

func (v_ ValueFunction) Init() ValueFunction {
	rv := objc.CallMethod[ValueFunction](v_, objc.GetSelector("init"))
	return rv
}

func ValueFunction_Init() ValueFunction {
	return ValueFunctionClass.Alloc().Init()
}

func (v_ ValueFunction) Name() ValueFunctionName {
	rv := objc.CallMethod[ValueFunctionName](v_, objc.GetSelector("name"))
	return rv
}
