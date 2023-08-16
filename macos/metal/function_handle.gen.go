// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/objc"
)

// An object representing a function that you can add to a visible function table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionhandle?language=objc
type PFunctionHandle interface {
	// optional
	Name() string
	HasName() bool

	// optional
	Device() PDevice
	HasDevice() bool

	// optional
	FunctionType() FunctionType
	HasFunctionType() bool
}

// A concrete type wrapper for the [PFunctionHandle] protocol.
type FunctionHandleWrapper struct {
	objc.Object
}

func (f_ FunctionHandleWrapper) HasName() bool {
	return f_.RespondsToSelector(objc.Sel("name"))
}

// The function’s name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionhandle/3554004-name?language=objc
func (f_ FunctionHandleWrapper) Name() string {
	rv := objc.Call[string](f_, objc.Sel("name"))
	return rv
}

func (f_ FunctionHandleWrapper) HasDevice() bool {
	return f_.RespondsToSelector(objc.Sel("device"))
}

// The device object that created the shader function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionhandle/3554002-device?language=objc
func (f_ FunctionHandleWrapper) Device() DeviceWrapper {
	rv := objc.Call[DeviceWrapper](f_, objc.Sel("device"))
	return rv
}

func (f_ FunctionHandleWrapper) HasFunctionType() bool {
	return f_.RespondsToSelector(objc.Sel("functionType"))
}

// The shader function’s type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionhandle/3554003-functiontype?language=objc
func (f_ FunctionHandleWrapper) FunctionType() FunctionType {
	rv := objc.Call[FunctionType](f_, objc.Sel("functionType"))
	return rv
}
