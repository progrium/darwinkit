// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var FormatterClass = _FormatterClass{objc.GetClass("NSFormatter")}

type _FormatterClass struct {
	objc.Class
}

type IFormatter interface {
	objc.IObject
	StringForObjectValue(obj objc.IObject) string
	AttributedStringForObjectValueWithDefaultAttributes(obj objc.IObject, attrs map[AttributedStringKey]objc.IObject) AttributedString
	EditingStringForObjectValue(obj objc.IObject) string
	IsPartialStringValidNewEditingStringErrorDescription(partialString string, newString *String, error *String) bool
	IsPartialStringValidProposedSelectedRangeOriginalStringOriginalSelectedRangeErrorDescription(partialStringPtr *String, proposedSelRangePtr *Range, origString string, origSelRange Range, error *String) bool
}

type Formatter struct {
	objc.Object
}

func MakeFormatter(ptr unsafe.Pointer) Formatter {
	return Formatter{
		Object: objc.MakeObject(ptr),
	}
}

func (fc _FormatterClass) Alloc() Formatter {
	rv := objc.CallMethod[Formatter](fc, objc.GetSelector("alloc"))
	return rv
}

func Formatter_Alloc() Formatter {
	return FormatterClass.Alloc()
}

func (fc _FormatterClass) New() Formatter {
	rv := objc.CallMethod[Formatter](fc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewFormatter() Formatter {
	return FormatterClass.New()
}

func Formatter_New() Formatter {
	return FormatterClass.New()
}

func (f_ Formatter) Init() Formatter {
	rv := objc.CallMethod[Formatter](f_, objc.GetSelector("init"))
	return rv
}

func Formatter_Init() Formatter {
	return FormatterClass.Alloc().Init()
}

func (f_ Formatter) StringForObjectValue(obj objc.IObject) string {
	rv := objc.CallMethod[string](f_, objc.GetSelector("stringForObjectValue:"), objc.ExtractPtr(obj))
	return rv
}

func (f_ Formatter) AttributedStringForObjectValueWithDefaultAttributes(obj objc.IObject, attrs map[AttributedStringKey]objc.IObject) AttributedString {
	rv := objc.CallMethod[AttributedString](f_, objc.GetSelector("attributedStringForObjectValue:withDefaultAttributes:"), objc.ExtractPtr(obj), attrs)
	return rv
}

func (f_ Formatter) EditingStringForObjectValue(obj objc.IObject) string {
	rv := objc.CallMethod[string](f_, objc.GetSelector("editingStringForObjectValue:"), objc.ExtractPtr(obj))
	return rv
}

func (f_ Formatter) IsPartialStringValidNewEditingStringErrorDescription(partialString string, newString *String, error *String) bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("isPartialStringValid:newEditingString:errorDescription:"), partialString, newString, error)
	return rv
}

func (f_ Formatter) IsPartialStringValidProposedSelectedRangeOriginalStringOriginalSelectedRangeErrorDescription(partialStringPtr *String, proposedSelRangePtr *Range, origString string, origSelRange Range, error *String) bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("isPartialStringValid:proposedSelectedRange:originalString:originalSelectedRange:errorDescription:"), partialStringPtr, proposedSelRangePtr, origString, origSelRange, error)
	return rv
}
