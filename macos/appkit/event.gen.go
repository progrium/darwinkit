// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var EventClass = _EventClass{objc.GetClass("NSEvent")}

type _EventClass struct {
	objc.Class
}

type IEvent interface {
	objc.IObject
	CharactersByApplyingModifiers(modifiers EventModifierFlags) string
	TrackSwipeEventWithOptionsDampenAmountThresholdMinMaxUsingHandler(options EventSwipeTrackingOptions, minDampenThreshold float64, maxDampenThreshold float64, trackingHandler func(gestureAmount float64, phase EventPhase, isComplete bool, stop *bool))
	TouchesMatchingPhaseInView(phase TouchPhase, view IView) foundation.Set
	AllTouches() foundation.Set
	TouchesForView(view IView) foundation.Set
	CoalescedTouchesForTouch(touch ITouch) []Touch
	Type() EventType
	Subtype() EventSubtype
	LocationInWindow() foundation.Point
	Timestamp() foundation.TimeInterval
	Window() Window
	WindowNumber() int
	EventRef() unsafe.Pointer
	CGEvent() coregraphics.EventRef
	ModifierFlags() EventModifierFlags
	Characters() string
	CharactersIgnoringModifiers() string
	KeyCode() uint16
	IsARepeat() bool
	ButtonNumber() int
	ClickCount() int
	AssociatedEventsMask() EventMask
	DeltaX() float64
	DeltaY() float64
	DeltaZ() float64
	HasPreciseScrollingDeltas() bool
	ScrollingDeltaX() float64
	ScrollingDeltaY() float64
	MomentumPhase() EventPhase
	IsDirectionInvertedFromDevice() bool
	Phase() EventPhase
	Magnification() float64
	Pressure() float32
	Stage() int
	StageTransition() float64
	PressureBehavior() PressureBehavior
	CapabilityMask() uint
	DeviceID() uint
	IsEnteringProximity() bool
	PointingDeviceID() uint
	PointingDeviceSerialNumber() uint
	PointingDeviceType() PointingDeviceType
	SystemTabletID() uint
	TabletID() uint
	UniqueID() uint64
	VendorID() uint
	VendorPointingDeviceType() uint
	AbsoluteX() int
	AbsoluteY() int
	AbsoluteZ() int
	ButtonMask() EventButtonMask
	Rotation() float32
	TangentialPressure() float32
	Tilt() foundation.Point
	VendorDefined() objc.Object
	EventNumber() int
	TrackingNumber() int
	TrackingArea() TrackingArea
	UserData() unsafe.Pointer
	Data1() int
	Data2() int
}

type Event struct {
	objc.Object
}

func MakeEvent(ptr unsafe.Pointer) Event {
	return Event{
		Object: objc.MakeObject(ptr),
	}
}

func (ec _EventClass) Alloc() Event {
	rv := objc.CallMethod[Event](ec, objc.GetSelector("alloc"))
	return rv
}

func Event_Alloc() Event {
	return EventClass.Alloc()
}

func (ec _EventClass) New() Event {
	rv := objc.CallMethod[Event](ec, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewEvent() Event {
	return EventClass.New()
}

func Event_New() Event {
	return EventClass.New()
}

func (e_ Event) Init() Event {
	rv := objc.CallMethod[Event](e_, objc.GetSelector("init"))
	return rv
}

func Event_Init() Event {
	return EventClass.Alloc().Init()
}

func (ec _EventClass) KeyEventWithTypeLocationModifierFlagsTimestampWindowNumberContextCharactersCharactersIgnoringModifiersIsARepeatKeyCode(type_ EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil IGraphicsContext, keys string, ukeys string, flag bool, code uint16) Event {
	rv := objc.CallMethod[Event](ec, objc.GetSelector("keyEventWithType:location:modifierFlags:timestamp:windowNumber:context:characters:charactersIgnoringModifiers:isARepeat:keyCode:"), type_, location, flags, time, wNum, objc.ExtractPtr(unusedPassNil), keys, ukeys, flag, code)
	return rv
}

func Event_KeyEventWithTypeLocationModifierFlagsTimestampWindowNumberContextCharactersCharactersIgnoringModifiersIsARepeatKeyCode(type_ EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil IGraphicsContext, keys string, ukeys string, flag bool, code uint16) Event {
	return EventClass.KeyEventWithTypeLocationModifierFlagsTimestampWindowNumberContextCharactersCharactersIgnoringModifiersIsARepeatKeyCode(type_, location, flags, time, wNum, unusedPassNil, keys, ukeys, flag, code)
}

func (ec _EventClass) MouseEventWithTypeLocationModifierFlagsTimestampWindowNumberContextEventNumberClickCountPressure(type_ EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil IGraphicsContext, eNum int, cNum int, pressure float32) Event {
	rv := objc.CallMethod[Event](ec, objc.GetSelector("mouseEventWithType:location:modifierFlags:timestamp:windowNumber:context:eventNumber:clickCount:pressure:"), type_, location, flags, time, wNum, objc.ExtractPtr(unusedPassNil), eNum, cNum, pressure)
	return rv
}

func Event_MouseEventWithTypeLocationModifierFlagsTimestampWindowNumberContextEventNumberClickCountPressure(type_ EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil IGraphicsContext, eNum int, cNum int, pressure float32) Event {
	return EventClass.MouseEventWithTypeLocationModifierFlagsTimestampWindowNumberContextEventNumberClickCountPressure(type_, location, flags, time, wNum, unusedPassNil, eNum, cNum, pressure)
}

func (ec _EventClass) EnterExitEventWithTypeLocationModifierFlagsTimestampWindowNumberContextEventNumberTrackingNumberUserData(type_ EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil IGraphicsContext, eNum int, tNum int, data unsafe.Pointer) Event {
	rv := objc.CallMethod[Event](ec, objc.GetSelector("enterExitEventWithType:location:modifierFlags:timestamp:windowNumber:context:eventNumber:trackingNumber:userData:"), type_, location, flags, time, wNum, objc.ExtractPtr(unusedPassNil), eNum, tNum, data)
	return rv
}

func Event_EnterExitEventWithTypeLocationModifierFlagsTimestampWindowNumberContextEventNumberTrackingNumberUserData(type_ EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil IGraphicsContext, eNum int, tNum int, data unsafe.Pointer) Event {
	return EventClass.EnterExitEventWithTypeLocationModifierFlagsTimestampWindowNumberContextEventNumberTrackingNumberUserData(type_, location, flags, time, wNum, unusedPassNil, eNum, tNum, data)
}

func (ec _EventClass) OtherEventWithTypeLocationModifierFlagsTimestampWindowNumberContextSubtypeData1Data2(type_ EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil IGraphicsContext, subtype int16, d1 int, d2 int) Event {
	rv := objc.CallMethod[Event](ec, objc.GetSelector("otherEventWithType:location:modifierFlags:timestamp:windowNumber:context:subtype:data1:data2:"), type_, location, flags, time, wNum, objc.ExtractPtr(unusedPassNil), subtype, d1, d2)
	return rv
}

func Event_OtherEventWithTypeLocationModifierFlagsTimestampWindowNumberContextSubtypeData1Data2(type_ EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil IGraphicsContext, subtype int16, d1 int, d2 int) Event {
	return EventClass.OtherEventWithTypeLocationModifierFlagsTimestampWindowNumberContextSubtypeData1Data2(type_, location, flags, time, wNum, unusedPassNil, subtype, d1, d2)
}

func (ec _EventClass) EventWithEventRef(eventRef unsafe.Pointer) Event {
	rv := objc.CallMethod[Event](ec, objc.GetSelector("eventWithEventRef:"), eventRef)
	return rv
}

func Event_EventWithEventRef(eventRef unsafe.Pointer) Event {
	return EventClass.EventWithEventRef(eventRef)
}

func (ec _EventClass) EventWithCGEvent(cgEvent coregraphics.EventRef) Event {
	rv := objc.CallMethod[Event](ec, objc.GetSelector("eventWithCGEvent:"), cgEvent)
	return rv
}

func Event_EventWithCGEvent(cgEvent coregraphics.EventRef) Event {
	return EventClass.EventWithCGEvent(cgEvent)
}

func (e_ Event) CharactersByApplyingModifiers(modifiers EventModifierFlags) string {
	rv := objc.CallMethod[string](e_, objc.GetSelector("charactersByApplyingModifiers:"), modifiers)
	return rv
}

func (e_ Event) TrackSwipeEventWithOptionsDampenAmountThresholdMinMaxUsingHandler(options EventSwipeTrackingOptions, minDampenThreshold float64, maxDampenThreshold float64, trackingHandler func(gestureAmount float64, phase EventPhase, isComplete bool, stop *bool)) {
	objc.CallMethod[objc.Void](e_, objc.GetSelector("trackSwipeEventWithOptions:dampenAmountThresholdMin:max:usingHandler:"), options, minDampenThreshold, maxDampenThreshold, trackingHandler)
}

func (e_ Event) TouchesMatchingPhaseInView(phase TouchPhase, view IView) foundation.Set {
	rv := objc.CallMethod[foundation.Set](e_, objc.GetSelector("touchesMatchingPhase:inView:"), phase, objc.ExtractPtr(view))
	return rv
}

func (e_ Event) AllTouches() foundation.Set {
	rv := objc.CallMethod[foundation.Set](e_, objc.GetSelector("allTouches"))
	return rv
}

func (e_ Event) TouchesForView(view IView) foundation.Set {
	rv := objc.CallMethod[foundation.Set](e_, objc.GetSelector("touchesForView:"), objc.ExtractPtr(view))
	return rv
}

func (e_ Event) CoalescedTouchesForTouch(touch ITouch) []Touch {
	rv := objc.CallMethod[[]Touch](e_, objc.GetSelector("coalescedTouchesForTouch:"), objc.ExtractPtr(touch))
	return rv
}

func (ec _EventClass) StartPeriodicEventsAfterDelayWithPeriod(delay foundation.TimeInterval, period foundation.TimeInterval) {
	objc.CallMethod[objc.Void](ec, objc.GetSelector("startPeriodicEventsAfterDelay:withPeriod:"), delay, period)
}

func Event_StartPeriodicEventsAfterDelayWithPeriod(delay foundation.TimeInterval, period foundation.TimeInterval) {
	EventClass.StartPeriodicEventsAfterDelayWithPeriod(delay, period)
}

func (ec _EventClass) StopPeriodicEvents() {
	objc.CallMethod[objc.Void](ec, objc.GetSelector("stopPeriodicEvents"))
}

func Event_StopPeriodicEvents() {
	EventClass.StopPeriodicEvents()
}

func (ec _EventClass) AddGlobalMonitorForEventsMatchingMaskHandler(mask EventMask, block func(event Event)) objc.Object {
	rv := objc.CallMethod[objc.Object](ec, objc.GetSelector("addGlobalMonitorForEventsMatchingMask:handler:"), mask, block)
	return rv
}

func Event_AddGlobalMonitorForEventsMatchingMaskHandler(mask EventMask, block func(event Event)) objc.Object {
	return EventClass.AddGlobalMonitorForEventsMatchingMaskHandler(mask, block)
}

func (ec _EventClass) AddLocalMonitorForEventsMatchingMaskHandler(mask EventMask, block func(event Event) Event) objc.Object {
	rv := objc.CallMethod[objc.Object](ec, objc.GetSelector("addLocalMonitorForEventsMatchingMask:handler:"), mask, block)
	return rv
}

func Event_AddLocalMonitorForEventsMatchingMaskHandler(mask EventMask, block func(event Event) Event) objc.Object {
	return EventClass.AddLocalMonitorForEventsMatchingMaskHandler(mask, block)
}

func (ec _EventClass) RemoveMonitor(eventMonitor objc.IObject) {
	objc.CallMethod[objc.Void](ec, objc.GetSelector("removeMonitor:"), objc.ExtractPtr(eventMonitor))
}

func Event_RemoveMonitor(eventMonitor objc.IObject) {
	EventClass.RemoveMonitor(eventMonitor)
}

func (e_ Event) Type() EventType {
	rv := objc.CallMethod[EventType](e_, objc.GetSelector("type"))
	return rv
}

func (e_ Event) Subtype() EventSubtype {
	rv := objc.CallMethod[EventSubtype](e_, objc.GetSelector("subtype"))
	return rv
}

func (e_ Event) LocationInWindow() foundation.Point {
	rv := objc.CallMethod[foundation.Point](e_, objc.GetSelector("locationInWindow"))
	return rv
}

func (e_ Event) Timestamp() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](e_, objc.GetSelector("timestamp"))
	return rv
}

func (e_ Event) Window() Window {
	rv := objc.CallMethod[Window](e_, objc.GetSelector("window"))
	return rv
}

func (e_ Event) WindowNumber() int {
	rv := objc.CallMethod[int](e_, objc.GetSelector("windowNumber"))
	return rv
}

func (e_ Event) EventRef() unsafe.Pointer {
	rv := objc.CallMethod[unsafe.Pointer](e_, objc.GetSelector("eventRef"))
	return rv
}

func (e_ Event) CGEvent() coregraphics.EventRef {
	rv := objc.CallMethod[coregraphics.EventRef](e_, objc.GetSelector("CGEvent"))
	return rv
}

func (e_ Event) ModifierFlags() EventModifierFlags {
	rv := objc.CallMethod[EventModifierFlags](e_, objc.GetSelector("modifierFlags"))
	return rv
}

func (e_ Event) Characters() string {
	rv := objc.CallMethod[string](e_, objc.GetSelector("characters"))
	return rv
}

func (e_ Event) CharactersIgnoringModifiers() string {
	rv := objc.CallMethod[string](e_, objc.GetSelector("charactersIgnoringModifiers"))
	return rv
}

func (e_ Event) KeyCode() uint16 {
	rv := objc.CallMethod[uint16](e_, objc.GetSelector("keyCode"))
	return rv
}

func (ec _EventClass) KeyRepeatDelay() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](ec, objc.GetSelector("keyRepeatDelay"))
	return rv
}

func Event_KeyRepeatDelay() foundation.TimeInterval {
	return EventClass.KeyRepeatDelay()
}

func (ec _EventClass) KeyRepeatInterval() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](ec, objc.GetSelector("keyRepeatInterval"))
	return rv
}

func Event_KeyRepeatInterval() foundation.TimeInterval {
	return EventClass.KeyRepeatInterval()
}

func (e_ Event) IsARepeat() bool {
	rv := objc.CallMethod[bool](e_, objc.GetSelector("isARepeat"))
	return rv
}

func (ec _EventClass) PressedMouseButtons() uint {
	rv := objc.CallMethod[uint](ec, objc.GetSelector("pressedMouseButtons"))
	return rv
}

func Event_PressedMouseButtons() uint {
	return EventClass.PressedMouseButtons()
}

func (ec _EventClass) DoubleClickInterval() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](ec, objc.GetSelector("doubleClickInterval"))
	return rv
}

func Event_DoubleClickInterval() foundation.TimeInterval {
	return EventClass.DoubleClickInterval()
}

func (ec _EventClass) MouseLocation() foundation.Point {
	rv := objc.CallMethod[foundation.Point](ec, objc.GetSelector("mouseLocation"))
	return rv
}

func Event_MouseLocation() foundation.Point {
	return EventClass.MouseLocation()
}

func (e_ Event) ButtonNumber() int {
	rv := objc.CallMethod[int](e_, objc.GetSelector("buttonNumber"))
	return rv
}

func (e_ Event) ClickCount() int {
	rv := objc.CallMethod[int](e_, objc.GetSelector("clickCount"))
	return rv
}

func (e_ Event) AssociatedEventsMask() EventMask {
	rv := objc.CallMethod[EventMask](e_, objc.GetSelector("associatedEventsMask"))
	return rv
}

func (e_ Event) DeltaX() float64 {
	rv := objc.CallMethod[float64](e_, objc.GetSelector("deltaX"))
	return rv
}

func (e_ Event) DeltaY() float64 {
	rv := objc.CallMethod[float64](e_, objc.GetSelector("deltaY"))
	return rv
}

func (e_ Event) DeltaZ() float64 {
	rv := objc.CallMethod[float64](e_, objc.GetSelector("deltaZ"))
	return rv
}

func (e_ Event) HasPreciseScrollingDeltas() bool {
	rv := objc.CallMethod[bool](e_, objc.GetSelector("hasPreciseScrollingDeltas"))
	return rv
}

func (e_ Event) ScrollingDeltaX() float64 {
	rv := objc.CallMethod[float64](e_, objc.GetSelector("scrollingDeltaX"))
	return rv
}

func (e_ Event) ScrollingDeltaY() float64 {
	rv := objc.CallMethod[float64](e_, objc.GetSelector("scrollingDeltaY"))
	return rv
}

func (e_ Event) MomentumPhase() EventPhase {
	rv := objc.CallMethod[EventPhase](e_, objc.GetSelector("momentumPhase"))
	return rv
}

func (e_ Event) IsDirectionInvertedFromDevice() bool {
	rv := objc.CallMethod[bool](e_, objc.GetSelector("isDirectionInvertedFromDevice"))
	return rv
}

func (ec _EventClass) IsSwipeTrackingFromScrollEventsEnabled() bool {
	rv := objc.CallMethod[bool](ec, objc.GetSelector("isSwipeTrackingFromScrollEventsEnabled"))
	return rv
}

func Event_IsSwipeTrackingFromScrollEventsEnabled() bool {
	return EventClass.IsSwipeTrackingFromScrollEventsEnabled()
}

func (e_ Event) Phase() EventPhase {
	rv := objc.CallMethod[EventPhase](e_, objc.GetSelector("phase"))
	return rv
}

func (e_ Event) Magnification() float64 {
	rv := objc.CallMethod[float64](e_, objc.GetSelector("magnification"))
	return rv
}

func (ec _EventClass) IsMouseCoalescingEnabled() bool {
	rv := objc.CallMethod[bool](ec, objc.GetSelector("isMouseCoalescingEnabled"))
	return rv
}

func Event_IsMouseCoalescingEnabled() bool {
	return EventClass.IsMouseCoalescingEnabled()
}

func (ec _EventClass) SetMouseCoalescingEnabled(value bool) {
	objc.CallMethod[objc.Void](ec, objc.GetSelector("setMouseCoalescingEnabled:"), value)
}

func Event_SetMouseCoalescingEnabled(value bool) {
	EventClass.SetMouseCoalescingEnabled(value)
}

func (e_ Event) Pressure() float32 {
	rv := objc.CallMethod[float32](e_, objc.GetSelector("pressure"))
	return rv
}

func (e_ Event) Stage() int {
	rv := objc.CallMethod[int](e_, objc.GetSelector("stage"))
	return rv
}

func (e_ Event) StageTransition() float64 {
	rv := objc.CallMethod[float64](e_, objc.GetSelector("stageTransition"))
	return rv
}

func (e_ Event) PressureBehavior() PressureBehavior {
	rv := objc.CallMethod[PressureBehavior](e_, objc.GetSelector("pressureBehavior"))
	return rv
}

func (e_ Event) CapabilityMask() uint {
	rv := objc.CallMethod[uint](e_, objc.GetSelector("capabilityMask"))
	return rv
}

func (e_ Event) DeviceID() uint {
	rv := objc.CallMethod[uint](e_, objc.GetSelector("deviceID"))
	return rv
}

func (e_ Event) IsEnteringProximity() bool {
	rv := objc.CallMethod[bool](e_, objc.GetSelector("isEnteringProximity"))
	return rv
}

func (e_ Event) PointingDeviceID() uint {
	rv := objc.CallMethod[uint](e_, objc.GetSelector("pointingDeviceID"))
	return rv
}

func (e_ Event) PointingDeviceSerialNumber() uint {
	rv := objc.CallMethod[uint](e_, objc.GetSelector("pointingDeviceSerialNumber"))
	return rv
}

func (e_ Event) PointingDeviceType() PointingDeviceType {
	rv := objc.CallMethod[PointingDeviceType](e_, objc.GetSelector("pointingDeviceType"))
	return rv
}

func (e_ Event) SystemTabletID() uint {
	rv := objc.CallMethod[uint](e_, objc.GetSelector("systemTabletID"))
	return rv
}

func (e_ Event) TabletID() uint {
	rv := objc.CallMethod[uint](e_, objc.GetSelector("tabletID"))
	return rv
}

func (e_ Event) UniqueID() uint64 {
	rv := objc.CallMethod[uint64](e_, objc.GetSelector("uniqueID"))
	return rv
}

func (e_ Event) VendorID() uint {
	rv := objc.CallMethod[uint](e_, objc.GetSelector("vendorID"))
	return rv
}

func (e_ Event) VendorPointingDeviceType() uint {
	rv := objc.CallMethod[uint](e_, objc.GetSelector("vendorPointingDeviceType"))
	return rv
}

func (e_ Event) AbsoluteX() int {
	rv := objc.CallMethod[int](e_, objc.GetSelector("absoluteX"))
	return rv
}

func (e_ Event) AbsoluteY() int {
	rv := objc.CallMethod[int](e_, objc.GetSelector("absoluteY"))
	return rv
}

func (e_ Event) AbsoluteZ() int {
	rv := objc.CallMethod[int](e_, objc.GetSelector("absoluteZ"))
	return rv
}

func (e_ Event) ButtonMask() EventButtonMask {
	rv := objc.CallMethod[EventButtonMask](e_, objc.GetSelector("buttonMask"))
	return rv
}

func (e_ Event) Rotation() float32 {
	rv := objc.CallMethod[float32](e_, objc.GetSelector("rotation"))
	return rv
}

func (e_ Event) TangentialPressure() float32 {
	rv := objc.CallMethod[float32](e_, objc.GetSelector("tangentialPressure"))
	return rv
}

func (e_ Event) Tilt() foundation.Point {
	rv := objc.CallMethod[foundation.Point](e_, objc.GetSelector("tilt"))
	return rv
}

func (e_ Event) VendorDefined() objc.Object {
	rv := objc.CallMethod[objc.Object](e_, objc.GetSelector("vendorDefined"))
	return rv
}

func (e_ Event) EventNumber() int {
	rv := objc.CallMethod[int](e_, objc.GetSelector("eventNumber"))
	return rv
}

func (e_ Event) TrackingNumber() int {
	rv := objc.CallMethod[int](e_, objc.GetSelector("trackingNumber"))
	return rv
}

func (e_ Event) TrackingArea() TrackingArea {
	rv := objc.CallMethod[TrackingArea](e_, objc.GetSelector("trackingArea"))
	return rv
}

func (e_ Event) UserData() unsafe.Pointer {
	rv := objc.CallMethod[unsafe.Pointer](e_, objc.GetSelector("userData"))
	return rv
}

func (e_ Event) Data1() int {
	rv := objc.CallMethod[int](e_, objc.GetSelector("data1"))
	return rv
}

func (e_ Event) Data2() int {
	rv := objc.CallMethod[int](e_, objc.GetSelector("data2"))
	return rv
}
