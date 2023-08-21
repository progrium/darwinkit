// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [OpenPanel] class.
var OpenPanelClass = _OpenPanelClass{objc.GetClass("NSOpenPanel")}

type _OpenPanelClass struct {
	objc.Class
}

// An interface definition for the [OpenPanel] class.
type IOpenPanel interface {
	ISavePanel
	IsAccessoryViewDisclosed() bool
	SetAccessoryViewDisclosed(value bool)
	ResolvesAliases() bool
	SetResolvesAliases(value bool)
	CanResolveUbiquitousConflicts() bool
	SetCanResolveUbiquitousConflicts(value bool)
	CanDownloadUbiquitousContents() bool
	SetCanDownloadUbiquitousContents(value bool)
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	URLs() []foundation.URL
	CanChooseFiles() bool
	SetCanChooseFiles(value bool)
	CanChooseDirectories() bool
	SetCanChooseDirectories(value bool)
}

// A panel that prompts the user to select a file to open. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel?language=objc
type OpenPanel struct {
	SavePanel
}

func OpenPanelFrom(ptr unsafe.Pointer) OpenPanel {
	return OpenPanel{
		SavePanel: SavePanelFrom(ptr),
	}
}

func (oc _OpenPanelClass) Alloc() OpenPanel {
	rv := objc.Call[OpenPanel](oc, objc.Sel("alloc"))
	return rv
}

func OpenPanel_Alloc() OpenPanel {
	return OpenPanelClass.Alloc()
}

func (oc _OpenPanelClass) New() OpenPanel {
	rv := objc.Call[OpenPanel](oc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewOpenPanel() OpenPanel {
	return OpenPanelClass.New()
}

func (o_ OpenPanel) Init() OpenPanel {
	rv := objc.Call[OpenPanel](o_, objc.Sel("init"))
	return rv
}

func (oc _OpenPanelClass) WindowWithContentViewController(contentViewController IViewController) OpenPanel {
	rv := objc.Call[OpenPanel](oc, objc.Sel("windowWithContentViewController:"), objc.Ptr(contentViewController))
	return rv
}

// Creates a titled window that contains the specified content view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419551-windowwithcontentviewcontroller?language=objc
func OpenPanel_WindowWithContentViewController(contentViewController IViewController) OpenPanel {
	return OpenPanelClass.WindowWithContentViewController(contentViewController)
}

func (o_ OpenPanel) InitWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) OpenPanel {
	rv := objc.Call[OpenPanel](o_, objc.Sel("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, objc.Ptr(screen))
	return rv
}

// Initializes an allocated window with the specified values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419755-initwithcontentrect?language=objc
func NewOpenPanelWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) OpenPanel {
	instance := OpenPanelClass.Alloc().InitWithContentRectStyleMaskBackingDeferScreen(contentRect, style, backingStoreType, flag, screen)
	instance.Autorelease()
	return instance
}

// Creates a new Open panel and initializes it with a default configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/1584365-openpanel?language=objc
func (oc _OpenPanelClass) OpenPanel() OpenPanel {
	rv := objc.Call[OpenPanel](oc, objc.Sel("openPanel"))
	return rv
}

// Creates a new Open panel and initializes it with a default configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/1584365-openpanel?language=objc
func OpenPanel_OpenPanel() OpenPanel {
	return OpenPanelClass.OpenPanel()
}

// A Boolean value that indicates whether the panel's accessory view is visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/1535036-accessoryviewdisclosed?language=objc
func (o_ OpenPanel) IsAccessoryViewDisclosed() bool {
	rv := objc.Call[bool](o_, objc.Sel("isAccessoryViewDisclosed"))
	return rv
}

// A Boolean value that indicates whether the panel's accessory view is visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/1535036-accessoryviewdisclosed?language=objc
func (o_ OpenPanel) SetAccessoryViewDisclosed(value bool) {
	objc.Call[objc.Void](o_, objc.Sel("setAccessoryViewDisclosed:"), value)
}

// A Boolean that indicates whether the panel resolves aliases. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/1533625-resolvesaliases?language=objc
func (o_ OpenPanel) ResolvesAliases() bool {
	rv := objc.Call[bool](o_, objc.Sel("resolvesAliases"))
	return rv
}

// A Boolean that indicates whether the panel resolves aliases. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/1533625-resolvesaliases?language=objc
func (o_ OpenPanel) SetResolvesAliases(value bool) {
	objc.Call[objc.Void](o_, objc.Sel("setResolvesAliases:"), value)
}

// A Boolean value that indicates how the panel responds to iCloud documents that have conflicting versions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/1533261-canresolveubiquitousconflicts?language=objc
func (o_ OpenPanel) CanResolveUbiquitousConflicts() bool {
	rv := objc.Call[bool](o_, objc.Sel("canResolveUbiquitousConflicts"))
	return rv
}

// A Boolean value that indicates how the panel responds to iCloud documents that have conflicting versions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/1533261-canresolveubiquitousconflicts?language=objc
func (o_ OpenPanel) SetCanResolveUbiquitousConflicts(value bool) {
	objc.Call[objc.Void](o_, objc.Sel("setCanResolveUbiquitousConflicts:"), value)
}

// A Boolean value that indicates how the panel responds to iCloud documents that aren't fully downloaded locally. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/1533418-candownloadubiquitouscontents?language=objc
func (o_ OpenPanel) CanDownloadUbiquitousContents() bool {
	rv := objc.Call[bool](o_, objc.Sel("canDownloadUbiquitousContents"))
	return rv
}

// A Boolean value that indicates how the panel responds to iCloud documents that aren't fully downloaded locally. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/1533418-candownloadubiquitouscontents?language=objc
func (o_ OpenPanel) SetCanDownloadUbiquitousContents(value bool) {
	objc.Call[objc.Void](o_, objc.Sel("setCanDownloadUbiquitousContents:"), value)
}

// A Boolean that indicates whether the user may select multiple files and directories. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/1530786-allowsmultipleselection?language=objc
func (o_ OpenPanel) AllowsMultipleSelection() bool {
	rv := objc.Call[bool](o_, objc.Sel("allowsMultipleSelection"))
	return rv
}

// A Boolean that indicates whether the user may select multiple files and directories. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/1530786-allowsmultipleselection?language=objc
func (o_ OpenPanel) SetAllowsMultipleSelection(value bool) {
	objc.Call[objc.Void](o_, objc.Sel("setAllowsMultipleSelection:"), value)
}

// An array of URLs, each of which contains the fully specified location of a selected file or directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/1529845-urls?language=objc
func (o_ OpenPanel) URLs() []foundation.URL {
	rv := objc.Call[[]foundation.URL](o_, objc.Sel("URLs"))
	return rv
}

// A Boolean that indicates whether the user can choose files in the panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/1527060-canchoosefiles?language=objc
func (o_ OpenPanel) CanChooseFiles() bool {
	rv := objc.Call[bool](o_, objc.Sel("canChooseFiles"))
	return rv
}

// A Boolean that indicates whether the user can choose files in the panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/1527060-canchoosefiles?language=objc
func (o_ OpenPanel) SetCanChooseFiles(value bool) {
	objc.Call[objc.Void](o_, objc.Sel("setCanChooseFiles:"), value)
}

// A Boolean that indicates whether the user can choose directories in the panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/1532668-canchoosedirectories?language=objc
func (o_ OpenPanel) CanChooseDirectories() bool {
	rv := objc.Call[bool](o_, objc.Sel("canChooseDirectories"))
	return rv
}

// A Boolean that indicates whether the user can choose directories in the panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/1532668-canchoosedirectories?language=objc
func (o_ OpenPanel) SetCanChooseDirectories(value bool) {
	objc.Call[objc.Void](o_, objc.Sel("setCanChooseDirectories:"), value)
}
