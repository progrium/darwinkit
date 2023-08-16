// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SegmentedControl] class.
var SegmentedControlClass = _SegmentedControlClass{objc.GetClass("NSSegmentedControl")}

type _SegmentedControlClass struct {
	objc.Class
}

// An interface definition for the [SegmentedControl] class.
type ISegmentedControl interface {
	IControl
	SetLabelForSegment(label string, segment int)
	ImageScalingForSegment(segment int) ImageScaling
	SetToolTipForSegment(toolTip string, segment int)
	TagForSegment(segment int) int
	ToolTipForSegment(segment int) string
	SetAlignmentForSegment(alignment TextAlignment, segment int)
	SetSelectedForSegment(selected bool, segment int)
	IsSelectedForSegment(segment int) bool
	SetImageForSegment(image IImage, segment int)
	SetShowsMenuIndicatorForSegment(showsMenuIndicator bool, segment int)
	MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) foundation.Size
	ImageForSegment(segment int) Image
	CompressWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions)
	SetEnabledForSegment(enabled bool, segment int)
	WidthForSegment(segment int) float64
	SelectSegmentWithTag(tag int) bool
	AlignmentForSegment(segment int) TextAlignment
	SetImageScalingForSegment(scaling ImageScaling, segment int)
	SetTagForSegment(tag int, segment int)
	MenuForSegment(segment int) Menu
	SetMenuForSegment(menu IMenu, segment int)
	ShowsMenuIndicatorForSegment(segment int) bool
	LabelForSegment(segment int) string
	SetWidthForSegment(width float64, segment int)
	IsEnabledForSegment(segment int) bool
	ActiveCompressionOptions() UserInterfaceCompressionOptions
	IndexOfSelectedItem() int
	SegmentStyle() SegmentStyle
	SetSegmentStyle(value SegmentStyle)
	SegmentCount() int
	SetSegmentCount(value int)
	SelectedSegment() int
	SetSelectedSegment(value int)
	SegmentDistribution() SegmentDistribution
	SetSegmentDistribution(value SegmentDistribution)
	TrackingMode() SegmentSwitchTracking
	SetTrackingMode(value SegmentSwitchTracking)
	DoubleValueForSelectedSegment() float64
	SelectedSegmentBezelColor() Color
	SetSelectedSegmentBezelColor(value IColor)
	IsSpringLoaded() bool
	SetSpringLoaded(value bool)
}

// Display one or more buttons in a single horizontal group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol?language=objc
type SegmentedControl struct {
	Control
}

func SegmentedControlFrom(ptr unsafe.Pointer) SegmentedControl {
	return SegmentedControl{
		Control: ControlFrom(ptr),
	}
}

func (sc _SegmentedControlClass) SegmentedControlWithImagesTrackingModeTargetAction(images []IImage, trackingMode SegmentSwitchTracking, target objc.IObject, action objc.Selector) SegmentedControl {
	rv := objc.Call[SegmentedControl](sc, objc.Sel("segmentedControlWithImages:trackingMode:target:action:"), images, trackingMode, target, action)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1644010-segmentedcontrolwithimages?language=objc
func SegmentedControl_SegmentedControlWithImagesTrackingModeTargetAction(images []IImage, trackingMode SegmentSwitchTracking, target objc.IObject, action objc.Selector) SegmentedControl {
	return SegmentedControlClass.SegmentedControlWithImagesTrackingModeTargetAction(images, trackingMode, target, action)
}

func (sc _SegmentedControlClass) SegmentedControlWithLabelsTrackingModeTargetAction(labels []string, trackingMode SegmentSwitchTracking, target objc.IObject, action objc.Selector) SegmentedControl {
	rv := objc.Call[SegmentedControl](sc, objc.Sel("segmentedControlWithLabels:trackingMode:target:action:"), labels, trackingMode, target, action)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1644011-segmentedcontrolwithlabels?language=objc
func SegmentedControl_SegmentedControlWithLabelsTrackingModeTargetAction(labels []string, trackingMode SegmentSwitchTracking, target objc.IObject, action objc.Selector) SegmentedControl {
	return SegmentedControlClass.SegmentedControlWithLabelsTrackingModeTargetAction(labels, trackingMode, target, action)
}

func (sc _SegmentedControlClass) Alloc() SegmentedControl {
	rv := objc.Call[SegmentedControl](sc, objc.Sel("alloc"))
	return rv
}

func SegmentedControl_Alloc() SegmentedControl {
	return SegmentedControlClass.Alloc()
}

func (sc _SegmentedControlClass) New() SegmentedControl {
	rv := objc.Call[SegmentedControl](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSegmentedControl() SegmentedControl {
	return SegmentedControlClass.New()
}

func (s_ SegmentedControl) Init() SegmentedControl {
	rv := objc.Call[SegmentedControl](s_, objc.Sel("init"))
	return rv
}

func (s_ SegmentedControl) InitWithFrame(frameRect foundation.Rect) SegmentedControl {
	rv := objc.Call[SegmentedControl](s_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes a control with the specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428900-initwithframe?language=objc
func SegmentedControl_InitWithFrame(frameRect foundation.Rect) SegmentedControl {
	return SegmentedControlClass.Alloc().InitWithFrame(frameRect)
}

// Sets the label for the specified segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1533073-setlabel?language=objc
func (s_ SegmentedControl) SetLabelForSegment(label string, segment int) {
	objc.Call[objc.Void](s_, objc.Sel("setLabel:forSegment:"), label, segment)
}

// Returns the scaling mode used to display the specified segment’s image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1532175-imagescalingforsegment?language=objc
func (s_ SegmentedControl) ImageScalingForSegment(segment int) ImageScaling {
	rv := objc.Call[ImageScaling](s_, objc.Sel("imageScalingForSegment:"), segment)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/2887101-settooltip?language=objc
func (s_ SegmentedControl) SetToolTipForSegment(toolTip string, segment int) {
	objc.Call[objc.Void](s_, objc.Sel("setToolTip:forSegment:"), toolTip, segment)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/2887113-tagforsegment?language=objc
func (s_ SegmentedControl) TagForSegment(segment int) int {
	rv := objc.Call[int](s_, objc.Sel("tagForSegment:"), segment)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/2887103-tooltipforsegment?language=objc
func (s_ SegmentedControl) ToolTipForSegment(segment int) string {
	rv := objc.Call[string](s_, objc.Sel("toolTipForSegment:"), segment)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/2887107-setalignment?language=objc
func (s_ SegmentedControl) SetAlignmentForSegment(alignment TextAlignment, segment int) {
	objc.Call[objc.Void](s_, objc.Sel("setAlignment:forSegment:"), alignment, segment)
}

// Sets the selection state of the specified segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1528334-setselected?language=objc
func (s_ SegmentedControl) SetSelectedForSegment(selected bool, segment int) {
	objc.Call[objc.Void](s_, objc.Sel("setSelected:forSegment:"), selected, segment)
}

// Returns a Boolean value indicating whether the specified segment is selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1525676-isselectedforsegment?language=objc
func (s_ SegmentedControl) IsSelectedForSegment(segment int) bool {
	rv := objc.Call[bool](s_, objc.Sel("isSelectedForSegment:"), segment)
	return rv
}

// Sets the image for the specified segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1525540-setimage?language=objc
func (s_ SegmentedControl) SetImageForSegment(image IImage, segment int) {
	objc.Call[objc.Void](s_, objc.Sel("setImage:forSegment:"), objc.Ptr(image), segment)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/2887112-setshowsmenuindicator?language=objc
func (s_ SegmentedControl) SetShowsMenuIndicatorForSegment(showsMenuIndicator bool, segment int) {
	objc.Call[objc.Void](s_, objc.Sel("setShowsMenuIndicator:forSegment:"), showsMenuIndicator, segment)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/2952063-minimumsizewithprioritizedcompre?language=objc
func (s_ SegmentedControl) MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) foundation.Size {
	rv := objc.Call[foundation.Size](s_, objc.Sel("minimumSizeWithPrioritizedCompressionOptions:"), prioritizedOptions)
	return rv
}

// Returns the image associated with the specified segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1527623-imageforsegment?language=objc
func (s_ SegmentedControl) ImageForSegment(segment int) Image {
	rv := objc.Call[Image](s_, objc.Sel("imageForSegment:"), segment)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/2952064-compresswithprioritizedcompressi?language=objc
func (s_ SegmentedControl) CompressWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) {
	objc.Call[objc.Void](s_, objc.Sel("compressWithPrioritizedCompressionOptions:"), prioritizedOptions)
}

// Sets the enabled state of the specified segment [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1526417-setenabled?language=objc
func (s_ SegmentedControl) SetEnabledForSegment(enabled bool, segment int) {
	objc.Call[objc.Void](s_, objc.Sel("setEnabled:forSegment:"), enabled, segment)
}

// Returns the width of the specified segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1533866-widthforsegment?language=objc
func (s_ SegmentedControl) WidthForSegment(segment int) float64 {
	rv := objc.Call[float64](s_, objc.Sel("widthForSegment:"), segment)
	return rv
}

// Selects the segment with the specified tag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1534016-selectsegmentwithtag?language=objc
func (s_ SegmentedControl) SelectSegmentWithTag(tag int) bool {
	rv := objc.Call[bool](s_, objc.Sel("selectSegmentWithTag:"), tag)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/2887114-alignmentforsegment?language=objc
func (s_ SegmentedControl) AlignmentForSegment(segment int) TextAlignment {
	rv := objc.Call[TextAlignment](s_, objc.Sel("alignmentForSegment:"), segment)
	return rv
}

// Sets the scaling mode used to display the specified segment’s image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1532914-setimagescaling?language=objc
func (s_ SegmentedControl) SetImageScalingForSegment(scaling ImageScaling, segment int) {
	objc.Call[objc.Void](s_, objc.Sel("setImageScaling:forSegment:"), scaling, segment)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/2887104-settag?language=objc
func (s_ SegmentedControl) SetTagForSegment(tag int, segment int) {
	objc.Call[objc.Void](s_, objc.Sel("setTag:forSegment:"), tag, segment)
}

// Returns the menu for the specified segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1525880-menuforsegment?language=objc
func (s_ SegmentedControl) MenuForSegment(segment int) Menu {
	rv := objc.Call[Menu](s_, objc.Sel("menuForSegment:"), segment)
	return rv
}

// Sets the menu for the specified segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1528853-setmenu?language=objc
func (s_ SegmentedControl) SetMenuForSegment(menu IMenu, segment int) {
	objc.Call[objc.Void](s_, objc.Sel("setMenu:forSegment:"), objc.Ptr(menu), segment)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/2887111-showsmenuindicatorforsegment?language=objc
func (s_ SegmentedControl) ShowsMenuIndicatorForSegment(segment int) bool {
	rv := objc.Call[bool](s_, objc.Sel("showsMenuIndicatorForSegment:"), segment)
	return rv
}

// Returns the label of the specified segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1524322-labelforsegment?language=objc
func (s_ SegmentedControl) LabelForSegment(segment int) string {
	rv := objc.Call[string](s_, objc.Sel("labelForSegment:"), segment)
	return rv
}

// Sets the width of the specified segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1534487-setwidth?language=objc
func (s_ SegmentedControl) SetWidthForSegment(width float64, segment int) {
	objc.Call[objc.Void](s_, objc.Sel("setWidth:forSegment:"), width, segment)
}

// Returns a Boolean value indicating whether the specified segment is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1524542-isenabledforsegment?language=objc
func (s_ SegmentedControl) IsEnabledForSegment(segment int) bool {
	rv := objc.Call[bool](s_, objc.Sel("isEnabledForSegment:"), segment)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/2952062-activecompressionoptions?language=objc
func (s_ SegmentedControl) ActiveCompressionOptions() UserInterfaceCompressionOptions {
	rv := objc.Call[UserInterfaceCompressionOptions](s_, objc.Sel("activeCompressionOptions"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/2887105-indexofselecteditem?language=objc
func (s_ SegmentedControl) IndexOfSelectedItem() int {
	rv := objc.Call[int](s_, objc.Sel("indexOfSelectedItem"))
	return rv
}

// The visual style used to display the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1528507-segmentstyle?language=objc
func (s_ SegmentedControl) SegmentStyle() SegmentStyle {
	rv := objc.Call[SegmentStyle](s_, objc.Sel("segmentStyle"))
	return rv
}

// The visual style used to display the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1528507-segmentstyle?language=objc
func (s_ SegmentedControl) SetSegmentStyle(value SegmentStyle) {
	objc.Call[objc.Void](s_, objc.Sel("setSegmentStyle:"), value)
}

// The number of segments in the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1527451-segmentcount?language=objc
func (s_ SegmentedControl) SegmentCount() int {
	rv := objc.Call[int](s_, objc.Sel("segmentCount"))
	return rv
}

// The number of segments in the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1527451-segmentcount?language=objc
func (s_ SegmentedControl) SetSegmentCount(value int) {
	objc.Call[objc.Void](s_, objc.Sel("setSegmentCount:"), value)
}

// The index of the selected segment of the control, or -1 if no segment is selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1534490-selectedsegment?language=objc
func (s_ SegmentedControl) SelectedSegment() int {
	rv := objc.Call[int](s_, objc.Sel("selectedSegment"))
	return rv
}

// The index of the selected segment of the control, or -1 if no segment is selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1534490-selectedsegment?language=objc
func (s_ SegmentedControl) SetSelectedSegment(value int) {
	objc.Call[objc.Void](s_, objc.Sel("setSelectedSegment:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/2887109-segmentdistribution?language=objc
func (s_ SegmentedControl) SegmentDistribution() SegmentDistribution {
	rv := objc.Call[SegmentDistribution](s_, objc.Sel("segmentDistribution"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/2887109-segmentdistribution?language=objc
func (s_ SegmentedControl) SetSegmentDistribution(value SegmentDistribution) {
	objc.Call[objc.Void](s_, objc.Sel("setSegmentDistribution:"), value)
}

// The type of tracking behavior the control exhibits. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1526285-trackingmode?language=objc
func (s_ SegmentedControl) TrackingMode() SegmentSwitchTracking {
	rv := objc.Call[SegmentSwitchTracking](s_, objc.Sel("trackingMode"))
	return rv
}

// The type of tracking behavior the control exhibits. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1526285-trackingmode?language=objc
func (s_ SegmentedControl) SetTrackingMode(value SegmentSwitchTracking) {
	objc.Call[objc.Void](s_, objc.Sel("setTrackingMode:"), value)
}

// When the tracking mode for the control is set to use a momentary accelerator, returns a value for the selected segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1529727-doublevalueforselectedsegment?language=objc
func (s_ SegmentedControl) DoubleValueForSelectedSegment() float64 {
	rv := objc.Call[float64](s_, objc.Sel("doubleValueForSelectedSegment"))
	return rv
}

// The color of the selected segment's bezel, in appearances that support it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/2561002-selectedsegmentbezelcolor?language=objc
func (s_ SegmentedControl) SelectedSegmentBezelColor() Color {
	rv := objc.Call[Color](s_, objc.Sel("selectedSegmentBezelColor"))
	return rv
}

// The color of the selected segment's bezel, in appearances that support it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/2561002-selectedsegmentbezelcolor?language=objc
func (s_ SegmentedControl) SetSelectedSegmentBezelColor(value IColor) {
	objc.Call[objc.Void](s_, objc.Sel("setSelectedSegmentBezelColor:"), objc.Ptr(value))
}

// A Boolean value that indicates whether spring loading is enabled for the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1534122-springloaded?language=objc
func (s_ SegmentedControl) IsSpringLoaded() bool {
	rv := objc.Call[bool](s_, objc.Sel("isSpringLoaded"))
	return rv
}

// A Boolean value that indicates whether spring loading is enabled for the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/1534122-springloaded?language=objc
func (s_ SegmentedControl) SetSpringLoaded(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setSpringLoaded:"), value)
}
