// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var GestureRecognizerClass = _GestureRecognizerClass{objc.GetClass("NSGestureRecognizer")}

type _GestureRecognizerClass struct {
	objc.Class
}

type IGestureRecognizer interface {
	objc.IObject
	LocationInView(view IView) foundation.Point
	Reset()
	MouseDown(event IEvent)
	MouseDragged(event IEvent)
	MouseUp(event IEvent)
	OtherMouseDown(event IEvent)
	OtherMouseDragged(event IEvent)
	OtherMouseUp(event IEvent)
	RightMouseDown(event IEvent)
	RightMouseDragged(event IEvent)
	RightMouseUp(event IEvent)
	MagnifyWithEvent(event IEvent)
	RotateWithEvent(event IEvent)
	CanBePreventedByGestureRecognizer(preventingGestureRecognizer IGestureRecognizer) bool
	CanPreventGestureRecognizer(preventedGestureRecognizer IGestureRecognizer) bool
	ShouldBeRequiredToFailByGestureRecognizer(otherGestureRecognizer IGestureRecognizer) bool
	ShouldRequireFailureOfGestureRecognizer(otherGestureRecognizer IGestureRecognizer) bool
	KeyDown(event IEvent)
	KeyUp(event IEvent)
	TabletPoint(event IEvent)
	FlagsChanged(event IEvent)
	PressureChangeWithEvent(event IEvent)
	TouchesBeganWithEvent(event IEvent)
	TouchesCancelledWithEvent(event IEvent)
	TouchesEndedWithEvent(event IEvent)
	TouchesMovedWithEvent(event IEvent)
	Action() objc.Selector
	SetAction(value objc.Selector)
	Target() objc.Object
	SetTarget(value objc.IObject)
	State() GestureRecognizerState
	View() View
	IsEnabled() bool
	SetEnabled(value bool)
	DelaysPrimaryMouseButtonEvents() bool
	SetDelaysPrimaryMouseButtonEvents(value bool)
	DelaysSecondaryMouseButtonEvents() bool
	SetDelaysSecondaryMouseButtonEvents(value bool)
	DelaysOtherMouseButtonEvents() bool
	SetDelaysOtherMouseButtonEvents(value bool)
	DelaysKeyEvents() bool
	SetDelaysKeyEvents(value bool)
	DelaysMagnificationEvents() bool
	SetDelaysMagnificationEvents(value bool)
	DelaysRotationEvents() bool
	SetDelaysRotationEvents(value bool)
	Delegate() GestureRecognizerDelegateWrapper
	SetDelegate(value IGestureRecognizerDelegate)
	SetDelegate0(value objc.IObject)
	PressureConfiguration() PressureConfiguration
	SetPressureConfiguration(value IPressureConfiguration)
	AllowedTouchTypes() TouchTypeMask
	SetAllowedTouchTypes(value TouchTypeMask)
}

type GestureRecognizer struct {
	objc.Object
}

func MakeGestureRecognizer(ptr unsafe.Pointer) GestureRecognizer {
	return GestureRecognizer{
		Object: objc.MakeObject(ptr),
	}
}

func (g_ GestureRecognizer) InitWithTargetAction(target objc.IObject, action objc.Selector) GestureRecognizer {
	rv := objc.CallMethod[GestureRecognizer](g_, objc.GetSelector("initWithTarget:action:"), objc.ExtractPtr(target), action)
	return rv
}

func GestureRecognizer_InitWithTargetAction(target objc.IObject, action objc.Selector) GestureRecognizer {
	return GestureRecognizerClass.Alloc().InitWithTargetAction(target, action)
}

func (gc _GestureRecognizerClass) Alloc() GestureRecognizer {
	rv := objc.CallMethod[GestureRecognizer](gc, objc.GetSelector("alloc"))
	return rv
}

func GestureRecognizer_Alloc() GestureRecognizer {
	return GestureRecognizerClass.Alloc()
}

func (gc _GestureRecognizerClass) New() GestureRecognizer {
	rv := objc.CallMethod[GestureRecognizer](gc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewGestureRecognizer() GestureRecognizer {
	return GestureRecognizerClass.New()
}

func GestureRecognizer_New() GestureRecognizer {
	return GestureRecognizerClass.New()
}

func (g_ GestureRecognizer) Init() GestureRecognizer {
	rv := objc.CallMethod[GestureRecognizer](g_, objc.GetSelector("init"))
	return rv
}

func GestureRecognizer_Init() GestureRecognizer {
	return GestureRecognizerClass.Alloc().Init()
}

func (g_ GestureRecognizer) LocationInView(view IView) foundation.Point {
	rv := objc.CallMethod[foundation.Point](g_, objc.GetSelector("locationInView:"), objc.ExtractPtr(view))
	return rv
}

func (g_ GestureRecognizer) Reset() {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("reset"))
}

func (g_ GestureRecognizer) MouseDown(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("mouseDown:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) MouseDragged(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("mouseDragged:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) MouseUp(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("mouseUp:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) OtherMouseDown(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("otherMouseDown:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) OtherMouseDragged(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("otherMouseDragged:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) OtherMouseUp(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("otherMouseUp:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) RightMouseDown(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("rightMouseDown:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) RightMouseDragged(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("rightMouseDragged:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) RightMouseUp(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("rightMouseUp:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) MagnifyWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("magnifyWithEvent:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) RotateWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("rotateWithEvent:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) CanBePreventedByGestureRecognizer(preventingGestureRecognizer IGestureRecognizer) bool {
	rv := objc.CallMethod[bool](g_, objc.GetSelector("canBePreventedByGestureRecognizer:"), objc.ExtractPtr(preventingGestureRecognizer))
	return rv
}

func (g_ GestureRecognizer) CanPreventGestureRecognizer(preventedGestureRecognizer IGestureRecognizer) bool {
	rv := objc.CallMethod[bool](g_, objc.GetSelector("canPreventGestureRecognizer:"), objc.ExtractPtr(preventedGestureRecognizer))
	return rv
}

func (g_ GestureRecognizer) ShouldBeRequiredToFailByGestureRecognizer(otherGestureRecognizer IGestureRecognizer) bool {
	rv := objc.CallMethod[bool](g_, objc.GetSelector("shouldBeRequiredToFailByGestureRecognizer:"), objc.ExtractPtr(otherGestureRecognizer))
	return rv
}

func (g_ GestureRecognizer) ShouldRequireFailureOfGestureRecognizer(otherGestureRecognizer IGestureRecognizer) bool {
	rv := objc.CallMethod[bool](g_, objc.GetSelector("shouldRequireFailureOfGestureRecognizer:"), objc.ExtractPtr(otherGestureRecognizer))
	return rv
}

func (g_ GestureRecognizer) KeyDown(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("keyDown:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) KeyUp(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("keyUp:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) TabletPoint(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("tabletPoint:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) FlagsChanged(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("flagsChanged:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) PressureChangeWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("pressureChangeWithEvent:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) TouchesBeganWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("touchesBeganWithEvent:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) TouchesCancelledWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("touchesCancelledWithEvent:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) TouchesEndedWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("touchesEndedWithEvent:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) TouchesMovedWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("touchesMovedWithEvent:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) Action() objc.Selector {
	rv := objc.CallMethod[objc.Selector](g_, objc.GetSelector("action"))
	return rv
}

func (g_ GestureRecognizer) SetAction(value objc.Selector) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setAction:"), value)
}

func (g_ GestureRecognizer) Target() objc.Object {
	rv := objc.CallMethod[objc.Object](g_, objc.GetSelector("target"))
	return rv
}

func (g_ GestureRecognizer) SetTarget(value objc.IObject) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setTarget:"), objc.ExtractPtr(value))
}

func (g_ GestureRecognizer) State() GestureRecognizerState {
	rv := objc.CallMethod[GestureRecognizerState](g_, objc.GetSelector("state"))
	return rv
}

func (g_ GestureRecognizer) View() View {
	rv := objc.CallMethod[View](g_, objc.GetSelector("view"))
	return rv
}

func (g_ GestureRecognizer) IsEnabled() bool {
	rv := objc.CallMethod[bool](g_, objc.GetSelector("isEnabled"))
	return rv
}

func (g_ GestureRecognizer) SetEnabled(value bool) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setEnabled:"), value)
}

func (g_ GestureRecognizer) DelaysPrimaryMouseButtonEvents() bool {
	rv := objc.CallMethod[bool](g_, objc.GetSelector("delaysPrimaryMouseButtonEvents"))
	return rv
}

func (g_ GestureRecognizer) SetDelaysPrimaryMouseButtonEvents(value bool) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setDelaysPrimaryMouseButtonEvents:"), value)
}

func (g_ GestureRecognizer) DelaysSecondaryMouseButtonEvents() bool {
	rv := objc.CallMethod[bool](g_, objc.GetSelector("delaysSecondaryMouseButtonEvents"))
	return rv
}

func (g_ GestureRecognizer) SetDelaysSecondaryMouseButtonEvents(value bool) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setDelaysSecondaryMouseButtonEvents:"), value)
}

func (g_ GestureRecognizer) DelaysOtherMouseButtonEvents() bool {
	rv := objc.CallMethod[bool](g_, objc.GetSelector("delaysOtherMouseButtonEvents"))
	return rv
}

func (g_ GestureRecognizer) SetDelaysOtherMouseButtonEvents(value bool) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setDelaysOtherMouseButtonEvents:"), value)
}

func (g_ GestureRecognizer) DelaysKeyEvents() bool {
	rv := objc.CallMethod[bool](g_, objc.GetSelector("delaysKeyEvents"))
	return rv
}

func (g_ GestureRecognizer) SetDelaysKeyEvents(value bool) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setDelaysKeyEvents:"), value)
}

func (g_ GestureRecognizer) DelaysMagnificationEvents() bool {
	rv := objc.CallMethod[bool](g_, objc.GetSelector("delaysMagnificationEvents"))
	return rv
}

func (g_ GestureRecognizer) SetDelaysMagnificationEvents(value bool) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setDelaysMagnificationEvents:"), value)
}

func (g_ GestureRecognizer) DelaysRotationEvents() bool {
	rv := objc.CallMethod[bool](g_, objc.GetSelector("delaysRotationEvents"))
	return rv
}

func (g_ GestureRecognizer) SetDelaysRotationEvents(value bool) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setDelaysRotationEvents:"), value)
}

func (g_ GestureRecognizer) Delegate() GestureRecognizerDelegateWrapper {
	rv := objc.CallMethod[GestureRecognizerDelegateWrapper](g_, objc.GetSelector("delegate"))
	return rv
}

func (g_ GestureRecognizer) SetDelegate(value IGestureRecognizerDelegate) {
	po := objc.WrapAsProtocol("NSGestureRecognizerDelegate", value)
	objc.SetAssociatedObject(g_, objc.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setDelegate:"), po)
}

func (g_ GestureRecognizer) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}

func (g_ GestureRecognizer) PressureConfiguration() PressureConfiguration {
	rv := objc.CallMethod[PressureConfiguration](g_, objc.GetSelector("pressureConfiguration"))
	return rv
}

func (g_ GestureRecognizer) SetPressureConfiguration(value IPressureConfiguration) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setPressureConfiguration:"), objc.ExtractPtr(value))
}

func (g_ GestureRecognizer) AllowedTouchTypes() TouchTypeMask {
	rv := objc.CallMethod[TouchTypeMask](g_, objc.GetSelector("allowedTouchTypes"))
	return rv
}

func (g_ GestureRecognizer) SetAllowedTouchTypes(value TouchTypeMask) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setAllowedTouchTypes:"), value)
}
