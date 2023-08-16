// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AlignmentFeedbackFilter] class.
var AlignmentFeedbackFilterClass = _AlignmentFeedbackFilterClass{objc.GetClass("NSAlignmentFeedbackFilter")}

type _AlignmentFeedbackFilterClass struct {
	objc.Class
}

// An interface definition for the [AlignmentFeedbackFilter] class.
type IAlignmentFeedbackFilter interface {
	objc.IObject
	UpdateWithEvent(event IEvent)
	UpdateWithPanRecognizer(panRecognizer IPanGestureRecognizer)
	PerformFeedbackPerformanceTime(alignmentFeedbackTokens []PAlignmentFeedbackToken, performanceTime HapticFeedbackPerformanceTime)
	AlignmentFeedbackTokenForVerticalMovementInViewPreviousYAlignedYDefaultY(view IView, previousY float64, alignedY float64, defaultY float64) AlignmentFeedbackTokenWrapper
	AlignmentFeedbackTokenForMovementInViewPreviousPointAlignedPointDefaultPoint(view IView, previousPoint foundation.Point, alignedPoint foundation.Point, defaultPoint foundation.Point) AlignmentFeedbackTokenWrapper
	AlignmentFeedbackTokenForHorizontalMovementInViewPreviousXAlignedXDefaultX(view IView, previousX float64, alignedX float64, defaultX float64) AlignmentFeedbackTokenWrapper
}

// An object that can filter the movement of an object and provides haptic feedback when alignment occurs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalignmentfeedbackfilter?language=objc
type AlignmentFeedbackFilter struct {
	objc.Object
}

func AlignmentFeedbackFilterFrom(ptr unsafe.Pointer) AlignmentFeedbackFilter {
	return AlignmentFeedbackFilter{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AlignmentFeedbackFilterClass) Alloc() AlignmentFeedbackFilter {
	rv := objc.Call[AlignmentFeedbackFilter](ac, objc.Sel("alloc"))
	return rv
}

func AlignmentFeedbackFilter_Alloc() AlignmentFeedbackFilter {
	return AlignmentFeedbackFilterClass.Alloc()
}

func (ac _AlignmentFeedbackFilterClass) New() AlignmentFeedbackFilter {
	rv := objc.Call[AlignmentFeedbackFilter](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAlignmentFeedbackFilter() AlignmentFeedbackFilter {
	return AlignmentFeedbackFilterClass.New()
}

func (a_ AlignmentFeedbackFilter) Init() AlignmentFeedbackFilter {
	rv := objc.Call[AlignmentFeedbackFilter](a_, objc.Sel("init"))
	return rv
}

// Informs the feedback filter about a new event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalignmentfeedbackfilter/1532257-updatewithevent?language=objc
func (a_ AlignmentFeedbackFilter) UpdateWithEvent(event IEvent) {
	objc.Call[objc.Void](a_, objc.Sel("updateWithEvent:"), objc.Ptr(event))
}

// Informs the feedback filter about a new pan (drag) gesture recognizer event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalignmentfeedbackfilter/1527844-updatewithpanrecognizer?language=objc
func (a_ AlignmentFeedbackFilter) UpdateWithPanRecognizer(panRecognizer IPanGestureRecognizer) {
	objc.Call[objc.Void](a_, objc.Sel("updateWithPanRecognizer:"), objc.Ptr(panRecognizer))
}

// Performs the haptic feedback described by one or more alignment feedback tokens. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalignmentfeedbackfilter/1534423-performfeedback?language=objc
func (a_ AlignmentFeedbackFilter) PerformFeedbackPerformanceTime(alignmentFeedbackTokens []PAlignmentFeedbackToken, performanceTime HapticFeedbackPerformanceTime) {
	objc.Call[objc.Void](a_, objc.Sel("performFeedback:performanceTime:"), alignmentFeedbackTokens, performanceTime)
}

// Requests a feedback token for the alignment of an object requiring vertical movement only. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalignmentfeedbackfilter/1531563-alignmentfeedbacktokenforvertica?language=objc
func (a_ AlignmentFeedbackFilter) AlignmentFeedbackTokenForVerticalMovementInViewPreviousYAlignedYDefaultY(view IView, previousY float64, alignedY float64, defaultY float64) AlignmentFeedbackTokenWrapper {
	rv := objc.Call[AlignmentFeedbackTokenWrapper](a_, objc.Sel("alignmentFeedbackTokenForVerticalMovementInView:previousY:alignedY:defaultY:"), objc.Ptr(view), previousY, alignedY, defaultY)
	return rv
}

// Requests a feedback token for the alignment of an object requiring horizontal and vertical movement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalignmentfeedbackfilter/1535346-alignmentfeedbacktokenformovemen?language=objc
func (a_ AlignmentFeedbackFilter) AlignmentFeedbackTokenForMovementInViewPreviousPointAlignedPointDefaultPoint(view IView, previousPoint foundation.Point, alignedPoint foundation.Point, defaultPoint foundation.Point) AlignmentFeedbackTokenWrapper {
	rv := objc.Call[AlignmentFeedbackTokenWrapper](a_, objc.Sel("alignmentFeedbackTokenForMovementInView:previousPoint:alignedPoint:defaultPoint:"), objc.Ptr(view), previousPoint, alignedPoint, defaultPoint)
	return rv
}

// Requests a feedback token for the alignment of an object requiring horizontal movement only. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalignmentfeedbackfilter/1534596-alignmentfeedbacktokenforhorizon?language=objc
func (a_ AlignmentFeedbackFilter) AlignmentFeedbackTokenForHorizontalMovementInViewPreviousXAlignedXDefaultX(view IView, previousX float64, alignedX float64, defaultX float64) AlignmentFeedbackTokenWrapper {
	rv := objc.Call[AlignmentFeedbackTokenWrapper](a_, objc.Sel("alignmentFeedbackTokenForHorizontalMovementInView:previousX:alignedX:defaultX:"), objc.Ptr(view), previousX, alignedX, defaultX)
	return rv
}

// Retrieves the event types the filter accepts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalignmentfeedbackfilter/1530921-inputeventmask?language=objc
func (ac _AlignmentFeedbackFilterClass) InputEventMask() EventMask {
	rv := objc.Call[EventMask](ac, objc.Sel("inputEventMask"))
	return rv
}

// Retrieves the event types the filter accepts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalignmentfeedbackfilter/1530921-inputeventmask?language=objc
func AlignmentFeedbackFilter_InputEventMask() EventMask {
	return AlignmentFeedbackFilterClass.InputEventMask()
}
