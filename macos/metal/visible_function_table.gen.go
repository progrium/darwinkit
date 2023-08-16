// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A table of shader functions visible to your app that you can pass into compute commands to customize the behavior of a shader. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvisiblefunctiontable?language=objc
type PVisibleFunctionTable interface {
	// optional
	SetFunctionAtIndex(function FunctionHandleWrapper, index uint)
	HasSetFunctionAtIndex() bool

	// optional
	SetFunctionsWithRange(functions FunctionHandleWrapper, range_ foundation.Range)
	HasSetFunctionsWithRange() bool
}

// A concrete type wrapper for the [PVisibleFunctionTable] protocol.
type VisibleFunctionTableWrapper struct {
	objc.Object
}

func (v_ VisibleFunctionTableWrapper) HasSetFunctionAtIndex() bool {
	return v_.RespondsToSelector(objc.Sel("setFunction:atIndex:"))
}

// Sets a table entry to point to a callable function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvisiblefunctiontable/3554056-setfunction?language=objc
func (v_ VisibleFunctionTableWrapper) SetFunctionAtIndex(function PFunctionHandle, index uint) {
	po0 := objc.WrapAsProtocol("MTLFunctionHandle", function)
	objc.Call[objc.Void](v_, objc.Sel("setFunction:atIndex:"), po0, index)
}

func (v_ VisibleFunctionTableWrapper) HasSetFunctionsWithRange() bool {
	return v_.RespondsToSelector(objc.Sel("setFunctions:withRange:"))
}

// Sets a range of table entries to point to an array of callable functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvisiblefunctiontable/3554057-setfunctions?language=objc
func (v_ VisibleFunctionTableWrapper) SetFunctionsWithRange(functions PFunctionHandle, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLFunctionHandle", functions)
	objc.Call[objc.Void](v_, objc.Sel("setFunctions:withRange:"), po0, range_)
}
