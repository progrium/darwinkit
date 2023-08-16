// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// Validation of a toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritemvalidation?language=objc
type PToolbarItemValidation interface {
	// optional
	ValidateToolbarItem(item ToolbarItem) bool
	HasValidateToolbarItem() bool
}

// A concrete type wrapper for the [PToolbarItemValidation] protocol.
type ToolbarItemValidationWrapper struct {
	objc.Object
}

func (t_ ToolbarItemValidationWrapper) HasValidateToolbarItem() bool {
	return t_.RespondsToSelector(objc.Sel("validateToolbarItem:"))
}

// If this method is implemented and returns NO, NSToolbar will disable theItem; returning YES causes theItem to be enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritemvalidation/3005292-validatetoolbaritem?language=objc
func (t_ ToolbarItemValidationWrapper) ValidateToolbarItem(item IToolbarItem) bool {
	rv := objc.Call[bool](t_, objc.Sel("validateToolbarItem:"), objc.Ptr(item))
	return rv
}
