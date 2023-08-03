// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var SegmentedCellClass = _SegmentedCellClass{objc.GetClass("NSSegmentedCell")}

type _SegmentedCellClass struct {
	objc.Class
}

type ISegmentedCell interface {
	IActionCell
	SetSelectedForSegment(selected bool, segment int)
	SelectSegmentWithTag(tag int) bool
	MakeNextSegmentKey()
	MakePreviousSegmentKey()
	IsSelectedForSegment(segment int) bool
	SetLabelForSegment(label string, segment int)
	LabelForSegment(segment int) string
	SetImageForSegment(image IImage, segment int)
	ImageForSegment(segment int) Image
	SetImageScalingForSegment(scaling ImageScaling, segment int)
	ImageScalingForSegment(segment int) ImageScaling
	SetWidthForSegment(width float64, segment int)
	WidthForSegment(segment int) float64
	SetEnabledForSegment(enabled bool, segment int)
	IsEnabledForSegment(segment int) bool
	SetMenuForSegment(menu IMenu, segment int)
	MenuForSegment(segment int) Menu
	SetToolTipForSegment(toolTip string, segment int)
	ToolTipForSegment(segment int) string
	SetTagForSegment(tag int, segment int)
	TagForSegment(segment int) int
	DrawSegmentInFrameWithView(segment int, frame foundation.Rect, controlView IView)
	InteriorBackgroundStyleForSegment(segment int) BackgroundStyle
	SegmentCount() int
	SetSegmentCount(value int)
	SelectedSegment() int
	SetSelectedSegment(value int)
	TrackingMode() SegmentSwitchTracking
	SetTrackingMode(value SegmentSwitchTracking)
	SegmentStyle() SegmentStyle
	SetSegmentStyle(value SegmentStyle)
}

type SegmentedCell struct {
	ActionCell
}

func MakeSegmentedCell(ptr unsafe.Pointer) SegmentedCell {
	return SegmentedCell{
		ActionCell: MakeActionCell(ptr),
	}
}

func (s_ SegmentedCell) InitImageCell(image IImage) SegmentedCell {
	rv := objc.CallMethod[SegmentedCell](s_, objc.GetSelector("initImageCell:"), objc.ExtractPtr(image))
	return rv
}

func SegmentedCell_InitImageCell(image IImage) SegmentedCell {
	return SegmentedCellClass.Alloc().InitImageCell(image)
}

func (s_ SegmentedCell) InitTextCell(string_ string) SegmentedCell {
	rv := objc.CallMethod[SegmentedCell](s_, objc.GetSelector("initTextCell:"), string_)
	return rv
}

func SegmentedCell_InitTextCell(string_ string) SegmentedCell {
	return SegmentedCellClass.Alloc().InitTextCell(string_)
}

func (s_ SegmentedCell) Init() SegmentedCell {
	rv := objc.CallMethod[SegmentedCell](s_, objc.GetSelector("init"))
	return rv
}

func SegmentedCell_Init() SegmentedCell {
	return SegmentedCellClass.Alloc().Init()
}

func (sc _SegmentedCellClass) Alloc() SegmentedCell {
	rv := objc.CallMethod[SegmentedCell](sc, objc.GetSelector("alloc"))
	return rv
}

func SegmentedCell_Alloc() SegmentedCell {
	return SegmentedCellClass.Alloc()
}

func (sc _SegmentedCellClass) New() SegmentedCell {
	rv := objc.CallMethod[SegmentedCell](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewSegmentedCell() SegmentedCell {
	return SegmentedCellClass.New()
}

func SegmentedCell_New() SegmentedCell {
	return SegmentedCellClass.New()
}

func (s_ SegmentedCell) SetSelectedForSegment(selected bool, segment int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setSelected:forSegment:"), selected, segment)
}

func (s_ SegmentedCell) SelectSegmentWithTag(tag int) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("selectSegmentWithTag:"), tag)
	return rv
}

func (s_ SegmentedCell) MakeNextSegmentKey() {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("makeNextSegmentKey"))
}

func (s_ SegmentedCell) MakePreviousSegmentKey() {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("makePreviousSegmentKey"))
}

func (s_ SegmentedCell) IsSelectedForSegment(segment int) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("isSelectedForSegment:"), segment)
	return rv
}

func (s_ SegmentedCell) SetLabelForSegment(label string, segment int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setLabel:forSegment:"), label, segment)
}

func (s_ SegmentedCell) LabelForSegment(segment int) string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("labelForSegment:"), segment)
	return rv
}

func (s_ SegmentedCell) SetImageForSegment(image IImage, segment int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setImage:forSegment:"), objc.ExtractPtr(image), segment)
}

func (s_ SegmentedCell) ImageForSegment(segment int) Image {
	rv := objc.CallMethod[Image](s_, objc.GetSelector("imageForSegment:"), segment)
	return rv
}

func (s_ SegmentedCell) SetImageScalingForSegment(scaling ImageScaling, segment int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setImageScaling:forSegment:"), scaling, segment)
}

func (s_ SegmentedCell) ImageScalingForSegment(segment int) ImageScaling {
	rv := objc.CallMethod[ImageScaling](s_, objc.GetSelector("imageScalingForSegment:"), segment)
	return rv
}

func (s_ SegmentedCell) SetWidthForSegment(width float64, segment int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setWidth:forSegment:"), width, segment)
}

func (s_ SegmentedCell) WidthForSegment(segment int) float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("widthForSegment:"), segment)
	return rv
}

func (s_ SegmentedCell) SetEnabledForSegment(enabled bool, segment int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setEnabled:forSegment:"), enabled, segment)
}

func (s_ SegmentedCell) IsEnabledForSegment(segment int) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("isEnabledForSegment:"), segment)
	return rv
}

func (s_ SegmentedCell) SetMenuForSegment(menu IMenu, segment int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setMenu:forSegment:"), objc.ExtractPtr(menu), segment)
}

func (s_ SegmentedCell) MenuForSegment(segment int) Menu {
	rv := objc.CallMethod[Menu](s_, objc.GetSelector("menuForSegment:"), segment)
	return rv
}

func (s_ SegmentedCell) SetToolTipForSegment(toolTip string, segment int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setToolTip:forSegment:"), toolTip, segment)
}

func (s_ SegmentedCell) ToolTipForSegment(segment int) string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("toolTipForSegment:"), segment)
	return rv
}

func (s_ SegmentedCell) SetTagForSegment(tag int, segment int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setTag:forSegment:"), tag, segment)
}

func (s_ SegmentedCell) TagForSegment(segment int) int {
	rv := objc.CallMethod[int](s_, objc.GetSelector("tagForSegment:"), segment)
	return rv
}

func (s_ SegmentedCell) DrawSegmentInFrameWithView(segment int, frame foundation.Rect, controlView IView) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("drawSegment:inFrame:withView:"), segment, frame, objc.ExtractPtr(controlView))
}

func (s_ SegmentedCell) InteriorBackgroundStyleForSegment(segment int) BackgroundStyle {
	rv := objc.CallMethod[BackgroundStyle](s_, objc.GetSelector("interiorBackgroundStyleForSegment:"), segment)
	return rv
}

func (s_ SegmentedCell) SegmentCount() int {
	rv := objc.CallMethod[int](s_, objc.GetSelector("segmentCount"))
	return rv
}

func (s_ SegmentedCell) SetSegmentCount(value int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setSegmentCount:"), value)
}

func (s_ SegmentedCell) SelectedSegment() int {
	rv := objc.CallMethod[int](s_, objc.GetSelector("selectedSegment"))
	return rv
}

func (s_ SegmentedCell) SetSelectedSegment(value int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setSelectedSegment:"), value)
}

func (s_ SegmentedCell) TrackingMode() SegmentSwitchTracking {
	rv := objc.CallMethod[SegmentSwitchTracking](s_, objc.GetSelector("trackingMode"))
	return rv
}

func (s_ SegmentedCell) SetTrackingMode(value SegmentSwitchTracking) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setTrackingMode:"), value)
}

func (s_ SegmentedCell) SegmentStyle() SegmentStyle {
	rv := objc.CallMethod[SegmentStyle](s_, objc.GetSelector("segmentStyle"))
	return rv
}

func (s_ SegmentedCell) SetSegmentStyle(value SegmentStyle) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setSegmentStyle:"), value)
}
