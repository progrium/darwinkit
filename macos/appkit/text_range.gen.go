// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var TextRangeClass = _TextRangeClass{objc.GetClass("NSTextRange")}

type _TextRangeClass struct {
	objc.Class
}

type ITextRange interface {
	objc.IObject
	IntersectsWithTextRange(textRange ITextRange) bool
	IsEqualToTextRange(textRange ITextRange) bool
	ContainsRange(textRange ITextRange) bool
	IsEmpty() bool
}

type TextRange struct {
	objc.Object
}

func MakeTextRange(ptr unsafe.Pointer) TextRange {
	return TextRange{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextRange) TextRangeByIntersectingWithTextRange(textRange ITextRange) TextRange {
	rv := objc.CallMethod[TextRange](t_, objc.GetSelector("textRangeByIntersectingWithTextRange:"), objc.ExtractPtr(textRange))
	return rv
}

func TextRange_TextRangeByIntersectingWithTextRange(textRange ITextRange) TextRange {
	return TextRangeClass.Alloc().TextRangeByIntersectingWithTextRange(textRange)
}

func (t_ TextRange) TextRangeByFormingUnionWithTextRange(textRange ITextRange) TextRange {
	rv := objc.CallMethod[TextRange](t_, objc.GetSelector("textRangeByFormingUnionWithTextRange:"), objc.ExtractPtr(textRange))
	return rv
}

func TextRange_TextRangeByFormingUnionWithTextRange(textRange ITextRange) TextRange {
	return TextRangeClass.Alloc().TextRangeByFormingUnionWithTextRange(textRange)
}

func (tc _TextRangeClass) Alloc() TextRange {
	rv := objc.CallMethod[TextRange](tc, objc.GetSelector("alloc"))
	return rv
}

func TextRange_Alloc() TextRange {
	return TextRangeClass.Alloc()
}

func (tc _TextRangeClass) New() TextRange {
	rv := objc.CallMethod[TextRange](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTextRange() TextRange {
	return TextRangeClass.New()
}

func TextRange_New() TextRange {
	return TextRangeClass.New()
}

func (t_ TextRange) Init() TextRange {
	rv := objc.CallMethod[TextRange](t_, objc.GetSelector("init"))
	return rv
}

func TextRange_Init() TextRange {
	return TextRangeClass.Alloc().Init()
}

func (t_ TextRange) IntersectsWithTextRange(textRange ITextRange) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("intersectsWithTextRange:"), objc.ExtractPtr(textRange))
	return rv
}

func (t_ TextRange) IsEqualToTextRange(textRange ITextRange) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isEqualToTextRange:"), objc.ExtractPtr(textRange))
	return rv
}

func (t_ TextRange) ContainsRange(textRange ITextRange) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("containsRange:"), objc.ExtractPtr(textRange))
	return rv
}

func (t_ TextRange) IsEmpty() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isEmpty"))
	return rv
}
