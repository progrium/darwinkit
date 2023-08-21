// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [LabeledValue] class.
var LabeledValueClass = _LabeledValueClass{objc.GetClass("CNLabeledValue")}

type _LabeledValueClass struct {
	objc.Class
}

// An interface definition for the [LabeledValue] class.
type ILabeledValue interface {
	objc.IObject
	Value() objc.Object
	Label() string
	Identifier() string
}

// An immutable object that combines a contact property value with a label that describes that property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabeledvalue?language=objc
type LabeledValue struct {
	objc.Object
}

func LabeledValueFrom(ptr unsafe.Pointer) LabeledValue {
	return LabeledValue{
		Object: objc.ObjectFrom(ptr),
	}
}

func (l_ LabeledValue) LabeledValueBySettingLabel(label string) LabeledValue {
	rv := objc.Call[LabeledValue](l_, objc.Sel("labeledValueBySettingLabel:"), label)
	return rv
}

// Returns a labeled value object with an existing value and identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabeledvalue/1403198-labeledvaluebysettinglabel?language=objc
func LabeledValue_LabeledValueBySettingLabel(label string) LabeledValue {
	instance := LabeledValueClass.Alloc().LabeledValueBySettingLabel(label)
	instance.Autorelease()
	return instance
}

func (l_ LabeledValue) LabeledValueBySettingValue(value objc.IObject) LabeledValue {
	rv := objc.Call[LabeledValue](l_, objc.Sel("labeledValueBySettingValue:"), objc.Ptr(value))
	return rv
}

// Returns a new value for an existing label and identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabeledvalue/1403084-labeledvaluebysettingvalue?language=objc
func LabeledValue_LabeledValueBySettingValue(value objc.IObject) LabeledValue {
	instance := LabeledValueClass.Alloc().LabeledValueBySettingValue(value)
	instance.Autorelease()
	return instance
}

func (lc _LabeledValueClass) LabeledValueWithLabelValue(label string, value objc.IObject) LabeledValue {
	rv := objc.Call[LabeledValue](lc, objc.Sel("labeledValueWithLabel:value:"), label, objc.Ptr(value))
	return rv
}

// Returns a new labeled value identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabeledvalue/1586531-labeledvaluewithlabel?language=objc
func LabeledValue_LabeledValueWithLabelValue(label string, value objc.IObject) LabeledValue {
	return LabeledValueClass.LabeledValueWithLabelValue(label, value)
}

func (l_ LabeledValue) InitWithLabelValue(label string, value objc.IObject) LabeledValue {
	rv := objc.Call[LabeledValue](l_, objc.Sel("initWithLabel:value:"), label, objc.Ptr(value))
	return rv
}

// Returns a new labeled value identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabeledvalue/1403404-initwithlabel?language=objc
func NewLabeledValueWithLabelValue(label string, value objc.IObject) LabeledValue {
	instance := LabeledValueClass.Alloc().InitWithLabelValue(label, value)
	instance.Autorelease()
	return instance
}

func (lc _LabeledValueClass) Alloc() LabeledValue {
	rv := objc.Call[LabeledValue](lc, objc.Sel("alloc"))
	return rv
}

func LabeledValue_Alloc() LabeledValue {
	return LabeledValueClass.Alloc()
}

func (lc _LabeledValueClass) New() LabeledValue {
	rv := objc.Call[LabeledValue](lc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewLabeledValue() LabeledValue {
	return LabeledValueClass.New()
}

func (l_ LabeledValue) Init() LabeledValue {
	rv := objc.Call[LabeledValue](l_, objc.Sel("init"))
	return rv
}

// Returns a localized string for the specified label. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabeledvalue/1403190-localizedstringforlabel?language=objc
func (lc _LabeledValueClass) LocalizedStringForLabel(label string) string {
	rv := objc.Call[string](lc, objc.Sel("localizedStringForLabel:"), label)
	return rv
}

// Returns a localized string for the specified label. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabeledvalue/1403190-localizedstringforlabel?language=objc
func LabeledValue_LocalizedStringForLabel(label string) string {
	return LabeledValueClass.LocalizedStringForLabel(label)
}

// A contact property value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabeledvalue/1403088-value?language=objc
func (l_ LabeledValue) Value() objc.Object {
	rv := objc.Call[objc.Object](l_, objc.Sel("value"))
	return rv
}

// The label for a contact property value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabeledvalue/1403392-label?language=objc
func (l_ LabeledValue) Label() string {
	rv := objc.Call[string](l_, objc.Sel("label"))
	return rv
}

// A unique identifier for the labeled value object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabeledvalue/1403408-identifier?language=objc
func (l_ LabeledValue) Identifier() string {
	rv := objc.Call[string](l_, objc.Sel("identifier"))
	return rv
}
