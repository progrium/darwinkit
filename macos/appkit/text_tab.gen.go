// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var TextTabClass = _TextTabClass{objc.GetClass("NSTextTab")}

type _TextTabClass struct {
	objc.Class
}

type ITextTab interface {
	objc.IObject
	Location() float64
	Alignment() TextAlignment
	Options() map[TextTabOptionKey]objc.Object
}

type TextTab struct {
	objc.Object
}

func MakeTextTab(ptr unsafe.Pointer) TextTab {
	return TextTab{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextTab) InitWithTextAlignmentLocationOptions(alignment TextAlignment, loc float64, options map[TextTabOptionKey]objc.IObject) TextTab {
	rv := objc.CallMethod[TextTab](t_, objc.GetSelector("initWithTextAlignment:location:options:"), alignment, loc, options)
	return rv
}

func TextTab_InitWithTextAlignmentLocationOptions(alignment TextAlignment, loc float64, options map[TextTabOptionKey]objc.IObject) TextTab {
	return TextTabClass.Alloc().InitWithTextAlignmentLocationOptions(alignment, loc, options)
}

func (t_ TextTab) InitWithTypeLocation(type_ TextTabType, loc float64) TextTab {
	rv := objc.CallMethod[TextTab](t_, objc.GetSelector("initWithType:location:"), type_, loc)
	return rv
}

func TextTab_InitWithTypeLocation(type_ TextTabType, loc float64) TextTab {
	return TextTabClass.Alloc().InitWithTypeLocation(type_, loc)
}

func (tc _TextTabClass) Alloc() TextTab {
	rv := objc.CallMethod[TextTab](tc, objc.GetSelector("alloc"))
	return rv
}

func TextTab_Alloc() TextTab {
	return TextTabClass.Alloc()
}

func (tc _TextTabClass) New() TextTab {
	rv := objc.CallMethod[TextTab](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTextTab() TextTab {
	return TextTabClass.New()
}

func TextTab_New() TextTab {
	return TextTabClass.New()
}

func (t_ TextTab) Init() TextTab {
	rv := objc.CallMethod[TextTab](t_, objc.GetSelector("init"))
	return rv
}

func TextTab_Init() TextTab {
	return TextTabClass.Alloc().Init()
}

func (tc _TextTabClass) ColumnTerminatorsForLocale(aLocale foundation.ILocale) foundation.CharacterSet {
	rv := objc.CallMethod[foundation.CharacterSet](tc, objc.GetSelector("columnTerminatorsForLocale:"), objc.ExtractPtr(aLocale))
	return rv
}

func TextTab_ColumnTerminatorsForLocale(aLocale foundation.ILocale) foundation.CharacterSet {
	return TextTabClass.ColumnTerminatorsForLocale(aLocale)
}

func (t_ TextTab) Location() float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("location"))
	return rv
}

func (t_ TextTab) Alignment() TextAlignment {
	rv := objc.CallMethod[TextAlignment](t_, objc.GetSelector("alignment"))
	return rv
}

func (t_ TextTab) Options() map[TextTabOptionKey]objc.Object {
	rv := objc.CallMethod[map[TextTabOptionKey]objc.Object](t_, objc.GetSelector("options"))
	return rv
}
