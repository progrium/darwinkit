// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RulerView] class.
var RulerViewClass = _RulerViewClass{objc.GetClass("NSRulerView")}

type _RulerViewClass struct {
	objc.Class
}

// An interface definition for the [RulerView] class.
type IRulerView interface {
	IView
	InvalidateHashMarks()
	TrackMarkerWithMouseEvent(marker IRulerMarker, event IEvent) bool
	MoveRulerlineFromLocationToLocation(oldLocation float64, newLocation float64)
	DrawHashMarksAndLabelsInRect(rect foundation.Rect)
	DrawMarkersInRect(rect foundation.Rect)
	RemoveMarker(marker IRulerMarker)
	AddMarker(marker IRulerMarker)
	BaselineLocation() float64
	ReservedThicknessForMarkers() float64
	SetReservedThicknessForMarkers(value float64)
	RuleThickness() float64
	SetRuleThickness(value float64)
	AccessoryView() View
	SetAccessoryView(value IView)
	OriginOffset() float64
	SetOriginOffset(value float64)
	RequiredThickness() float64
	ClientView() View
	SetClientView(value IView)
	ScrollView() ScrollView
	SetScrollView(value IScrollView)
	MeasurementUnits() RulerViewUnitName
	SetMeasurementUnits(value RulerViewUnitName)
	ReservedThicknessForAccessoryView() float64
	SetReservedThicknessForAccessoryView(value float64)
	Markers() []RulerMarker
	SetMarkers(value []IRulerMarker)
	Orientation() RulerOrientation
	SetOrientation(value RulerOrientation)
}

// A ruler and the markers above or to the side of a scroll view’s document view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview?language=objc
type RulerView struct {
	View
}

func RulerViewFrom(ptr unsafe.Pointer) RulerView {
	return RulerView{
		View: ViewFrom(ptr),
	}
}

func (r_ RulerView) InitWithScrollViewOrientation(scrollView IScrollView, orientation RulerOrientation) RulerView {
	rv := objc.Call[RulerView](r_, objc.Sel("initWithScrollView:orientation:"), objc.Ptr(scrollView), orientation)
	return rv
}

// Initializes a newly allocated NSRulerView to have orientation (NSHorizontalRuler or NSVerticalRuler) within aScrollView. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1535316-initwithscrollview?language=objc
func RulerView_InitWithScrollViewOrientation(scrollView IScrollView, orientation RulerOrientation) RulerView {
	return RulerViewClass.Alloc().InitWithScrollViewOrientation(scrollView, orientation)
}

func (rc _RulerViewClass) Alloc() RulerView {
	rv := objc.Call[RulerView](rc, objc.Sel("alloc"))
	return rv
}

func RulerView_Alloc() RulerView {
	return RulerViewClass.Alloc()
}

func (rc _RulerViewClass) New() RulerView {
	rv := objc.Call[RulerView](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRulerView() RulerView {
	return RulerViewClass.New()
}

func (r_ RulerView) Init() RulerView {
	rv := objc.Call[RulerView](r_, objc.Sel("init"))
	return rv
}

func (r_ RulerView) InitWithFrame(frameRect foundation.Rect) RulerView {
	rv := objc.Call[RulerView](r_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func RulerView_InitWithFrame(frameRect foundation.Rect) RulerView {
	return RulerViewClass.Alloc().InitWithFrame(frameRect)
}

// Registers a new unit of measurement with the NSRulerView class, making it available to all instances of NSRulerView. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1534137-registerunitwithname?language=objc
func (rc _RulerViewClass) RegisterUnitWithNameAbbreviationUnitToPointsConversionFactorStepUpCycleStepDownCycle(unitName RulerViewUnitName, abbreviation string, conversionFactor float64, stepUpCycle []foundation.INumber, stepDownCycle []foundation.INumber) {
	objc.Call[objc.Void](rc, objc.Sel("registerUnitWithName:abbreviation:unitToPointsConversionFactor:stepUpCycle:stepDownCycle:"), unitName, abbreviation, conversionFactor, stepUpCycle, stepDownCycle)
}

// Registers a new unit of measurement with the NSRulerView class, making it available to all instances of NSRulerView. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1534137-registerunitwithname?language=objc
func RulerView_RegisterUnitWithNameAbbreviationUnitToPointsConversionFactorStepUpCycleStepDownCycle(unitName RulerViewUnitName, abbreviation string, conversionFactor float64, stepUpCycle []foundation.INumber, stepDownCycle []foundation.INumber) {
	RulerViewClass.RegisterUnitWithNameAbbreviationUnitToPointsConversionFactorStepUpCycleStepDownCycle(unitName, abbreviation, conversionFactor, stepUpCycle, stepDownCycle)
}

// Forces recalculation of the hash mark spacing for the next time the receiver is displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1530838-invalidatehashmarks?language=objc
func (r_ RulerView) InvalidateHashMarks() {
	objc.Call[objc.Void](r_, objc.Sel("invalidateHashMarks"))
}

// Tracks the mouse to add aMarker based on the initial mouse-down or mouse-dragged event theEvent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1535375-trackmarker?language=objc
func (r_ RulerView) TrackMarkerWithMouseEvent(marker IRulerMarker, event IEvent) bool {
	rv := objc.Call[bool](r_, objc.Sel("trackMarker:withMouseEvent:"), objc.Ptr(marker), objc.Ptr(event))
	return rv
}

// Draws temporary lines in the ruler area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1527414-moverulerlinefromlocation?language=objc
func (r_ RulerView) MoveRulerlineFromLocationToLocation(oldLocation float64, newLocation float64) {
	objc.Call[objc.Void](r_, objc.Sel("moveRulerlineFromLocation:toLocation:"), oldLocation, newLocation)
}

// Draws the receiver’s hash marks and labels in aRect, which is expressed in the receiver’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1529177-drawhashmarksandlabelsinrect?language=objc
func (r_ RulerView) DrawHashMarksAndLabelsInRect(rect foundation.Rect) {
	objc.Call[objc.Void](r_, objc.Sel("drawHashMarksAndLabelsInRect:"), rect)
}

// Draws the receiver’s markers in aRect, which is expressed in the receiver’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1525997-drawmarkersinrect?language=objc
func (r_ RulerView) DrawMarkersInRect(rect foundation.Rect) {
	objc.Call[objc.Void](r_, objc.Sel("drawMarkersInRect:"), rect)
}

// Removes aMarker from the receiver, without consulting the client view for approval. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1528230-removemarker?language=objc
func (r_ RulerView) RemoveMarker(marker IRulerMarker) {
	objc.Call[objc.Void](r_, objc.Sel("removeMarker:"), objc.Ptr(marker))
}

// Adds aMarker to the receiver, without consulting the client view for approval. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1524916-addmarker?language=objc
func (r_ RulerView) AddMarker(marker IRulerMarker) {
	objc.Call[objc.Void](r_, objc.Sel("addMarker:"), objc.Ptr(marker))
}

// The location of the receiver’s baseline, in its own coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1534886-baselinelocation?language=objc
func (r_ RulerView) BaselineLocation() float64 {
	rv := objc.Call[float64](r_, objc.Sel("baselineLocation"))
	return rv
}

// The room available for ruler markers to thickness. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1535112-reservedthicknessformarkers?language=objc
func (r_ RulerView) ReservedThicknessForMarkers() float64 {
	rv := objc.Call[float64](r_, objc.Sel("reservedThicknessForMarkers"))
	return rv
}

// The room available for ruler markers to thickness. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1535112-reservedthicknessformarkers?language=objc
func (r_ RulerView) SetReservedThicknessForMarkers(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setReservedThicknessForMarkers:"), value)
}

// The thickness of the area where ruler hash marks and labels are drawn. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1527872-rulethickness?language=objc
func (r_ RulerView) RuleThickness() float64 {
	rv := objc.Call[float64](r_, objc.Sel("ruleThickness"))
	return rv
}

// The thickness of the area where ruler hash marks and labels are drawn. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1527872-rulethickness?language=objc
func (r_ RulerView) SetRuleThickness(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setRuleThickness:"), value)
}

// The receiver’s accessory view to aView. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1532804-accessoryview?language=objc
func (r_ RulerView) AccessoryView() View {
	rv := objc.Call[View](r_, objc.Sel("accessoryView"))
	return rv
}

// The receiver’s accessory view to aView. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1532804-accessoryview?language=objc
func (r_ RulerView) SetAccessoryView(value IView) {
	objc.Call[objc.Void](r_, objc.Sel("setAccessoryView:"), objc.Ptr(value))
}

// The distance to the zero hash mark from the bounds origin of the NSScrollView’s document view (not of the receiver’s client view), in the document view’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1535432-originoffset?language=objc
func (r_ RulerView) OriginOffset() float64 {
	rv := objc.Call[float64](r_, objc.Sel("originOffset"))
	return rv
}

// The distance to the zero hash mark from the bounds origin of the NSScrollView’s document view (not of the receiver’s client view), in the document view’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1535432-originoffset?language=objc
func (r_ RulerView) SetOriginOffset(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setOriginOffset:"), value)
}

// The thickness needed for proper tiling of the receiver within an NSScrollView. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1526479-requiredthickness?language=objc
func (r_ RulerView) RequiredThickness() float64 {
	rv := objc.Call[float64](r_, objc.Sel("requiredThickness"))
	return rv
}

// The receiver’s client view, if it has one. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1533483-clientview?language=objc
func (r_ RulerView) ClientView() View {
	rv := objc.Call[View](r_, objc.Sel("clientView"))
	return rv
}

// The receiver’s client view, if it has one. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1533483-clientview?language=objc
func (r_ RulerView) SetClientView(value IView) {
	objc.Call[objc.Void](r_, objc.Sel("setClientView:"), objc.Ptr(value))
}

// The NSScrollView that owns the receiver to scrollView, without retaining it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1533741-scrollview?language=objc
func (r_ RulerView) ScrollView() ScrollView {
	rv := objc.Call[ScrollView](r_, objc.Sel("scrollView"))
	return rv
}

// The NSScrollView that owns the receiver to scrollView, without retaining it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1533741-scrollview?language=objc
func (r_ RulerView) SetScrollView(value IScrollView) {
	objc.Call[objc.Void](r_, objc.Sel("setScrollView:"), objc.Ptr(value))
}

// The measurement units used by the ruler to unitName. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1531493-measurementunits?language=objc
func (r_ RulerView) MeasurementUnits() RulerViewUnitName {
	rv := objc.Call[RulerViewUnitName](r_, objc.Sel("measurementUnits"))
	return rv
}

// The measurement units used by the ruler to unitName. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1531493-measurementunits?language=objc
func (r_ RulerView) SetMeasurementUnits(value RulerViewUnitName) {
	objc.Call[objc.Void](r_, objc.Sel("setMeasurementUnits:"), value)
}

// The room available for the receiver’s accessory view to thickness. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1530160-reservedthicknessforaccessoryvie?language=objc
func (r_ RulerView) ReservedThicknessForAccessoryView() float64 {
	rv := objc.Call[float64](r_, objc.Sel("reservedThicknessForAccessoryView"))
	return rv
}

// The room available for the receiver’s accessory view to thickness. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1530160-reservedthicknessforaccessoryvie?language=objc
func (r_ RulerView) SetReservedThicknessForAccessoryView(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setReservedThicknessForAccessoryView:"), value)
}

// The receiver’s ruler markers to markers, removing any existing ruler markers and not consulting with the client view about the new markers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1535213-markers?language=objc
func (r_ RulerView) Markers() []RulerMarker {
	rv := objc.Call[[]RulerMarker](r_, objc.Sel("markers"))
	return rv
}

// The receiver’s ruler markers to markers, removing any existing ruler markers and not consulting with the client view about the new markers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1535213-markers?language=objc
func (r_ RulerView) SetMarkers(value []IRulerMarker) {
	objc.Call[objc.Void](r_, objc.Sel("setMarkers:"), value)
}

// The orientation of the receiver to orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1530596-orientation?language=objc
func (r_ RulerView) Orientation() RulerOrientation {
	rv := objc.Call[RulerOrientation](r_, objc.Sel("orientation"))
	return rv
}

// The orientation of the receiver to orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/1530596-orientation?language=objc
func (r_ RulerView) SetOrientation(value RulerOrientation) {
	objc.Call[objc.Void](r_, objc.Sel("setOrientation:"), value)
}
