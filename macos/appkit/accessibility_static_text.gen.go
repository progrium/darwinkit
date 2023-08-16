// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as static text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitystatictext?language=objc
type PAccessibilityStaticText interface {
	// optional
	AccessibilityAttributedStringForRange(range_ foundation.Range) foundation.IAttributedString
	HasAccessibilityAttributedStringForRange() bool

	// optional
	AccessibilityVisibleCharacterRange() foundation.Range
	HasAccessibilityVisibleCharacterRange() bool

	// optional
	AccessibilityValue() string
	HasAccessibilityValue() bool
}

// A concrete type wrapper for the [PAccessibilityStaticText] protocol.
type AccessibilityStaticTextWrapper struct {
	objc.Object
}

func (a_ AccessibilityStaticTextWrapper) HasAccessibilityAttributedStringForRange() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityAttributedStringForRange:"))
}

// Returns the attributed substring for the specified range of characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitystatictext/1535001-accessibilityattributedstringfor?language=objc
func (a_ AccessibilityStaticTextWrapper) AccessibilityAttributedStringForRange(range_ foundation.Range) foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](a_, objc.Sel("accessibilityAttributedStringForRange:"), range_)
	return rv
}

func (a_ AccessibilityStaticTextWrapper) HasAccessibilityVisibleCharacterRange() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityVisibleCharacterRange"))
}

// Returns the range of visible characters in the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitystatictext/1532230-accessibilityvisiblecharacterran?language=objc
func (a_ AccessibilityStaticTextWrapper) AccessibilityVisibleCharacterRange() foundation.Range {
	rv := objc.Call[foundation.Range](a_, objc.Sel("accessibilityVisibleCharacterRange"))
	return rv
}

func (a_ AccessibilityStaticTextWrapper) HasAccessibilityValue() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityValue"))
}

// Returns the text that the accessibility element displays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitystatictext/1528730-accessibilityvalue?language=objc
func (a_ AccessibilityStaticTextWrapper) AccessibilityValue() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityValue"))
	return rv
}
