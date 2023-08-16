// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitybutton?language=objc
type PAccessibilityButton interface {
	// optional
	AccessibilityLabel() string
	HasAccessibilityLabel() bool

	// optional
	AccessibilityPerformPress() bool
	HasAccessibilityPerformPress() bool
}

// A concrete type wrapper for the [PAccessibilityButton] protocol.
type AccessibilityButtonWrapper struct {
	objc.Object
}

func (a_ AccessibilityButtonWrapper) HasAccessibilityLabel() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityLabel"))
}

// Returns a short description of the button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitybutton/1524910-accessibilitylabel?language=objc
func (a_ AccessibilityButtonWrapper) AccessibilityLabel() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityLabel"))
	return rv
}

func (a_ AccessibilityButtonWrapper) HasAccessibilityPerformPress() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformPress"))
}

// Simulates clicking the button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitybutton/1525542-accessibilityperformpress?language=objc
func (a_ AccessibilityButtonWrapper) AccessibilityPerformPress() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformPress"))
	return rv
}
