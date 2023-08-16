// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [GestureRecognizer] class.
var GestureRecognizerClass = _GestureRecognizerClass{objc.GetClass("NSGestureRecognizer")}

type _GestureRecognizerClass struct {
	objc.Class
}

// An interface definition for the [GestureRecognizer] class.
type IGestureRecognizer interface {
	objc.IObject
	TouchesMovedWithEvent(event IEvent)
	MouseUp(event IEvent)
	CanBePreventedByGestureRecognizer(preventingGestureRecognizer IGestureRecognizer) bool
	PressureChangeWithEvent(event IEvent)
	ShouldBeRequiredToFailByGestureRecognizer(otherGestureRecognizer IGestureRecognizer) bool
	OtherMouseDragged(event IEvent)
	RightMouseDragged(event IEvent)
	CanPreventGestureRecognizer(preventedGestureRecognizer IGestureRecognizer) bool
	ShouldRequireFailureOfGestureRecognizer(otherGestureRecognizer IGestureRecognizer) bool
	FlagsChanged(event IEvent)
	TouchesBeganWithEvent(event IEvent)
	TouchesCancelledWithEvent(event IEvent)
	RotateWithEvent(event IEvent)
	KeyDown(event IEvent)
	TabletPoint(event IEvent)
	TouchesEndedWithEvent(event IEvent)
	OtherMouseDown(event IEvent)
	OtherMouseUp(event IEvent)
	MagnifyWithEvent(event IEvent)
	MouseDragged(event IEvent)
	LocationInView(view IView) foundation.Point
	MouseDown(event IEvent)
	KeyUp(event IEvent)
	RightMouseDown(event IEvent)
	Reset()
	RightMouseUp(event IEvent)
	DelaysOtherMouseButtonEvents() bool
	SetDelaysOtherMouseButtonEvents(value bool)
	DelaysSecondaryMouseButtonEvents() bool
	SetDelaysSecondaryMouseButtonEvents(value bool)
	Target() objc.Object
	SetTarget(value objc.IObject)
	State() GestureRecognizerState
	Action() objc.Selector
	SetAction(value objc.Selector)
	PressureConfiguration() PressureConfiguration
	SetPressureConfiguration(value IPressureConfiguration)
	View() View
	DelaysMagnificationEvents() bool
	SetDelaysMagnificationEvents(value bool)
	Delegate() GestureRecognizerDelegateWrapper
	SetDelegate(value PGestureRecognizerDelegate)
	SetDelegateObject(valueObject objc.IObject)
	DelaysPrimaryMouseButtonEvents() bool
	SetDelaysPrimaryMouseButtonEvents(value bool)
	DelaysKeyEvents() bool
	SetDelaysKeyEvents(value bool)
	DelaysRotationEvents() bool
	SetDelaysRotationEvents(value bool)
	AllowedTouchTypes() TouchTypeMask
	SetAllowedTouchTypes(value TouchTypeMask)
	IsEnabled() bool
	SetEnabled(value bool)
}

// An object that monitors events and calls its action method when a predefined sequence of events occur. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer?language=objc
type GestureRecognizer struct {
	objc.Object
}

func GestureRecognizerFrom(ptr unsafe.Pointer) GestureRecognizer {
	return GestureRecognizer{
		Object: objc.ObjectFrom(ptr),
	}
}

func (g_ GestureRecognizer) InitWithTargetAction(target objc.IObject, action objc.Selector) GestureRecognizer {
	rv := objc.Call[GestureRecognizer](g_, objc.Sel("initWithTarget:action:"), target, action)
	return rv
}

// Initializes the gesture recognizer with the specified target and action information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1535012-initwithtarget?language=objc
func GestureRecognizer_InitWithTargetAction(target objc.IObject, action objc.Selector) GestureRecognizer {
	return GestureRecognizerClass.Alloc().InitWithTargetAction(target, action)
}

func (gc _GestureRecognizerClass) Alloc() GestureRecognizer {
	rv := objc.Call[GestureRecognizer](gc, objc.Sel("alloc"))
	return rv
}

func GestureRecognizer_Alloc() GestureRecognizer {
	return GestureRecognizerClass.Alloc()
}

func (gc _GestureRecognizerClass) New() GestureRecognizer {
	rv := objc.Call[GestureRecognizer](gc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewGestureRecognizer() GestureRecognizer {
	return GestureRecognizerClass.New()
}

func (g_ GestureRecognizer) Init() GestureRecognizer {
	rv := objc.Call[GestureRecognizer](g_, objc.Sel("init"))
	return rv
}

// Called when one or more fingers, associated with an in-progress event, move within an NSTouchBar instance on the Touch Bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/2544863-touchesmovedwithevent?language=objc
func (g_ GestureRecognizer) TouchesMovedWithEvent(event IEvent) {
	objc.Call[objc.Void](g_, objc.Sel("touchesMovedWithEvent:"), objc.Ptr(event))
}

// Informs the gesture recognizer that the user released the left mouse button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1526116-mouseup?language=objc
func (g_ GestureRecognizer) MouseUp(event IEvent) {
	objc.Call[objc.Void](g_, objc.Sel("mouseUp:"), objc.Ptr(event))
}

// Overridden to indicate that the specified gesture recognizer can prevent the current object from recognizing a gesture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1534587-canbepreventedbygesturerecognize?language=objc
func (g_ GestureRecognizer) CanBePreventedByGestureRecognizer(preventingGestureRecognizer IGestureRecognizer) bool {
	rv := objc.Call[bool](g_, objc.Sel("canBePreventedByGestureRecognizer:"), objc.Ptr(preventingGestureRecognizer))
	return rv
}

// Informs the current object that a pressure change occurred on a system that supports pressure sensitivity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1527009-pressurechangewithevent?language=objc
func (g_ GestureRecognizer) PressureChangeWithEvent(event IEvent) {
	objc.Call[objc.Void](g_, objc.Sel("pressureChangeWithEvent:"), objc.Ptr(event))
}

// Overridden to indicate that the current object must fail before the specified gesture recognizer begins recognizing its gesture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1530755-shouldberequiredtofailbygesturer?language=objc
func (g_ GestureRecognizer) ShouldBeRequiredToFailByGestureRecognizer(otherGestureRecognizer IGestureRecognizer) bool {
	rv := objc.Call[bool](g_, objc.Sel("shouldBeRequiredToFailByGestureRecognizer:"), objc.Ptr(otherGestureRecognizer))
	return rv
}

// Informs the gesture recognizer that the user moved the mouse with a button other than the left or right one pressed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1534208-othermousedragged?language=objc
func (g_ GestureRecognizer) OtherMouseDragged(event IEvent) {
	objc.Call[objc.Void](g_, objc.Sel("otherMouseDragged:"), objc.Ptr(event))
}

// Informs the gesture recognizer that the user moved the mouse with the right button pressed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1529778-rightmousedragged?language=objc
func (g_ GestureRecognizer) RightMouseDragged(event IEvent) {
	objc.Call[objc.Void](g_, objc.Sel("rightMouseDragged:"), objc.Ptr(event))
}

// Overridden to indicate that the current object can prevent the specified gesture recognizer from recognizing its gesture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1534503-canpreventgesturerecognizer?language=objc
func (g_ GestureRecognizer) CanPreventGestureRecognizer(preventedGestureRecognizer IGestureRecognizer) bool {
	rv := objc.Call[bool](g_, objc.Sel("canPreventGestureRecognizer:"), objc.Ptr(preventedGestureRecognizer))
	return rv
}

// Overridden to indicate that the specified gesture recognizer must fail before the current object begins recognizing its gesture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1532939-shouldrequirefailureofgesturerec?language=objc
func (g_ GestureRecognizer) ShouldRequireFailureOfGestureRecognizer(otherGestureRecognizer IGestureRecognizer) bool {
	rv := objc.Call[bool](g_, objc.Sel("shouldRequireFailureOfGestureRecognizer:"), objc.Ptr(otherGestureRecognizer))
	return rv
}

// Informs the current object that the user pressed or released a modifier key (Shift, Control, and so on). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1532604-flagschanged?language=objc
func (g_ GestureRecognizer) FlagsChanged(event IEvent) {
	objc.Call[objc.Void](g_, objc.Sel("flagsChanged:"), objc.Ptr(event))
}

// Called when one or more fingers first make contact with an NSTouchBar instance on the Touch Bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/2544794-touchesbeganwithevent?language=objc
func (g_ GestureRecognizer) TouchesBeganWithEvent(event IEvent) {
	objc.Call[objc.Void](g_, objc.Sel("touchesBeganWithEvent:"), objc.Ptr(event))
}

// Called when a system event, such as a low-memory warning, cancels an in-progress touch event in an NSTouchBar object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/2544757-touchescancelledwithevent?language=objc
func (g_ GestureRecognizer) TouchesCancelledWithEvent(event IEvent) {
	objc.Call[objc.Void](g_, objc.Sel("touchesCancelledWithEvent:"), objc.Ptr(event))
}

// Informs the gesture recognizer that the user is performing a rotation gesture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1531401-rotatewithevent?language=objc
func (g_ GestureRecognizer) RotateWithEvent(event IEvent) {
	objc.Call[objc.Void](g_, objc.Sel("rotateWithEvent:"), objc.Ptr(event))
}

// Informs the gesture recognizer that the user has pressed a key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1532158-keydown?language=objc
func (g_ GestureRecognizer) KeyDown(event IEvent) {
	objc.Call[objc.Void](g_, objc.Sel("keyDown:"), objc.Ptr(event))
}

// Informs the user that a tablet-point event occurred. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1529889-tabletpoint?language=objc
func (g_ GestureRecognizer) TabletPoint(event IEvent) {
	objc.Call[objc.Void](g_, objc.Sel("tabletPoint:"), objc.Ptr(event))
}

// Called when one or more fingers are removed from contact with an NSTouchBar instance on the Touch Bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/2544740-touchesendedwithevent?language=objc
func (g_ GestureRecognizer) TouchesEndedWithEvent(event IEvent) {
	objc.Call[objc.Void](g_, objc.Sel("touchesEndedWithEvent:"), objc.Ptr(event))
}

// Informs the gesture recognizer that the user pressed a mouse button other than the left or right one. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1534433-othermousedown?language=objc
func (g_ GestureRecognizer) OtherMouseDown(event IEvent) {
	objc.Call[objc.Void](g_, objc.Sel("otherMouseDown:"), objc.Ptr(event))
}

// Informs the gesture recognizer that the user released a mouse button other than the left or right one. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1527874-othermouseup?language=objc
func (g_ GestureRecognizer) OtherMouseUp(event IEvent) {
	objc.Call[objc.Void](g_, objc.Sel("otherMouseUp:"), objc.Ptr(event))
}

// Informs the gesture recognizer that the user is performing a pinch gesture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1528828-magnifywithevent?language=objc
func (g_ GestureRecognizer) MagnifyWithEvent(event IEvent) {
	objc.Call[objc.Void](g_, objc.Sel("magnifyWithEvent:"), objc.Ptr(event))
}

// Informs the gesture recognizer that the user moved the mouse with the left button pressed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1534535-mousedragged?language=objc
func (g_ GestureRecognizer) MouseDragged(event IEvent) {
	objc.Call[objc.Void](g_, objc.Sel("mouseDragged:"), objc.Ptr(event))
}

// Returns the point computed as the location of the gesture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1524261-locationinview?language=objc
func (g_ GestureRecognizer) LocationInView(view IView) foundation.Point {
	rv := objc.Call[foundation.Point](g_, objc.Sel("locationInView:"), objc.Ptr(view))
	return rv
}

// Informs the gesture recognizer that the user pressed the left mouse button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1524901-mousedown?language=objc
func (g_ GestureRecognizer) MouseDown(event IEvent) {
	objc.Call[objc.Void](g_, objc.Sel("mouseDown:"), objc.Ptr(event))
}

// Informs the gesture recognizer that the user released a key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1526578-keyup?language=objc
func (g_ GestureRecognizer) KeyUp(event IEvent) {
	objc.Call[objc.Void](g_, objc.Sel("keyUp:"), objc.Ptr(event))
}

// Informs the gesture recognizer that the user pressed the right mouse button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1527421-rightmousedown?language=objc
func (g_ GestureRecognizer) RightMouseDown(event IEvent) {
	objc.Call[objc.Void](g_, objc.Sel("rightMouseDown:"), objc.Ptr(event))
}

// Overridden to reset the internal state of the gesture recognizer when an attempt completes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1527219-reset?language=objc
func (g_ GestureRecognizer) Reset() {
	objc.Call[objc.Void](g_, objc.Sel("reset"))
}

// Informs the gesture recognizer that the user released the right mouse button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1525729-rightmouseup?language=objc
func (g_ GestureRecognizer) RightMouseUp(event IEvent) {
	objc.Call[objc.Void](g_, objc.Sel("rightMouseUp:"), objc.Ptr(event))
}

// A Boolean value that indicates whether other mouse button events are delivered only after gesture recognition fails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1530375-delaysothermousebuttonevents?language=objc
func (g_ GestureRecognizer) DelaysOtherMouseButtonEvents() bool {
	rv := objc.Call[bool](g_, objc.Sel("delaysOtherMouseButtonEvents"))
	return rv
}

// A Boolean value that indicates whether other mouse button events are delivered only after gesture recognition fails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1530375-delaysothermousebuttonevents?language=objc
func (g_ GestureRecognizer) SetDelaysOtherMouseButtonEvents(value bool) {
	objc.Call[objc.Void](g_, objc.Sel("setDelaysOtherMouseButtonEvents:"), value)
}

// A Boolean value that indicates whether secondary mouse button events are delivered only after gesture recognition fails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1535204-delayssecondarymousebuttonevents?language=objc
func (g_ GestureRecognizer) DelaysSecondaryMouseButtonEvents() bool {
	rv := objc.Call[bool](g_, objc.Sel("delaysSecondaryMouseButtonEvents"))
	return rv
}

// A Boolean value that indicates whether secondary mouse button events are delivered only after gesture recognition fails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1535204-delayssecondarymousebuttonevents?language=objc
func (g_ GestureRecognizer) SetDelaysSecondaryMouseButtonEvents(value bool) {
	objc.Call[objc.Void](g_, objc.Sel("setDelaysSecondaryMouseButtonEvents:"), value)
}

// The object that implements the action method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1530932-target?language=objc
func (g_ GestureRecognizer) Target() objc.Object {
	rv := objc.Call[objc.Object](g_, objc.Sel("target"))
	return rv
}

// The object that implements the action method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1530932-target?language=objc
func (g_ GestureRecognizer) SetTarget(value objc.IObject) {
	objc.Call[objc.Void](g_, objc.Sel("setTarget:"), value)
}

// The current state of the gesture recognizer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1535694-state?language=objc
func (g_ GestureRecognizer) State() GestureRecognizerState {
	rv := objc.Call[GestureRecognizerState](g_, objc.Sel("state"))
	return rv
}

// The action method to call when the gesture is recognized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1529280-action?language=objc
func (g_ GestureRecognizer) Action() objc.Selector {
	rv := objc.Call[objc.Selector](g_, objc.Sel("action"))
	return rv
}

// The action method to call when the gesture is recognized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1529280-action?language=objc
func (g_ GestureRecognizer) SetAction(value objc.Selector) {
	objc.Call[objc.Void](g_, objc.Sel("setAction:"), value)
}

// Configures the behavior and progression of the Force Touch trackpad when responding to recognized pressure gestures. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1535895-pressureconfiguration?language=objc
func (g_ GestureRecognizer) PressureConfiguration() PressureConfiguration {
	rv := objc.Call[PressureConfiguration](g_, objc.Sel("pressureConfiguration"))
	return rv
}

// Configures the behavior and progression of the Force Touch trackpad when responding to recognized pressure gestures. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1535895-pressureconfiguration?language=objc
func (g_ GestureRecognizer) SetPressureConfiguration(value IPressureConfiguration) {
	objc.Call[objc.Void](g_, objc.Sel("setPressureConfiguration:"), objc.Ptr(value))
}

// The view to which the gesture recognizer is attached. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1527192-view?language=objc
func (g_ GestureRecognizer) View() View {
	rv := objc.Call[View](g_, objc.Sel("view"))
	return rv
}

// A Boolean value that indicates whether magnification events are delivered only after gesture recognition fails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1529974-delaysmagnificationevents?language=objc
func (g_ GestureRecognizer) DelaysMagnificationEvents() bool {
	rv := objc.Call[bool](g_, objc.Sel("delaysMagnificationEvents"))
	return rv
}

// A Boolean value that indicates whether magnification events are delivered only after gesture recognition fails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1529974-delaysmagnificationevents?language=objc
func (g_ GestureRecognizer) SetDelaysMagnificationEvents(value bool) {
	objc.Call[objc.Void](g_, objc.Sel("setDelaysMagnificationEvents:"), value)
}

// The delegate of the gesture recognizer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1529879-delegate?language=objc
func (g_ GestureRecognizer) Delegate() GestureRecognizerDelegateWrapper {
	rv := objc.Call[GestureRecognizerDelegateWrapper](g_, objc.Sel("delegate"))
	return rv
}

// The delegate of the gesture recognizer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1529879-delegate?language=objc
func (g_ GestureRecognizer) SetDelegate(value PGestureRecognizerDelegate) {
	po0 := objc.WrapAsProtocol("NSGestureRecognizerDelegate", value)
	objc.SetAssociatedObject(g_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](g_, objc.Sel("setDelegate:"), po0)
}

// The delegate of the gesture recognizer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1529879-delegate?language=objc
func (g_ GestureRecognizer) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](g_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// A Boolean value that indicates whether primary mouse button events are delivered only after gesture recognition fails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1533619-delaysprimarymousebuttonevents?language=objc
func (g_ GestureRecognizer) DelaysPrimaryMouseButtonEvents() bool {
	rv := objc.Call[bool](g_, objc.Sel("delaysPrimaryMouseButtonEvents"))
	return rv
}

// A Boolean value that indicates whether primary mouse button events are delivered only after gesture recognition fails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1533619-delaysprimarymousebuttonevents?language=objc
func (g_ GestureRecognizer) SetDelaysPrimaryMouseButtonEvents(value bool) {
	objc.Call[objc.Void](g_, objc.Sel("setDelaysPrimaryMouseButtonEvents:"), value)
}

// A Boolean value that indicates whether key events are delivered only after gesture recognition fails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1529788-delayskeyevents?language=objc
func (g_ GestureRecognizer) DelaysKeyEvents() bool {
	rv := objc.Call[bool](g_, objc.Sel("delaysKeyEvents"))
	return rv
}

// A Boolean value that indicates whether key events are delivered only after gesture recognition fails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1529788-delayskeyevents?language=objc
func (g_ GestureRecognizer) SetDelaysKeyEvents(value bool) {
	objc.Call[objc.Void](g_, objc.Sel("setDelaysKeyEvents:"), value)
}

// A Boolean value that indicates whether rotation events are delivered only after gesture recognition fails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1535384-delaysrotationevents?language=objc
func (g_ GestureRecognizer) DelaysRotationEvents() bool {
	rv := objc.Call[bool](g_, objc.Sel("delaysRotationEvents"))
	return rv
}

// A Boolean value that indicates whether rotation events are delivered only after gesture recognition fails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1535384-delaysrotationevents?language=objc
func (g_ GestureRecognizer) SetDelaysRotationEvents(value bool) {
	objc.Call[objc.Void](g_, objc.Sel("setDelaysRotationEvents:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/2588249-allowedtouchtypes?language=objc
func (g_ GestureRecognizer) AllowedTouchTypes() TouchTypeMask {
	rv := objc.Call[TouchTypeMask](g_, objc.Sel("allowedTouchTypes"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/2588249-allowedtouchtypes?language=objc
func (g_ GestureRecognizer) SetAllowedTouchTypes(value TouchTypeMask) {
	objc.Call[objc.Void](g_, objc.Sel("setAllowedTouchTypes:"), value)
}

// A Boolean value indicating whether the gesture recognizer is able to handle events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1535263-enabled?language=objc
func (g_ GestureRecognizer) IsEnabled() bool {
	rv := objc.Call[bool](g_, objc.Sel("isEnabled"))
	return rv
}

// A Boolean value indicating whether the gesture recognizer is able to handle events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1535263-enabled?language=objc
func (g_ GestureRecognizer) SetEnabled(value bool) {
	objc.Call[objc.Void](g_, objc.Sel("setEnabled:"), value)
}
