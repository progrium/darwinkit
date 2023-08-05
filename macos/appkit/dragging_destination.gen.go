// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type IDraggingDestination interface {
	ImplementsDraggingEntered() bool
	// optional
	DraggingEntered(sender DraggingInfoWrapper) DragOperation
	ImplementsWantsPeriodicDraggingUpdates() bool
	// optional
	WantsPeriodicDraggingUpdates() bool
	ImplementsDraggingUpdated() bool
	// optional
	DraggingUpdated(sender DraggingInfoWrapper) DragOperation
	ImplementsDraggingEnded() bool
	// optional
	DraggingEnded(sender DraggingInfoWrapper)
	ImplementsDraggingExited() bool
	// optional
	DraggingExited(sender DraggingInfoWrapper)
	ImplementsPrepareForDragOperation() bool
	// optional
	PrepareForDragOperation(sender DraggingInfoWrapper) bool
	ImplementsPerformDragOperation() bool
	// optional
	PerformDragOperation(sender DraggingInfoWrapper) bool
	ImplementsConcludeDragOperation() bool
	// optional
	ConcludeDragOperation(sender DraggingInfoWrapper)
	ImplementsUpdateDraggingItemsForDrag() bool
	// optional
	UpdateDraggingItemsForDrag(sender DraggingInfoWrapper)
}

type DraggingDestinationWrapper struct {
	objc.Object
}

func (d_ DraggingDestinationWrapper) ImplementsDraggingEntered() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingEntered:"))
}

func (d_ DraggingDestinationWrapper) DraggingEntered(sender IDraggingInfo) DragOperation {
	po := objc.WrapAsProtocol("NSDraggingInfo", sender)
	rv := objc.CallMethod[DragOperation](d_, objc.GetSelector("draggingEntered:"), po)
	return rv
}

func (d_ DraggingDestinationWrapper) ImplementsWantsPeriodicDraggingUpdates() bool {
	return d_.RespondsToSelector(objc.GetSelector("wantsPeriodicDraggingUpdates"))
}

func (d_ DraggingDestinationWrapper) WantsPeriodicDraggingUpdates() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("wantsPeriodicDraggingUpdates"))
	return rv
}

func (d_ DraggingDestinationWrapper) ImplementsDraggingUpdated() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingUpdated:"))
}

func (d_ DraggingDestinationWrapper) DraggingUpdated(sender IDraggingInfo) DragOperation {
	po := objc.WrapAsProtocol("NSDraggingInfo", sender)
	rv := objc.CallMethod[DragOperation](d_, objc.GetSelector("draggingUpdated:"), po)
	return rv
}

func (d_ DraggingDestinationWrapper) ImplementsDraggingEnded() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingEnded:"))
}

func (d_ DraggingDestinationWrapper) DraggingEnded(sender IDraggingInfo) {
	po := objc.WrapAsProtocol("NSDraggingInfo", sender)
	objc.CallMethod[objc.Void](d_, objc.GetSelector("draggingEnded:"), po)
}

func (d_ DraggingDestinationWrapper) ImplementsDraggingExited() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingExited:"))
}

func (d_ DraggingDestinationWrapper) DraggingExited(sender IDraggingInfo) {
	po := objc.WrapAsProtocol("NSDraggingInfo", sender)
	objc.CallMethod[objc.Void](d_, objc.GetSelector("draggingExited:"), po)
}

func (d_ DraggingDestinationWrapper) ImplementsPrepareForDragOperation() bool {
	return d_.RespondsToSelector(objc.GetSelector("prepareForDragOperation:"))
}

func (d_ DraggingDestinationWrapper) PrepareForDragOperation(sender IDraggingInfo) bool {
	po := objc.WrapAsProtocol("NSDraggingInfo", sender)
	rv := objc.CallMethod[bool](d_, objc.GetSelector("prepareForDragOperation:"), po)
	return rv
}

func (d_ DraggingDestinationWrapper) ImplementsPerformDragOperation() bool {
	return d_.RespondsToSelector(objc.GetSelector("performDragOperation:"))
}

func (d_ DraggingDestinationWrapper) PerformDragOperation(sender IDraggingInfo) bool {
	po := objc.WrapAsProtocol("NSDraggingInfo", sender)
	rv := objc.CallMethod[bool](d_, objc.GetSelector("performDragOperation:"), po)
	return rv
}

func (d_ DraggingDestinationWrapper) ImplementsConcludeDragOperation() bool {
	return d_.RespondsToSelector(objc.GetSelector("concludeDragOperation:"))
}

func (d_ DraggingDestinationWrapper) ConcludeDragOperation(sender IDraggingInfo) {
	po := objc.WrapAsProtocol("NSDraggingInfo", sender)
	objc.CallMethod[objc.Void](d_, objc.GetSelector("concludeDragOperation:"), po)
}

func (d_ DraggingDestinationWrapper) ImplementsUpdateDraggingItemsForDrag() bool {
	return d_.RespondsToSelector(objc.GetSelector("updateDraggingItemsForDrag:"))
}

func (d_ DraggingDestinationWrapper) UpdateDraggingItemsForDrag(sender IDraggingInfo) {
	po := objc.WrapAsProtocol("NSDraggingInfo", sender)
	objc.CallMethod[objc.Void](d_, objc.GetSelector("updateDraggingItemsForDrag:"), po)
}
