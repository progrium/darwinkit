// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var SegmentedControlClass = _SegmentedControlClass{objc.GetClass("NSSegmentedControl")}

type _SegmentedControlClass struct {
	objc.Class
}

type ISegmentedControl interface {
	IControl
	LabelForSegment(segment int) string
	SetLabelForSegment(label string, segment int)
	SetAlignmentForSegment(alignment TextAlignment, segment int)
	AlignmentForSegment(segment int) TextAlignment
	SetImageForSegment(image IImage, segment int)
	ImageForSegment(segment int) Image
	SetImageScalingForSegment(scaling ImageScaling, segment int)
	ImageScalingForSegment(segment int) ImageScaling
	SetMenuForSegment(menu IMenu, segment int)
	MenuForSegment(segment int) Menu
	SetShowsMenuIndicatorForSegment(showsMenuIndicator bool, segment int)
	ShowsMenuIndicatorForSegment(segment int) bool
	SelectSegmentWithTag(tag int) bool
	SetSelectedForSegment(selected bool, segment int)
	IsSelectedForSegment(segment int) bool
	SetWidthForSegment(width float64, segment int)
	WidthForSegment(segment int) float64
	CompressWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions)
	MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) foundation.Size
	SetEnabledForSegment(enabled bool, segment int)
	IsEnabledForSegment(segment int) bool
	TagForSegment(segment int) int
	SetTagForSegment(tag int, segment int)
	SetToolTipForSegment(toolTip string, segment int)
	ToolTipForSegment(segment int) string
	TrackingMode() SegmentSwitchTracking
	SetTrackingMode(value SegmentSwitchTracking)
	SegmentStyle() SegmentStyle
	SetSegmentStyle(value SegmentStyle)
	SegmentCount() int
	SetSegmentCount(value int)
	IsSpringLoaded() bool
	SetSpringLoaded(value bool)
	SelectedSegment() int
	SetSelectedSegment(value int)
	IndexOfSelectedItem() int
	SelectedSegmentBezelColor() Color
	SetSelectedSegmentBezelColor(value IColor)
	DoubleValueForSelectedSegment() float64
	SegmentDistribution() SegmentDistribution
	SetSegmentDistribution(value SegmentDistribution)
	ActiveCompressionOptions() UserInterfaceCompressionOptions
}

type SegmentedControl struct {
	Control
}

func MakeSegmentedControl(ptr unsafe.Pointer) SegmentedControl {
	return SegmentedControl{
		Control: MakeControl(ptr),
	}
}

func (sc _SegmentedControlClass) SegmentedControlWithImagesTrackingModeTargetAction(images []IImage, trackingMode SegmentSwitchTracking, target objc.IObject, action objc.Selector) SegmentedControl {
	rv := objc.CallMethod[SegmentedControl](sc, objc.GetSelector("segmentedControlWithImages:trackingMode:target:action:"), images, trackingMode, objc.ExtractPtr(target), action)
	return rv
}

func SegmentedControl_SegmentedControlWithImagesTrackingModeTargetAction(images []IImage, trackingMode SegmentSwitchTracking, target objc.IObject, action objc.Selector) SegmentedControl {
	return SegmentedControlClass.SegmentedControlWithImagesTrackingModeTargetAction(images, trackingMode, target, action)
}

func (sc _SegmentedControlClass) SegmentedControlWithLabelsTrackingModeTargetAction(labels []string, trackingMode SegmentSwitchTracking, target objc.IObject, action objc.Selector) SegmentedControl {
	rv := objc.CallMethod[SegmentedControl](sc, objc.GetSelector("segmentedControlWithLabels:trackingMode:target:action:"), labels, trackingMode, objc.ExtractPtr(target), action)
	return rv
}

func SegmentedControl_SegmentedControlWithLabelsTrackingModeTargetAction(labels []string, trackingMode SegmentSwitchTracking, target objc.IObject, action objc.Selector) SegmentedControl {
	return SegmentedControlClass.SegmentedControlWithLabelsTrackingModeTargetAction(labels, trackingMode, target, action)
}

func (s_ SegmentedControl) InitWithFrame(frameRect foundation.Rect) SegmentedControl {
	rv := objc.CallMethod[SegmentedControl](s_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func SegmentedControl_InitWithFrame(frameRect foundation.Rect) SegmentedControl {
	return SegmentedControlClass.Alloc().InitWithFrame(frameRect)
}

func (s_ SegmentedControl) Init() SegmentedControl {
	rv := objc.CallMethod[SegmentedControl](s_, objc.GetSelector("init"))
	return rv
}

func SegmentedControl_Init() SegmentedControl {
	return SegmentedControlClass.Alloc().Init()
}

func (sc _SegmentedControlClass) Alloc() SegmentedControl {
	rv := objc.CallMethod[SegmentedControl](sc, objc.GetSelector("alloc"))
	return rv
}

func SegmentedControl_Alloc() SegmentedControl {
	return SegmentedControlClass.Alloc()
}

func (sc _SegmentedControlClass) New() SegmentedControl {
	rv := objc.CallMethod[SegmentedControl](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewSegmentedControl() SegmentedControl {
	return SegmentedControlClass.New()
}

func SegmentedControl_New() SegmentedControl {
	return SegmentedControlClass.New()
}

func (s_ SegmentedControl) LabelForSegment(segment int) string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("labelForSegment:"), segment)
	return rv
}

func (s_ SegmentedControl) SetLabelForSegment(label string, segment int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setLabel:forSegment:"), label, segment)
}

func (s_ SegmentedControl) SetAlignmentForSegment(alignment TextAlignment, segment int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setAlignment:forSegment:"), alignment, segment)
}

func (s_ SegmentedControl) AlignmentForSegment(segment int) TextAlignment {
	rv := objc.CallMethod[TextAlignment](s_, objc.GetSelector("alignmentForSegment:"), segment)
	return rv
}

func (s_ SegmentedControl) SetImageForSegment(image IImage, segment int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setImage:forSegment:"), objc.ExtractPtr(image), segment)
}

func (s_ SegmentedControl) ImageForSegment(segment int) Image {
	rv := objc.CallMethod[Image](s_, objc.GetSelector("imageForSegment:"), segment)
	return rv
}

func (s_ SegmentedControl) SetImageScalingForSegment(scaling ImageScaling, segment int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setImageScaling:forSegment:"), scaling, segment)
}

func (s_ SegmentedControl) ImageScalingForSegment(segment int) ImageScaling {
	rv := objc.CallMethod[ImageScaling](s_, objc.GetSelector("imageScalingForSegment:"), segment)
	return rv
}

func (s_ SegmentedControl) SetMenuForSegment(menu IMenu, segment int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setMenu:forSegment:"), objc.ExtractPtr(menu), segment)
}

func (s_ SegmentedControl) MenuForSegment(segment int) Menu {
	rv := objc.CallMethod[Menu](s_, objc.GetSelector("menuForSegment:"), segment)
	return rv
}

func (s_ SegmentedControl) SetShowsMenuIndicatorForSegment(showsMenuIndicator bool, segment int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setShowsMenuIndicator:forSegment:"), showsMenuIndicator, segment)
}

func (s_ SegmentedControl) ShowsMenuIndicatorForSegment(segment int) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("showsMenuIndicatorForSegment:"), segment)
	return rv
}

func (s_ SegmentedControl) SelectSegmentWithTag(tag int) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("selectSegmentWithTag:"), tag)
	return rv
}

func (s_ SegmentedControl) SetSelectedForSegment(selected bool, segment int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setSelected:forSegment:"), selected, segment)
}

func (s_ SegmentedControl) IsSelectedForSegment(segment int) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("isSelectedForSegment:"), segment)
	return rv
}

func (s_ SegmentedControl) SetWidthForSegment(width float64, segment int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setWidth:forSegment:"), width, segment)
}

func (s_ SegmentedControl) WidthForSegment(segment int) float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("widthForSegment:"), segment)
	return rv
}

func (s_ SegmentedControl) CompressWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("compressWithPrioritizedCompressionOptions:"), prioritizedOptions)
}

func (s_ SegmentedControl) MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) foundation.Size {
	rv := objc.CallMethod[foundation.Size](s_, objc.GetSelector("minimumSizeWithPrioritizedCompressionOptions:"), prioritizedOptions)
	return rv
}

func (s_ SegmentedControl) SetEnabledForSegment(enabled bool, segment int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setEnabled:forSegment:"), enabled, segment)
}

func (s_ SegmentedControl) IsEnabledForSegment(segment int) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("isEnabledForSegment:"), segment)
	return rv
}

func (s_ SegmentedControl) TagForSegment(segment int) int {
	rv := objc.CallMethod[int](s_, objc.GetSelector("tagForSegment:"), segment)
	return rv
}

func (s_ SegmentedControl) SetTagForSegment(tag int, segment int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setTag:forSegment:"), tag, segment)
}

func (s_ SegmentedControl) SetToolTipForSegment(toolTip string, segment int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setToolTip:forSegment:"), toolTip, segment)
}

func (s_ SegmentedControl) ToolTipForSegment(segment int) string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("toolTipForSegment:"), segment)
	return rv
}

func (s_ SegmentedControl) TrackingMode() SegmentSwitchTracking {
	rv := objc.CallMethod[SegmentSwitchTracking](s_, objc.GetSelector("trackingMode"))
	return rv
}

func (s_ SegmentedControl) SetTrackingMode(value SegmentSwitchTracking) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setTrackingMode:"), value)
}

func (s_ SegmentedControl) SegmentStyle() SegmentStyle {
	rv := objc.CallMethod[SegmentStyle](s_, objc.GetSelector("segmentStyle"))
	return rv
}

func (s_ SegmentedControl) SetSegmentStyle(value SegmentStyle) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setSegmentStyle:"), value)
}

func (s_ SegmentedControl) SegmentCount() int {
	rv := objc.CallMethod[int](s_, objc.GetSelector("segmentCount"))
	return rv
}

func (s_ SegmentedControl) SetSegmentCount(value int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setSegmentCount:"), value)
}

func (s_ SegmentedControl) IsSpringLoaded() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("isSpringLoaded"))
	return rv
}

func (s_ SegmentedControl) SetSpringLoaded(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setSpringLoaded:"), value)
}

func (s_ SegmentedControl) SelectedSegment() int {
	rv := objc.CallMethod[int](s_, objc.GetSelector("selectedSegment"))
	return rv
}

func (s_ SegmentedControl) SetSelectedSegment(value int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setSelectedSegment:"), value)
}

func (s_ SegmentedControl) IndexOfSelectedItem() int {
	rv := objc.CallMethod[int](s_, objc.GetSelector("indexOfSelectedItem"))
	return rv
}

func (s_ SegmentedControl) SelectedSegmentBezelColor() Color {
	rv := objc.CallMethod[Color](s_, objc.GetSelector("selectedSegmentBezelColor"))
	return rv
}

func (s_ SegmentedControl) SetSelectedSegmentBezelColor(value IColor) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setSelectedSegmentBezelColor:"), objc.ExtractPtr(value))
}

func (s_ SegmentedControl) DoubleValueForSelectedSegment() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("doubleValueForSelectedSegment"))
	return rv
}

func (s_ SegmentedControl) SegmentDistribution() SegmentDistribution {
	rv := objc.CallMethod[SegmentDistribution](s_, objc.GetSelector("segmentDistribution"))
	return rv
}

func (s_ SegmentedControl) SetSegmentDistribution(value SegmentDistribution) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setSegmentDistribution:"), value)
}

func (s_ SegmentedControl) ActiveCompressionOptions() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](s_, objc.GetSelector("activeCompressionOptions"))
	return rv
}
