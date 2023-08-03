// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var StatusItemClass = _StatusItemClass{objc.GetClass("NSStatusItem")}

type _StatusItemClass struct {
	objc.Class
}

type IStatusItem interface {
	objc.IObject
	StatusBar() StatusBar
	Behavior() StatusItemBehavior
	SetBehavior(value StatusItemBehavior)
	Button() StatusBarButton
	Menu() Menu
	SetMenu(value IMenu)
	IsVisible() bool
	SetVisible(value bool)
	Length() float64
	SetLength(value float64)
	AutosaveName() StatusItemAutosaveName
	SetAutosaveName(value StatusItemAutosaveName)
}

type StatusItem struct {
	objc.Object
}

func MakeStatusItem(ptr unsafe.Pointer) StatusItem {
	return StatusItem{
		Object: objc.MakeObject(ptr),
	}
}

func (sc _StatusItemClass) Alloc() StatusItem {
	rv := objc.CallMethod[StatusItem](sc, objc.GetSelector("alloc"))
	return rv
}

func StatusItem_Alloc() StatusItem {
	return StatusItemClass.Alloc()
}

func (sc _StatusItemClass) New() StatusItem {
	rv := objc.CallMethod[StatusItem](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewStatusItem() StatusItem {
	return StatusItemClass.New()
}

func StatusItem_New() StatusItem {
	return StatusItemClass.New()
}

func (s_ StatusItem) Init() StatusItem {
	rv := objc.CallMethod[StatusItem](s_, objc.GetSelector("init"))
	return rv
}

func StatusItem_Init() StatusItem {
	return StatusItemClass.Alloc().Init()
}

func (s_ StatusItem) StatusBar() StatusBar {
	rv := objc.CallMethod[StatusBar](s_, objc.GetSelector("statusBar"))
	return rv
}

func (s_ StatusItem) Behavior() StatusItemBehavior {
	rv := objc.CallMethod[StatusItemBehavior](s_, objc.GetSelector("behavior"))
	return rv
}

func (s_ StatusItem) SetBehavior(value StatusItemBehavior) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setBehavior:"), value)
}

func (s_ StatusItem) Button() StatusBarButton {
	rv := objc.CallMethod[StatusBarButton](s_, objc.GetSelector("button"))
	return rv
}

func (s_ StatusItem) Menu() Menu {
	rv := objc.CallMethod[Menu](s_, objc.GetSelector("menu"))
	return rv
}

func (s_ StatusItem) SetMenu(value IMenu) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setMenu:"), objc.ExtractPtr(value))
}

func (s_ StatusItem) IsVisible() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("isVisible"))
	return rv
}

func (s_ StatusItem) SetVisible(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setVisible:"), value)
}

func (s_ StatusItem) Length() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("length"))
	return rv
}

func (s_ StatusItem) SetLength(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setLength:"), value)
}

func (s_ StatusItem) AutosaveName() StatusItemAutosaveName {
	rv := objc.CallMethod[StatusItemAutosaveName](s_, objc.GetSelector("autosaveName"))
	return rv
}

func (s_ StatusItem) SetAutosaveName(value StatusItemAutosaveName) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setAutosaveName:"), value)
}
