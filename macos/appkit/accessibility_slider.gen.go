// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a slider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityslider?language=objc
type PAccessibilitySlider interface {
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

// A concrete type wrapper for the [PAccessibilitySlider] protocol.
type AccessibilitySliderWrapper struct {
	objc.Object
}

func (a_ AccessibilitySliderWrapper) HasAccessibilityPerformDecrement() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformDecrement"))
}

// Decrements the slider’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityslider/1534967-accessibilityperformdecrement?language=objc
func (a_ AccessibilitySliderWrapper) AccessibilityPerformDecrement() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformDecrement"))
	return rv
}

func (a_ AccessibilitySliderWrapper) HasAccessibilityLabel() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityLabel"))
}

// Returns a short description of the slider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityslider/1530176-accessibilitylabel?language=objc
func (a_ AccessibilitySliderWrapper) AccessibilityLabel() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityLabel"))
	return rv
}

func (a_ AccessibilitySliderWrapper) HasAccessibilityPerformIncrement() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformIncrement"))
}

// Increments the slider’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityslider/1528478-accessibilityperformincrement?language=objc
func (a_ AccessibilitySliderWrapper) AccessibilityPerformIncrement() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformIncrement"))
	return rv
}

func (a_ AccessibilitySliderWrapper) HasAccessibilityValue() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityValue"))
}

// Returns the slider’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityslider/1530335-accessibilityvalue?language=objc
func (a_ AccessibilitySliderWrapper) AccessibilityValue() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityValue"))
	return rv
}
