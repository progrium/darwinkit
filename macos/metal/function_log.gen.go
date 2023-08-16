// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/objc"
)

// A log entry a Metal device generates when the it runs a command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionlog?language=objc
type PFunctionLog interface {
	// optional
	Function() PFunction
	HasFunction() bool

	// optional
	DebugLocation() PFunctionLogDebugLocation
	HasDebugLocation() bool

	// optional
	EncoderLabel() string
	HasEncoderLabel() bool

	// optional
	Type() FunctionLogType
	HasType() bool
}

// A concrete type wrapper for the [PFunctionLog] protocol.
type FunctionLogWrapper struct {
	objc.Object
}

func (f_ FunctionLogWrapper) HasFunction() bool {
	return f_.RespondsToSelector(objc.Sel("function"))
}

// When known, the function object corresponding to the logged message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionlog/3554010-function?language=objc
func (f_ FunctionLogWrapper) Function() FunctionWrapper {
	rv := objc.Call[FunctionWrapper](f_, objc.Sel("function"))
	return rv
}

func (f_ FunctionLogWrapper) HasDebugLocation() bool {
	return f_.RespondsToSelector(objc.Sel("debugLocation"))
}

// If known, the location of the logging command within a shader source file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionlog/3554007-debuglocation?language=objc
func (f_ FunctionLogWrapper) DebugLocation() FunctionLogDebugLocationWrapper {
	rv := objc.Call[FunctionLogDebugLocationWrapper](f_, objc.Sel("debugLocation"))
	return rv
}

func (f_ FunctionLogWrapper) HasEncoderLabel() bool {
	return f_.RespondsToSelector(objc.Sel("encoderLabel"))
}

// The label for the encoder that logged the message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionlog/3577673-encoderlabel?language=objc
func (f_ FunctionLogWrapper) EncoderLabel() string {
	rv := objc.Call[string](f_, objc.Sel("encoderLabel"))
	return rv
}

func (f_ FunctionLogWrapper) HasType() bool {
	return f_.RespondsToSelector(objc.Sel("type"))
}

// The type of message that was logged. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionlog/3554016-type?language=objc
func (f_ FunctionLogWrapper) Type() FunctionLogType {
	rv := objc.Call[FunctionLogType](f_, objc.Sel("type"))
	return rv
}
