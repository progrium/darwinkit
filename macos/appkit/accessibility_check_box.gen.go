// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a checkbox. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycheckbox?language=objc
type PAccessibilityCheckBox interface {
	// optional
	AccessibilityValue() foundation.INumber
	HasAccessibilityValue() bool
}

// A concrete type wrapper for the [PAccessibilityCheckBox] protocol.
type AccessibilityCheckBoxWrapper struct {
	objc.Object
}

func (a_ AccessibilityCheckBoxWrapper) HasAccessibilityValue() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityValue"))
}

// Returns the checkbox’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycheckbox/1524299-accessibilityvalue?language=objc
func (a_ AccessibilityCheckBoxWrapper) AccessibilityValue() foundation.Number {
	rv := objc.Call[foundation.Number](a_, objc.Sel("accessibilityValue"))
	return rv
}
