// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var ColorWellClass = _ColorWellClass{objc.GetClass("NSColorWell")}

type _ColorWellClass struct {
	objc.Class
}

type IColorWell interface {
	IControl
	TakeColorFrom(sender objc.IObject)
	Activate(exclusive bool)
	Deactivate()
	DrawWellInside(insideRect foundation.Rect)
	Color() Color
	SetColor(value IColor)
	ColorWellStyle() ColorWellStyle
	SetColorWellStyle(value ColorWellStyle)
	Image() Image
	SetImage(value IImage)
	IsActive() bool
	PulldownAction() objc.Selector
	SetPulldownAction(value objc.Selector)
	PulldownTarget() objc.Object
	SetPulldownTarget(value objc.IObject)
}

type ColorWell struct {
	Control
}

func MakeColorWell(ptr unsafe.Pointer) ColorWell {
	return ColorWell{
		Control: MakeControl(ptr),
	}
}

func (cc _ColorWellClass) ColorWellWithStyle(style ColorWellStyle) ColorWell {
	rv := objc.CallMethod[ColorWell](cc, objc.GetSelector("colorWellWithStyle:"), style)
	return rv
}

func ColorWell_ColorWellWithStyle(style ColorWellStyle) ColorWell {
	return ColorWellClass.ColorWellWithStyle(style)
}

func (c_ ColorWell) InitWithFrame(frameRect foundation.Rect) ColorWell {
	rv := objc.CallMethod[ColorWell](c_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func ColorWell_InitWithFrame(frameRect foundation.Rect) ColorWell {
	return ColorWellClass.Alloc().InitWithFrame(frameRect)
}

func (c_ ColorWell) Init() ColorWell {
	rv := objc.CallMethod[ColorWell](c_, objc.GetSelector("init"))
	return rv
}

func ColorWell_Init() ColorWell {
	return ColorWellClass.Alloc().Init()
}

func (cc _ColorWellClass) Alloc() ColorWell {
	rv := objc.CallMethod[ColorWell](cc, objc.GetSelector("alloc"))
	return rv
}

func ColorWell_Alloc() ColorWell {
	return ColorWellClass.Alloc()
}

func (cc _ColorWellClass) New() ColorWell {
	rv := objc.CallMethod[ColorWell](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewColorWell() ColorWell {
	return ColorWellClass.New()
}

func ColorWell_New() ColorWell {
	return ColorWellClass.New()
}

func (c_ ColorWell) TakeColorFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("takeColorFrom:"), objc.ExtractPtr(sender))
}

func (c_ ColorWell) Activate(exclusive bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("activate:"), exclusive)
}

func (c_ ColorWell) Deactivate() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("deactivate"))
}

func (c_ ColorWell) DrawWellInside(insideRect foundation.Rect) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("drawWellInside:"), insideRect)
}

func (c_ ColorWell) Color() Color {
	rv := objc.CallMethod[Color](c_, objc.GetSelector("color"))
	return rv
}

func (c_ ColorWell) SetColor(value IColor) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setColor:"), objc.ExtractPtr(value))
}

func (c_ ColorWell) ColorWellStyle() ColorWellStyle {
	rv := objc.CallMethod[ColorWellStyle](c_, objc.GetSelector("colorWellStyle"))
	return rv
}

func (c_ ColorWell) SetColorWellStyle(value ColorWellStyle) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setColorWellStyle:"), value)
}

func (c_ ColorWell) Image() Image {
	rv := objc.CallMethod[Image](c_, objc.GetSelector("image"))
	return rv
}

func (c_ ColorWell) SetImage(value IImage) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setImage:"), objc.ExtractPtr(value))
}

func (c_ ColorWell) IsActive() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isActive"))
	return rv
}

func (c_ ColorWell) PulldownAction() objc.Selector {
	rv := objc.CallMethod[objc.Selector](c_, objc.GetSelector("pulldownAction"))
	return rv
}

func (c_ ColorWell) SetPulldownAction(value objc.Selector) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setPulldownAction:"), value)
}

func (c_ ColorWell) PulldownTarget() objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.GetSelector("pulldownTarget"))
	return rv
}

func (c_ ColorWell) SetPulldownTarget(value objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setPulldownTarget:"), objc.ExtractPtr(value))
}
