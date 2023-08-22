// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageEditPanel] class.
var ImageEditPanelClass = _ImageEditPanelClass{objc.GetClass("IKImageEditPanel")}

type _ImageEditPanelClass struct {
	objc.Class
}

// An interface definition for the [ImageEditPanel] class.
type IImageEditPanel interface {
	appkit.IPanel
	ReloadData()
	DataSource() ImageEditPanelDataSourceWrapper
	SetDataSource(value PImageEditPanelDataSource)
	SetDataSourceObject(valueObject objc.IObject)
	FilterArray() []objc.Object
}

// The IKImageEditPanel class provides a panel, that is, a utility window that floats on top of document windows, optimized for image editing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageeditpanel?language=objc
type ImageEditPanel struct {
	appkit.Panel
}

func ImageEditPanelFrom(ptr unsafe.Pointer) ImageEditPanel {
	return ImageEditPanel{
		Panel: appkit.PanelFrom(ptr),
	}
}

func (ic _ImageEditPanelClass) Alloc() ImageEditPanel {
	rv := objc.Call[ImageEditPanel](ic, objc.Sel("alloc"))
	return rv
}

func ImageEditPanel_Alloc() ImageEditPanel {
	return ImageEditPanelClass.Alloc()
}

func (ic _ImageEditPanelClass) New() ImageEditPanel {
	rv := objc.Call[ImageEditPanel](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageEditPanel() ImageEditPanel {
	return ImageEditPanelClass.New()
}

func (i_ ImageEditPanel) Init() ImageEditPanel {
	rv := objc.Call[ImageEditPanel](i_, objc.Sel("init"))
	return rv
}

func (ic _ImageEditPanelClass) WindowWithContentViewController(contentViewController appkit.IViewController) ImageEditPanel {
	rv := objc.Call[ImageEditPanel](ic, objc.Sel("windowWithContentViewController:"), objc.Ptr(contentViewController))
	return rv
}

// Creates a titled window that contains the specified content view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419551-windowwithcontentviewcontroller?language=objc
func ImageEditPanel_WindowWithContentViewController(contentViewController appkit.IViewController) ImageEditPanel {
	return ImageEditPanelClass.WindowWithContentViewController(contentViewController)
}

func (i_ ImageEditPanel) InitWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style appkit.WindowStyleMask, backingStoreType appkit.BackingStoreType, flag bool, screen appkit.IScreen) ImageEditPanel {
	rv := objc.Call[ImageEditPanel](i_, objc.Sel("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, objc.Ptr(screen))
	return rv
}

// Initializes an allocated window with the specified values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419755-initwithcontentrect?language=objc
func NewImageEditPanelWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style appkit.WindowStyleMask, backingStoreType appkit.BackingStoreType, flag bool, screen appkit.IScreen) ImageEditPanel {
	instance := ImageEditPanelClass.Alloc().InitWithContentRectStyleMaskBackingDeferScreen(contentRect, style, backingStoreType, flag, screen)
	instance.Autorelease()
	return instance
}

// Creates a shared instance of an image editing panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageeditpanel/1503827-sharedimageeditpanel?language=objc
func (ic _ImageEditPanelClass) SharedImageEditPanel() ImageEditPanel {
	rv := objc.Call[ImageEditPanel](ic, objc.Sel("sharedImageEditPanel"))
	return rv
}

// Creates a shared instance of an image editing panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageeditpanel/1503827-sharedimageeditpanel?language=objc
func ImageEditPanel_SharedImageEditPanel() ImageEditPanel {
	return ImageEditPanelClass.SharedImageEditPanel()
}

// Reloads the data from the data associated with an image editing panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageeditpanel/1503419-reloaddata?language=objc
func (i_ ImageEditPanel) ReloadData() {
	objc.Call[objc.Void](i_, objc.Sel("reloadData"))
}

// Specifies the edit panel’s dataSource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageeditpanel/1503831-datasource?language=objc
func (i_ ImageEditPanel) DataSource() ImageEditPanelDataSourceWrapper {
	rv := objc.Call[ImageEditPanelDataSourceWrapper](i_, objc.Sel("dataSource"))
	return rv
}

// Specifies the edit panel’s dataSource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageeditpanel/1503831-datasource?language=objc
func (i_ ImageEditPanel) SetDataSource(value PImageEditPanelDataSource) {
	po0 := objc.WrapAsProtocol("IKImageEditPanelDataSource", value)
	objc.Call[objc.Void](i_, objc.Sel("setDataSource:"), po0)
}

// Specifies the edit panel’s dataSource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageeditpanel/1503831-datasource?language=objc
func (i_ ImageEditPanel) SetDataSourceObject(valueObject objc.IObject) {
	objc.Call[objc.Void](i_, objc.Sel("setDataSource:"), objc.Ptr(valueObject))
}

// Returns the current array of user adjustments to effects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageeditpanel/1504436-filterarray?language=objc
func (i_ ImageEditPanel) FilterArray() []objc.Object {
	rv := objc.Call[[]objc.Object](i_, objc.Sel("filterArray"))
	return rv
}
