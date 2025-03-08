package javascriptcore

import (
	"github.com/progrium/darwinkit/objc"
	"github.com/progrium/darwinkit/macos/foundation"
)

// Context represents a JavaScript execution context
type Context struct {
	objc.Object
}

// ContextClass is the class instance for Context
var ContextClass = objc.GetClass("JSContext")

// NewContext creates a new JavaScript context
func NewContext() Context {
	alloc := objc.Call[objc.Object](ContextClass, objc.Sel("alloc"))
	initialized := objc.Call[objc.Object](alloc, objc.Sel("init"))
	return Context{initialized}
}

// EvaluateScript evaluates a JavaScript script in the context
func (c Context) EvaluateScript(script foundation.String, exception *Value) Value {
	return Value{objc.Call[objc.Object](c, objc.Sel("evaluateScript:exception:"), script, exception)}
}

// SetExceptionHandler sets the exception handler for the context
func (c Context) SetExceptionHandler(handler objc.Object) {
	objc.Call[objc.Void](c, objc.Sel("setExceptionHandler:"), handler)
}

// GlobalObject returns the global object for the context
func (c Context) GlobalObject() Value {
	return Value{objc.Call[objc.Object](c, objc.Sel("globalObject"))}
}

// Value represents a JavaScript value
type Value struct {
	objc.Object
}

// ValueClass is the class instance for Value
var ValueClass = objc.GetClass("JSValue")

// NewValueWithBool creates a new JavaScript boolean value
func NewValueWithBool(context Context, value bool) Value {
	return Value{objc.Call[objc.Object](ValueClass, objc.Sel("valueWithBool:inContext:"), value, context)}
}

// NewValueWithDouble creates a new JavaScript number value
func NewValueWithDouble(context Context, value float64) Value {
	return Value{objc.Call[objc.Object](ValueClass, objc.Sel("valueWithDouble:inContext:"), value, context)}
}

// NewValueWithInt32 creates a new JavaScript number value
func NewValueWithInt32(context Context, value int32) Value {
	return Value{objc.Call[objc.Object](ValueClass, objc.Sel("valueWithInt32:inContext:"), value, context)}
}

// NewValueWithUInt32 creates a new JavaScript number value
func NewValueWithUInt32(context Context, value uint32) Value {
	return Value{objc.Call[objc.Object](ValueClass, objc.Sel("valueWithUInt32:inContext:"), value, context)}
}

// NewValueWithString creates a new JavaScript string value
func NewValueWithString(context Context, value foundation.String) Value {
	return Value{objc.Call[objc.Object](ValueClass, objc.Sel("valueWithObject:inContext:"), value, context)}
}

// NewValueWithNewObject creates a new JavaScript object value
func NewValueWithNewObject(context Context) Value {
	return Value{objc.Call[objc.Object](ValueClass, objc.Sel("valueWithNewObjectInContext:"), context)}
}

// NewValueWithNewArray creates a new JavaScript array value
func NewValueWithNewArray(context Context, values []Value) Value {
	array := foundation.ArrayClass.Array()
	for _, value := range values {
		objc.Call[objc.Void](array, objc.Sel("addObject:"), value)
	}
	return Value{objc.Call[objc.Object](ValueClass, objc.Sel("valueWithNewArrayInContext:"), context)}
}

// ToBool converts the JavaScript value to a boolean
func (v Value) ToBool() bool {
	return objc.Call[bool](v, objc.Sel("toBool"))
}

// ToDouble converts the JavaScript value to a double
func (v Value) ToDouble() float64 {
	return objc.Call[float64](v, objc.Sel("toDouble"))
}

// ToInt32 converts the JavaScript value to an int32
func (v Value) ToInt32() int32 {
	return objc.Call[int32](v, objc.Sel("toInt32"))
}

// ToUInt32 converts the JavaScript value to a uint32
func (v Value) ToUInt32() uint32 {
	return objc.Call[uint32](v, objc.Sel("toUInt32"))
}

// ToString converts the JavaScript value to a string
func (v Value) ToString() foundation.String {
	return objc.Call[foundation.String](v, objc.Sel("toString"))
}

// IsString returns whether the JavaScript value is a string
func (v Value) IsString() bool {
	return objc.Call[bool](v, objc.Sel("isString"))
}

// IsNumber returns whether the JavaScript value is a number
func (v Value) IsNumber() bool {
	return objc.Call[bool](v, objc.Sel("isNumber"))
}

// IsBoolean returns whether the JavaScript value is a boolean
func (v Value) IsBoolean() bool {
	return objc.Call[bool](v, objc.Sel("isBoolean"))
}

// IsObject returns whether the JavaScript value is an object
func (v Value) IsObject() bool {
	return objc.Call[bool](v, objc.Sel("isObject"))
}

// IsArray returns whether the JavaScript value is an array
func (v Value) IsArray() bool {
	return objc.Call[bool](v, objc.Sel("isArray"))
}

// IsDate returns whether the JavaScript value is a date
func (v Value) IsDate() bool {
	return objc.Call[bool](v, objc.Sel("isDate"))
}

// IsNull returns whether the JavaScript value is null
func (v Value) IsNull() bool {
	return objc.Call[bool](v, objc.Sel("isNull"))
}

// IsUndefined returns whether the JavaScript value is undefined
func (v Value) IsUndefined() bool {
	return objc.Call[bool](v, objc.Sel("isUndefined"))
}

// VirtualMachine represents a JavaScript virtual machine
type VirtualMachine struct {
	objc.Object
}

// VirtualMachineClass is the class instance for VirtualMachine
var VirtualMachineClass = objc.GetClass("JSVirtualMachine")

// NewVirtualMachine creates a new JavaScript virtual machine
func NewVirtualMachine() VirtualMachine {
	alloc := objc.Call[objc.Object](VirtualMachineClass, objc.Sel("alloc"))
	initialized := objc.Call[objc.Object](alloc, objc.Sel("init"))
	return VirtualMachine{initialized}
}

// Context creates a new JavaScript context for the virtual machine
func (vm VirtualMachine) Context() Context {
	alloc := objc.Call[objc.Object](ContextClass, objc.Sel("alloc"))
	initialized := objc.Call[objc.Object](alloc, objc.Sel("initWithVirtualMachine:"), vm)
	return Context{initialized}
}

// AddManagedReference adds a managed reference to the virtual machine
func (vm VirtualMachine) AddManagedReference(object objc.Object, owner objc.Object) {
	objc.Call[objc.Void](vm, objc.Sel("addManagedReference:withOwner:"), object, owner)
}

// RemoveManagedReference removes a managed reference from the virtual machine
func (vm VirtualMachine) RemoveManagedReference(object objc.Object, owner objc.Object) {
	objc.Call[objc.Void](vm, objc.Sel("removeManagedReference:withOwner:"), object, owner)
}