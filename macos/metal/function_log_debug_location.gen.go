// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The source code that logged a debug message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionlogdebuglocation?language=objc
type PFunctionLogDebugLocation interface {
	// optional
	Line() uint
	HasLine() bool

	// optional
	FunctionName() string
	HasFunctionName() bool

	// optional
	Column() uint
	HasColumn() bool

	// optional
	URL() foundation.IURL
	HasURL() bool
}

// A concrete type wrapper for the [PFunctionLogDebugLocation] protocol.
type FunctionLogDebugLocationWrapper struct {
	objc.Object
}

func (f_ FunctionLogDebugLocationWrapper) HasLine() bool {
	return f_.RespondsToSelector(objc.Sel("line"))
}

// The line that the log message appears on. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionlogdebuglocation/3554022-line?language=objc
func (f_ FunctionLogDebugLocationWrapper) Line() uint {
	rv := objc.Call[uint](f_, objc.Sel("line"))
	return rv
}

func (f_ FunctionLogDebugLocationWrapper) HasFunctionName() bool {
	return f_.RespondsToSelector(objc.Sel("functionName"))
}

// The name of the shader function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionlogdebuglocation/3554020-functionname?language=objc
func (f_ FunctionLogDebugLocationWrapper) FunctionName() string {
	rv := objc.Call[string](f_, objc.Sel("functionName"))
	return rv
}

func (f_ FunctionLogDebugLocationWrapper) HasColumn() bool {
	return f_.RespondsToSelector(objc.Sel("column"))
}

// The column where the log message appears. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionlogdebuglocation/3554019-column?language=objc
func (f_ FunctionLogDebugLocationWrapper) Column() uint {
	rv := objc.Call[uint](f_, objc.Sel("column"))
	return rv
}

func (f_ FunctionLogDebugLocationWrapper) HasURL() bool {
	return f_.RespondsToSelector(objc.Sel("URL"))
}

// The URL of the file that contains the shader function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionlogdebuglocation/3554018-url?language=objc
func (f_ FunctionLogDebugLocationWrapper) URL() foundation.URL {
	rv := objc.Call[foundation.URL](f_, objc.Sel("URL"))
	return rv
}
