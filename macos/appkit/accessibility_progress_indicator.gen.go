// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a progress indicator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityprogressindicator?language=objc
type PAccessibilityProgressIndicator interface {
	// optional
	AccessibilityValue() foundation.INumber
	HasAccessibilityValue() bool
}

// A concrete type wrapper for the [PAccessibilityProgressIndicator] protocol.
type AccessibilityProgressIndicatorWrapper struct {
	objc.Object
}

func (a_ AccessibilityProgressIndicatorWrapper) HasAccessibilityValue() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityValue"))
}

// Returns the progress indicator’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityprogressindicator/1531500-accessibilityvalue?language=objc
func (a_ AccessibilityProgressIndicatorWrapper) AccessibilityValue() foundation.Number {
	rv := objc.Call[foundation.Number](a_, objc.Sel("accessibilityValue"))
	return rv
}
