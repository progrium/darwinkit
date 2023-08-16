// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A command buffer containing reusable commands, encoded either on the CPU or GPU. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcommandbuffer?language=objc
type PIndirectCommandBuffer interface {
	// optional
	IndirectComputeCommandAtIndex(commandIndex uint) PIndirectComputeCommand
	HasIndirectComputeCommandAtIndex() bool

	// optional
	IndirectRenderCommandAtIndex(commandIndex uint) PIndirectRenderCommand
	HasIndirectRenderCommandAtIndex() bool

	// optional
	ResetWithRange(range_ foundation.Range)
	HasResetWithRange() bool

	// optional
	Size() uint
	HasSize() bool
}

// A concrete type wrapper for the [PIndirectCommandBuffer] protocol.
type IndirectCommandBufferWrapper struct {
	objc.Object
}

func (i_ IndirectCommandBufferWrapper) HasIndirectComputeCommandAtIndex() bool {
	return i_.RespondsToSelector(objc.Sel("indirectComputeCommandAtIndex:"))
}

// Gets the compute command at the given index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcommandbuffer/2967430-indirectcomputecommandatindex?language=objc
func (i_ IndirectCommandBufferWrapper) IndirectComputeCommandAtIndex(commandIndex uint) IndirectComputeCommandWrapper {
	rv := objc.Call[IndirectComputeCommandWrapper](i_, objc.Sel("indirectComputeCommandAtIndex:"), commandIndex)
	return rv
}

func (i_ IndirectCommandBufferWrapper) HasIndirectRenderCommandAtIndex() bool {
	return i_.RespondsToSelector(objc.Sel("indirectRenderCommandAtIndex:"))
}

// Gets the render command at the given index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcommandbuffer/2967431-indirectrendercommandatindex?language=objc
func (i_ IndirectCommandBufferWrapper) IndirectRenderCommandAtIndex(commandIndex uint) IndirectRenderCommandWrapper {
	rv := objc.Call[IndirectRenderCommandWrapper](i_, objc.Sel("indirectRenderCommandAtIndex:"), commandIndex)
	return rv
}

func (i_ IndirectCommandBufferWrapper) HasResetWithRange() bool {
	return i_.RespondsToSelector(objc.Sel("resetWithRange:"))
}

// Resets a range of commands to their default state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcommandbuffer/3012870-resetwithrange?language=objc
func (i_ IndirectCommandBufferWrapper) ResetWithRange(range_ foundation.Range) {
	objc.Call[objc.Void](i_, objc.Sel("resetWithRange:"), range_)
}

func (i_ IndirectCommandBufferWrapper) HasSize() bool {
	return i_.RespondsToSelector(objc.Sel("size"))
}

// The number of commands contained in the indirect command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcommandbuffer/2981028-size?language=objc
func (i_ IndirectCommandBufferWrapper) Size() uint {
	rv := objc.Call[uint](i_, objc.Sel("size"))
	return rv
}
