// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A table of intersection functions that Metal calls to perform ray-tracing intersection tests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlintersectionfunctiontable?language=objc
type PIntersectionFunctionTable interface {
	// optional
	SetOpaqueTriangleIntersectionFunctionWithSignatureAtIndex(signature IntersectionFunctionSignature, index uint)
	HasSetOpaqueTriangleIntersectionFunctionWithSignatureAtIndex() bool

	// optional
	SetBuffersOffsetsWithRange(buffers BufferWrapper, offsets *uint, range_ foundation.Range)
	HasSetBuffersOffsetsWithRange() bool

	// optional
	SetVisibleFunctionTableAtBufferIndex(functionTable VisibleFunctionTableWrapper, bufferIndex uint)
	HasSetVisibleFunctionTableAtBufferIndex() bool

	// optional
	SetFunctionAtIndex(function FunctionHandleWrapper, index uint)
	HasSetFunctionAtIndex() bool

	// optional
	SetVisibleFunctionTablesWithBufferRange(functionTables VisibleFunctionTableWrapper, bufferRange foundation.Range)
	HasSetVisibleFunctionTablesWithBufferRange() bool

	// optional
	SetFunctionsWithRange(functions FunctionHandleWrapper, range_ foundation.Range)
	HasSetFunctionsWithRange() bool

	// optional
	SetBufferOffsetAtIndex(buffer BufferWrapper, offset uint, index uint)
	HasSetBufferOffsetAtIndex() bool
}

// A concrete type wrapper for the [PIntersectionFunctionTable] protocol.
type IntersectionFunctionTableWrapper struct {
	objc.Object
}

func (i_ IntersectionFunctionTableWrapper) HasSetOpaqueTriangleIntersectionFunctionWithSignatureAtIndex() bool {
	return i_.RespondsToSelector(objc.Sel("setOpaqueTriangleIntersectionFunctionWithSignature:atIndex:"))
}

// Sets an entry in the intersection table to point to a system-defined opaque triangle intersection function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlintersectionfunctiontable/3667484-setopaquetriangleintersectionfun?language=objc
func (i_ IntersectionFunctionTableWrapper) SetOpaqueTriangleIntersectionFunctionWithSignatureAtIndex(signature IntersectionFunctionSignature, index uint) {
	objc.Call[objc.Void](i_, objc.Sel("setOpaqueTriangleIntersectionFunctionWithSignature:atIndex:"), signature, index)
}

func (i_ IntersectionFunctionTableWrapper) HasSetBuffersOffsetsWithRange() bool {
	return i_.RespondsToSelector(objc.Sel("setBuffers:offsets:withRange:"))
}

// Sets a range of buffers for the intersection functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlintersectionfunctiontable/3578280-setbuffers?language=objc
func (i_ IntersectionFunctionTableWrapper) SetBuffersOffsetsWithRange(buffers PBuffer, offsets *uint, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffers)
	objc.Call[objc.Void](i_, objc.Sel("setBuffers:offsets:withRange:"), po0, offsets, range_)
}

func (i_ IntersectionFunctionTableWrapper) HasSetVisibleFunctionTableAtBufferIndex() bool {
	return i_.RespondsToSelector(objc.Sel("setVisibleFunctionTable:atBufferIndex:"))
}

// Sets a visible function table for the intersection functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlintersectionfunctiontable/3674230-setvisiblefunctiontable?language=objc
func (i_ IntersectionFunctionTableWrapper) SetVisibleFunctionTableAtBufferIndex(functionTable PVisibleFunctionTable, bufferIndex uint) {
	po0 := objc.WrapAsProtocol("MTLVisibleFunctionTable", functionTable)
	objc.Call[objc.Void](i_, objc.Sel("setVisibleFunctionTable:atBufferIndex:"), po0, bufferIndex)
}

func (i_ IntersectionFunctionTableWrapper) HasSetFunctionAtIndex() bool {
	return i_.RespondsToSelector(objc.Sel("setFunction:atIndex:"))
}

// Sets an entry in the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlintersectionfunctiontable/3566554-setfunction?language=objc
func (i_ IntersectionFunctionTableWrapper) SetFunctionAtIndex(function PFunctionHandle, index uint) {
	po0 := objc.WrapAsProtocol("MTLFunctionHandle", function)
	objc.Call[objc.Void](i_, objc.Sel("setFunction:atIndex:"), po0, index)
}

func (i_ IntersectionFunctionTableWrapper) HasSetVisibleFunctionTablesWithBufferRange() bool {
	return i_.RespondsToSelector(objc.Sel("setVisibleFunctionTables:withBufferRange:"))
}

// Sets a range of visible function tables for the intersection functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlintersectionfunctiontable/3674231-setvisiblefunctiontables?language=objc
func (i_ IntersectionFunctionTableWrapper) SetVisibleFunctionTablesWithBufferRange(functionTables PVisibleFunctionTable, bufferRange foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLVisibleFunctionTable", functionTables)
	objc.Call[objc.Void](i_, objc.Sel("setVisibleFunctionTables:withBufferRange:"), po0, bufferRange)
}

func (i_ IntersectionFunctionTableWrapper) HasSetFunctionsWithRange() bool {
	return i_.RespondsToSelector(objc.Sel("setFunctions:withRange:"))
}

// Sets a range of entries in the table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlintersectionfunctiontable/3566555-setfunctions?language=objc
func (i_ IntersectionFunctionTableWrapper) SetFunctionsWithRange(functions PFunctionHandle, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLFunctionHandle", functions)
	objc.Call[objc.Void](i_, objc.Sel("setFunctions:withRange:"), po0, range_)
}

func (i_ IntersectionFunctionTableWrapper) HasSetBufferOffsetAtIndex() bool {
	return i_.RespondsToSelector(objc.Sel("setBuffer:offset:atIndex:"))
}

// Sets a buffer for the intersection functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlintersectionfunctiontable/3578279-setbuffer?language=objc
func (i_ IntersectionFunctionTableWrapper) SetBufferOffsetAtIndex(buffer PBuffer, offset uint, index uint) {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffer)
	objc.Call[objc.Void](i_, objc.Sel("setBuffer:offset:atIndex:"), po0, offset, index)
}
