// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of methods that are implemented by the source object in a dragging session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingsource?language=objc
type PDraggingSource interface {
	// optional
	DraggingSessionMovedToPoint(session DraggingSession, screenPoint foundation.Point)
	HasDraggingSessionMovedToPoint() bool

	// optional
	IgnoreModifierKeysForDraggingSession(session DraggingSession) bool
	HasIgnoreModifierKeysForDraggingSession() bool
}

// A concrete type wrapper for the [PDraggingSource] protocol.
type DraggingSourceWrapper struct {
	objc.Object
}

func (d_ DraggingSourceWrapper) HasDraggingSessionMovedToPoint() bool {
	return d_.RespondsToSelector(objc.Sel("draggingSession:movedToPoint:"))
}

// Invoked when the drag moves on the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingsource/1416079-draggingsession?language=objc
func (d_ DraggingSourceWrapper) DraggingSessionMovedToPoint(session IDraggingSession, screenPoint foundation.Point) {
	objc.Call[objc.Void](d_, objc.Sel("draggingSession:movedToPoint:"), objc.Ptr(session), screenPoint)
}

func (d_ DraggingSourceWrapper) HasIgnoreModifierKeysForDraggingSession() bool {
	return d_.RespondsToSelector(objc.Sel("ignoreModifierKeysForDraggingSession:"))
}

// Returns whether the modifier keys will be ignored for this dragging session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingsource/1415974-ignoremodifierkeysfordraggingses?language=objc
func (d_ DraggingSourceWrapper) IgnoreModifierKeysForDraggingSession(session IDraggingSession) bool {
	rv := objc.Call[bool](d_, objc.Sel("ignoreModifierKeysForDraggingSession:"), objc.Ptr(session))
	return rv
}
