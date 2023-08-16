// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PathControlItem] class.
var PathControlItemClass = _PathControlItemClass{objc.GetClass("NSPathControlItem")}

type _PathControlItemClass struct {
	objc.Class
}

// An interface definition for the [PathControlItem] class.
type IPathControlItem interface {
	objc.IObject
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.IAttributedString)
	URL() foundation.URL
	Title() string
	SetTitle(value string)
	Image() Image
	SetImage(value IImage)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrolitem?language=objc
type PathControlItem struct {
	objc.Object
}

func PathControlItemFrom(ptr unsafe.Pointer) PathControlItem {
	return PathControlItem{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PathControlItemClass) Alloc() PathControlItem {
	rv := objc.Call[PathControlItem](pc, objc.Sel("alloc"))
	return rv
}

func PathControlItem_Alloc() PathControlItem {
	return PathControlItemClass.Alloc()
}

func (pc _PathControlItemClass) New() PathControlItem {
	rv := objc.Call[PathControlItem](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPathControlItem() PathControlItem {
	return PathControlItemClass.New()
}

func (p_ PathControlItem) Init() PathControlItem {
	rv := objc.Call[PathControlItem](p_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrolitem/1388287-attributedtitle?language=objc
func (p_ PathControlItem) AttributedTitle() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](p_, objc.Sel("attributedTitle"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrolitem/1388287-attributedtitle?language=objc
func (p_ PathControlItem) SetAttributedTitle(value foundation.IAttributedString) {
	objc.Call[objc.Void](p_, objc.Sel("setAttributedTitle:"), objc.Ptr(value))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrolitem/1388289-url?language=objc
func (p_ PathControlItem) URL() foundation.URL {
	rv := objc.Call[foundation.URL](p_, objc.Sel("URL"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrolitem/1388293-title?language=objc
func (p_ PathControlItem) Title() string {
	rv := objc.Call[string](p_, objc.Sel("title"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrolitem/1388293-title?language=objc
func (p_ PathControlItem) SetTitle(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setTitle:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrolitem/1388295-image?language=objc
func (p_ PathControlItem) Image() Image {
	rv := objc.Call[Image](p_, objc.Sel("image"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrolitem/1388295-image?language=objc
func (p_ PathControlItem) SetImage(value IImage) {
	objc.Call[objc.Void](p_, objc.Sel("setImage:"), objc.Ptr(value))
}
