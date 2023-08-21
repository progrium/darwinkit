// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SegmentedCell] class.
var SegmentedCellClass = _SegmentedCellClass{objc.GetClass("NSSegmentedCell")}

type _SegmentedCellClass struct {
	objc.Class
}

// An interface definition for the [SegmentedCell] class.
type ISegmentedCell interface {
	IActionCell
	SetLabelForSegment(label string, segment int)
	ImageScalingForSegment(segment int) ImageScaling
	SetToolTipForSegment(toolTip string, segment int)
	TagForSegment(segment int) int
	ToolTipForSegment(segment int) string
	SetSelectedForSegment(selected bool, segment int)
	IsSelectedForSegment(segment int) bool
	SetImageForSegment(image IImage, segment int)
	MakeNextSegmentKey()
	ImageForSegment(segment int) Image
	InteriorBackgroundStyleForSegment(segment int) BackgroundStyle
	SetEnabledForSegment(enabled bool, segment int)
	WidthForSegment(segment int) float64
	MakePreviousSegmentKey()
	SelectSegmentWithTag(tag int) bool
	SetImageScalingForSegment(scaling ImageScaling, segment int)
	SetTagForSegment(tag int, segment int)
	MenuForSegment(segment int) Menu
	SetMenuForSegment(menu IMenu, segment int)
	DrawSegmentInFrameWithView(segment int, frame foundation.Rect, controlView IView)
	LabelForSegment(segment int) string
	SetWidthForSegment(width float64, segment int)
	IsEnabledForSegment(segment int) bool
	SegmentStyle() SegmentStyle
	SetSegmentStyle(value SegmentStyle)
	SegmentCount() int
	SetSegmentCount(value int)
	SelectedSegment() int
	SetSelectedSegment(value int)
	TrackingMode() SegmentSwitchTracking
	SetTrackingMode(value SegmentSwitchTracking)
}

// An NSSegmentedCell object implements the appearance and behavior of a horizontal button divided into multiple segments. This class is used in conjunction with the NSSegmentedControl class to implement a segmented control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell?language=objc
type SegmentedCell struct {
	ActionCell
}

func SegmentedCellFrom(ptr unsafe.Pointer) SegmentedCell {
	return SegmentedCell{
		ActionCell: ActionCellFrom(ptr),
	}
}

func (sc _SegmentedCellClass) Alloc() SegmentedCell {
	rv := objc.Call[SegmentedCell](sc, objc.Sel("alloc"))
	return rv
}

func SegmentedCell_Alloc() SegmentedCell {
	return SegmentedCellClass.Alloc()
}

func (sc _SegmentedCellClass) New() SegmentedCell {
	rv := objc.Call[SegmentedCell](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSegmentedCell() SegmentedCell {
	return SegmentedCellClass.New()
}

func (s_ SegmentedCell) Init() SegmentedCell {
	rv := objc.Call[SegmentedCell](s_, objc.Sel("init"))
	return rv
}

func (s_ SegmentedCell) InitImageCell(image IImage) SegmentedCell {
	rv := objc.Call[SegmentedCell](s_, objc.Sel("initImageCell:"), objc.Ptr(image))
	return rv
}

// Returns an NSCell object initialized with the specified image and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533898-initimagecell?language=objc
func NewSegmentedCellImageCell(image IImage) SegmentedCell {
	instance := SegmentedCellClass.Alloc().InitImageCell(image)
	instance.Autorelease()
	return instance
}

func (s_ SegmentedCell) InitTextCell(string_ string) SegmentedCell {
	rv := objc.Call[SegmentedCell](s_, objc.Sel("initTextCell:"), string_)
	return rv
}

// Returns an NSCell object initialized with the specified string and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1530851-inittextcell?language=objc
func NewSegmentedCellTextCell(string_ string) SegmentedCell {
	instance := SegmentedCellClass.Alloc().InitTextCell(string_)
	instance.Autorelease()
	return instance
}

// Sets the label for the specified segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500213-setlabel?language=objc
func (s_ SegmentedCell) SetLabelForSegment(label string, segment int) {
	objc.Call[objc.Void](s_, objc.Sel("setLabel:forSegment:"), label, segment)
}

// Returns the image scaling mode associated with the specified segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500217-imagescalingforsegment?language=objc
func (s_ SegmentedCell) ImageScalingForSegment(segment int) ImageScaling {
	rv := objc.Call[ImageScaling](s_, objc.Sel("imageScalingForSegment:"), segment)
	return rv
}

// Sets the tooltip for the specified segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500186-settooltip?language=objc
func (s_ SegmentedCell) SetToolTipForSegment(toolTip string, segment int) {
	objc.Call[objc.Void](s_, objc.Sel("setToolTip:forSegment:"), toolTip, segment)
}

// Returns the tag of the specified segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500207-tagforsegment?language=objc
func (s_ SegmentedCell) TagForSegment(segment int) int {
	rv := objc.Call[int](s_, objc.Sel("tagForSegment:"), segment)
	return rv
}

// Returns the tooltip of the specified segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500183-tooltipforsegment?language=objc
func (s_ SegmentedCell) ToolTipForSegment(segment int) string {
	rv := objc.Call[string](s_, objc.Sel("toolTipForSegment:"), segment)
	return rv
}

// Sets the selection state of the specified segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500188-setselected?language=objc
func (s_ SegmentedCell) SetSelectedForSegment(selected bool, segment int) {
	objc.Call[objc.Void](s_, objc.Sel("setSelected:forSegment:"), selected, segment)
}

// Returns a Boolean value indicating whether the specified segment is selected, [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500177-isselectedforsegment?language=objc
func (s_ SegmentedCell) IsSelectedForSegment(segment int) bool {
	rv := objc.Call[bool](s_, objc.Sel("isSelectedForSegment:"), segment)
	return rv
}

// Sets the image for the specified segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500174-setimage?language=objc
func (s_ SegmentedCell) SetImageForSegment(image IImage, segment int) {
	objc.Call[objc.Void](s_, objc.Sel("setImage:forSegment:"), objc.Ptr(image), segment)
}

// Selects the next segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500190-makenextsegmentkey?language=objc
func (s_ SegmentedCell) MakeNextSegmentKey() {
	objc.Call[objc.Void](s_, objc.Sel("makeNextSegmentKey"))
}

// Returns the image associated with the specified segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500204-imageforsegment?language=objc
func (s_ SegmentedCell) ImageForSegment(segment int) Image {
	rv := objc.Call[Image](s_, objc.Sel("imageForSegment:"), segment)
	return rv
}

// Returns the interior background style for the specified segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500198-interiorbackgroundstyleforsegmen?language=objc
func (s_ SegmentedCell) InteriorBackgroundStyleForSegment(segment int) BackgroundStyle {
	rv := objc.Call[BackgroundStyle](s_, objc.Sel("interiorBackgroundStyleForSegment:"), segment)
	return rv
}

// Sets the enabled state of the specified segment [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500219-setenabled?language=objc
func (s_ SegmentedCell) SetEnabledForSegment(enabled bool, segment int) {
	objc.Call[objc.Void](s_, objc.Sel("setEnabled:forSegment:"), enabled, segment)
}

// Returns the width of the specified segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500211-widthforsegment?language=objc
func (s_ SegmentedCell) WidthForSegment(segment int) float64 {
	rv := objc.Call[float64](s_, objc.Sel("widthForSegment:"), segment)
	return rv
}

// Selects the previous segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500170-makeprevioussegmentkey?language=objc
func (s_ SegmentedCell) MakePreviousSegmentKey() {
	objc.Call[objc.Void](s_, objc.Sel("makePreviousSegmentKey"))
}

// Selects the segment with the specified tag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500196-selectsegmentwithtag?language=objc
func (s_ SegmentedCell) SelectSegmentWithTag(tag int) bool {
	rv := objc.Call[bool](s_, objc.Sel("selectSegmentWithTag:"), tag)
	return rv
}

// Sets the image scaling mode for the specified segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500209-setimagescaling?language=objc
func (s_ SegmentedCell) SetImageScalingForSegment(scaling ImageScaling, segment int) {
	objc.Call[objc.Void](s_, objc.Sel("setImageScaling:forSegment:"), scaling, segment)
}

// Sets the tag for the specified segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500181-settag?language=objc
func (s_ SegmentedCell) SetTagForSegment(tag int, segment int) {
	objc.Call[objc.Void](s_, objc.Sel("setTag:forSegment:"), tag, segment)
}

// Returns the menu for the specified segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500175-menuforsegment?language=objc
func (s_ SegmentedCell) MenuForSegment(segment int) Menu {
	rv := objc.Call[Menu](s_, objc.Sel("menuForSegment:"), segment)
	return rv
}

// Sets the menu for the specified segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500194-setmenu?language=objc
func (s_ SegmentedCell) SetMenuForSegment(menu IMenu, segment int) {
	objc.Call[objc.Void](s_, objc.Sel("setMenu:forSegment:"), objc.Ptr(menu), segment)
}

// Draws the image and label of the segment in the specified view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500191-drawsegment?language=objc
func (s_ SegmentedCell) DrawSegmentInFrameWithView(segment int, frame foundation.Rect, controlView IView) {
	objc.Call[objc.Void](s_, objc.Sel("drawSegment:inFrame:withView:"), segment, frame, objc.Ptr(controlView))
}

// Returns the label of the specified segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500215-labelforsegment?language=objc
func (s_ SegmentedCell) LabelForSegment(segment int) string {
	rv := objc.Call[string](s_, objc.Sel("labelForSegment:"), segment)
	return rv
}

// Sets the width of the specified segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500202-setwidth?language=objc
func (s_ SegmentedCell) SetWidthForSegment(width float64, segment int) {
	objc.Call[objc.Void](s_, objc.Sel("setWidth:forSegment:"), width, segment)
}

// Returns a Boolean value indicating whether the specified segment is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500206-isenabledforsegment?language=objc
func (s_ SegmentedCell) IsEnabledForSegment(segment int) bool {
	rv := objc.Call[bool](s_, objc.Sel("isEnabledForSegment:"), segment)
	return rv
}

// The visual style used to display the segmented control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500172-segmentstyle?language=objc
func (s_ SegmentedCell) SegmentStyle() SegmentStyle {
	rv := objc.Call[SegmentStyle](s_, objc.Sel("segmentStyle"))
	return rv
}

// The visual style used to display the segmented control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500172-segmentstyle?language=objc
func (s_ SegmentedCell) SetSegmentStyle(value SegmentStyle) {
	objc.Call[objc.Void](s_, objc.Sel("setSegmentStyle:"), value)
}

// The number of segments in the segmented control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500185-segmentcount?language=objc
func (s_ SegmentedCell) SegmentCount() int {
	rv := objc.Call[int](s_, objc.Sel("segmentCount"))
	return rv
}

// The number of segments in the segmented control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500185-segmentcount?language=objc
func (s_ SegmentedCell) SetSegmentCount(value int) {
	objc.Call[objc.Void](s_, objc.Sel("setSegmentCount:"), value)
}

// The index of the selected segment of the control, or -1 if no segment is selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500193-selectedsegment?language=objc
func (s_ SegmentedCell) SelectedSegment() int {
	rv := objc.Call[int](s_, objc.Sel("selectedSegment"))
	return rv
}

// The index of the selected segment of the control, or -1 if no segment is selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500193-selectedsegment?language=objc
func (s_ SegmentedCell) SetSelectedSegment(value int) {
	objc.Call[objc.Void](s_, objc.Sel("setSelectedSegment:"), value)
}

// The tracking mode used for the segments of the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500200-trackingmode?language=objc
func (s_ SegmentedCell) TrackingMode() SegmentSwitchTracking {
	rv := objc.Call[SegmentSwitchTracking](s_, objc.Sel("trackingMode"))
	return rv
}

// The tracking mode used for the segments of the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/1500200-trackingmode?language=objc
func (s_ SegmentedCell) SetTrackingMode(value SegmentSwitchTracking) {
	objc.Call[objc.Void](s_, objc.Sel("setTrackingMode:"), value)
}
