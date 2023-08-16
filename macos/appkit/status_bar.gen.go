// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [StatusBar] class.
var StatusBarClass = _StatusBarClass{objc.GetClass("NSStatusBar")}

type _StatusBarClass struct {
	objc.Class
}

// An interface definition for the [StatusBar] class.
type IStatusBar interface {
	objc.IObject
	StatusItemWithLength(length float64) StatusItem
	RemoveStatusItem(item IStatusItem)
	IsVertical() bool
	Thickness() float64
}

// An object that manages a collection of status items displayed within the system-wide menu bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusbar?language=objc
type StatusBar struct {
	objc.Object
}

func StatusBarFrom(ptr unsafe.Pointer) StatusBar {
	return StatusBar{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _StatusBarClass) Alloc() StatusBar {
	rv := objc.Call[StatusBar](sc, objc.Sel("alloc"))
	return rv
}

func StatusBar_Alloc() StatusBar {
	return StatusBarClass.Alloc()
}

func (sc _StatusBarClass) New() StatusBar {
	rv := objc.Call[StatusBar](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewStatusBar() StatusBar {
	return StatusBarClass.New()
}

func (s_ StatusBar) Init() StatusBar {
	rv := objc.Call[StatusBar](s_, objc.Sel("init"))
	return rv
}

// Returns a newly created status item that has been allotted a specified space within the status bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusbar/1532895-statusitemwithlength?language=objc
func (s_ StatusBar) StatusItemWithLength(length float64) StatusItem {
	rv := objc.Call[StatusItem](s_, objc.Sel("statusItemWithLength:"), length)
	return rv
}

// Removes the specified status item from the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusbar/1530377-removestatusitem?language=objc
func (s_ StatusBar) RemoveStatusItem(item IStatusItem) {
	objc.Call[objc.Void](s_, objc.Sel("removeStatusItem:"), objc.Ptr(item))
}

// Returns the system-wide status bar located in the menu bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusbar/1530619-systemstatusbar?language=objc
func (sc _StatusBarClass) SystemStatusBar() StatusBar {
	rv := objc.Call[StatusBar](sc, objc.Sel("systemStatusBar"))
	return rv
}

// Returns the system-wide status bar located in the menu bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusbar/1530619-systemstatusbar?language=objc
func StatusBar_SystemStatusBar() StatusBar {
	return StatusBarClass.SystemStatusBar()
}

// A Boolean value indicating whether the status bar has a vertical orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusbar/1530580-vertical?language=objc
func (s_ StatusBar) IsVertical() bool {
	rv := objc.Call[bool](s_, objc.Sel("isVertical"))
	return rv
}

// The thickness of the status bar, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusbar/1534591-thickness?language=objc
func (s_ StatusBar) Thickness() float64 {
	rv := objc.Call[float64](s_, objc.Sel("thickness"))
	return rv
}
