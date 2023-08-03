// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type IDraggingSource interface {
	// required
	DraggingSessionSourceOperationMaskForDraggingContext(session DraggingSession, context DraggingContext) DragOperation
	ImplementsDraggingSessionEndedAtPointOperation() bool
	// optional
	DraggingSessionEndedAtPointOperation(session DraggingSession, screenPoint foundation.Point, operation DragOperation)
	ImplementsDraggingSessionMovedToPoint() bool
	// optional
	DraggingSessionMovedToPoint(session DraggingSession, screenPoint foundation.Point)
	ImplementsDraggingSessionWillBeginAtPoint() bool
	// optional
	DraggingSessionWillBeginAtPoint(session DraggingSession, screenPoint foundation.Point)
	ImplementsIgnoreModifierKeysForDraggingSession() bool
	// optional
	IgnoreModifierKeysForDraggingSession(session DraggingSession) bool
}

type DraggingSourceWrapper struct {
	objc.Object
}

func (d_ DraggingSourceWrapper) DraggingSessionSourceOperationMaskForDraggingContext(session IDraggingSession, context DraggingContext) DragOperation {
	rv := objc.CallMethod[DragOperation](d_, objc.GetSelector("draggingSession:sourceOperationMaskForDraggingContext:"), objc.ExtractPtr(session), context)
	return rv
}

func (d_ DraggingSourceWrapper) ImplementsDraggingSessionEndedAtPointOperation() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingSession:endedAtPoint:operation:"))
}

func (d_ DraggingSourceWrapper) DraggingSessionEndedAtPointOperation(session IDraggingSession, screenPoint foundation.Point, operation DragOperation) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("draggingSession:endedAtPoint:operation:"), objc.ExtractPtr(session), screenPoint, operation)
}

func (d_ DraggingSourceWrapper) ImplementsDraggingSessionMovedToPoint() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingSession:movedToPoint:"))
}

func (d_ DraggingSourceWrapper) DraggingSessionMovedToPoint(session IDraggingSession, screenPoint foundation.Point) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("draggingSession:movedToPoint:"), objc.ExtractPtr(session), screenPoint)
}

func (d_ DraggingSourceWrapper) ImplementsDraggingSessionWillBeginAtPoint() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingSession:willBeginAtPoint:"))
}

func (d_ DraggingSourceWrapper) DraggingSessionWillBeginAtPoint(session IDraggingSession, screenPoint foundation.Point) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("draggingSession:willBeginAtPoint:"), objc.ExtractPtr(session), screenPoint)
}

func (d_ DraggingSourceWrapper) ImplementsIgnoreModifierKeysForDraggingSession() bool {
	return d_.RespondsToSelector(objc.GetSelector("ignoreModifierKeysForDraggingSession:"))
}

func (d_ DraggingSourceWrapper) IgnoreModifierKeysForDraggingSession(session IDraggingSession) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("ignoreModifierKeysForDraggingSession:"), objc.ExtractPtr(session))
	return rv
}
