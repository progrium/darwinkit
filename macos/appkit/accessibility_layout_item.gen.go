// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a layout item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitylayoutitem?language=objc
type PAccessibilityLayoutItem interface {
	// optional
	SetAccessibilityFrame(frame foundation.Rect)
	HasSetAccessibilityFrame() bool
}

// A concrete type wrapper for the [PAccessibilityLayoutItem] protocol.
type AccessibilityLayoutItemWrapper struct {
	objc.Object
}

func (a_ AccessibilityLayoutItemWrapper) HasSetAccessibilityFrame() bool {
	return a_.RespondsToSelector(objc.Sel("setAccessibilityFrame:"))
}

// Sets the accessibility element’s frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitylayoutitem/1533160-setaccessibilityframe?language=objc
func (a_ AccessibilityLayoutItemWrapper) SetAccessibilityFrame(frame foundation.Rect) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityFrame:"), frame)
}
