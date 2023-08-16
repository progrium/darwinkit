// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to support loading. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityelementloading?language=objc
type PAccessibilityElementLoading interface {
	// optional
	AccessibilityElementWithToken(token AccessibilityLoadingToken) objc.IObject
	HasAccessibilityElementWithToken() bool

	// optional
	AccessibilityRangeInTargetElementWithToken(token AccessibilityLoadingToken) foundation.Range
	HasAccessibilityRangeInTargetElementWithToken() bool
}

// A concrete type wrapper for the [PAccessibilityElementLoading] protocol.
type AccessibilityElementLoadingWrapper struct {
	objc.Object
}

func (a_ AccessibilityElementLoadingWrapper) HasAccessibilityElementWithToken() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityElementWithToken:"))
}

// Loads the target accessibility element with the specified load token. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityelementloading/2890815-accessibilityelementwithtoken?language=objc
func (a_ AccessibilityElementLoadingWrapper) AccessibilityElementWithToken(token AccessibilityLoadingToken) objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityElementWithToken:"), token)
	return rv
}

func (a_ AccessibilityElementLoadingWrapper) HasAccessibilityRangeInTargetElementWithToken() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityRangeInTargetElementWithToken:"))
}

// Returns the range that specifies the area of interest in text-based accessibility elements with the specified load token. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityelementloading/2890818-accessibilityrangeintargetelemen?language=objc
func (a_ AccessibilityElementLoadingWrapper) AccessibilityRangeInTargetElementWithToken(token AccessibilityLoadingToken) foundation.Range {
	rv := objc.Call[foundation.Range](a_, objc.Sel("accessibilityRangeInTargetElementWithToken:"), token)
	return rv
}
