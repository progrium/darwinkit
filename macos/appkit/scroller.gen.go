// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Scroller] class.
var ScrollerClass = _ScrollerClass{objc.GetClass("NSScroller")}

type _ScrollerClass struct {
	objc.Class
}

// An interface definition for the [Scroller] class.
type IScroller interface {
	IControl
	CheckSpaceForParts()
	TrackKnob(event IEvent)
	SetKnobProportion(proportion float64)
	TestPart(point foundation.Point) ScrollerPart
	DrawKnob()
	DrawKnobSlotInRectHighlight(slotRect foundation.Rect, flag bool)
	RectForPart(partCode ScrollerPart) foundation.Rect
	ScrollerStyle() ScrollerStyle
	SetScrollerStyle(value ScrollerStyle)
	KnobStyle() ScrollerKnobStyle
	SetKnobStyle(value ScrollerKnobStyle)
	UsableParts() UsableScrollerParts
	HitPart() ScrollerPart
	KnobProportion() float64
}

// An object that controls scrolling of a document view within a scroll view or other type of container view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller?language=objc
type Scroller struct {
	Control
}

func ScrollerFrom(ptr unsafe.Pointer) Scroller {
	return Scroller{
		Control: ControlFrom(ptr),
	}
}

func (sc _ScrollerClass) Alloc() Scroller {
	rv := objc.Call[Scroller](sc, objc.Sel("alloc"))
	return rv
}

func Scroller_Alloc() Scroller {
	return ScrollerClass.Alloc()
}

func (sc _ScrollerClass) New() Scroller {
	rv := objc.Call[Scroller](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScroller() Scroller {
	return ScrollerClass.New()
}

func (s_ Scroller) Init() Scroller {
	rv := objc.Call[Scroller](s_, objc.Sel("init"))
	return rv
}

func (s_ Scroller) InitWithFrame(frameRect foundation.Rect) Scroller {
	rv := objc.Call[Scroller](s_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes a control with the specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428900-initwithframe?language=objc
func Scroller_InitWithFrame(frameRect foundation.Rect) Scroller {
	return ScrollerClass.Alloc().InitWithFrame(frameRect)
}

// Checks to see if there is enough room in the receiver to display the knob and buttons. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/1523625-checkspaceforparts?language=objc
func (s_ Scroller) CheckSpaceForParts() {
	objc.Call[objc.Void](s_, objc.Sel("checkSpaceForParts"))
}

// Tracks the knob and sends action messages to the receiver’s target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/1523594-trackknob?language=objc
func (s_ Scroller) TrackKnob(event IEvent) {
	objc.Call[objc.Void](s_, objc.Sel("trackKnob:"), objc.Ptr(event))
}

// The proportion of the knob slot that the knob should fill. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/1523626-setknobproportion?language=objc
func (s_ Scroller) SetKnobProportion(proportion float64) {
	objc.Call[objc.Void](s_, objc.Sel("setKnobProportion:"), proportion)
}

// Returns the part that would be hit by a mouse-down event at aPoint (expressed in the window’s coordinate system). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/1523645-testpart?language=objc
func (s_ Scroller) TestPart(point foundation.Point) ScrollerPart {
	rv := objc.Call[ScrollerPart](s_, objc.Sel("testPart:"), point)
	return rv
}

// Draws the knob. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/1523663-drawknob?language=objc
func (s_ Scroller) DrawKnob() {
	objc.Call[objc.Void](s_, objc.Sel("drawKnob"))
}

// Draws the portion of the scroller’s track, possibly including the line increment and decrement arrow buttons, that falls in the given rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/1523622-drawknobslotinrect?language=objc
func (s_ Scroller) DrawKnobSlotInRectHighlight(slotRect foundation.Rect, flag bool) {
	objc.Call[objc.Void](s_, objc.Sel("drawKnobSlotInRect:highlight:"), slotRect, flag)
}

// Returns the rectangle occupied by aPart, which for this method is interpreted literally rather than as an indicator of scrolling direction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/1523647-rectforpart?language=objc
func (s_ Scroller) RectForPart(partCode ScrollerPart) foundation.Rect {
	rv := objc.Call[foundation.Rect](s_, objc.Sel("rectForPart:"), partCode)
	return rv
}

// The scroller style for this scroller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/1523591-scrollerstyle?language=objc
func (s_ Scroller) ScrollerStyle() ScrollerStyle {
	rv := objc.Call[ScrollerStyle](s_, objc.Sel("scrollerStyle"))
	return rv
}

// The scroller style for this scroller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/1523591-scrollerstyle?language=objc
func (s_ Scroller) SetScrollerStyle(value ScrollerStyle) {
	objc.Call[objc.Void](s_, objc.Sel("setScrollerStyle:"), value)
}

// The scroller’s knob style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/1523666-knobstyle?language=objc
func (s_ Scroller) KnobStyle() ScrollerKnobStyle {
	rv := objc.Call[ScrollerKnobStyle](s_, objc.Sel("knobStyle"))
	return rv
}

// The scroller’s knob style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/1523666-knobstyle?language=objc
func (s_ Scroller) SetKnobStyle(value ScrollerKnobStyle) {
	objc.Call[objc.Void](s_, objc.Sel("setKnobStyle:"), value)
}

// Returns the style of scrollers that applications should use wherever possible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/1523620-preferredscrollerstyle?language=objc
func (sc _ScrollerClass) PreferredScrollerStyle() ScrollerStyle {
	rv := objc.Call[ScrollerStyle](sc, objc.Sel("preferredScrollerStyle"))
	return rv
}

// Returns the style of scrollers that applications should use wherever possible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/1523620-preferredscrollerstyle?language=objc
func Scroller_PreferredScrollerStyle() ScrollerStyle {
	return ScrollerClass.PreferredScrollerStyle()
}

// A value that indicates which parts of the receiver are displayed and usable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/1523648-usableparts?language=objc
func (s_ Scroller) UsableParts() UsableScrollerParts {
	rv := objc.Call[UsableScrollerParts](s_, objc.Sel("usableParts"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/2870071-compatiblewithoverlayscrollers?language=objc
func (sc _ScrollerClass) CompatibleWithOverlayScrollers() bool {
	rv := objc.Call[bool](sc, objc.Sel("compatibleWithOverlayScrollers"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/2870071-compatiblewithoverlayscrollers?language=objc
func Scroller_CompatibleWithOverlayScrollers() bool {
	return ScrollerClass.CompatibleWithOverlayScrollers()
}

// A part code indicating the manner in which the scrolling should be performed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/1523596-hitpart?language=objc
func (s_ Scroller) HitPart() ScrollerPart {
	rv := objc.Call[ScrollerPart](s_, objc.Sel("hitPart"))
	return rv
}

// The proportion of the knob slot that the knob should fill. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/1523593-knobproportion?language=objc
func (s_ Scroller) KnobProportion() float64 {
	rv := objc.Call[float64](s_, objc.Sel("knobProportion"))
	return rv
}
