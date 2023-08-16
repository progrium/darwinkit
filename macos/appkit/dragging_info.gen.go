// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of methods that supply information about a dragging session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdragginginfo?language=objc
type PDraggingInfo interface {
	// optional
	EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock(enumOpts DraggingItemEnumerationOptions, view View, classArray []objc.Class, searchOptions map[PasteboardReadingOptionKey]objc.Object, block func(draggingItem DraggingItem, idx int, stop *bool))
	HasEnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock() bool

	// optional
	ResetSpringLoading()
	HasResetSpringLoading() bool

	// optional
	SlideDraggedImageTo(screenPoint foundation.Point)
	HasSlideDraggedImageTo() bool

	// optional
	DraggingDestinationWindow() IWindow
	HasDraggingDestinationWindow() bool

	// optional
	DraggingPasteboard() IPasteboard
	HasDraggingPasteboard() bool

	// optional
	DraggingSourceOperationMask() DragOperation
	HasDraggingSourceOperationMask() bool

	// optional
	SetNumberOfValidItemsForDrop(value int)
	HasSetNumberOfValidItemsForDrop() bool

	// optional
	NumberOfValidItemsForDrop() int
	HasNumberOfValidItemsForDrop() bool

	// optional
	SetDraggingFormation(value DraggingFormation)
	HasSetDraggingFormation() bool

	// optional
	DraggingFormation() DraggingFormation
	HasDraggingFormation() bool

	// optional
	SpringLoadingHighlight() SpringLoadingHighlight
	HasSpringLoadingHighlight() bool

	// optional
	DraggedImageLocation() foundation.Point
	HasDraggedImageLocation() bool

	// optional
	DraggingLocation() foundation.Point
	HasDraggingLocation() bool

	// optional
	DraggingSource() objc.IObject
	HasDraggingSource() bool

	// optional
	SetAnimatesToDestination(value bool)
	HasSetAnimatesToDestination() bool

	// optional
	AnimatesToDestination() bool
	HasAnimatesToDestination() bool

	// optional
	DraggingSequenceNumber() int
	HasDraggingSequenceNumber() bool
}

// A concrete type wrapper for the [PDraggingInfo] protocol.
type DraggingInfoWrapper struct {
	objc.Object
}

func (d_ DraggingInfoWrapper) HasEnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock() bool {
	return d_.RespondsToSelector(objc.Sel("enumerateDraggingItemsWithOptions:forView:classes:searchOptions:usingBlock:"))
}

// Enumerates through each dragging item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdragginginfo/1416074-enumeratedraggingitemswithoption?language=objc
func (d_ DraggingInfoWrapper) EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock(enumOpts DraggingItemEnumerationOptions, view IView, classArray []objc.IClass, searchOptions map[PasteboardReadingOptionKey]objc.IObject, block func(draggingItem DraggingItem, idx int, stop *bool)) {
	objc.Call[objc.Void](d_, objc.Sel("enumerateDraggingItemsWithOptions:forView:classes:searchOptions:usingBlock:"), enumOpts, objc.Ptr(view), classArray, searchOptions, block)
}

func (d_ DraggingInfoWrapper) HasResetSpringLoading() bool {
	return d_.RespondsToSelector(objc.Sel("resetSpringLoading"))
}

// Resets a spring-loading operation to its initial state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdragginginfo/1416012-resetspringloading?language=objc
func (d_ DraggingInfoWrapper) ResetSpringLoading() {
	objc.Call[objc.Void](d_, objc.Sel("resetSpringLoading"))
}

func (d_ DraggingInfoWrapper) HasSlideDraggedImageTo() bool {
	return d_.RespondsToSelector(objc.Sel("slideDraggedImageTo:"))
}

// Slides the image to a specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdragginginfo/1416045-slidedraggedimageto?language=objc
func (d_ DraggingInfoWrapper) SlideDraggedImageTo(screenPoint foundation.Point) {
	objc.Call[objc.Void](d_, objc.Sel("slideDraggedImageTo:"), screenPoint)
}

func (d_ DraggingInfoWrapper) HasDraggingDestinationWindow() bool {
	return d_.RespondsToSelector(objc.Sel("draggingDestinationWindow"))
}

// The destination window for the dragging operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdragginginfo/1416088-draggingdestinationwindow?language=objc
func (d_ DraggingInfoWrapper) DraggingDestinationWindow() Window {
	rv := objc.Call[Window](d_, objc.Sel("draggingDestinationWindow"))
	return rv
}

func (d_ DraggingInfoWrapper) HasDraggingPasteboard() bool {
	return d_.RespondsToSelector(objc.Sel("draggingPasteboard"))
}

// The pasteboard object that holds the dragged data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdragginginfo/1416092-draggingpasteboard?language=objc
func (d_ DraggingInfoWrapper) DraggingPasteboard() Pasteboard {
	rv := objc.Call[Pasteboard](d_, objc.Sel("draggingPasteboard"))
	return rv
}

func (d_ DraggingInfoWrapper) HasDraggingSourceOperationMask() bool {
	return d_.RespondsToSelector(objc.Sel("draggingSourceOperationMask"))
}

// Information about the dragging operation and the data it contains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdragginginfo/1415966-draggingsourceoperationmask?language=objc
func (d_ DraggingInfoWrapper) DraggingSourceOperationMask() DragOperation {
	rv := objc.Call[DragOperation](d_, objc.Sel("draggingSourceOperationMask"))
	return rv
}

func (d_ DraggingInfoWrapper) HasSetNumberOfValidItemsForDrop() bool {
	return d_.RespondsToSelector(objc.Sel("setNumberOfValidItemsForDrop:"))
}

// The number of valid items for a drop operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdragginginfo/1416033-numberofvaliditemsfordrop?language=objc
func (d_ DraggingInfoWrapper) SetNumberOfValidItemsForDrop(value int) {
	objc.Call[objc.Void](d_, objc.Sel("setNumberOfValidItemsForDrop:"), value)
}

func (d_ DraggingInfoWrapper) HasNumberOfValidItemsForDrop() bool {
	return d_.RespondsToSelector(objc.Sel("numberOfValidItemsForDrop"))
}

// The number of valid items for a drop operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdragginginfo/1416033-numberofvaliditemsfordrop?language=objc
func (d_ DraggingInfoWrapper) NumberOfValidItemsForDrop() int {
	rv := objc.Call[int](d_, objc.Sel("numberOfValidItemsForDrop"))
	return rv
}

func (d_ DraggingInfoWrapper) HasSetDraggingFormation() bool {
	return d_.RespondsToSelector(objc.Sel("setDraggingFormation:"))
}

// The formation of the dragging items while the drag is over the destination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdragginginfo/1416041-draggingformation?language=objc
func (d_ DraggingInfoWrapper) SetDraggingFormation(value DraggingFormation) {
	objc.Call[objc.Void](d_, objc.Sel("setDraggingFormation:"), value)
}

func (d_ DraggingInfoWrapper) HasDraggingFormation() bool {
	return d_.RespondsToSelector(objc.Sel("draggingFormation"))
}

// The formation of the dragging items while the drag is over the destination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdragginginfo/1416041-draggingformation?language=objc
func (d_ DraggingInfoWrapper) DraggingFormation() DraggingFormation {
	rv := objc.Call[DraggingFormation](d_, objc.Sel("draggingFormation"))
	return rv
}

func (d_ DraggingInfoWrapper) HasSpringLoadingHighlight() bool {
	return d_.RespondsToSelector(objc.Sel("springLoadingHighlight"))
}

// A highlighting style for your app’s user interface to display during a spring-loading operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdragginginfo/1416077-springloadinghighlight?language=objc
func (d_ DraggingInfoWrapper) SpringLoadingHighlight() SpringLoadingHighlight {
	rv := objc.Call[SpringLoadingHighlight](d_, objc.Sel("springLoadingHighlight"))
	return rv
}

func (d_ DraggingInfoWrapper) HasDraggedImageLocation() bool {
	return d_.RespondsToSelector(objc.Sel("draggedImageLocation"))
}

// The current location of the dragged image’s origin, in the base coordinate system of the destination object’s window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdragginginfo/1415990-draggedimagelocation?language=objc
func (d_ DraggingInfoWrapper) DraggedImageLocation() foundation.Point {
	rv := objc.Call[foundation.Point](d_, objc.Sel("draggedImageLocation"))
	return rv
}

func (d_ DraggingInfoWrapper) HasDraggingLocation() bool {
	return d_.RespondsToSelector(objc.Sel("draggingLocation"))
}

// The current location of the mouse pointer in the base coordinate system of the destination object’s window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdragginginfo/1416070-dragginglocation?language=objc
func (d_ DraggingInfoWrapper) DraggingLocation() foundation.Point {
	rv := objc.Call[foundation.Point](d_, objc.Sel("draggingLocation"))
	return rv
}

func (d_ DraggingInfoWrapper) HasDraggingSource() bool {
	return d_.RespondsToSelector(objc.Sel("draggingSource"))
}

// The source, or owner, of the dragged data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdragginginfo/1416068-draggingsource?language=objc
func (d_ DraggingInfoWrapper) DraggingSource() objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("draggingSource"))
	return rv
}

func (d_ DraggingInfoWrapper) HasSetAnimatesToDestination() bool {
	return d_.RespondsToSelector(objc.Sel("setAnimatesToDestination:"))
}

// A Boolean value that indicates whether the dragging formation animates while the drag is over the destination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdragginginfo/1416098-animatestodestination?language=objc
func (d_ DraggingInfoWrapper) SetAnimatesToDestination(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setAnimatesToDestination:"), value)
}

func (d_ DraggingInfoWrapper) HasAnimatesToDestination() bool {
	return d_.RespondsToSelector(objc.Sel("animatesToDestination"))
}

// A Boolean value that indicates whether the dragging formation animates while the drag is over the destination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdragginginfo/1416098-animatestodestination?language=objc
func (d_ DraggingInfoWrapper) AnimatesToDestination() bool {
	rv := objc.Call[bool](d_, objc.Sel("animatesToDestination"))
	return rv
}

func (d_ DraggingInfoWrapper) HasDraggingSequenceNumber() bool {
	return d_.RespondsToSelector(objc.Sel("draggingSequenceNumber"))
}

// A number that uniquely identifies the dragging session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdragginginfo/1416039-draggingsequencenumber?language=objc
func (d_ DraggingInfoWrapper) DraggingSequenceNumber() int {
	rv := objc.Call[int](d_, objc.Sel("draggingSequenceNumber"))
	return rv
}
