// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/objc"
)

// An encoder that writes GPU commands into a command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandencoder?language=objc
type PCommandEncoder interface {
	// optional
	PushDebugGroup(string_ string)
	HasPushDebugGroup() bool

	// optional
	PopDebugGroup()
	HasPopDebugGroup() bool

	// optional
	InsertDebugSignpost(string_ string)
	HasInsertDebugSignpost() bool

	// optional
	EndEncoding()
	HasEndEncoding() bool

	// optional
	Device() PDevice
	HasDevice() bool

	// optional
	SetLabel(value string)
	HasSetLabel() bool

	// optional
	Label() string
	HasLabel() bool
}

// A concrete type wrapper for the [PCommandEncoder] protocol.
type CommandEncoderWrapper struct {
	objc.Object
}

func (c_ CommandEncoderWrapper) HasPushDebugGroup() bool {
	return c_.RespondsToSelector(objc.Sel("pushDebugGroup:"))
}

// Pushes a specific string onto a stack of debug group strings for the command encoder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandencoder/1458041-pushdebuggroup?language=objc
func (c_ CommandEncoderWrapper) PushDebugGroup(string_ string) {
	objc.Call[objc.Void](c_, objc.Sel("pushDebugGroup:"), string_)
}

func (c_ CommandEncoderWrapper) HasPopDebugGroup() bool {
	return c_.RespondsToSelector(objc.Sel("popDebugGroup"))
}

// Pops the latest string off of a stack of debug group strings for the command encoder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandencoder/1458040-popdebuggroup?language=objc
func (c_ CommandEncoderWrapper) PopDebugGroup() {
	objc.Call[objc.Void](c_, objc.Sel("popDebugGroup"))
}

func (c_ CommandEncoderWrapper) HasInsertDebugSignpost() bool {
	return c_.RespondsToSelector(objc.Sel("insertDebugSignpost:"))
}

// Inserts a debug string into the captured frame data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandencoder/1458034-insertdebugsignpost?language=objc
func (c_ CommandEncoderWrapper) InsertDebugSignpost(string_ string) {
	objc.Call[objc.Void](c_, objc.Sel("insertDebugSignpost:"), string_)
}

func (c_ CommandEncoderWrapper) HasEndEncoding() bool {
	return c_.RespondsToSelector(objc.Sel("endEncoding"))
}

// Declares that all command generation from the encoder is completed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandencoder/1458038-endencoding?language=objc
func (c_ CommandEncoderWrapper) EndEncoding() {
	objc.Call[objc.Void](c_, objc.Sel("endEncoding"))
}

func (c_ CommandEncoderWrapper) HasDevice() bool {
	return c_.RespondsToSelector(objc.Sel("device"))
}

// The Metal device from which the command encoder was created. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandencoder/1458032-device?language=objc
func (c_ CommandEncoderWrapper) Device() DeviceWrapper {
	rv := objc.Call[DeviceWrapper](c_, objc.Sel("device"))
	return rv
}

func (c_ CommandEncoderWrapper) HasSetLabel() bool {
	return c_.RespondsToSelector(objc.Sel("setLabel:"))
}

// A string that labels the command encoder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandencoder/1458036-label?language=objc
func (c_ CommandEncoderWrapper) SetLabel(value string) {
	objc.Call[objc.Void](c_, objc.Sel("setLabel:"), value)
}

func (c_ CommandEncoderWrapper) HasLabel() bool {
	return c_.RespondsToSelector(objc.Sel("label"))
}

// A string that labels the command encoder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandencoder/1458036-label?language=objc
func (c_ CommandEncoderWrapper) Label() string {
	rv := objc.Call[string](c_, objc.Sel("label"))
	return rv
}
