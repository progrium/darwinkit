// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextTab] class.
var TextTabClass = _TextTabClass{objc.GetClass("NSTextTab")}

type _TextTabClass struct {
	objc.Class
}

// An interface definition for the [TextTab] class.
type ITextTab interface {
	objc.IObject
	Options() map[TextTabOptionKey]objc.Object
	Location() float64
	Alignment() TextAlignment
}

// A tab in a paragraph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstexttab?language=objc
type TextTab struct {
	objc.Object
}

func TextTabFrom(ptr unsafe.Pointer) TextTab {
	return TextTab{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TextTab) InitWithTextAlignmentLocationOptions(alignment TextAlignment, loc float64, options map[TextTabOptionKey]objc.IObject) TextTab {
	rv := objc.Call[TextTab](t_, objc.Sel("initWithTextAlignment:location:options:"), alignment, loc, options)
	return rv
}

// Initializes a text tab with the specified text alignment, location, and options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstexttab/1526080-initwithtextalignment?language=objc
func NewTextTabWithTextAlignmentLocationOptions(alignment TextAlignment, loc float64, options map[TextTabOptionKey]objc.IObject) TextTab {
	instance := TextTabClass.Alloc().InitWithTextAlignmentLocationOptions(alignment, loc, options)
	instance.Autorelease()
	return instance
}

func (tc _TextTabClass) Alloc() TextTab {
	rv := objc.Call[TextTab](tc, objc.Sel("alloc"))
	return rv
}

func TextTab_Alloc() TextTab {
	return TextTabClass.Alloc()
}

func (tc _TextTabClass) New() TextTab {
	rv := objc.Call[TextTab](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextTab() TextTab {
	return TextTabClass.New()
}

func (t_ TextTab) Init() TextTab {
	rv := objc.Call[TextTab](t_, objc.Sel("init"))
	return rv
}

// Returns the column terminators for the specified locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstexttab/1535107-columnterminatorsforlocale?language=objc
func (tc _TextTabClass) ColumnTerminatorsForLocale(aLocale foundation.ILocale) foundation.CharacterSet {
	rv := objc.Call[foundation.CharacterSet](tc, objc.Sel("columnTerminatorsForLocale:"), objc.Ptr(aLocale))
	return rv
}

// Returns the column terminators for the specified locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstexttab/1535107-columnterminatorsforlocale?language=objc
func TextTab_ColumnTerminatorsForLocale(aLocale foundation.ILocale) foundation.CharacterSet {
	return TextTabClass.ColumnTerminatorsForLocale(aLocale)
}

// The dictionary of attributes for the text tab. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstexttab/1534965-options?language=objc
func (t_ TextTab) Options() map[TextTabOptionKey]objc.Object {
	rv := objc.Call[map[TextTabOptionKey]objc.Object](t_, objc.Sel("options"))
	return rv
}

// The text tab’s ruler location relative to the back margin. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstexttab/1527968-location?language=objc
func (t_ TextTab) Location() float64 {
	rv := objc.Call[float64](t_, objc.Sel("location"))
	return rv
}

// The text alignment of the text tab. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstexttab/1527212-alignment?language=objc
func (t_ TextTab) Alignment() TextAlignment {
	rv := objc.Call[TextAlignment](t_, objc.Sel("alignment"))
	return rv
}
