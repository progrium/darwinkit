// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type IUserInterfaceCompression interface {
	// required
	CompressWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions)
	// required
	MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions) foundation.Size
	ImplementsActiveCompressionOptions() bool
	// optional
	ActiveCompressionOptions() IUserInterfaceCompressionOptions
}

type UserInterfaceCompressionWrapper struct {
	objc.Object
}

func (u_ UserInterfaceCompressionWrapper) CompressWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("compressWithPrioritizedCompressionOptions:"), prioritizedOptions)
}

func (u_ UserInterfaceCompressionWrapper) MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) foundation.Size {
	rv := objc.CallMethod[foundation.Size](u_, objc.GetSelector("minimumSizeWithPrioritizedCompressionOptions:"), prioritizedOptions)
	return rv
}

func (u_ UserInterfaceCompressionWrapper) ImplementsActiveCompressionOptions() bool {
	return u_.RespondsToSelector(objc.GetSelector("activeCompressionOptions"))
}

func (u_ UserInterfaceCompressionWrapper) ActiveCompressionOptions() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](u_, objc.GetSelector("activeCompressionOptions"))
	return rv
}
