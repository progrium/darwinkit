// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NavigationAction] class.
var NavigationActionClass = _NavigationActionClass{objc.GetClass("WKNavigationAction")}

type _NavigationActionClass struct {
	objc.Class
}

// An interface definition for the [NavigationAction] class.
type INavigationAction interface {
	objc.IObject
	Request() foundation.URLRequest
	ModifierFlags() appkit.EventModifierFlags
	ButtonNumber() int
	SourceFrame() FrameInfo
	TargetFrame() FrameInfo
	ShouldPerformDownload() bool
	NavigationType() NavigationType
}

// An object that contains information about an action that causes navigation to occur. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction?language=objc
type NavigationAction struct {
	objc.Object
}

func NavigationActionFrom(ptr unsafe.Pointer) NavigationAction {
	return NavigationAction{
		Object: objc.ObjectFrom(ptr),
	}
}

func (nc _NavigationActionClass) Alloc() NavigationAction {
	rv := objc.Call[NavigationAction](nc, objc.Sel("alloc"))
	return rv
}

func NavigationAction_Alloc() NavigationAction {
	return NavigationActionClass.Alloc()
}

func (nc _NavigationActionClass) New() NavigationAction {
	rv := objc.Call[NavigationAction](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNavigationAction() NavigationAction {
	return NavigationActionClass.New()
}

func (n_ NavigationAction) Init() NavigationAction {
	rv := objc.Call[NavigationAction](n_, objc.Sel("init"))
	return rv
}

// The URL request object associated with the navigation action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/1401910-request?language=objc
func (n_ NavigationAction) Request() foundation.URLRequest {
	rv := objc.Call[foundation.URLRequest](n_, objc.Sel("request"))
	return rv
}

// The modifier keys that were pressed at the time of the navigation request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/1401934-modifierflags?language=objc
func (n_ NavigationAction) ModifierFlags() appkit.EventModifierFlags {
	rv := objc.Call[appkit.EventModifierFlags](n_, objc.Sel("modifierFlags"))
	return rv
}

// The number of the mouse button that caused the navigation request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/1401916-buttonnumber?language=objc
func (n_ NavigationAction) ButtonNumber() int {
	rv := objc.Call[int](n_, objc.Sel("buttonNumber"))
	return rv
}

// The frame that requested the navigation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/1401926-sourceframe?language=objc
func (n_ NavigationAction) SourceFrame() FrameInfo {
	rv := objc.Call[FrameInfo](n_, objc.Sel("sourceFrame"))
	return rv
}

// The frame in which to display the new content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/1401918-targetframe?language=objc
func (n_ NavigationAction) TargetFrame() FrameInfo {
	rv := objc.Call[FrameInfo](n_, objc.Sel("targetFrame"))
	return rv
}

// A Boolean value that indicates whether the web content provided an attribute that indicates a download. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/3727357-shouldperformdownload?language=objc
func (n_ NavigationAction) ShouldPerformDownload() bool {
	rv := objc.Call[bool](n_, objc.Sel("shouldPerformDownload"))
	return rv
}

// The type of action that triggered the navigation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/1401914-navigationtype?language=objc
func (n_ NavigationAction) NavigationType() NavigationType {
	rv := objc.Call[NavigationType](n_, objc.Sel("navigationType"))
	return rv
}
