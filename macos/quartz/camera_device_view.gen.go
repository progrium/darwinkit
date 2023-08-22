// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CameraDeviceView] class.
var CameraDeviceViewClass = _CameraDeviceViewClass{objc.GetClass("IKCameraDeviceView")}

type _CameraDeviceViewClass struct {
	objc.Class
}

// An interface definition for the [CameraDeviceView] class.
type ICameraDeviceView interface {
	appkit.IView
	SetShowStatusInfoAsWindowSubtitle(value bool)
	RotateRight(sender objc.IObject) objc.Object
	SelectIndexesByExtendingSelection(indexes foundation.IIndexSet, extend bool)
	SetCustomModeControl(control appkit.ISegmentedControl)
	DownloadAllItems(sender objc.IObject) objc.Object
	SetCustomActionControl(control appkit.ISegmentedControl)
	RotateLeft(sender objc.IObject) objc.Object
	DeleteSelectedItems(sender objc.IObject) objc.Object
	SetCustomDeleteControl(control appkit.ISegmentedControl)
	SetCustomIconSizeSlider(slider appkit.ISlider)
	SetCustomRotateControl(control appkit.ISegmentedControl)
	SelectedIndexes() foundation.IndexSet
	DownloadSelectedItems(sender objc.IObject) objc.Object
	DownloadsDirectory() foundation.URL
	SetDownloadsDirectory(value foundation.IURL)
	DownloadSelectedControlLabel() string
	SetDownloadSelectedControlLabel(value string)
	DisplaysDownloadsDirectoryControl() bool
	SetDisplaysDownloadsDirectoryControl(value bool)
	CanDownloadSelectedItems() bool
	CanRotateSelectedItemsLeft() bool
	DisplaysPostProcessApplicationControl() bool
	SetDisplaysPostProcessApplicationControl(value bool)
	DownloadAllControlLabel() string
	SetDownloadAllControlLabel(value string)
	IconSize() uint
	SetIconSize(value uint)
	TransferMode() CameraDeviceViewTransferMode
	SetTransferMode(value CameraDeviceViewTransferMode)
	Delegate() CameraDeviceViewDelegateWrapper
	SetDelegate(value PCameraDeviceViewDelegate)
	SetDelegateObject(valueObject objc.IObject)
	PostProcessApplication() foundation.URL
	SetPostProcessApplication(value foundation.IURL)
	HasDisplayModeIcon() bool
	SetHasDisplayModeIcon(value bool)
	Mode() CameraDeviceViewDisplayMode
	SetMode(value CameraDeviceViewDisplayMode)
	CanRotateSelectedItemsRight() bool
	CanDeleteSelectedItems() bool
	HasDisplayModeTable() bool
	SetHasDisplayModeTable(value bool)
}

// The IKCameraDeviceView class displays the contents of the selected camera. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview?language=objc
type CameraDeviceView struct {
	appkit.View
}

func CameraDeviceViewFrom(ptr unsafe.Pointer) CameraDeviceView {
	return CameraDeviceView{
		View: appkit.ViewFrom(ptr),
	}
}

func (cc _CameraDeviceViewClass) Alloc() CameraDeviceView {
	rv := objc.Call[CameraDeviceView](cc, objc.Sel("alloc"))
	return rv
}

func CameraDeviceView_Alloc() CameraDeviceView {
	return CameraDeviceViewClass.Alloc()
}

func (cc _CameraDeviceViewClass) New() CameraDeviceView {
	rv := objc.Call[CameraDeviceView](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCameraDeviceView() CameraDeviceView {
	return CameraDeviceViewClass.New()
}

func (c_ CameraDeviceView) Init() CameraDeviceView {
	rv := objc.Call[CameraDeviceView](c_, objc.Sel("init"))
	return rv
}

func (c_ CameraDeviceView) InitWithFrame(frameRect foundation.Rect) CameraDeviceView {
	rv := objc.Call[CameraDeviceView](c_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func NewCameraDeviceViewWithFrame(frameRect foundation.Rect) CameraDeviceView {
	instance := CameraDeviceViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/3852628-setshowstatusinfoaswindowsubtitl?language=objc
func (c_ CameraDeviceView) SetShowStatusInfoAsWindowSubtitle(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setShowStatusInfoAsWindowSubtitle:"), value)
}

// Rotates the selected image to the right. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1505123-rotateright?language=objc
func (c_ CameraDeviceView) RotateRight(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("rotateRight:"), sender)
	return rv
}

// Invoked to select the specified files, extending the selection if specified. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1503804-selectindexes?language=objc
func (c_ CameraDeviceView) SelectIndexesByExtendingSelection(indexes foundation.IIndexSet, extend bool) {
	objc.Call[objc.Void](c_, objc.Sel("selectIndexes:byExtendingSelection:"), objc.Ptr(indexes), extend)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/3852626-setcustommodecontrol?language=objc
func (c_ CameraDeviceView) SetCustomModeControl(control appkit.ISegmentedControl) {
	objc.Call[objc.Void](c_, objc.Sel("setCustomModeControl:"), objc.Ptr(control))
}

// Downloads all the items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1504326-downloadallitems?language=objc
func (c_ CameraDeviceView) DownloadAllItems(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("downloadAllItems:"), sender)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/3852623-setcustomactioncontrol?language=objc
func (c_ CameraDeviceView) SetCustomActionControl(control appkit.ISegmentedControl) {
	objc.Call[objc.Void](c_, objc.Sel("setCustomActionControl:"), objc.Ptr(control))
}

// Rotates the selected image to the left. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1503662-rotateleft?language=objc
func (c_ CameraDeviceView) RotateLeft(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("rotateLeft:"), sender)
	return rv
}

// Deletes the currently selected items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1504333-deleteselecteditems?language=objc
func (c_ CameraDeviceView) DeleteSelectedItems(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("deleteSelectedItems:"), sender)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/3852624-setcustomdeletecontrol?language=objc
func (c_ CameraDeviceView) SetCustomDeleteControl(control appkit.ISegmentedControl) {
	objc.Call[objc.Void](c_, objc.Sel("setCustomDeleteControl:"), objc.Ptr(control))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/3852625-setcustomiconsizeslider?language=objc
func (c_ CameraDeviceView) SetCustomIconSizeSlider(slider appkit.ISlider) {
	objc.Call[objc.Void](c_, objc.Sel("setCustomIconSizeSlider:"), objc.Ptr(slider))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/3852627-setcustomrotatecontrol?language=objc
func (c_ CameraDeviceView) SetCustomRotateControl(control appkit.ISegmentedControl) {
	objc.Call[objc.Void](c_, objc.Sel("setCustomRotateControl:"), objc.Ptr(control))
}

// The selected indexes of the camera files. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1504546-selectedindexes?language=objc
func (c_ CameraDeviceView) SelectedIndexes() foundation.IndexSet {
	rv := objc.Call[foundation.IndexSet](c_, objc.Sel("selectedIndexes"))
	return rv
}

// Deletes the selected items from the camera. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1504833-downloadselecteditems?language=objc
func (c_ CameraDeviceView) DownloadSelectedItems(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("downloadSelectedItems:"), sender)
	return rv
}

// Specifies the directory where files are downloaded [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1503702-downloadsdirectory?language=objc
func (c_ CameraDeviceView) DownloadsDirectory() foundation.URL {
	rv := objc.Call[foundation.URL](c_, objc.Sel("downloadsDirectory"))
	return rv
}

// Specifies the directory where files are downloaded [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1503702-downloadsdirectory?language=objc
func (c_ CameraDeviceView) SetDownloadsDirectory(value foundation.IURL) {
	objc.Call[objc.Void](c_, objc.Sel("setDownloadsDirectory:"), objc.Ptr(value))
}

// Allows the “Download Selected” control to be renamed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1503822-downloadselectedcontrollabel?language=objc
func (c_ CameraDeviceView) DownloadSelectedControlLabel() string {
	rv := objc.Call[string](c_, objc.Sel("downloadSelectedControlLabel"))
	return rv
}

// Allows the “Download Selected” control to be renamed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1503822-downloadselectedcontrollabel?language=objc
func (c_ CameraDeviceView) SetDownloadSelectedControlLabel(value string) {
	objc.Call[objc.Void](c_, objc.Sel("setDownloadSelectedControlLabel:"), value)
}

// Specifies whether the downloads directory control should be displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1505006-displaysdownloadsdirectorycontro?language=objc
func (c_ CameraDeviceView) DisplaysDownloadsDirectoryControl() bool {
	rv := objc.Call[bool](c_, objc.Sel("displaysDownloadsDirectoryControl"))
	return rv
}

// Specifies whether the downloads directory control should be displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1505006-displaysdownloadsdirectorycontro?language=objc
func (c_ CameraDeviceView) SetDisplaysDownloadsDirectoryControl(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setDisplaysDownloadsDirectoryControl:"), value)
}

// Returns whether the selected items can be downloaded [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1504639-candownloadselecteditems?language=objc
func (c_ CameraDeviceView) CanDownloadSelectedItems() bool {
	rv := objc.Call[bool](c_, objc.Sel("canDownloadSelectedItems"))
	return rv
}

// Returns whether the selected items can be rotated left. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1503947-canrotateselecteditemsleft?language=objc
func (c_ CameraDeviceView) CanRotateSelectedItemsLeft() bool {
	rv := objc.Call[bool](c_, objc.Sel("canRotateSelectedItemsLeft"))
	return rv
}

// Displays whether the post process application control should be displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1503644-displayspostprocessapplicationco?language=objc
func (c_ CameraDeviceView) DisplaysPostProcessApplicationControl() bool {
	rv := objc.Call[bool](c_, objc.Sel("displaysPostProcessApplicationControl"))
	return rv
}

// Displays whether the post process application control should be displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1503644-displayspostprocessapplicationco?language=objc
func (c_ CameraDeviceView) SetDisplaysPostProcessApplicationControl(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setDisplaysPostProcessApplicationControl:"), value)
}

// Allows the “Download All” control to be renamed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1504281-downloadallcontrollabel?language=objc
func (c_ CameraDeviceView) DownloadAllControlLabel() string {
	rv := objc.Call[string](c_, objc.Sel("downloadAllControlLabel"))
	return rv
}

// Allows the “Download All” control to be renamed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1504281-downloadallcontrollabel?language=objc
func (c_ CameraDeviceView) SetDownloadAllControlLabel(value string) {
	objc.Call[objc.Void](c_, objc.Sel("setDownloadAllControlLabel:"), value)
}

// Specifies the icon size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1504002-iconsize?language=objc
func (c_ CameraDeviceView) IconSize() uint {
	rv := objc.Call[uint](c_, objc.Sel("iconSize"))
	return rv
}

// Specifies the icon size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1504002-iconsize?language=objc
func (c_ CameraDeviceView) SetIconSize(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setIconSize:"), value)
}

// Determines how the contents are saved by the delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1503771-transfermode?language=objc
func (c_ CameraDeviceView) TransferMode() CameraDeviceViewTransferMode {
	rv := objc.Call[CameraDeviceViewTransferMode](c_, objc.Sel("transferMode"))
	return rv
}

// Determines how the contents are saved by the delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1503771-transfermode?language=objc
func (c_ CameraDeviceView) SetTransferMode(value CameraDeviceViewTransferMode) {
	objc.Call[objc.Void](c_, objc.Sel("setTransferMode:"), value)
}

// The camera device view delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1504315-delegate?language=objc
func (c_ CameraDeviceView) Delegate() CameraDeviceViewDelegateWrapper {
	rv := objc.Call[CameraDeviceViewDelegateWrapper](c_, objc.Sel("delegate"))
	return rv
}

// The camera device view delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1504315-delegate?language=objc
func (c_ CameraDeviceView) SetDelegate(value PCameraDeviceViewDelegate) {
	po0 := objc.WrapAsProtocol("IKCameraDeviceViewDelegate", value)
	objc.Call[objc.Void](c_, objc.Sel("setDelegate:"), po0)
}

// The camera device view delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1504315-delegate?language=objc
func (c_ CameraDeviceView) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The URL of the application used to post process the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1504227-postprocessapplication?language=objc
func (c_ CameraDeviceView) PostProcessApplication() foundation.URL {
	rv := objc.Call[foundation.URL](c_, objc.Sel("postProcessApplication"))
	return rv
}

// The URL of the application used to post process the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1504227-postprocessapplication?language=objc
func (c_ CameraDeviceView) SetPostProcessApplication(value foundation.IURL) {
	objc.Call[objc.Void](c_, objc.Sel("setPostProcessApplication:"), objc.Ptr(value))
}

// Returns whether the device view is being displayed in icon mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1504417-hasdisplaymodeicon?language=objc
func (c_ CameraDeviceView) HasDisplayModeIcon() bool {
	rv := objc.Call[bool](c_, objc.Sel("hasDisplayModeIcon"))
	return rv
}

// Returns whether the device view is being displayed in icon mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1504417-hasdisplaymodeicon?language=objc
func (c_ CameraDeviceView) SetHasDisplayModeIcon(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setHasDisplayModeIcon:"), value)
}

// Specifies the display mode of the camera device view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1504881-mode?language=objc
func (c_ CameraDeviceView) Mode() CameraDeviceViewDisplayMode {
	rv := objc.Call[CameraDeviceViewDisplayMode](c_, objc.Sel("mode"))
	return rv
}

// Specifies the display mode of the camera device view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1504881-mode?language=objc
func (c_ CameraDeviceView) SetMode(value CameraDeviceViewDisplayMode) {
	objc.Call[objc.Void](c_, objc.Sel("setMode:"), value)
}

// Returns whether the selected items can be rotated right. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1503525-canrotateselecteditemsright?language=objc
func (c_ CameraDeviceView) CanRotateSelectedItemsRight() bool {
	rv := objc.Call[bool](c_, objc.Sel("canRotateSelectedItemsRight"))
	return rv
}

// Returns whether the selected items can be deleted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1504949-candeleteselecteditems?language=objc
func (c_ CameraDeviceView) CanDeleteSelectedItems() bool {
	rv := objc.Call[bool](c_, objc.Sel("canDeleteSelectedItems"))
	return rv
}

// Returns whether the device view is being displayed in table mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1503413-hasdisplaymodetable?language=objc
func (c_ CameraDeviceView) HasDisplayModeTable() bool {
	rv := objc.Call[bool](c_, objc.Sel("hasDisplayModeTable"))
	return rv
}

// Returns whether the device view is being displayed in table mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/1503413-hasdisplaymodetable?language=objc
func (c_ CameraDeviceView) SetHasDisplayModeTable(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setHasDisplayModeTable:"), value)
}
