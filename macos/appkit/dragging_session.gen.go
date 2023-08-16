// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DraggingSession] class.
var DraggingSessionClass = _DraggingSessionClass{objc.GetClass("NSDraggingSession")}

type _DraggingSessionClass struct {
	objc.Class
}

// An interface definition for the [DraggingSession] class.
type IDraggingSession interface {
	objc.IObject
	EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock(enumOpts DraggingItemEnumerationOptions, view IView, classArray []objc.IClass, searchOptions map[PasteboardReadingOptionKey]objc.IObject, block func(draggingItem DraggingItem, idx int, stop *bool))
	DraggingPasteboard() Pasteboard
	DraggingFormation() DraggingFormation
	SetDraggingFormation(value DraggingFormation)
	DraggingLeaderIndex() int
	SetDraggingLeaderIndex(value int)
	AnimatesToStartingPositionsOnCancelOrFail() bool
	SetAnimatesToStartingPositionsOnCancelOrFail(value bool)
	DraggingLocation() foundation.Point
	DraggingSequenceNumber() int
}

// The encapsulation of a drag-and-drop action that supports modification of the drag while in progress. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingsession?language=objc
type DraggingSession struct {
	objc.Object
}

func DraggingSessionFrom(ptr unsafe.Pointer) DraggingSession {
	return DraggingSession{
		Object: objc.ObjectFrom(ptr),
	}
}

func (dc _DraggingSessionClass) Alloc() DraggingSession {
	rv := objc.Call[DraggingSession](dc, objc.Sel("alloc"))
	return rv
}

func DraggingSession_Alloc() DraggingSession {
	return DraggingSessionClass.Alloc()
}

func (dc _DraggingSessionClass) New() DraggingSession {
	rv := objc.Call[DraggingSession](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDraggingSession() DraggingSession {
	return DraggingSessionClass.New()
}

func (d_ DraggingSession) Init() DraggingSession {
	rv := objc.Call[DraggingSession](d_, objc.Sel("init"))
	return rv
}

// Enumerates through each dragging item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingsession/1532445-enumeratedraggingitemswithoption?language=objc
func (d_ DraggingSession) EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock(enumOpts DraggingItemEnumerationOptions, view IView, classArray []objc.IClass, searchOptions map[PasteboardReadingOptionKey]objc.IObject, block func(draggingItem DraggingItem, idx int, stop *bool)) {
	objc.Call[objc.Void](d_, objc.Sel("enumerateDraggingItemsWithOptions:forView:classes:searchOptions:usingBlock:"), enumOpts, objc.Ptr(view), classArray, searchOptions, block)
}

// Returns the pasteboard object that contains the data being dragged. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingsession/1534103-draggingpasteboard?language=objc
func (d_ DraggingSession) DraggingPasteboard() Pasteboard {
	rv := objc.Call[Pasteboard](d_, objc.Sel("draggingPasteboard"))
	return rv
}

// Controls the dragging formation when the drag is not over the source or a valid destination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingsession/1524544-draggingformation?language=objc
func (d_ DraggingSession) DraggingFormation() DraggingFormation {
	rv := objc.Call[DraggingFormation](d_, objc.Sel("draggingFormation"))
	return rv
}

// Controls the dragging formation when the drag is not over the source or a valid destination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingsession/1524544-draggingformation?language=objc
func (d_ DraggingSession) SetDraggingFormation(value DraggingFormation) {
	objc.Call[objc.Void](d_, objc.Sel("setDraggingFormation:"), value)
}

// The index of the dragging item under the cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingsession/1533729-draggingleaderindex?language=objc
func (d_ DraggingSession) DraggingLeaderIndex() int {
	rv := objc.Call[int](d_, objc.Sel("draggingLeaderIndex"))
	return rv
}

// The index of the dragging item under the cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingsession/1533729-draggingleaderindex?language=objc
func (d_ DraggingSession) SetDraggingLeaderIndex(value int) {
	objc.Call[objc.Void](d_, objc.Sel("setDraggingLeaderIndex:"), value)
}

// Controls whether the dragging image animates back to its starting point on a cancelled or failed drag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingsession/1531277-animatestostartingpositionsoncan?language=objc
func (d_ DraggingSession) AnimatesToStartingPositionsOnCancelOrFail() bool {
	rv := objc.Call[bool](d_, objc.Sel("animatesToStartingPositionsOnCancelOrFail"))
	return rv
}

// Controls whether the dragging image animates back to its starting point on a cancelled or failed drag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingsession/1531277-animatestostartingpositionsoncan?language=objc
func (d_ DraggingSession) SetAnimatesToStartingPositionsOnCancelOrFail(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setAnimatesToStartingPositionsOnCancelOrFail:"), value)
}

// The current cursor location of the drag in screen coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingsession/1529395-dragginglocation?language=objc
func (d_ DraggingSession) DraggingLocation() foundation.Point {
	rv := objc.Call[foundation.Point](d_, objc.Sel("draggingLocation"))
	return rv
}

// Returns a number that uniquely identifies the dragging session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingsession/1533229-draggingsequencenumber?language=objc
func (d_ DraggingSession) DraggingSequenceNumber() int {
	rv := objc.Call[int](d_, objc.Sel("draggingSequenceNumber"))
	return rv
}
