// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MediaLibraryBrowserController] class.
var MediaLibraryBrowserControllerClass = _MediaLibraryBrowserControllerClass{objc.GetClass("NSMediaLibraryBrowserController")}

type _MediaLibraryBrowserControllerClass struct {
	objc.Class
}

// An interface definition for the [MediaLibraryBrowserController] class.
type IMediaLibraryBrowserController interface {
	objc.IObject
	TogglePanel(sender objc.IObject) objc.Object
	IsVisible() bool
	SetVisible(value bool)
	Frame() foundation.Rect
	SetFrame(value foundation.Rect)
	MediaLibraries() MediaLibrary
	SetMediaLibraries(value MediaLibrary)
}

// An object that configures and displays a Media Library Browser panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmedialibrarybrowsercontroller?language=objc
type MediaLibraryBrowserController struct {
	objc.Object
}

func MediaLibraryBrowserControllerFrom(ptr unsafe.Pointer) MediaLibraryBrowserController {
	return MediaLibraryBrowserController{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _MediaLibraryBrowserControllerClass) Alloc() MediaLibraryBrowserController {
	rv := objc.Call[MediaLibraryBrowserController](mc, objc.Sel("alloc"))
	return rv
}

func MediaLibraryBrowserController_Alloc() MediaLibraryBrowserController {
	return MediaLibraryBrowserControllerClass.Alloc()
}

func (mc _MediaLibraryBrowserControllerClass) New() MediaLibraryBrowserController {
	rv := objc.Call[MediaLibraryBrowserController](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMediaLibraryBrowserController() MediaLibraryBrowserController {
	return MediaLibraryBrowserControllerClass.New()
}

func (m_ MediaLibraryBrowserController) Init() MediaLibraryBrowserController {
	rv := objc.Call[MediaLibraryBrowserController](m_, objc.Sel("init"))
	return rv
}

// Toggles the visibility of the Media Library Browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmedialibrarybrowsercontroller/1423479-togglepanel?language=objc
func (m_ MediaLibraryBrowserController) TogglePanel(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("togglePanel:"), sender)
	return rv
}

// A Boolean value that determines whether the Media Library Browser panel is visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmedialibrarybrowsercontroller/1423473-visible?language=objc
func (m_ MediaLibraryBrowserController) IsVisible() bool {
	rv := objc.Call[bool](m_, objc.Sel("isVisible"))
	return rv
}

// A Boolean value that determines whether the Media Library Browser panel is visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmedialibrarybrowsercontroller/1423473-visible?language=objc
func (m_ MediaLibraryBrowserController) SetVisible(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setVisible:"), value)
}

// Returns the shared Media Library Browser instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmedialibrarybrowsercontroller/1423485-sharedmedialibrarybrowsercontrol?language=objc
func (mc _MediaLibraryBrowserControllerClass) SharedMediaLibraryBrowserController() MediaLibraryBrowserController {
	rv := objc.Call[MediaLibraryBrowserController](mc, objc.Sel("sharedMediaLibraryBrowserController"))
	return rv
}

// Returns the shared Media Library Browser instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmedialibrarybrowsercontroller/1423485-sharedmedialibrarybrowsercontrol?language=objc
func MediaLibraryBrowserController_SharedMediaLibraryBrowserController() MediaLibraryBrowserController {
	return MediaLibraryBrowserControllerClass.SharedMediaLibraryBrowserController()
}

// The frame, in global coordinates, used to display the Media Library Browser panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmedialibrarybrowsercontroller/1423477-frame?language=objc
func (m_ MediaLibraryBrowserController) Frame() foundation.Rect {
	rv := objc.Call[foundation.Rect](m_, objc.Sel("frame"))
	return rv
}

// The frame, in global coordinates, used to display the Media Library Browser panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmedialibrarybrowsercontroller/1423477-frame?language=objc
func (m_ MediaLibraryBrowserController) SetFrame(value foundation.Rect) {
	objc.Call[objc.Void](m_, objc.Sel("setFrame:"), value)
}

// The media library that is in use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmedialibrarybrowsercontroller/1423481-medialibraries?language=objc
func (m_ MediaLibraryBrowserController) MediaLibraries() MediaLibrary {
	rv := objc.Call[MediaLibrary](m_, objc.Sel("mediaLibraries"))
	return rv
}

// The media library that is in use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmedialibrarybrowsercontroller/1423481-medialibraries?language=objc
func (m_ MediaLibraryBrowserController) SetMediaLibraries(value MediaLibrary) {
	objc.Call[objc.Void](m_, objc.Sel("setMediaLibraries:"), value)
}
