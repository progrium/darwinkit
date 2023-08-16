// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a row for a table, list, or outline view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityrow?language=objc
type PAccessibilityRow interface {
	// optional
	AccessibilityIndex() int
	HasAccessibilityIndex() bool

	// optional
	AccessibilityDisclosureLevel() int
	HasAccessibilityDisclosureLevel() bool
}

// A concrete type wrapper for the [PAccessibilityRow] protocol.
type AccessibilityRowWrapper struct {
	objc.Object
}

func (a_ AccessibilityRowWrapper) HasAccessibilityIndex() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityIndex"))
}

// Returns the index for the row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityrow/1526746-accessibilityindex?language=objc
func (a_ AccessibilityRowWrapper) AccessibilityIndex() int {
	rv := objc.Call[int](a_, objc.Sel("accessibilityIndex"))
	return rv
}

func (a_ AccessibilityRowWrapper) HasAccessibilityDisclosureLevel() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityDisclosureLevel"))
}

// Returns the indention level for the row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityrow/1531837-accessibilitydisclosurelevel?language=objc
func (a_ AccessibilityRowWrapper) AccessibilityDisclosureLevel() int {
	rv := objc.Call[int](a_, objc.Sel("accessibilityDisclosureLevel"))
	return rv
}
