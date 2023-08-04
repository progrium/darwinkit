// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var RulerViewClass = _RulerViewClass{objc.GetClass("NSRulerView")}

type _RulerViewClass struct {
	objc.Class
}

type IRulerView interface {
	IView
	AddMarker(marker IRulerMarker)
	RemoveMarker(marker IRulerMarker)
	TrackMarkerWithMouseEvent(marker IRulerMarker, event IEvent) bool
	MoveRulerlineFromLocationToLocation(oldLocation float64, newLocation float64)
	DrawHashMarksAndLabelsInRect(rect foundation.Rect)
	DrawMarkersInRect(rect foundation.Rect)
	InvalidateHashMarks()
	MeasurementUnits() RulerViewUnitName
	SetMeasurementUnits(value RulerViewUnitName)
	ClientView() View
	SetClientView(value IView)
	AccessoryView() View
	SetAccessoryView(value IView)
	OriginOffset() float64
	SetOriginOffset(value float64)
	Markers() []RulerMarker
	SetMarkers(value []IRulerMarker)
	ScrollView() ScrollView
	SetScrollView(value IScrollView)
	Orientation() RulerOrientation
	SetOrientation(value RulerOrientation)
	ReservedThicknessForAccessoryView() float64
	SetReservedThicknessForAccessoryView(value float64)
	ReservedThicknessForMarkers() float64
	SetReservedThicknessForMarkers(value float64)
	RuleThickness() float64
	SetRuleThickness(value float64)
	RequiredThickness() float64
	BaselineLocation() float64
}

type RulerView struct {
	View
}

func MakeRulerView(ptr unsafe.Pointer) RulerView {
	return RulerView{
		View: MakeView(ptr),
	}
}

func (r_ RulerView) InitWithScrollViewOrientation(scrollView IScrollView, orientation RulerOrientation) RulerView {
	rv := objc.CallMethod[RulerView](r_, objc.GetSelector("initWithScrollView:orientation:"), objc.ExtractPtr(scrollView), orientation)
	return rv
}

func RulerView_InitWithScrollViewOrientation(scrollView IScrollView, orientation RulerOrientation) RulerView {
	return RulerViewClass.Alloc().InitWithScrollViewOrientation(scrollView, orientation)
}

func (r_ RulerView) InitWithFrame(frameRect foundation.Rect) RulerView {
	rv := objc.CallMethod[RulerView](r_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func RulerView_InitWithFrame(frameRect foundation.Rect) RulerView {
	return RulerViewClass.Alloc().InitWithFrame(frameRect)
}

func (r_ RulerView) Init() RulerView {
	rv := objc.CallMethod[RulerView](r_, objc.GetSelector("init"))
	return rv
}

func RulerView_Init() RulerView {
	return RulerViewClass.Alloc().Init()
}

func (rc _RulerViewClass) Alloc() RulerView {
	rv := objc.CallMethod[RulerView](rc, objc.GetSelector("alloc"))
	return rv
}

func RulerView_Alloc() RulerView {
	return RulerViewClass.Alloc()
}

func (rc _RulerViewClass) New() RulerView {
	rv := objc.CallMethod[RulerView](rc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewRulerView() RulerView {
	return RulerViewClass.New()
}

func RulerView_New() RulerView {
	return RulerViewClass.New()
}

func (rc _RulerViewClass) RegisterUnitWithNameAbbreviationUnitToPointsConversionFactorStepUpCycleStepDownCycle(unitName RulerViewUnitName, abbreviation string, conversionFactor float64, stepUpCycle []foundation.INumber, stepDownCycle []foundation.INumber) {
	objc.CallMethod[objc.Void](rc, objc.GetSelector("registerUnitWithName:abbreviation:unitToPointsConversionFactor:stepUpCycle:stepDownCycle:"), unitName, abbreviation, conversionFactor, stepUpCycle, stepDownCycle)
}

func RulerView_RegisterUnitWithNameAbbreviationUnitToPointsConversionFactorStepUpCycleStepDownCycle(unitName RulerViewUnitName, abbreviation string, conversionFactor float64, stepUpCycle []foundation.INumber, stepDownCycle []foundation.INumber) {
	RulerViewClass.RegisterUnitWithNameAbbreviationUnitToPointsConversionFactorStepUpCycleStepDownCycle(unitName, abbreviation, conversionFactor, stepUpCycle, stepDownCycle)
}

func (r_ RulerView) AddMarker(marker IRulerMarker) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("addMarker:"), objc.ExtractPtr(marker))
}

func (r_ RulerView) RemoveMarker(marker IRulerMarker) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("removeMarker:"), objc.ExtractPtr(marker))
}

func (r_ RulerView) TrackMarkerWithMouseEvent(marker IRulerMarker, event IEvent) bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("trackMarker:withMouseEvent:"), objc.ExtractPtr(marker), objc.ExtractPtr(event))
	return rv
}

func (r_ RulerView) MoveRulerlineFromLocationToLocation(oldLocation float64, newLocation float64) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("moveRulerlineFromLocation:toLocation:"), oldLocation, newLocation)
}

func (r_ RulerView) DrawHashMarksAndLabelsInRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("drawHashMarksAndLabelsInRect:"), rect)
}

func (r_ RulerView) DrawMarkersInRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("drawMarkersInRect:"), rect)
}

func (r_ RulerView) InvalidateHashMarks() {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("invalidateHashMarks"))
}

func (r_ RulerView) MeasurementUnits() RulerViewUnitName {
	rv := objc.CallMethod[RulerViewUnitName](r_, objc.GetSelector("measurementUnits"))
	return rv
}

func (r_ RulerView) SetMeasurementUnits(value RulerViewUnitName) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setMeasurementUnits:"), value)
}

func (r_ RulerView) ClientView() View {
	rv := objc.CallMethod[View](r_, objc.GetSelector("clientView"))
	return rv
}

func (r_ RulerView) SetClientView(value IView) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setClientView:"), objc.ExtractPtr(value))
}

func (r_ RulerView) AccessoryView() View {
	rv := objc.CallMethod[View](r_, objc.GetSelector("accessoryView"))
	return rv
}

func (r_ RulerView) SetAccessoryView(value IView) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setAccessoryView:"), objc.ExtractPtr(value))
}

func (r_ RulerView) OriginOffset() float64 {
	rv := objc.CallMethod[float64](r_, objc.GetSelector("originOffset"))
	return rv
}

func (r_ RulerView) SetOriginOffset(value float64) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setOriginOffset:"), value)
}

func (r_ RulerView) Markers() []RulerMarker {
	rv := objc.CallMethod[[]RulerMarker](r_, objc.GetSelector("markers"))
	return rv
}

func (r_ RulerView) SetMarkers(value []IRulerMarker) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setMarkers:"), value)
}

func (r_ RulerView) ScrollView() ScrollView {
	rv := objc.CallMethod[ScrollView](r_, objc.GetSelector("scrollView"))
	return rv
}

func (r_ RulerView) SetScrollView(value IScrollView) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setScrollView:"), objc.ExtractPtr(value))
}

func (r_ RulerView) Orientation() RulerOrientation {
	rv := objc.CallMethod[RulerOrientation](r_, objc.GetSelector("orientation"))
	return rv
}

func (r_ RulerView) SetOrientation(value RulerOrientation) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setOrientation:"), value)
}

func (r_ RulerView) ReservedThicknessForAccessoryView() float64 {
	rv := objc.CallMethod[float64](r_, objc.GetSelector("reservedThicknessForAccessoryView"))
	return rv
}

func (r_ RulerView) SetReservedThicknessForAccessoryView(value float64) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setReservedThicknessForAccessoryView:"), value)
}

func (r_ RulerView) ReservedThicknessForMarkers() float64 {
	rv := objc.CallMethod[float64](r_, objc.GetSelector("reservedThicknessForMarkers"))
	return rv
}

func (r_ RulerView) SetReservedThicknessForMarkers(value float64) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setReservedThicknessForMarkers:"), value)
}

func (r_ RulerView) RuleThickness() float64 {
	rv := objc.CallMethod[float64](r_, objc.GetSelector("ruleThickness"))
	return rv
}

func (r_ RulerView) SetRuleThickness(value float64) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("setRuleThickness:"), value)
}

func (r_ RulerView) RequiredThickness() float64 {
	rv := objc.CallMethod[float64](r_, objc.GetSelector("requiredThickness"))
	return rv
}

func (r_ RulerView) BaselineLocation() float64 {
	rv := objc.CallMethod[float64](r_, objc.GetSelector("baselineLocation"))
	return rv
}
