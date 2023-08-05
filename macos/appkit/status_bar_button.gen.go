// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var StatusBarButtonClass = _StatusBarButtonClass{objc.GetClass("NSStatusBarButton")}

type _StatusBarButtonClass struct {
	objc.Class
}

type IStatusBarButton interface {
	IButton
	AppearsDisabled() bool
	SetAppearsDisabled(value bool)
}

type StatusBarButton struct {
	Button
}

func MakeStatusBarButton(ptr unsafe.Pointer) StatusBarButton {
	return StatusBarButton{
		Button: MakeButton(ptr),
	}
}

func (sc _StatusBarButtonClass) CheckboxWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) StatusBarButton {
	rv := objc.CallMethod[StatusBarButton](sc, objc.GetSelector("checkboxWithTitle:target:action:"), title, objc.ExtractPtr(target), action)
	return rv
}

func StatusBarButton_CheckboxWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) StatusBarButton {
	return StatusBarButtonClass.CheckboxWithTitleTargetAction(title, target, action)
}

func (sc _StatusBarButtonClass) ButtonWithImageTargetAction(image IImage, target objc.IObject, action objc.Selector) StatusBarButton {
	rv := objc.CallMethod[StatusBarButton](sc, objc.GetSelector("buttonWithImage:target:action:"), objc.ExtractPtr(image), objc.ExtractPtr(target), action)
	return rv
}

func StatusBarButton_ButtonWithImageTargetAction(image IImage, target objc.IObject, action objc.Selector) StatusBarButton {
	return StatusBarButtonClass.ButtonWithImageTargetAction(image, target, action)
}

func (sc _StatusBarButtonClass) RadioButtonWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) StatusBarButton {
	rv := objc.CallMethod[StatusBarButton](sc, objc.GetSelector("radioButtonWithTitle:target:action:"), title, objc.ExtractPtr(target), action)
	return rv
}

func StatusBarButton_RadioButtonWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) StatusBarButton {
	return StatusBarButtonClass.RadioButtonWithTitleTargetAction(title, target, action)
}

func (sc _StatusBarButtonClass) ButtonWithTitleImageTargetAction(title string, image IImage, target objc.IObject, action objc.Selector) StatusBarButton {
	rv := objc.CallMethod[StatusBarButton](sc, objc.GetSelector("buttonWithTitle:image:target:action:"), title, objc.ExtractPtr(image), objc.ExtractPtr(target), action)
	return rv
}

func StatusBarButton_ButtonWithTitleImageTargetAction(title string, image IImage, target objc.IObject, action objc.Selector) StatusBarButton {
	return StatusBarButtonClass.ButtonWithTitleImageTargetAction(title, image, target, action)
}

func (sc _StatusBarButtonClass) ButtonWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) StatusBarButton {
	rv := objc.CallMethod[StatusBarButton](sc, objc.GetSelector("buttonWithTitle:target:action:"), title, objc.ExtractPtr(target), action)
	return rv
}

func StatusBarButton_ButtonWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) StatusBarButton {
	return StatusBarButtonClass.ButtonWithTitleTargetAction(title, target, action)
}

func (s_ StatusBarButton) InitWithFrame(frameRect foundation.Rect) StatusBarButton {
	rv := objc.CallMethod[StatusBarButton](s_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func StatusBarButton_InitWithFrame(frameRect foundation.Rect) StatusBarButton {
	return StatusBarButtonClass.Alloc().InitWithFrame(frameRect)
}

func (s_ StatusBarButton) Init() StatusBarButton {
	rv := objc.CallMethod[StatusBarButton](s_, objc.GetSelector("init"))
	return rv
}

func StatusBarButton_Init() StatusBarButton {
	return StatusBarButtonClass.Alloc().Init()
}

func (sc _StatusBarButtonClass) Alloc() StatusBarButton {
	rv := objc.CallMethod[StatusBarButton](sc, objc.GetSelector("alloc"))
	return rv
}

func StatusBarButton_Alloc() StatusBarButton {
	return StatusBarButtonClass.Alloc()
}

func (sc _StatusBarButtonClass) New() StatusBarButton {
	rv := objc.CallMethod[StatusBarButton](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewStatusBarButton() StatusBarButton {
	return StatusBarButtonClass.New()
}

func StatusBarButton_New() StatusBarButton {
	return StatusBarButtonClass.New()
}

func (s_ StatusBarButton) AppearsDisabled() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("appearsDisabled"))
	return rv
}

func (s_ StatusBarButton) SetAppearsDisabled(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setAppearsDisabled:"), value)
}
