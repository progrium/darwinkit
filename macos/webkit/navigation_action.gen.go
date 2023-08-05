// AUTO-GENERATED CODE, DO NOT MODIFY
package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var NavigationActionClass = _NavigationActionClass{objc.GetClass("WKNavigationAction")}

type _NavigationActionClass struct {
	objc.Class
}

type INavigationAction interface {
	objc.IObject
	NavigationType() NavigationType
	Request() foundation.URLRequest
	SourceFrame() FrameInfo
	TargetFrame() FrameInfo
	ButtonNumber() int
	ModifierFlags() appkit.EventModifierFlags
	ShouldPerformDownload() bool
}

type NavigationAction struct {
	objc.Object
}

func MakeNavigationAction(ptr unsafe.Pointer) NavigationAction {
	return NavigationAction{
		Object: objc.MakeObject(ptr),
	}
}

func (nc _NavigationActionClass) Alloc() NavigationAction {
	rv := objc.CallMethod[NavigationAction](nc, objc.GetSelector("alloc"))
	return rv
}

func NavigationAction_Alloc() NavigationAction {
	return NavigationActionClass.Alloc()
}

func (nc _NavigationActionClass) New() NavigationAction {
	rv := objc.CallMethod[NavigationAction](nc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewNavigationAction() NavigationAction {
	return NavigationActionClass.New()
}

func NavigationAction_New() NavigationAction {
	return NavigationActionClass.New()
}

func (n_ NavigationAction) Init() NavigationAction {
	rv := objc.CallMethod[NavigationAction](n_, objc.GetSelector("init"))
	return rv
}

func NavigationAction_Init() NavigationAction {
	return NavigationActionClass.Alloc().Init()
}

func (n_ NavigationAction) NavigationType() NavigationType {
	rv := objc.CallMethod[NavigationType](n_, objc.GetSelector("navigationType"))
	return rv
}

func (n_ NavigationAction) Request() foundation.URLRequest {
	rv := objc.CallMethod[foundation.URLRequest](n_, objc.GetSelector("request"))
	return rv
}

func (n_ NavigationAction) SourceFrame() FrameInfo {
	rv := objc.CallMethod[FrameInfo](n_, objc.GetSelector("sourceFrame"))
	return rv
}

func (n_ NavigationAction) TargetFrame() FrameInfo {
	rv := objc.CallMethod[FrameInfo](n_, objc.GetSelector("targetFrame"))
	return rv
}

func (n_ NavigationAction) ButtonNumber() int {
	rv := objc.CallMethod[int](n_, objc.GetSelector("buttonNumber"))
	return rv
}

func (n_ NavigationAction) ModifierFlags() appkit.EventModifierFlags {
	rv := objc.CallMethod[appkit.EventModifierFlags](n_, objc.GetSelector("modifierFlags"))
	return rv
}

func (n_ NavigationAction) ShouldPerformDownload() bool {
	rv := objc.CallMethod[bool](n_, objc.GetSelector("shouldPerformDownload"))
	return rv
}
