// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a stepper. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitystepper?language=objc
type PAccessibilityStepper interface {
	// optional
	AccessibilityPerformDecrement() bool
	HasAccessibilityPerformDecrement() bool

	// optional
	AccessibilityLabel() string
	HasAccessibilityLabel() bool

	// optional
	AccessibilityPerformIncrement() bool
	HasAccessibilityPerformIncrement() bool

	// optional
	AccessibilityValue() objc.IObject
	HasAccessibilityValue() bool
}

// A concrete type wrapper for the [PAccessibilityStepper] protocol.
type AccessibilityStepperWrapper struct {
	objc.Object
}

func (a_ AccessibilityStepperWrapper) HasAccessibilityPerformDecrement() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformDecrement"))
}

// Decrements the stepper’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitystepper/1525327-accessibilityperformdecrement?language=objc
func (a_ AccessibilityStepperWrapper) AccessibilityPerformDecrement() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformDecrement"))
	return rv
}

func (a_ AccessibilityStepperWrapper) HasAccessibilityLabel() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityLabel"))
}

// Returns a short description of the stepper. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitystepper/1528702-accessibilitylabel?language=objc
func (a_ AccessibilityStepperWrapper) AccessibilityLabel() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityLabel"))
	return rv
}

func (a_ AccessibilityStepperWrapper) HasAccessibilityPerformIncrement() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformIncrement"))
}

// Increments the stepper’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitystepper/1533764-accessibilityperformincrement?language=objc
func (a_ AccessibilityStepperWrapper) AccessibilityPerformIncrement() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformIncrement"))
	return rv
}

func (a_ AccessibilityStepperWrapper) HasAccessibilityValue() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityValue"))
}

// Returns the stepper’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitystepper/1528167-accessibilityvalue?language=objc
func (a_ AccessibilityStepperWrapper) AccessibilityValue() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityValue"))
	return rv
}
