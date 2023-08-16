// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityimage?language=objc
type PAccessibilityImage interface {
	// optional
	AccessibilityLabel() string
	HasAccessibilityLabel() bool
}

// A concrete type wrapper for the [PAccessibilityImage] protocol.
type AccessibilityImageWrapper struct {
	objc.Object
}

func (a_ AccessibilityImageWrapper) HasAccessibilityLabel() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityLabel"))
}

// Returns a short description of the image’s label. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityimage/1531608-accessibilitylabel?language=objc
func (a_ AccessibilityImageWrapper) AccessibilityLabel() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityLabel"))
	return rv
}
