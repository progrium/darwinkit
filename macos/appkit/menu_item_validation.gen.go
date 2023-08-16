// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemvalidation?language=objc
type PMenuItemValidation interface {
	// optional
	ValidateMenuItem(menuItem MenuItem) bool
	HasValidateMenuItem() bool
}

// A concrete type wrapper for the [PMenuItemValidation] protocol.
type MenuItemValidationWrapper struct {
	objc.Object
}

func (m_ MenuItemValidationWrapper) HasValidateMenuItem() bool {
	return m_.RespondsToSelector(objc.Sel("validateMenuItem:"))
}

// Implemented to override the default action of enabling or disabling a specific menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemvalidation/3005191-validatemenuitem?language=objc
func (m_ MenuItemValidationWrapper) ValidateMenuItem(menuItem IMenuItem) bool {
	rv := objc.Call[bool](m_, objc.Sel("validateMenuItem:"), objc.Ptr(menuItem))
	return rv
}
