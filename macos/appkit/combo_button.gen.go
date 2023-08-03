// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var ComboButtonClass = _ComboButtonClass{objc.GetClass("NSComboButton")}

type _ComboButtonClass struct {
	objc.Class
}

type IComboButton interface {
	IControl
	Style() ComboButtonStyle
	SetStyle(value ComboButtonStyle)
	Title() string
	SetTitle(value string)
	Image() Image
	SetImage(value IImage)
	ImageScaling() ImageScaling
	SetImageScaling(value ImageScaling)
}

type ComboButton struct {
	Control
}

func MakeComboButton(ptr unsafe.Pointer) ComboButton {
	return ComboButton{
		Control: MakeControl(ptr),
	}
}

func (cc _ComboButtonClass) ComboButtonWithTitleImageMenuTargetAction(title string, image IImage, menu IMenu, target objc.IObject, action objc.Selector) ComboButton {
	rv := objc.CallMethod[ComboButton](cc, objc.GetSelector("comboButtonWithTitle:image:menu:target:action:"), title, objc.ExtractPtr(image), objc.ExtractPtr(menu), objc.ExtractPtr(target), action)
	return rv
}

func ComboButton_ComboButtonWithTitleImageMenuTargetAction(title string, image IImage, menu IMenu, target objc.IObject, action objc.Selector) ComboButton {
	return ComboButtonClass.ComboButtonWithTitleImageMenuTargetAction(title, image, menu, target, action)
}

func (cc _ComboButtonClass) ComboButtonWithTitleMenuTargetAction(title string, menu IMenu, target objc.IObject, action objc.Selector) ComboButton {
	rv := objc.CallMethod[ComboButton](cc, objc.GetSelector("comboButtonWithTitle:menu:target:action:"), title, objc.ExtractPtr(menu), objc.ExtractPtr(target), action)
	return rv
}

func ComboButton_ComboButtonWithTitleMenuTargetAction(title string, menu IMenu, target objc.IObject, action objc.Selector) ComboButton {
	return ComboButtonClass.ComboButtonWithTitleMenuTargetAction(title, menu, target, action)
}

func (cc _ComboButtonClass) ComboButtonWithImageMenuTargetAction(image IImage, menu IMenu, target objc.IObject, action objc.Selector) ComboButton {
	rv := objc.CallMethod[ComboButton](cc, objc.GetSelector("comboButtonWithImage:menu:target:action:"), objc.ExtractPtr(image), objc.ExtractPtr(menu), objc.ExtractPtr(target), action)
	return rv
}

func ComboButton_ComboButtonWithImageMenuTargetAction(image IImage, menu IMenu, target objc.IObject, action objc.Selector) ComboButton {
	return ComboButtonClass.ComboButtonWithImageMenuTargetAction(image, menu, target, action)
}

func (c_ ComboButton) InitWithFrame(frameRect foundation.Rect) ComboButton {
	rv := objc.CallMethod[ComboButton](c_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func ComboButton_InitWithFrame(frameRect foundation.Rect) ComboButton {
	return ComboButtonClass.Alloc().InitWithFrame(frameRect)
}

func (c_ ComboButton) Init() ComboButton {
	rv := objc.CallMethod[ComboButton](c_, objc.GetSelector("init"))
	return rv
}

func ComboButton_Init() ComboButton {
	return ComboButtonClass.Alloc().Init()
}

func (cc _ComboButtonClass) Alloc() ComboButton {
	rv := objc.CallMethod[ComboButton](cc, objc.GetSelector("alloc"))
	return rv
}

func ComboButton_Alloc() ComboButton {
	return ComboButtonClass.Alloc()
}

func (cc _ComboButtonClass) New() ComboButton {
	rv := objc.CallMethod[ComboButton](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewComboButton() ComboButton {
	return ComboButtonClass.New()
}

func ComboButton_New() ComboButton {
	return ComboButtonClass.New()
}

func (c_ ComboButton) Style() ComboButtonStyle {
	rv := objc.CallMethod[ComboButtonStyle](c_, objc.GetSelector("style"))
	return rv
}

func (c_ ComboButton) SetStyle(value ComboButtonStyle) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setStyle:"), value)
}

func (c_ ComboButton) Title() string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("title"))
	return rv
}

func (c_ ComboButton) SetTitle(value string) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setTitle:"), value)
}

func (c_ ComboButton) Image() Image {
	rv := objc.CallMethod[Image](c_, objc.GetSelector("image"))
	return rv
}

func (c_ ComboButton) SetImage(value IImage) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setImage:"), objc.ExtractPtr(value))
}

func (c_ ComboButton) ImageScaling() ImageScaling {
	rv := objc.CallMethod[ImageScaling](c_, objc.GetSelector("imageScaling"))
	return rv
}

func (c_ ComboButton) SetImageScaling(value ImageScaling) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setImageScaling:"), value)
}
