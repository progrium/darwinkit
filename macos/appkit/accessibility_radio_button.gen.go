// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a radio button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityradiobutton?language=objc
type PAccessibilityRadioButton interface {
	// optional
	AccessibilityValue() foundation.INumber
	HasAccessibilityValue() bool
}

// A concrete type wrapper for the [PAccessibilityRadioButton] protocol.
type AccessibilityRadioButtonWrapper struct {
	objc.Object
}

func (a_ AccessibilityRadioButtonWrapper) HasAccessibilityValue() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityValue"))
}

// Returns the radio button’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityradiobutton/1526534-accessibilityvalue?language=objc
func (a_ AccessibilityRadioButtonWrapper) AccessibilityValue() foundation.Number {
	rv := objc.Call[foundation.Number](a_, objc.Sel("accessibilityValue"))
	return rv
}
