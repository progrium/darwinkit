// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that a custom class can adopt to manage the enabled state of a UI element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacevalidations?language=objc
type PUserInterfaceValidations interface {
	// optional
	ValidateUserInterfaceItem(item ValidatedUserInterfaceItemWrapper) bool
	HasValidateUserInterfaceItem() bool
}

// A concrete type wrapper for the [PUserInterfaceValidations] protocol.
type UserInterfaceValidationsWrapper struct {
	objc.Object
}

func (u_ UserInterfaceValidationsWrapper) HasValidateUserInterfaceItem() bool {
	return u_.RespondsToSelector(objc.Sel("validateUserInterfaceItem:"))
}

// Returns a Boolean value that indicates whether the sender should be enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacevalidations/1528162-validateuserinterfaceitem?language=objc
func (u_ UserInterfaceValidationsWrapper) ValidateUserInterfaceItem(item PValidatedUserInterfaceItem) bool {
	po0 := objc.WrapAsProtocol("NSValidatedUserInterfaceItem", item)
	rv := objc.Call[bool](u_, objc.Sel("validateUserInterfaceItem:"), po0)
	return rv
}
