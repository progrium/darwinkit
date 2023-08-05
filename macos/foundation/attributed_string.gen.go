// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var AttributedStringClass = _AttributedStringClass{objc.GetClass("NSAttributedString")}

type _AttributedStringClass struct {
	objc.Class
}

type IAttributedString interface {
	objc.IObject
	AttributesAtIndexEffectiveRange(location uint, range_ *Range) map[AttributedStringKey]objc.Object
	AttributesAtIndexLongestEffectiveRangeInRange(location uint, range_ *Range, rangeLimit Range) map[AttributedStringKey]objc.Object
	AttributeAtIndexEffectiveRange(attrName AttributedStringKey, location uint, range_ *Range) objc.Object
	AttributeAtIndexLongestEffectiveRangeInRange(attrName AttributedStringKey, location uint, range_ *Range, rangeLimit Range) objc.Object
	IsEqualToAttributedString(other IAttributedString) bool
	AttributedSubstringFromRange(range_ Range) AttributedString
	EnumerateAttributeInRangeOptionsUsingBlock(attrName AttributedStringKey, enumerationRange Range, opts AttributedStringEnumerationOptions, block func(value objc.Object, range_ Range, stop *bool))
	AttributedStringByInflectingString() AttributedString
	String() string
	Length() uint
}

type AttributedString struct {
	objc.Object
}

func MakeAttributedString(ptr unsafe.Pointer) AttributedString {
	return AttributedString{
		Object: objc.MakeObject(ptr),
	}
}

func (a_ AttributedString) InitWithString(str string) AttributedString {
	rv := objc.CallMethod[AttributedString](a_, objc.GetSelector("initWithString:"), str)
	return rv
}

func AttributedString_InitWithString(str string) AttributedString {
	return AttributedStringClass.Alloc().InitWithString(str)
}

func (a_ AttributedString) InitWithStringAttributes(str string, attrs map[AttributedStringKey]objc.IObject) AttributedString {
	rv := objc.CallMethod[AttributedString](a_, objc.GetSelector("initWithString:attributes:"), str, attrs)
	return rv
}

func AttributedString_InitWithStringAttributes(str string, attrs map[AttributedStringKey]objc.IObject) AttributedString {
	return AttributedStringClass.Alloc().InitWithStringAttributes(str, attrs)
}

func (a_ AttributedString) InitWithAttributedString(attrStr IAttributedString) AttributedString {
	rv := objc.CallMethod[AttributedString](a_, objc.GetSelector("initWithAttributedString:"), objc.ExtractPtr(attrStr))
	return rv
}

func AttributedString_InitWithAttributedString(attrStr IAttributedString) AttributedString {
	return AttributedStringClass.Alloc().InitWithAttributedString(attrStr)
}

func (ac _AttributedStringClass) Alloc() AttributedString {
	rv := objc.CallMethod[AttributedString](ac, objc.GetSelector("alloc"))
	return rv
}

func AttributedString_Alloc() AttributedString {
	return AttributedStringClass.Alloc()
}

func (ac _AttributedStringClass) New() AttributedString {
	rv := objc.CallMethod[AttributedString](ac, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewAttributedString() AttributedString {
	return AttributedStringClass.New()
}

func AttributedString_New() AttributedString {
	return AttributedStringClass.New()
}

func (a_ AttributedString) Init() AttributedString {
	rv := objc.CallMethod[AttributedString](a_, objc.GetSelector("init"))
	return rv
}

func AttributedString_Init() AttributedString {
	return AttributedStringClass.Alloc().Init()
}

func (a_ AttributedString) AttributesAtIndexEffectiveRange(location uint, range_ *Range) map[AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[AttributedStringKey]objc.Object](a_, objc.GetSelector("attributesAtIndex:effectiveRange:"), location, range_)
	return rv
}

func (a_ AttributedString) AttributesAtIndexLongestEffectiveRangeInRange(location uint, range_ *Range, rangeLimit Range) map[AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[AttributedStringKey]objc.Object](a_, objc.GetSelector("attributesAtIndex:longestEffectiveRange:inRange:"), location, range_, rangeLimit)
	return rv
}

func (a_ AttributedString) AttributeAtIndexEffectiveRange(attrName AttributedStringKey, location uint, range_ *Range) objc.Object {
	rv := objc.CallMethod[objc.Object](a_, objc.GetSelector("attribute:atIndex:effectiveRange:"), attrName, location, range_)
	return rv
}

func (a_ AttributedString) AttributeAtIndexLongestEffectiveRangeInRange(attrName AttributedStringKey, location uint, range_ *Range, rangeLimit Range) objc.Object {
	rv := objc.CallMethod[objc.Object](a_, objc.GetSelector("attribute:atIndex:longestEffectiveRange:inRange:"), attrName, location, range_, rangeLimit)
	return rv
}

func (a_ AttributedString) IsEqualToAttributedString(other IAttributedString) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("isEqualToAttributedString:"), objc.ExtractPtr(other))
	return rv
}

func (a_ AttributedString) AttributedSubstringFromRange(range_ Range) AttributedString {
	rv := objc.CallMethod[AttributedString](a_, objc.GetSelector("attributedSubstringFromRange:"), range_)
	return rv
}

func (a_ AttributedString) EnumerateAttributeInRangeOptionsUsingBlock(attrName AttributedStringKey, enumerationRange Range, opts AttributedStringEnumerationOptions, block func(value objc.Object, range_ Range, stop *bool)) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("enumerateAttribute:inRange:options:usingBlock:"), attrName, enumerationRange, opts, block)
}

func (a_ AttributedString) AttributedStringByInflectingString() AttributedString {
	rv := objc.CallMethod[AttributedString](a_, objc.GetSelector("attributedStringByInflectingString"))
	return rv
}

func (a_ AttributedString) String() string {
	rv := objc.CallMethod[string](a_, objc.GetSelector("string"))
	return rv
}

func (a_ AttributedString) Length() uint {
	rv := objc.CallMethod[uint](a_, objc.GetSelector("length"))
	return rv
}
