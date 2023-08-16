// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [StencilDescriptor] class.
var StencilDescriptorClass = _StencilDescriptorClass{objc.GetClass("MTLStencilDescriptor")}

type _StencilDescriptorClass struct {
	objc.Class
}

// An interface definition for the [StencilDescriptor] class.
type IStencilDescriptor interface {
	objc.IObject
	StencilFailureOperation() StencilOperation
	SetStencilFailureOperation(value StencilOperation)
	DepthFailureOperation() StencilOperation
	SetDepthFailureOperation(value StencilOperation)
	WriteMask() uint32
	SetWriteMask(value uint32)
	ReadMask() uint32
	SetReadMask(value uint32)
	StencilCompareFunction() CompareFunction
	SetStencilCompareFunction(value CompareFunction)
	DepthStencilPassOperation() StencilOperation
	SetDepthStencilPassOperation(value StencilOperation)
}

// An object that defines the front-facing or back-facing stencil operations of a depth and stencil state object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstencildescriptor?language=objc
type StencilDescriptor struct {
	objc.Object
}

func StencilDescriptorFrom(ptr unsafe.Pointer) StencilDescriptor {
	return StencilDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _StencilDescriptorClass) Alloc() StencilDescriptor {
	rv := objc.Call[StencilDescriptor](sc, objc.Sel("alloc"))
	return rv
}

func StencilDescriptor_Alloc() StencilDescriptor {
	return StencilDescriptorClass.Alloc()
}

func (sc _StencilDescriptorClass) New() StencilDescriptor {
	rv := objc.Call[StencilDescriptor](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewStencilDescriptor() StencilDescriptor {
	return StencilDescriptorClass.New()
}

func (s_ StencilDescriptor) Init() StencilDescriptor {
	rv := objc.Call[StencilDescriptor](s_, objc.Sel("init"))
	return rv
}

// The operation that is performed to update the values in the stencil attachment when the stencil test fails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstencildescriptor/1462471-stencilfailureoperation?language=objc
func (s_ StencilDescriptor) StencilFailureOperation() StencilOperation {
	rv := objc.Call[StencilOperation](s_, objc.Sel("stencilFailureOperation"))
	return rv
}

// The operation that is performed to update the values in the stencil attachment when the stencil test fails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstencildescriptor/1462471-stencilfailureoperation?language=objc
func (s_ StencilDescriptor) SetStencilFailureOperation(value StencilOperation) {
	objc.Call[objc.Void](s_, objc.Sel("setStencilFailureOperation:"), value)
}

// The operation that is performed to update the values in the stencil attachment when the stencil test passes, but the depth test fails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstencildescriptor/1462500-depthfailureoperation?language=objc
func (s_ StencilDescriptor) DepthFailureOperation() StencilOperation {
	rv := objc.Call[StencilOperation](s_, objc.Sel("depthFailureOperation"))
	return rv
}

// The operation that is performed to update the values in the stencil attachment when the stencil test passes, but the depth test fails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstencildescriptor/1462500-depthfailureoperation?language=objc
func (s_ StencilDescriptor) SetDepthFailureOperation(value StencilOperation) {
	objc.Call[objc.Void](s_, objc.Sel("setDepthFailureOperation:"), value)
}

// A bitmask that determines to which bits that stencil operations can write. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstencildescriptor/1462496-writemask?language=objc
func (s_ StencilDescriptor) WriteMask() uint32 {
	rv := objc.Call[uint32](s_, objc.Sel("writeMask"))
	return rv
}

// A bitmask that determines to which bits that stencil operations can write. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstencildescriptor/1462496-writemask?language=objc
func (s_ StencilDescriptor) SetWriteMask(value uint32) {
	objc.Call[objc.Void](s_, objc.Sel("setWriteMask:"), value)
}

// A bitmask that determines from which bits that stencil comparison tests can read. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstencildescriptor/1462465-readmask?language=objc
func (s_ StencilDescriptor) ReadMask() uint32 {
	rv := objc.Call[uint32](s_, objc.Sel("readMask"))
	return rv
}

// A bitmask that determines from which bits that stencil comparison tests can read. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstencildescriptor/1462465-readmask?language=objc
func (s_ StencilDescriptor) SetReadMask(value uint32) {
	objc.Call[objc.Void](s_, objc.Sel("setReadMask:"), value)
}

// The comparison that is performed between the masked reference value and a masked value in the stencil attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstencildescriptor/1462455-stencilcomparefunction?language=objc
func (s_ StencilDescriptor) StencilCompareFunction() CompareFunction {
	rv := objc.Call[CompareFunction](s_, objc.Sel("stencilCompareFunction"))
	return rv
}

// The comparison that is performed between the masked reference value and a masked value in the stencil attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstencildescriptor/1462455-stencilcomparefunction?language=objc
func (s_ StencilDescriptor) SetStencilCompareFunction(value CompareFunction) {
	objc.Call[objc.Void](s_, objc.Sel("setStencilCompareFunction:"), value)
}

// The operation that is performed to update the values in the stencil attachment when both the stencil test and the depth test pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstencildescriptor/1462486-depthstencilpassoperation?language=objc
func (s_ StencilDescriptor) DepthStencilPassOperation() StencilOperation {
	rv := objc.Call[StencilOperation](s_, objc.Sel("depthStencilPassOperation"))
	return rv
}

// The operation that is performed to update the values in the stencil attachment when both the stencil test and the depth test pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstencildescriptor/1462486-depthstencilpassoperation?language=objc
func (s_ StencilDescriptor) SetDepthStencilPassOperation(value StencilOperation) {
	objc.Call[objc.Void](s_, objc.Sel("setDepthStencilPassOperation:"), value)
}
