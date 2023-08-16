// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/objc"
)

// A container that provides additional information about a runtime failure a GPU encounters as it runs the commands in a command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbufferencoderinfo?language=objc
type PCommandBufferEncoderInfo interface {
	// optional
	ErrorState() CommandEncoderErrorState
	HasErrorState() bool

	// optional
	Label() string
	HasLabel() bool

	// optional
	DebugSignposts() []string
	HasDebugSignposts() bool
}

// A concrete type wrapper for the [PCommandBufferEncoderInfo] protocol.
type CommandBufferEncoderInfoWrapper struct {
	objc.Object
}

func (c_ CommandBufferEncoderInfoWrapper) HasErrorState() bool {
	return c_.RespondsToSelector(objc.Sel("errorState"))
}

// The execution status of the command encoder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbufferencoderinfo/3553945-errorstate?language=objc
func (c_ CommandBufferEncoderInfoWrapper) ErrorState() CommandEncoderErrorState {
	rv := objc.Call[CommandEncoderErrorState](c_, objc.Sel("errorState"))
	return rv
}

func (c_ CommandBufferEncoderInfoWrapper) HasLabel() bool {
	return c_.RespondsToSelector(objc.Sel("label"))
}

// The name of the encoder that generates the error information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbufferencoderinfo/3553946-label?language=objc
func (c_ CommandBufferEncoderInfoWrapper) Label() string {
	rv := objc.Call[string](c_, objc.Sel("label"))
	return rv
}

func (c_ CommandBufferEncoderInfoWrapper) HasDebugSignposts() bool {
	return c_.RespondsToSelector(objc.Sel("debugSignposts"))
}

// An array of debug signposts that Metal records as the GPU executes the commands of the encoder’s pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbufferencoderinfo/3553944-debugsignposts?language=objc
func (c_ CommandBufferEncoderInfoWrapper) DebugSignposts() []string {
	rv := objc.Call[[]string](c_, objc.Sel("debugSignposts"))
	return rv
}
