// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var DraggingSessionClass = _DraggingSessionClass{objc.GetClass("NSDraggingSession")}

type _DraggingSessionClass struct {
	objc.Class
}

type IDraggingSession interface {
	objc.IObject
	EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock(enumOpts DraggingItemEnumerationOptions, view IView, classArray []objc.IClass, searchOptions map[PasteboardReadingOptionKey]objc.IObject, block func(draggingItem DraggingItem, idx int, stop *bool))
	DraggingPasteboard() Pasteboard
	AnimatesToStartingPositionsOnCancelOrFail() bool
	SetAnimatesToStartingPositionsOnCancelOrFail(value bool)
	DraggingFormation() DraggingFormation
	SetDraggingFormation(value DraggingFormation)
	DraggingSequenceNumber() int
	DraggingLocation() foundation.Point
	DraggingLeaderIndex() int
	SetDraggingLeaderIndex(value int)
}

type DraggingSession struct {
	objc.Object
}

func MakeDraggingSession(ptr unsafe.Pointer) DraggingSession {
	return DraggingSession{
		Object: objc.MakeObject(ptr),
	}
}

func (dc _DraggingSessionClass) Alloc() DraggingSession {
	rv := objc.CallMethod[DraggingSession](dc, objc.GetSelector("alloc"))
	return rv
}

func DraggingSession_Alloc() DraggingSession {
	return DraggingSessionClass.Alloc()
}

func (dc _DraggingSessionClass) New() DraggingSession {
	rv := objc.CallMethod[DraggingSession](dc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewDraggingSession() DraggingSession {
	return DraggingSessionClass.New()
}

func DraggingSession_New() DraggingSession {
	return DraggingSessionClass.New()
}

func (d_ DraggingSession) Init() DraggingSession {
	rv := objc.CallMethod[DraggingSession](d_, objc.GetSelector("init"))
	return rv
}

func DraggingSession_Init() DraggingSession {
	return DraggingSessionClass.Alloc().Init()
}

func (d_ DraggingSession) EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock(enumOpts DraggingItemEnumerationOptions, view IView, classArray []objc.IClass, searchOptions map[PasteboardReadingOptionKey]objc.IObject, block func(draggingItem DraggingItem, idx int, stop *bool)) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("enumerateDraggingItemsWithOptions:forView:classes:searchOptions:usingBlock:"), enumOpts, objc.ExtractPtr(view), classArray, searchOptions, block)
}

func (d_ DraggingSession) DraggingPasteboard() Pasteboard {
	rv := objc.CallMethod[Pasteboard](d_, objc.GetSelector("draggingPasteboard"))
	return rv
}

func (d_ DraggingSession) AnimatesToStartingPositionsOnCancelOrFail() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("animatesToStartingPositionsOnCancelOrFail"))
	return rv
}

func (d_ DraggingSession) SetAnimatesToStartingPositionsOnCancelOrFail(value bool) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setAnimatesToStartingPositionsOnCancelOrFail:"), value)
}

func (d_ DraggingSession) DraggingFormation() DraggingFormation {
	rv := objc.CallMethod[DraggingFormation](d_, objc.GetSelector("draggingFormation"))
	return rv
}

func (d_ DraggingSession) SetDraggingFormation(value DraggingFormation) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDraggingFormation:"), value)
}

func (d_ DraggingSession) DraggingSequenceNumber() int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("draggingSequenceNumber"))
	return rv
}

func (d_ DraggingSession) DraggingLocation() foundation.Point {
	rv := objc.CallMethod[foundation.Point](d_, objc.GetSelector("draggingLocation"))
	return rv
}

func (d_ DraggingSession) DraggingLeaderIndex() int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("draggingLeaderIndex"))
	return rv
}

func (d_ DraggingSession) SetDraggingLeaderIndex(value int) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDraggingLeaderIndex:"), value)
}
