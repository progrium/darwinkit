// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a layout area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitylayoutarea?language=objc
type PAccessibilityLayoutArea interface {
	// optional
	AccessibilityChildren() []objc.IObject
	HasAccessibilityChildren() bool

	// optional
	AccessibilityLabel() string
	HasAccessibilityLabel() bool

	// optional
	AccessibilitySelectedChildren() []objc.IObject
	HasAccessibilitySelectedChildren() bool

	// optional
	AccessibilityFocusedUIElement() objc.IObject
	HasAccessibilityFocusedUIElement() bool
}

// A concrete type wrapper for the [PAccessibilityLayoutArea] protocol.
type AccessibilityLayoutAreaWrapper struct {
	objc.Object
}

func (a_ AccessibilityLayoutAreaWrapper) HasAccessibilityChildren() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityChildren"))
}

// Returns the accessibility element’s children in the accessibility hierarchy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitylayoutarea/1534997-accessibilitychildren?language=objc
func (a_ AccessibilityLayoutAreaWrapper) AccessibilityChildren() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilityChildren"))
	return rv
}

func (a_ AccessibilityLayoutAreaWrapper) HasAccessibilityLabel() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityLabel"))
}

// Returns a short description of the layout area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitylayoutarea/1527051-accessibilitylabel?language=objc
func (a_ AccessibilityLayoutAreaWrapper) AccessibilityLabel() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityLabel"))
	return rv
}

func (a_ AccessibilityLayoutAreaWrapper) HasAccessibilitySelectedChildren() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilitySelectedChildren"))
}

// Returns the layout area’s currently selected children. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitylayoutarea/1528883-accessibilityselectedchildren?language=objc
func (a_ AccessibilityLayoutAreaWrapper) AccessibilitySelectedChildren() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("accessibilitySelectedChildren"))
	return rv
}

func (a_ AccessibilityLayoutAreaWrapper) HasAccessibilityFocusedUIElement() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityFocusedUIElement"))
}

// The child accessibility element with the current focus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitylayoutarea/1533902-accessibilityfocuseduielement?language=objc
func (a_ AccessibilityLayoutAreaWrapper) AccessibilityFocusedUIElement() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityFocusedUIElement"))
	return rv
}
