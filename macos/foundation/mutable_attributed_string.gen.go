// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var MutableAttributedStringClass = _MutableAttributedStringClass{objc.GetClass("NSMutableAttributedString")}

type _MutableAttributedStringClass struct {
	objc.Class
}

type IMutableAttributedString interface {
	IAttributedString
	ReplaceCharactersInRangeWithString(range_ Range, str string)
	DeleteCharactersInRange(range_ Range)
	SetAttributesRange(attrs map[AttributedStringKey]objc.IObject, range_ Range)
	AddAttributeValueRange(name AttributedStringKey, value objc.IObject, range_ Range)
	AddAttributesRange(attrs map[AttributedStringKey]objc.IObject, range_ Range)
	RemoveAttributeRange(name AttributedStringKey, range_ Range)
	AppendAttributedString(attrString IAttributedString)
	InsertAttributedStringAtIndex(attrString IAttributedString, loc uint)
	ReplaceCharactersInRangeWithAttributedString(range_ Range, attrString IAttributedString)
	SetAttributedString(attrString IAttributedString)
	BeginEditing()
	EndEditing()
}

type MutableAttributedString struct {
	AttributedString
}

func MakeMutableAttributedString(ptr unsafe.Pointer) MutableAttributedString {
	return MutableAttributedString{
		AttributedString: MakeAttributedString(ptr),
	}
}

func (m_ MutableAttributedString) InitWithString(str string) MutableAttributedString {
	rv := objc.CallMethod[MutableAttributedString](m_, objc.GetSelector("initWithString:"), str)
	return rv
}

func MutableAttributedString_InitWithString(str string) MutableAttributedString {
	return MutableAttributedStringClass.Alloc().InitWithString(str)
}

func (m_ MutableAttributedString) InitWithStringAttributes(str string, attrs map[AttributedStringKey]objc.IObject) MutableAttributedString {
	rv := objc.CallMethod[MutableAttributedString](m_, objc.GetSelector("initWithString:attributes:"), str, attrs)
	return rv
}

func MutableAttributedString_InitWithStringAttributes(str string, attrs map[AttributedStringKey]objc.IObject) MutableAttributedString {
	return MutableAttributedStringClass.Alloc().InitWithStringAttributes(str, attrs)
}

func (m_ MutableAttributedString) InitWithAttributedString(attrStr IAttributedString) MutableAttributedString {
	rv := objc.CallMethod[MutableAttributedString](m_, objc.GetSelector("initWithAttributedString:"), objc.ExtractPtr(attrStr))
	return rv
}

func MutableAttributedString_InitWithAttributedString(attrStr IAttributedString) MutableAttributedString {
	return MutableAttributedStringClass.Alloc().InitWithAttributedString(attrStr)
}

func (mc _MutableAttributedStringClass) Alloc() MutableAttributedString {
	rv := objc.CallMethod[MutableAttributedString](mc, objc.GetSelector("alloc"))
	return rv
}

func MutableAttributedString_Alloc() MutableAttributedString {
	return MutableAttributedStringClass.Alloc()
}

func (mc _MutableAttributedStringClass) New() MutableAttributedString {
	rv := objc.CallMethod[MutableAttributedString](mc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewMutableAttributedString() MutableAttributedString {
	return MutableAttributedStringClass.New()
}

func MutableAttributedString_New() MutableAttributedString {
	return MutableAttributedStringClass.New()
}

func (m_ MutableAttributedString) Init() MutableAttributedString {
	rv := objc.CallMethod[MutableAttributedString](m_, objc.GetSelector("init"))
	return rv
}

func MutableAttributedString_Init() MutableAttributedString {
	return MutableAttributedStringClass.Alloc().Init()
}

func (m_ MutableAttributedString) ReplaceCharactersInRangeWithString(range_ Range, str string) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("replaceCharactersInRange:withString:"), range_, str)
}

func (m_ MutableAttributedString) DeleteCharactersInRange(range_ Range) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("deleteCharactersInRange:"), range_)
}

func (m_ MutableAttributedString) SetAttributesRange(attrs map[AttributedStringKey]objc.IObject, range_ Range) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAttributes:range:"), attrs, range_)
}

func (m_ MutableAttributedString) AddAttributeValueRange(name AttributedStringKey, value objc.IObject, range_ Range) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("addAttribute:value:range:"), name, objc.ExtractPtr(value), range_)
}

func (m_ MutableAttributedString) AddAttributesRange(attrs map[AttributedStringKey]objc.IObject, range_ Range) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("addAttributes:range:"), attrs, range_)
}

func (m_ MutableAttributedString) RemoveAttributeRange(name AttributedStringKey, range_ Range) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("removeAttribute:range:"), name, range_)
}

func (m_ MutableAttributedString) AppendAttributedString(attrString IAttributedString) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("appendAttributedString:"), objc.ExtractPtr(attrString))
}

func (m_ MutableAttributedString) InsertAttributedStringAtIndex(attrString IAttributedString, loc uint) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("insertAttributedString:atIndex:"), objc.ExtractPtr(attrString), loc)
}

func (m_ MutableAttributedString) ReplaceCharactersInRangeWithAttributedString(range_ Range, attrString IAttributedString) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("replaceCharactersInRange:withAttributedString:"), range_, objc.ExtractPtr(attrString))
}

func (m_ MutableAttributedString) SetAttributedString(attrString IAttributedString) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAttributedString:"), objc.ExtractPtr(attrString))
}

func (m_ MutableAttributedString) BeginEditing() {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("beginEditing"))
}

func (m_ MutableAttributedString) EndEditing() {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("endEditing"))
}
