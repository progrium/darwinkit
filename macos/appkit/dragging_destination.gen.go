// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods that the destination object (or recipient) of a dragged image must implement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingdestination?language=objc
type PDraggingDestination interface {
	// optional
	WantsPeriodicDraggingUpdates() bool
	HasWantsPeriodicDraggingUpdates() bool

	// optional
	PrepareForDragOperation(sender DraggingInfoWrapper) bool
	HasPrepareForDragOperation() bool

	// optional
	DraggingEntered(sender DraggingInfoWrapper) DragOperation
	HasDraggingEntered() bool

	// optional
	PerformDragOperation(sender DraggingInfoWrapper) bool
	HasPerformDragOperation() bool

	// optional
	ConcludeDragOperation(sender DraggingInfoWrapper)
	HasConcludeDragOperation() bool

	// optional
	DraggingUpdated(sender DraggingInfoWrapper) DragOperation
	HasDraggingUpdated() bool

	// optional
	DraggingEnded(sender DraggingInfoWrapper)
	HasDraggingEnded() bool

	// optional
	UpdateDraggingItemsForDrag(sender DraggingInfoWrapper)
	HasUpdateDraggingItemsForDrag() bool

	// optional
	DraggingExited(sender DraggingInfoWrapper)
	HasDraggingExited() bool
}

// A concrete type wrapper for the [PDraggingDestination] protocol.
type DraggingDestinationWrapper struct {
	objc.Object
}

func (d_ DraggingDestinationWrapper) HasWantsPeriodicDraggingUpdates() bool {
	return d_.RespondsToSelector(objc.Sel("wantsPeriodicDraggingUpdates"))
}

// Asks the destination object whether it wants to receive periodic draggingUpdated: messages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingdestination/1416049-wantsperiodicdraggingupdates?language=objc
func (d_ DraggingDestinationWrapper) WantsPeriodicDraggingUpdates() bool {
	rv := objc.Call[bool](d_, objc.Sel("wantsPeriodicDraggingUpdates"))
	return rv
}

func (d_ DraggingDestinationWrapper) HasPrepareForDragOperation() bool {
	return d_.RespondsToSelector(objc.Sel("prepareForDragOperation:"))
}

// Invoked when the image is released, allowing the receiver to agree to or refuse drag operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingdestination/1416066-preparefordragoperation?language=objc
func (d_ DraggingDestinationWrapper) PrepareForDragOperation(sender PDraggingInfo) bool {
	po0 := objc.WrapAsProtocol("NSDraggingInfo", sender)
	rv := objc.Call[bool](d_, objc.Sel("prepareForDragOperation:"), po0)
	return rv
}

func (d_ DraggingDestinationWrapper) HasDraggingEntered() bool {
	return d_.RespondsToSelector(objc.Sel("draggingEntered:"))
}

// Invoked when the dragged image enters destination bounds or frame; delegate returns dragging operation to perform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingdestination/1416019-draggingentered?language=objc
func (d_ DraggingDestinationWrapper) DraggingEntered(sender PDraggingInfo) DragOperation {
	po0 := objc.WrapAsProtocol("NSDraggingInfo", sender)
	rv := objc.Call[DragOperation](d_, objc.Sel("draggingEntered:"), po0)
	return rv
}

func (d_ DraggingDestinationWrapper) HasPerformDragOperation() bool {
	return d_.RespondsToSelector(objc.Sel("performDragOperation:"))
}

// Invoked after the released image has been removed from the screen, signaling the receiver to import the pasteboard data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingdestination/1415970-performdragoperation?language=objc
func (d_ DraggingDestinationWrapper) PerformDragOperation(sender PDraggingInfo) bool {
	po0 := objc.WrapAsProtocol("NSDraggingInfo", sender)
	rv := objc.Call[bool](d_, objc.Sel("performDragOperation:"), po0)
	return rv
}

func (d_ DraggingDestinationWrapper) HasConcludeDragOperation() bool {
	return d_.RespondsToSelector(objc.Sel("concludeDragOperation:"))
}

// Invoked when the dragging operation is complete, signaling the receiver to perform any necessary clean-up. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingdestination/1416010-concludedragoperation?language=objc
func (d_ DraggingDestinationWrapper) ConcludeDragOperation(sender PDraggingInfo) {
	po0 := objc.WrapAsProtocol("NSDraggingInfo", sender)
	objc.Call[objc.Void](d_, objc.Sel("concludeDragOperation:"), po0)
}

func (d_ DraggingDestinationWrapper) HasDraggingUpdated() bool {
	return d_.RespondsToSelector(objc.Sel("draggingUpdated:"))
}

// Invoked periodically as the image is held within the destination area, allowing modification of the dragging operation or mouse-pointer position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingdestination/1415998-draggingupdated?language=objc
func (d_ DraggingDestinationWrapper) DraggingUpdated(sender PDraggingInfo) DragOperation {
	po0 := objc.WrapAsProtocol("NSDraggingInfo", sender)
	rv := objc.Call[DragOperation](d_, objc.Sel("draggingUpdated:"), po0)
	return rv
}

func (d_ DraggingDestinationWrapper) HasDraggingEnded() bool {
	return d_.RespondsToSelector(objc.Sel("draggingEnded:"))
}

// Called when a drag operation ends. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingdestination/1416096-draggingended?language=objc
func (d_ DraggingDestinationWrapper) DraggingEnded(sender PDraggingInfo) {
	po0 := objc.WrapAsProtocol("NSDraggingInfo", sender)
	objc.Call[objc.Void](d_, objc.Sel("draggingEnded:"), po0)
}

func (d_ DraggingDestinationWrapper) HasUpdateDraggingItemsForDrag() bool {
	return d_.RespondsToSelector(objc.Sel("updateDraggingItemsForDrag:"))
}

// Invoked when the dragging images should be changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingdestination/1416050-updatedraggingitemsfordrag?language=objc
func (d_ DraggingDestinationWrapper) UpdateDraggingItemsForDrag(sender PDraggingInfo) {
	po0 := objc.WrapAsProtocol("NSDraggingInfo", sender)
	objc.Call[objc.Void](d_, objc.Sel("updateDraggingItemsForDrag:"), po0)
}

func (d_ DraggingDestinationWrapper) HasDraggingExited() bool {
	return d_.RespondsToSelector(objc.Sel("draggingExited:"))
}

// Invoked when the dragged image exits the destination’s bounds rectangle (in the case of a view object) or its frame rectangle (in the case of a window object). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingdestination/1416056-draggingexited?language=objc
func (d_ DraggingDestinationWrapper) DraggingExited(sender PDraggingInfo) {
	po0 := objc.WrapAsProtocol("NSDraggingInfo", sender)
	objc.Call[objc.Void](d_, objc.Sel("draggingExited:"), po0)
}
