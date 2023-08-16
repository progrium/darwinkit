// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PostalAddressFormatter] class.
var PostalAddressFormatterClass = _PostalAddressFormatterClass{objc.GetClass("CNPostalAddressFormatter")}

type _PostalAddressFormatterClass struct {
	objc.Class
}

// An interface definition for the [PostalAddressFormatter] class.
type IPostalAddressFormatter interface {
	foundation.IFormatter
	StringFromPostalAddress(postalAddress IPostalAddress) string
	AttributedStringFromPostalAddressWithDefaultAttributes(postalAddress IPostalAddress, attributes foundation.Dictionary) foundation.AttributedString
	Style() PostalAddressFormatterStyle
	SetStyle(value PostalAddressFormatterStyle)
}

// An object that you use to format a contact's postal addresses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdressformatter?language=objc
type PostalAddressFormatter struct {
	foundation.Formatter
}

func PostalAddressFormatterFrom(ptr unsafe.Pointer) PostalAddressFormatter {
	return PostalAddressFormatter{
		Formatter: foundation.FormatterFrom(ptr),
	}
}

func (pc _PostalAddressFormatterClass) Alloc() PostalAddressFormatter {
	rv := objc.Call[PostalAddressFormatter](pc, objc.Sel("alloc"))
	return rv
}

func PostalAddressFormatter_Alloc() PostalAddressFormatter {
	return PostalAddressFormatterClass.Alloc()
}

func (pc _PostalAddressFormatterClass) New() PostalAddressFormatter {
	rv := objc.Call[PostalAddressFormatter](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPostalAddressFormatter() PostalAddressFormatter {
	return PostalAddressFormatterClass.New()
}

func (p_ PostalAddressFormatter) Init() PostalAddressFormatter {
	rv := objc.Call[PostalAddressFormatter](p_, objc.Sel("init"))
	return rv
}

// Returns a formatted postal address. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdressformatter/1403011-stringfrompostaladdress?language=objc
func (p_ PostalAddressFormatter) StringFromPostalAddress(postalAddress IPostalAddress) string {
	rv := objc.Call[string](p_, objc.Sel("stringFromPostalAddress:"), objc.Ptr(postalAddress))
	return rv
}

// Returns a formatted postal address as an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdressformatter/1402831-attributedstringfrompostaladdres?language=objc
func (p_ PostalAddressFormatter) AttributedStringFromPostalAddressWithDefaultAttributes(postalAddress IPostalAddress, attributes foundation.Dictionary) foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](p_, objc.Sel("attributedStringFromPostalAddress:withDefaultAttributes:"), objc.Ptr(postalAddress), attributes)
	return rv
}

// The style to apply when formatting strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdressformatter/1402781-style?language=objc
func (p_ PostalAddressFormatter) Style() PostalAddressFormatterStyle {
	rv := objc.Call[PostalAddressFormatterStyle](p_, objc.Sel("style"))
	return rv
}

// The style to apply when formatting strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdressformatter/1402781-style?language=objc
func (p_ PostalAddressFormatter) SetStyle(value PostalAddressFormatterStyle) {
	objc.Call[objc.Void](p_, objc.Sel("setStyle:"), value)
}
