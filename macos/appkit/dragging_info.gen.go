// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type IDraggingInfo interface {
	// required
	// deprecated
	NamesOfPromisedFilesDroppedAtDestination(dropDestination foundation.URL) []string
	// required
	SlideDraggedImageTo(screenPoint foundation.Point)
	// required
	EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock(enumOpts DraggingItemEnumerationOptions, view View, classArray []objc.Class, searchOptions map[PasteboardReadingOptionKey]objc.Object, block func(draggingItem DraggingItem, idx int, stop *bool))
	// required
	ResetSpringLoading()
	ImplementsDraggingPasteboard() bool
	// optional
	DraggingPasteboard() IPasteboard
	ImplementsDraggingSequenceNumber() bool
	// optional
	DraggingSequenceNumber() int
	ImplementsDraggingSource() bool
	// optional
	DraggingSource() objc.IObject
	ImplementsDraggingSourceOperationMask() bool
	// optional
	DraggingSourceOperationMask() DragOperation
	ImplementsDraggingLocation() bool
	// optional
	DraggingLocation() foundation.Point
	ImplementsDraggingDestinationWindow() bool
	// optional
	DraggingDestinationWindow() IWindow
	ImplementsSetNumberOfValidItemsForDrop() bool
	// optional
	SetNumberOfValidItemsForDrop(value int)
	ImplementsNumberOfValidItemsForDrop() bool
	// optional
	NumberOfValidItemsForDrop() int
	ImplementsDraggedImage() bool
	// optional
	// deprecated
	DraggedImage() IImage
	ImplementsDraggedImageLocation() bool
	// optional
	DraggedImageLocation() foundation.Point
	ImplementsSetAnimatesToDestination() bool
	// optional
	SetAnimatesToDestination(value bool)
	ImplementsAnimatesToDestination() bool
	// optional
	AnimatesToDestination() bool
	ImplementsSetDraggingFormation() bool
	// optional
	SetDraggingFormation(value DraggingFormation)
	ImplementsDraggingFormation() bool
	// optional
	DraggingFormation() DraggingFormation
	ImplementsSpringLoadingHighlight() bool
	// optional
	SpringLoadingHighlight() SpringLoadingHighlight
}

type DraggingInfoWrapper struct {
	objc.Object
}

func (d_ DraggingInfoWrapper) SlideDraggedImageTo(screenPoint foundation.Point) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("slideDraggedImageTo:"), screenPoint)
}

func (d_ DraggingInfoWrapper) EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock(enumOpts DraggingItemEnumerationOptions, view IView, classArray []objc.IClass, searchOptions map[PasteboardReadingOptionKey]objc.IObject, block func(draggingItem DraggingItem, idx int, stop *bool)) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("enumerateDraggingItemsWithOptions:forView:classes:searchOptions:usingBlock:"), enumOpts, objc.ExtractPtr(view), classArray, searchOptions, block)
}

func (d_ DraggingInfoWrapper) ResetSpringLoading() {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("resetSpringLoading"))
}

func (d_ DraggingInfoWrapper) ImplementsDraggingPasteboard() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingPasteboard"))
}

func (d_ DraggingInfoWrapper) DraggingPasteboard() Pasteboard {
	rv := objc.CallMethod[Pasteboard](d_, objc.GetSelector("draggingPasteboard"))
	return rv
}

func (d_ DraggingInfoWrapper) ImplementsDraggingSequenceNumber() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingSequenceNumber"))
}

func (d_ DraggingInfoWrapper) DraggingSequenceNumber() int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("draggingSequenceNumber"))
	return rv
}

func (d_ DraggingInfoWrapper) ImplementsDraggingSource() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingSource"))
}

func (d_ DraggingInfoWrapper) DraggingSource() objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.GetSelector("draggingSource"))
	return rv
}

func (d_ DraggingInfoWrapper) ImplementsDraggingSourceOperationMask() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingSourceOperationMask"))
}

func (d_ DraggingInfoWrapper) DraggingSourceOperationMask() DragOperation {
	rv := objc.CallMethod[DragOperation](d_, objc.GetSelector("draggingSourceOperationMask"))
	return rv
}

func (d_ DraggingInfoWrapper) ImplementsDraggingLocation() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingLocation"))
}

func (d_ DraggingInfoWrapper) DraggingLocation() foundation.Point {
	rv := objc.CallMethod[foundation.Point](d_, objc.GetSelector("draggingLocation"))
	return rv
}

func (d_ DraggingInfoWrapper) ImplementsDraggingDestinationWindow() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingDestinationWindow"))
}

func (d_ DraggingInfoWrapper) DraggingDestinationWindow() Window {
	rv := objc.CallMethod[Window](d_, objc.GetSelector("draggingDestinationWindow"))
	return rv
}

func (d_ DraggingInfoWrapper) ImplementsSetNumberOfValidItemsForDrop() bool {
	return d_.RespondsToSelector(objc.GetSelector("setNumberOfValidItemsForDrop:"))
}

func (d_ DraggingInfoWrapper) SetNumberOfValidItemsForDrop(value int) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setNumberOfValidItemsForDrop:"), value)
}

func (d_ DraggingInfoWrapper) ImplementsNumberOfValidItemsForDrop() bool {
	return d_.RespondsToSelector(objc.GetSelector("numberOfValidItemsForDrop"))
}

func (d_ DraggingInfoWrapper) NumberOfValidItemsForDrop() int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("numberOfValidItemsForDrop"))
	return rv
}

func (d_ DraggingInfoWrapper) ImplementsDraggedImage() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggedImage"))
}

func (d_ DraggingInfoWrapper) ImplementsDraggedImageLocation() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggedImageLocation"))
}

func (d_ DraggingInfoWrapper) DraggedImageLocation() foundation.Point {
	rv := objc.CallMethod[foundation.Point](d_, objc.GetSelector("draggedImageLocation"))
	return rv
}

func (d_ DraggingInfoWrapper) ImplementsSetAnimatesToDestination() bool {
	return d_.RespondsToSelector(objc.GetSelector("setAnimatesToDestination:"))
}

func (d_ DraggingInfoWrapper) SetAnimatesToDestination(value bool) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setAnimatesToDestination:"), value)
}

func (d_ DraggingInfoWrapper) ImplementsAnimatesToDestination() bool {
	return d_.RespondsToSelector(objc.GetSelector("animatesToDestination"))
}

func (d_ DraggingInfoWrapper) AnimatesToDestination() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("animatesToDestination"))
	return rv
}

func (d_ DraggingInfoWrapper) ImplementsSetDraggingFormation() bool {
	return d_.RespondsToSelector(objc.GetSelector("setDraggingFormation:"))
}

func (d_ DraggingInfoWrapper) SetDraggingFormation(value DraggingFormation) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDraggingFormation:"), value)
}

func (d_ DraggingInfoWrapper) ImplementsDraggingFormation() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingFormation"))
}

func (d_ DraggingInfoWrapper) DraggingFormation() DraggingFormation {
	rv := objc.CallMethod[DraggingFormation](d_, objc.GetSelector("draggingFormation"))
	return rv
}

func (d_ DraggingInfoWrapper) ImplementsSpringLoadingHighlight() bool {
	return d_.RespondsToSelector(objc.GetSelector("springLoadingHighlight"))
}

func (d_ DraggingInfoWrapper) SpringLoadingHighlight() SpringLoadingHighlight {
	rv := objc.CallMethod[SpringLoadingHighlight](d_, objc.GetSelector("springLoadingHighlight"))
	return rv
}
