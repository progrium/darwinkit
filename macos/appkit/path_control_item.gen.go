// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var PathControlItemClass = _PathControlItemClass{objc.GetClass("NSPathControlItem")}

type _PathControlItemClass struct {
	objc.Class
}

type IPathControlItem interface {
	objc.IObject
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.IAttributedString)
	Image() Image
	SetImage(value IImage)
	Title() string
	SetTitle(value string)
	URL() foundation.URL
}

type PathControlItem struct {
	objc.Object
}

func MakePathControlItem(ptr unsafe.Pointer) PathControlItem {
	return PathControlItem{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PathControlItemClass) Alloc() PathControlItem {
	rv := objc.CallMethod[PathControlItem](pc, objc.GetSelector("alloc"))
	return rv
}

func PathControlItem_Alloc() PathControlItem {
	return PathControlItemClass.Alloc()
}

func (pc _PathControlItemClass) New() PathControlItem {
	rv := objc.CallMethod[PathControlItem](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewPathControlItem() PathControlItem {
	return PathControlItemClass.New()
}

func PathControlItem_New() PathControlItem {
	return PathControlItemClass.New()
}

func (p_ PathControlItem) Init() PathControlItem {
	rv := objc.CallMethod[PathControlItem](p_, objc.GetSelector("init"))
	return rv
}

func PathControlItem_Init() PathControlItem {
	return PathControlItemClass.Alloc().Init()
}

func (p_ PathControlItem) AttributedTitle() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](p_, objc.GetSelector("attributedTitle"))
	return rv
}

func (p_ PathControlItem) SetAttributedTitle(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setAttributedTitle:"), objc.ExtractPtr(value))
}

func (p_ PathControlItem) Image() Image {
	rv := objc.CallMethod[Image](p_, objc.GetSelector("image"))
	return rv
}

func (p_ PathControlItem) SetImage(value IImage) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setImage:"), objc.ExtractPtr(value))
}

func (p_ PathControlItem) Title() string {
	rv := objc.CallMethod[string](p_, objc.GetSelector("title"))
	return rv
}

func (p_ PathControlItem) SetTitle(value string) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setTitle:"), value)
}

func (p_ PathControlItem) URL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](p_, objc.GetSelector("URL"))
	return rv
}
