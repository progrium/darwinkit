// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var TextSelectionClass = _TextSelectionClass{objc.GetClass("NSTextSelection")}

type _TextSelectionClass struct {
	objc.Class
}

type ITextSelection interface {
	objc.IObject
	TextSelectionWithTextRanges(textRanges []ITextRange) TextSelection
	Affinity() TextSelectionAffinity
	AnchorPositionOffset() float64
	SetAnchorPositionOffset(value float64)
	Granularity() TextSelectionGranularity
	IsLogical() bool
	SetLogical(value bool)
	IsTransient() bool
	TextRanges() []TextRange
	TypingAttributes() map[foundation.AttributedStringKey]objc.Object
	SetTypingAttributes(value map[foundation.AttributedStringKey]objc.IObject)
}

type TextSelection struct {
	objc.Object
}

func MakeTextSelection(ptr unsafe.Pointer) TextSelection {
	return TextSelection{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextSelection) InitWithRangeAffinityGranularity(range_ ITextRange, affinity TextSelectionAffinity, granularity TextSelectionGranularity) TextSelection {
	rv := objc.CallMethod[TextSelection](t_, objc.GetSelector("initWithRange:affinity:granularity:"), objc.ExtractPtr(range_), affinity, granularity)
	return rv
}

func TextSelection_InitWithRangeAffinityGranularity(range_ ITextRange, affinity TextSelectionAffinity, granularity TextSelectionGranularity) TextSelection {
	return TextSelectionClass.Alloc().InitWithRangeAffinityGranularity(range_, affinity, granularity)
}

func (t_ TextSelection) InitWithRangesAffinityGranularity(textRanges []ITextRange, affinity TextSelectionAffinity, granularity TextSelectionGranularity) TextSelection {
	rv := objc.CallMethod[TextSelection](t_, objc.GetSelector("initWithRanges:affinity:granularity:"), textRanges, affinity, granularity)
	return rv
}

func TextSelection_InitWithRangesAffinityGranularity(textRanges []ITextRange, affinity TextSelectionAffinity, granularity TextSelectionGranularity) TextSelection {
	return TextSelectionClass.Alloc().InitWithRangesAffinityGranularity(textRanges, affinity, granularity)
}

func (tc _TextSelectionClass) Alloc() TextSelection {
	rv := objc.CallMethod[TextSelection](tc, objc.GetSelector("alloc"))
	return rv
}

func TextSelection_Alloc() TextSelection {
	return TextSelectionClass.Alloc()
}

func (tc _TextSelectionClass) New() TextSelection {
	rv := objc.CallMethod[TextSelection](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTextSelection() TextSelection {
	return TextSelectionClass.New()
}

func TextSelection_New() TextSelection {
	return TextSelectionClass.New()
}

func (t_ TextSelection) Init() TextSelection {
	rv := objc.CallMethod[TextSelection](t_, objc.GetSelector("init"))
	return rv
}

func TextSelection_Init() TextSelection {
	return TextSelectionClass.Alloc().Init()
}

func (t_ TextSelection) TextSelectionWithTextRanges(textRanges []ITextRange) TextSelection {
	rv := objc.CallMethod[TextSelection](t_, objc.GetSelector("textSelectionWithTextRanges:"), textRanges)
	return rv
}

func (t_ TextSelection) Affinity() TextSelectionAffinity {
	rv := objc.CallMethod[TextSelectionAffinity](t_, objc.GetSelector("affinity"))
	return rv
}

func (t_ TextSelection) AnchorPositionOffset() float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("anchorPositionOffset"))
	return rv
}

func (t_ TextSelection) SetAnchorPositionOffset(value float64) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAnchorPositionOffset:"), value)
}

func (t_ TextSelection) Granularity() TextSelectionGranularity {
	rv := objc.CallMethod[TextSelectionGranularity](t_, objc.GetSelector("granularity"))
	return rv
}

func (t_ TextSelection) IsLogical() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isLogical"))
	return rv
}

func (t_ TextSelection) SetLogical(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setLogical:"), value)
}

func (t_ TextSelection) IsTransient() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isTransient"))
	return rv
}

func (t_ TextSelection) TextRanges() []TextRange {
	rv := objc.CallMethod[[]TextRange](t_, objc.GetSelector("textRanges"))
	return rv
}

func (t_ TextSelection) TypingAttributes() map[foundation.AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[foundation.AttributedStringKey]objc.Object](t_, objc.GetSelector("typingAttributes"))
	return rv
}

func (t_ TextSelection) SetTypingAttributes(value map[foundation.AttributedStringKey]objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTypingAttributes:"), value)
}
