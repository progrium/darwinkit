// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as navigable static text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitynavigablestatictext?language=objc
type PAccessibilityNavigableStaticText interface {
	// optional
	AccessibilityFrameForRange(range_ foundation.Range) foundation.Rect
	HasAccessibilityFrameForRange() bool

	// optional
	AccessibilityRangeForLine(lineNumber int) foundation.Range
	HasAccessibilityRangeForLine() bool

	// optional
	AccessibilityLineForIndex(index int) int
	HasAccessibilityLineForIndex() bool

	// optional
	AccessibilityStringForRange(range_ foundation.Range) string
	HasAccessibilityStringForRange() bool
}

// A concrete type wrapper for the [PAccessibilityNavigableStaticText] protocol.
type AccessibilityNavigableStaticTextWrapper struct {
	objc.Object
}

func (a_ AccessibilityNavigableStaticTextWrapper) HasAccessibilityFrameForRange() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityFrameForRange:"))
}

// Returns the rectangle that encloses the specified range of characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitynavigablestatictext/1524702-accessibilityframeforrange?language=objc
func (a_ AccessibilityNavigableStaticTextWrapper) AccessibilityFrameForRange(range_ foundation.Range) foundation.Rect {
	rv := objc.Call[foundation.Rect](a_, objc.Sel("accessibilityFrameForRange:"), range_)
	return rv
}

func (a_ AccessibilityNavigableStaticTextWrapper) HasAccessibilityRangeForLine() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityRangeForLine:"))
}

// Returns the range of characters in the specified line. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitynavigablestatictext/1527015-accessibilityrangeforline?language=objc
func (a_ AccessibilityNavigableStaticTextWrapper) AccessibilityRangeForLine(lineNumber int) foundation.Range {
	rv := objc.Call[foundation.Range](a_, objc.Sel("accessibilityRangeForLine:"), lineNumber)
	return rv
}

func (a_ AccessibilityNavigableStaticTextWrapper) HasAccessibilityLineForIndex() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityLineForIndex:"))
}

// Returns the line number for the line that contains the specified character index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitynavigablestatictext/1534931-accessibilitylineforindex?language=objc
func (a_ AccessibilityNavigableStaticTextWrapper) AccessibilityLineForIndex(index int) int {
	rv := objc.Call[int](a_, objc.Sel("accessibilityLineForIndex:"), index)
	return rv
}

func (a_ AccessibilityNavigableStaticTextWrapper) HasAccessibilityStringForRange() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityStringForRange:"))
}

// Returns the substring for the specified range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitynavigablestatictext/1525402-accessibilitystringforrange?language=objc
func (a_ AccessibilityNavigableStaticTextWrapper) AccessibilityStringForRange(range_ foundation.Range) string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityStringForRange:"), range_)
	return rv
}
