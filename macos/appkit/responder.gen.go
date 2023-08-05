// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var ResponderClass = _ResponderClass{objc.GetClass("NSResponder")}

type _ResponderClass struct {
	objc.Class
}

type IResponder interface {
	objc.IObject
	BecomeFirstResponder() bool
	ResignFirstResponder() bool
	ValidateProposedFirstResponderForEvent(responder IResponder, event IEvent) bool
	MouseDown(event IEvent)
	MouseDragged(event IEvent)
	MouseUp(event IEvent)
	MouseMoved(event IEvent)
	MouseEntered(event IEvent)
	MouseExited(event IEvent)
	RightMouseDown(event IEvent)
	RightMouseDragged(event IEvent)
	RightMouseUp(event IEvent)
	OtherMouseDown(event IEvent)
	OtherMouseDragged(event IEvent)
	OtherMouseUp(event IEvent)
	KeyDown(event IEvent)
	KeyUp(event IEvent)
	InterpretKeyEvents(eventArray []IEvent)
	PerformKeyEquivalent(event IEvent) bool
	FlushBufferedKeyEvents()
	PressureChangeWithEvent(event IEvent)
	CursorUpdate(event IEvent)
	FlagsChanged(event IEvent)
	TabletPoint(event IEvent)
	TabletProximity(event IEvent)
	HelpRequested(eventPtr IEvent)
	ScrollWheel(event IEvent)
	QuickLookWithEvent(event IEvent)
	ChangeModeWithEvent(event IEvent)
	SupplementalTargetForActionSender(action objc.Selector, sender objc.IObject) objc.Object
	EncodeRestorableStateWithCoder(coder foundation.ICoder)
	EncodeRestorableStateWithCoderBackgroundQueue(coder foundation.ICoder, queue foundation.IOperationQueue)
	RestoreStateWithCoder(coder foundation.ICoder)
	InvalidateRestorableState()
	UpdateUserActivityState(userActivity foundation.IUserActivity)
	PresentError(error foundation.IError) bool
	PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error foundation.IError, window IWindow, delegate objc.IObject, didPresentSelector objc.Selector, contextInfo unsafe.Pointer)
	WillPresentError(error foundation.IError) foundation.Error
	TryToPerformWith(action objc.Selector, object objc.IObject) bool
	ValidRequestorForSendTypeReturnType(sendType PasteboardType, returnType PasteboardType) objc.Object
	ShouldBeTreatedAsInkEvent(event IEvent) bool
	NoResponderFor(eventSelector objc.Selector)
	BeginGestureWithEvent(event IEvent)
	EndGestureWithEvent(event IEvent)
	MagnifyWithEvent(event IEvent)
	RotateWithEvent(event IEvent)
	SwipeWithEvent(event IEvent)
	TouchesBeganWithEvent(event IEvent)
	TouchesMovedWithEvent(event IEvent)
	TouchesCancelledWithEvent(event IEvent)
	TouchesEndedWithEvent(event IEvent)
	WantsForwardedScrollEventsForAxis(axis EventGestureAxis) bool
	SmartMagnifyWithEvent(event IEvent)
	WantsScrollEventsForSwipeTrackingOnAxis(axis EventGestureAxis) bool
	MakeTouchBar() TouchBar
	PerformTextFinderAction(sender objc.IObject)
	NewWindowForTab(sender objc.IObject)
	ShowContextHelp(sender objc.IObject)
	AcceptsFirstResponder() bool
	NextResponder() Responder
	SetNextResponder(value IResponder)
	UserActivity() foundation.UserActivity
	SetUserActivity(value foundation.IUserActivity)
	Menu() Menu
	SetMenu(value IMenu)
	UndoManager() foundation.UndoManager
	TouchBar() TouchBar
	SetTouchBar(value ITouchBar)
}

type Responder struct {
	objc.Object
}

func MakeResponder(ptr unsafe.Pointer) Responder {
	return Responder{
		Object: objc.MakeObject(ptr),
	}
}

func (r_ Responder) Init() Responder {
	rv := objc.CallMethod[Responder](r_, objc.GetSelector("init"))
	return rv
}

func Responder_Init() Responder {
	return ResponderClass.Alloc().Init()
}

func (rc _ResponderClass) Alloc() Responder {
	rv := objc.CallMethod[Responder](rc, objc.GetSelector("alloc"))
	return rv
}

func Responder_Alloc() Responder {
	return ResponderClass.Alloc()
}

func (rc _ResponderClass) New() Responder {
	rv := objc.CallMethod[Responder](rc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewResponder() Responder {
	return ResponderClass.New()
}

func Responder_New() Responder {
	return ResponderClass.New()
}

func (r_ Responder) BecomeFirstResponder() bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("becomeFirstResponder"))
	return rv
}

func (r_ Responder) ResignFirstResponder() bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("resignFirstResponder"))
	return rv
}

func (r_ Responder) ValidateProposedFirstResponderForEvent(responder IResponder, event IEvent) bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("validateProposedFirstResponder:forEvent:"), objc.ExtractPtr(responder), objc.ExtractPtr(event))
	return rv
}

func (r_ Responder) MouseDown(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("mouseDown:"), objc.ExtractPtr(event))
}

func (r_ Responder) MouseDragged(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("mouseDragged:"), objc.ExtractPtr(event))
}

func (r_ Responder) MouseUp(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("mouseUp:"), objc.ExtractPtr(event))
}

func (r_ Responder) MouseMoved(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("mouseMoved:"), objc.ExtractPtr(event))
}

func (r_ Responder) MouseEntered(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("mouseEntered:"), objc.ExtractPtr(event))
}

func (r_ Responder) MouseExited(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("mouseExited:"), objc.ExtractPtr(event))
}

func (r_ Responder) RightMouseDown(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("rightMouseDown:"), objc.ExtractPtr(event))
}

func (r_ Responder) RightMouseDragged(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("rightMouseDragged:"), objc.ExtractPtr(event))
}

func (r_ Responder) RightMouseUp(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("rightMouseUp:"), objc.ExtractPtr(event))
}

func (r_ Responder) OtherMouseDown(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("otherMouseDown:"), objc.ExtractPtr(event))
}

func (r_ Responder) OtherMouseDragged(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("otherMouseDragged:"), objc.ExtractPtr(event))
}

func (r_ Responder) OtherMouseUp(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("otherMouseUp:"), objc.ExtractPtr(event))
}

func (r_ Responder) KeyDown(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("keyDown:"), objc.ExtractPtr(event))
}

func (r_ Responder) KeyUp(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("keyUp:"), objc.ExtractPtr(event))
}

func (r_ Responder) InterpretKeyEvents(eventArray []IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("interpretKeyEvents:"), eventArray)
}

func (r_ Responder) PerformKeyEquivalent(event IEvent) bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("performKeyEquivalent:"), objc.ExtractPtr(event))
	return rv
}

func (r_ Responder) FlushBufferedKeyEvents() {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("flushBufferedKeyEvents"))
}

func (r_ Responder) PressureChangeWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("pressureChangeWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) CursorUpdate(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("cursorUpdate:"), objc.ExtractPtr(event))
}

func (r_ Responder) FlagsChanged(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("flagsChanged:"), objc.ExtractPtr(event))
}

func (r_ Responder) TabletPoint(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("tabletPoint:"), objc.ExtractPtr(event))
}

func (r_ Responder) TabletProximity(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("tabletProximity:"), objc.ExtractPtr(event))
}

func (r_ Responder) HelpRequested(eventPtr IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("helpRequested:"), objc.ExtractPtr(eventPtr))
}

func (r_ Responder) ScrollWheel(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("scrollWheel:"), objc.ExtractPtr(event))
}

func (r_ Responder) QuickLookWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("quickLookWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) ChangeModeWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("changeModeWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) SupplementalTargetForActionSender(action objc.Selector, sender objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](r_, objc.GetSelector("supplementalTargetForAction:sender:"), action, objc.ExtractPtr(sender))
	return rv
}

func (rc _ResponderClass) AllowedClassesForRestorableStateKeyPath(keyPath string) []objc.Class {
	rv := objc.CallMethod[[]objc.Class](rc, objc.GetSelector("allowedClassesForRestorableStateKeyPath:"), keyPath)
	return rv
}

func Responder_AllowedClassesForRestorableStateKeyPath(keyPath string) []objc.Class {
	return ResponderClass.AllowedClassesForRestorableStateKeyPath(keyPath)
}

func (r_ Responder) EncodeRestorableStateWithCoder(coder foundation.ICoder) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("encodeRestorableStateWithCoder:"), objc.ExtractPtr(coder))
}

func (r_ Responder) EncodeRestorableStateWithCoderBackgroundQueue(coder foundation.ICoder, queue foundation.IOperationQueue) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("encodeRestorableStateWithCoder:backgroundQueue:"), objc.ExtractPtr(coder), objc.ExtractPtr(queue))
}

func (r_ Responder) RestoreStateWithCoder(coder foundation.ICoder) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("restoreStateWithCoder:"), objc.ExtractPtr(coder))
}

func (r_ Responder) InvalidateRestorableState() {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("invalidateRestorableState"))
}

func (r_ Responder) UpdateUserActivityState(userActivity foundation.IUserActivity) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("updateUserActivityState:"), objc.ExtractPtr(userActivity))
}

func (r_ Responder) PresentError(error foundation.IError) bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("presentError:"), objc.ExtractPtr(error))
	return rv
}

func (r_ Responder) PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error foundation.IError, window IWindow, delegate objc.IObject, didPresentSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("presentError:modalForWindow:delegate:didPresentSelector:contextInfo:"), objc.ExtractPtr(error), objc.ExtractPtr(window), objc.ExtractPtr(delegate), didPresentSelector, contextInfo)
}

func (r_ Responder) WillPresentError(error foundation.IError) foundation.Error {
	rv := objc.CallMethod[foundation.Error](r_, objc.GetSelector("willPresentError:"), objc.ExtractPtr(error))
	return rv
}

func (r_ Responder) TryToPerformWith(action objc.Selector, object objc.IObject) bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("tryToPerform:with:"), action, objc.ExtractPtr(object))
	return rv
}

func (r_ Responder) ValidRequestorForSendTypeReturnType(sendType PasteboardType, returnType PasteboardType) objc.Object {
	rv := objc.CallMethod[objc.Object](r_, objc.GetSelector("validRequestorForSendType:returnType:"), sendType, returnType)
	return rv
}

func (r_ Responder) ShouldBeTreatedAsInkEvent(event IEvent) bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("shouldBeTreatedAsInkEvent:"), objc.ExtractPtr(event))
	return rv
}

func (r_ Responder) NoResponderFor(eventSelector objc.Selector) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("noResponderFor:"), eventSelector)
}

func (r_ Responder) BeginGestureWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("beginGestureWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) EndGestureWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("endGestureWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) MagnifyWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("magnifyWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) RotateWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("rotateWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) SwipeWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("swipeWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) TouchesBeganWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("touchesBeganWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) TouchesMovedWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("touchesMovedWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) TouchesCancelledWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("touchesCancelledWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) TouchesEndedWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("touchesEndedWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) WantsForwardedScrollEventsForAxis(axis EventGestureAxis) bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("wantsForwardedScrollEventsForAxis:"), axis)
	return rv
}

func (r_ Responder) SmartMagnifyWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("smartMagnifyWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) WantsScrollEventsForSwipeTrackingOnAxis(axis EventGestureAxis) bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("wantsScrollEventsForSwipeTrackingOnAxis:"), axis)
	return rv
}

func (r_ Responder) MakeTouchBar() TouchBar {
	rv := objc.CallMethod[TouchBar](r_, objc.GetSelector("makeTouchBar"))
	return rv
}

func (r_ Responder) PerformTextFinderAction(sender objc.IObject) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("performTextFinderAction:"), objc.ExtractPtr(sender))
}

func (r_ Responder) NewWindowForTab(sender objc.IObject) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("newWindowForTab:"), objc.ExtractPtr(sender))
}

func (r_ Responder) ShowContextHelp(sender objc.IObject) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("showContextHelp:"), objc.ExtractPtr(sender))
}

func (r_ Responder) AcceptsFirstResponder() bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("acceptsFirstResponder"))
	return rv
}

func (r_ Responder) NextResponder() Responder {
	rv := objc.CallMethod[Responder](r_, objc.GetSelector("nextResponder"))
	return rv
}

func (r_ Responder) SetNextResponder(value IResponder) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setNextResponder:"), objc.ExtractPtr(value))
}

func (rc _ResponderClass) RestorableStateKeyPaths() []string {
	rv := objc.CallMethod[[]string](rc, objc.GetSelector("restorableStateKeyPaths"))
	return rv
}

func Responder_RestorableStateKeyPaths() []string {
	return ResponderClass.RestorableStateKeyPaths()
}

func (r_ Responder) UserActivity() foundation.UserActivity {
	rv := objc.CallMethod[foundation.UserActivity](r_, objc.GetSelector("userActivity"))
	return rv
}

func (r_ Responder) SetUserActivity(value foundation.IUserActivity) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setUserActivity:"), objc.ExtractPtr(value))
}

func (r_ Responder) Menu() Menu {
	rv := objc.CallMethod[Menu](r_, objc.GetSelector("menu"))
	return rv
}

func (r_ Responder) SetMenu(value IMenu) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setMenu:"), objc.ExtractPtr(value))
}

func (r_ Responder) UndoManager() foundation.UndoManager {
	rv := objc.CallMethod[foundation.UndoManager](r_, objc.GetSelector("undoManager"))
	return rv
}

func (r_ Responder) TouchBar() TouchBar {
	rv := objc.CallMethod[TouchBar](r_, objc.GetSelector("touchBar"))
	return rv
}

func (r_ Responder) SetTouchBar(value ITouchBar) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setTouchBar:"), objc.ExtractPtr(value))
}
