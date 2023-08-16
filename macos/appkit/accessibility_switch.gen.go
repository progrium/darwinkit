// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a switch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityswitch?language=objc
type PAccessibilitySwitch interface {
	// optional
	AccessibilityPerformDecrement() bool
	HasAccessibilityPerformDecrement() bool

	// optional
	AccessibilityPerformIncrement() bool
	HasAccessibilityPerformIncrement() bool

	// optional
	AccessibilityValue() string
	HasAccessibilityValue() bool
}

// A concrete type wrapper for the [PAccessibilitySwitch] protocol.
type AccessibilitySwitchWrapper struct {
	objc.Object
}

func (a_ AccessibilitySwitchWrapper) HasAccessibilityPerformDecrement() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformDecrement"))
}

// Decrements the switch’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityswitch/1528290-accessibilityperformdecrement?language=objc
func (a_ AccessibilitySwitchWrapper) AccessibilityPerformDecrement() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformDecrement"))
	return rv
}

func (a_ AccessibilitySwitchWrapper) HasAccessibilityPerformIncrement() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformIncrement"))
}

// Increments the switch’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityswitch/1533985-accessibilityperformincrement?language=objc
func (a_ AccessibilitySwitchWrapper) AccessibilityPerformIncrement() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformIncrement"))
	return rv
}

func (a_ AccessibilitySwitchWrapper) HasAccessibilityValue() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityValue"))
}

// Returns the switch’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityswitch/1533946-accessibilityvalue?language=objc
func (a_ AccessibilitySwitchWrapper) AccessibilityValue() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityValue"))
	return rv
}
