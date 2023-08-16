// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Formatter] class.
var FormatterClass = _FormatterClass{objc.GetClass("NSFormatter")}

type _FormatterClass struct {
	objc.Class
}

// An interface definition for the [Formatter] class.
type IFormatter interface {
	objc.IObject
	EditingStringForObjectValue(obj objc.IObject) string
	StringForObjectValue(obj objc.IObject) string
	GetObjectValueForStringErrorDescription(obj objc.IObject, string_ string, error string) bool
	IsPartialStringValidNewEditingStringErrorDescription(partialString string, newString string, error string) bool
	AttributedStringForObjectValueWithDefaultAttributes(obj objc.IObject, attrs map[AttributedStringKey]objc.IObject) AttributedString
}

// An abstract class that declares an interface for objects that create, interpret, and validate the textual representation of values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsformatter?language=objc
type Formatter struct {
	objc.Object
}

func FormatterFrom(ptr unsafe.Pointer) Formatter {
	return Formatter{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FormatterClass) Alloc() Formatter {
	rv := objc.Call[Formatter](fc, objc.Sel("alloc"))
	return rv
}

func Formatter_Alloc() Formatter {
	return FormatterClass.Alloc()
}

func (fc _FormatterClass) New() Formatter {
	rv := objc.Call[Formatter](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFormatter() Formatter {
	return FormatterClass.New()
}

func (f_ Formatter) Init() Formatter {
	rv := objc.Call[Formatter](f_, objc.Sel("init"))
	return rv
}

// The default implementation of this method invokes stringForObjectValue:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsformatter/1416333-editingstringforobjectvalue?language=objc
func (f_ Formatter) EditingStringForObjectValue(obj objc.IObject) string {
	rv := objc.Call[string](f_, objc.Sel("editingStringForObjectValue:"), obj)
	return rv
}

// The default implementation of this method raises an exception. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsformatter/1415993-stringforobjectvalue?language=objc
func (f_ Formatter) StringForObjectValue(obj objc.IObject) string {
	rv := objc.Call[string](f_, objc.Sel("stringForObjectValue:"), obj)
	return rv
}

// The default implementation of this method raises an exception. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsformatter/1408927-getobjectvalue?language=objc
func (f_ Formatter) GetObjectValueForStringErrorDescription(obj objc.IObject, string_ string, error string) bool {
	rv := objc.Call[bool](f_, objc.Sel("getObjectValue:forString:errorDescription:"), obj, string_, error)
	return rv
}

// Returns a Boolean value that indicates whether a partial string is valid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsformatter/1417993-ispartialstringvalid?language=objc
func (f_ Formatter) IsPartialStringValidNewEditingStringErrorDescription(partialString string, newString string, error string) bool {
	rv := objc.Call[bool](f_, objc.Sel("isPartialStringValid:newEditingString:errorDescription:"), partialString, newString, error)
	return rv
}

// The default implementation returns nil to indicate that the formatter object does not provide an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsformatter/1409478-attributedstringforobjectvalue?language=objc
func (f_ Formatter) AttributedStringForObjectValueWithDefaultAttributes(obj objc.IObject, attrs map[AttributedStringKey]objc.IObject) AttributedString {
	rv := objc.Call[AttributedString](f_, objc.Sel("attributedStringForObjectValue:withDefaultAttributes:"), obj, attrs)
	return rv
}
