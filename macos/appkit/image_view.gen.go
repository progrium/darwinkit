// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var ImageViewClass = _ImageViewClass{objc.GetClass("NSImageView")}

type _ImageViewClass struct {
	objc.Class
}

type IImageView interface {
	IControl
	Image() Image
	SetImage(value IImage)
	ImageFrameStyle() ImageFrameStyle
	SetImageFrameStyle(value ImageFrameStyle)
	ImageAlignment() ImageAlignment
	SetImageAlignment(value ImageAlignment)
	ImageScaling() ImageScaling
	SetImageScaling(value ImageScaling)
	Animates() bool
	SetAnimates(value bool)
	IsEditable() bool
	SetEditable(value bool)
	AllowsCutCopyPaste() bool
	SetAllowsCutCopyPaste(value bool)
	ContentTintColor() Color
	SetContentTintColor(value IColor)
	SymbolConfiguration() ImageSymbolConfiguration
	SetSymbolConfiguration(value IImageSymbolConfiguration)
}

type ImageView struct {
	Control
}

func MakeImageView(ptr unsafe.Pointer) ImageView {
	return ImageView{
		Control: MakeControl(ptr),
	}
}

func (ic _ImageViewClass) ImageViewWithImage(image IImage) ImageView {
	rv := objc.CallMethod[ImageView](ic, objc.GetSelector("imageViewWithImage:"), objc.ExtractPtr(image))
	return rv
}

func ImageView_ImageViewWithImage(image IImage) ImageView {
	return ImageViewClass.ImageViewWithImage(image)
}

func (i_ ImageView) InitWithFrame(frameRect foundation.Rect) ImageView {
	rv := objc.CallMethod[ImageView](i_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func ImageView_InitWithFrame(frameRect foundation.Rect) ImageView {
	return ImageViewClass.Alloc().InitWithFrame(frameRect)
}

func (i_ ImageView) Init() ImageView {
	rv := objc.CallMethod[ImageView](i_, objc.GetSelector("init"))
	return rv
}

func ImageView_Init() ImageView {
	return ImageViewClass.Alloc().Init()
}

func (ic _ImageViewClass) Alloc() ImageView {
	rv := objc.CallMethod[ImageView](ic, objc.GetSelector("alloc"))
	return rv
}

func ImageView_Alloc() ImageView {
	return ImageViewClass.Alloc()
}

func (ic _ImageViewClass) New() ImageView {
	rv := objc.CallMethod[ImageView](ic, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewImageView() ImageView {
	return ImageViewClass.New()
}

func ImageView_New() ImageView {
	return ImageViewClass.New()
}

func (i_ ImageView) Image() Image {
	rv := objc.CallMethod[Image](i_, objc.GetSelector("image"))
	return rv
}

func (i_ ImageView) SetImage(value IImage) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setImage:"), objc.ExtractPtr(value))
}

func (i_ ImageView) ImageFrameStyle() ImageFrameStyle {
	rv := objc.CallMethod[ImageFrameStyle](i_, objc.GetSelector("imageFrameStyle"))
	return rv
}

func (i_ ImageView) SetImageFrameStyle(value ImageFrameStyle) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setImageFrameStyle:"), value)
}

func (i_ ImageView) ImageAlignment() ImageAlignment {
	rv := objc.CallMethod[ImageAlignment](i_, objc.GetSelector("imageAlignment"))
	return rv
}

func (i_ ImageView) SetImageAlignment(value ImageAlignment) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setImageAlignment:"), value)
}

func (i_ ImageView) ImageScaling() ImageScaling {
	rv := objc.CallMethod[ImageScaling](i_, objc.GetSelector("imageScaling"))
	return rv
}

func (i_ ImageView) SetImageScaling(value ImageScaling) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setImageScaling:"), value)
}

func (i_ ImageView) Animates() bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("animates"))
	return rv
}

func (i_ ImageView) SetAnimates(value bool) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setAnimates:"), value)
}

func (i_ ImageView) IsEditable() bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("isEditable"))
	return rv
}

func (i_ ImageView) SetEditable(value bool) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setEditable:"), value)
}

func (i_ ImageView) AllowsCutCopyPaste() bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("allowsCutCopyPaste"))
	return rv
}

func (i_ ImageView) SetAllowsCutCopyPaste(value bool) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setAllowsCutCopyPaste:"), value)
}

func (i_ ImageView) ContentTintColor() Color {
	rv := objc.CallMethod[Color](i_, objc.GetSelector("contentTintColor"))
	return rv
}

func (i_ ImageView) SetContentTintColor(value IColor) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setContentTintColor:"), objc.ExtractPtr(value))
}

func (i_ ImageView) SymbolConfiguration() ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](i_, objc.GetSelector("symbolConfiguration"))
	return rv
}

func (i_ ImageView) SetSymbolConfiguration(value IImageSymbolConfiguration) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setSymbolConfiguration:"), objc.ExtractPtr(value))
}
