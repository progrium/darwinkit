// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [StatusItem] class.
var StatusItemClass = _StatusItemClass{objc.GetClass("NSStatusItem")}

type _StatusItemClass struct {
	objc.Class
}

// An interface definition for the [StatusItem] class.
type IStatusItem interface {
	objc.IObject
	IsVisible() bool
	SetVisible(value bool)
	Behavior() StatusItemBehavior
	SetBehavior(value StatusItemBehavior)
	StatusBar() StatusBar
	AutosaveName() StatusItemAutosaveName
	SetAutosaveName(value StatusItemAutosaveName)
	Length() float64
	SetLength(value float64)
	Menu() Menu
	SetMenu(value IMenu)
	Button() StatusBarButton
}

// An individual element displayed in the system menu bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem?language=objc
type StatusItem struct {
	objc.Object
}

func StatusItemFrom(ptr unsafe.Pointer) StatusItem {
	return StatusItem{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _StatusItemClass) Alloc() StatusItem {
	rv := objc.Call[StatusItem](sc, objc.Sel("alloc"))
	return rv
}

func StatusItem_Alloc() StatusItem {
	return StatusItemClass.Alloc()
}

func (sc _StatusItemClass) New() StatusItem {
	rv := objc.Call[StatusItem](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewStatusItem() StatusItem {
	return StatusItemClass.New()
}

func (s_ StatusItem) Init() StatusItem {
	rv := objc.Call[StatusItem](s_, objc.Sel("init"))
	return rv
}

// A Boolean value indicating if the menu bar currently displays the status item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/1644025-visible?language=objc
func (s_ StatusItem) IsVisible() bool {
	rv := objc.Call[bool](s_, objc.Sel("isVisible"))
	return rv
}

// A Boolean value indicating if the menu bar currently displays the status item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/1644025-visible?language=objc
func (s_ StatusItem) SetVisible(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setVisible:"), value)
}

// The set of allowed behaviors for the status item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/1644024-behavior?language=objc
func (s_ StatusItem) Behavior() StatusItemBehavior {
	rv := objc.Call[StatusItemBehavior](s_, objc.Sel("behavior"))
	return rv
}

// The set of allowed behaviors for the status item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/1644024-behavior?language=objc
func (s_ StatusItem) SetBehavior(value StatusItemBehavior) {
	objc.Call[objc.Void](s_, objc.Sel("setBehavior:"), value)
}

// The status bar that displays the status item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/1525951-statusbar?language=objc
func (s_ StatusItem) StatusBar() StatusBar {
	rv := objc.Call[StatusBar](s_, objc.Sel("statusBar"))
	return rv
}

// A unique name for saving and restoring information about a status item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/1644022-autosavename?language=objc
func (s_ StatusItem) AutosaveName() StatusItemAutosaveName {
	rv := objc.Call[StatusItemAutosaveName](s_, objc.Sel("autosaveName"))
	return rv
}

// A unique name for saving and restoring information about a status item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/1644022-autosavename?language=objc
func (s_ StatusItem) SetAutosaveName(value StatusItemAutosaveName) {
	objc.Call[objc.Void](s_, objc.Sel("setAutosaveName:"), value)
}

// The amount of space in the status bar that should be allocated to the status item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/1529402-length?language=objc
func (s_ StatusItem) Length() float64 {
	rv := objc.Call[float64](s_, objc.Sel("length"))
	return rv
}

// The amount of space in the status bar that should be allocated to the status item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/1529402-length?language=objc
func (s_ StatusItem) SetLength(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setLength:"), value)
}

// The pull-down menu displayed when the user clicks the status item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/1535918-menu?language=objc
func (s_ StatusItem) Menu() Menu {
	rv := objc.Call[Menu](s_, objc.Sel("menu"))
	return rv
}

// The pull-down menu displayed when the user clicks the status item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/1535918-menu?language=objc
func (s_ StatusItem) SetMenu(value IMenu) {
	objc.Call[objc.Void](s_, objc.Sel("setMenu:"), objc.Ptr(value))
}

// The button displayed in the status bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusitem/1535056-button?language=objc
func (s_ StatusItem) Button() StatusBarButton {
	rv := objc.Call[StatusBarButton](s_, objc.Sel("button"))
	return rv
}
