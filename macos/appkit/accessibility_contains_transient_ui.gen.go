// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to support dynamic UI changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycontainstransientui?language=objc
type PAccessibilityContainsTransientUI interface {
	// optional
	AccessibilityPerformShowDefaultUI() bool
	HasAccessibilityPerformShowDefaultUI() bool

	// optional
	AccessibilityPerformShowAlternateUI() bool
	HasAccessibilityPerformShowAlternateUI() bool

	// optional
	IsAccessibilityAlternateUIVisible() bool
	HasIsAccessibilityAlternateUIVisible() bool
}

// A concrete type wrapper for the [PAccessibilityContainsTransientUI] protocol.
type AccessibilityContainsTransientUIWrapper struct {
	objc.Object
}

func (a_ AccessibilityContainsTransientUIWrapper) HasAccessibilityPerformShowDefaultUI() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformShowDefaultUI"))
}

// Returns to the accessibility element’s original UI. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycontainstransientui/1529235-accessibilityperformshowdefaultu?language=objc
func (a_ AccessibilityContainsTransientUIWrapper) AccessibilityPerformShowDefaultUI() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformShowDefaultUI"))
	return rv
}

func (a_ AccessibilityContainsTransientUIWrapper) HasAccessibilityPerformShowAlternateUI() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformShowAlternateUI"))
}

// Displays the accessibility element’s alternative UI. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycontainstransientui/1535134-accessibilityperformshowalternat?language=objc
func (a_ AccessibilityContainsTransientUIWrapper) AccessibilityPerformShowAlternateUI() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformShowAlternateUI"))
	return rv
}

func (a_ AccessibilityContainsTransientUIWrapper) HasIsAccessibilityAlternateUIVisible() bool {
	return a_.RespondsToSelector(objc.Sel("isAccessibilityAlternateUIVisible"))
}

// Returns a Boolean value that determines whether the accessibility element’s alternative UI is currently visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycontainstransientui/1526272-isaccessibilityalternateuivisibl?language=objc
func (a_ AccessibilityContainsTransientUIWrapper) IsAccessibilityAlternateUIVisible() bool {
	rv := objc.Call[bool](a_, objc.Sel("isAccessibilityAlternateUIVisible"))
	return rv
}
