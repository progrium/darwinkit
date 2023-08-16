// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ContactFormatter] class.
var ContactFormatterClass = _ContactFormatterClass{objc.GetClass("CNContactFormatter")}

type _ContactFormatterClass struct {
	objc.Class
}

// An interface definition for the [ContactFormatter] class.
type IContactFormatter interface {
	foundation.IFormatter
	StringFromContact(contact IContact) string
	AttributedStringFromContactDefaultAttributes(contact IContact, attributes foundation.Dictionary) foundation.AttributedString
	Style() ContactFormatterStyle
	SetStyle(value ContactFormatterStyle)
}

// An object that you use to format contact information before displaying it to the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactformatter?language=objc
type ContactFormatter struct {
	foundation.Formatter
}

func ContactFormatterFrom(ptr unsafe.Pointer) ContactFormatter {
	return ContactFormatter{
		Formatter: foundation.FormatterFrom(ptr),
	}
}

func (cc _ContactFormatterClass) Alloc() ContactFormatter {
	rv := objc.Call[ContactFormatter](cc, objc.Sel("alloc"))
	return rv
}

func ContactFormatter_Alloc() ContactFormatter {
	return ContactFormatterClass.Alloc()
}

func (cc _ContactFormatterClass) New() ContactFormatter {
	rv := objc.Call[ContactFormatter](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewContactFormatter() ContactFormatter {
	return ContactFormatterClass.New()
}

func (c_ ContactFormatter) Init() ContactFormatter {
	rv := objc.Call[ContactFormatter](c_, objc.Sel("init"))
	return rv
}

// Formats the contact name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactformatter/1397754-stringfromcontact?language=objc
func (c_ ContactFormatter) StringFromContact(contact IContact) string {
	rv := objc.Call[string](c_, objc.Sel("stringFromContact:"), objc.Ptr(contact))
	return rv
}

// Returns the required key descriptor for the specified formatting style of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactformatter/1397746-descriptorforrequiredkeysforstyl?language=objc
func (cc _ContactFormatterClass) DescriptorForRequiredKeysForStyle(style ContactFormatterStyle) objc.Object {
	rv := objc.Call[objc.Object](cc, objc.Sel("descriptorForRequiredKeysForStyle:"), style)
	return rv
}

// Returns the required key descriptor for the specified formatting style of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactformatter/1397746-descriptorforrequiredkeysforstyl?language=objc
func ContactFormatter_DescriptorForRequiredKeysForStyle(style ContactFormatterStyle) objc.Object {
	return ContactFormatterClass.DescriptorForRequiredKeysForStyle(style)
}

// Returns the display name order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactformatter/1397731-nameorderforcontact?language=objc
func (cc _ContactFormatterClass) NameOrderForContact(contact IContact) ContactDisplayNameOrder {
	rv := objc.Call[ContactDisplayNameOrder](cc, objc.Sel("nameOrderForContact:"), objc.Ptr(contact))
	return rv
}

// Returns the display name order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactformatter/1397731-nameorderforcontact?language=objc
func ContactFormatter_NameOrderForContact(contact IContact) ContactDisplayNameOrder {
	return ContactFormatterClass.NameOrderForContact(contact)
}

// Formats the contact name as an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactformatter/1397733-attributedstringfromcontact?language=objc
func (c_ ContactFormatter) AttributedStringFromContactDefaultAttributes(contact IContact, attributes foundation.Dictionary) foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](c_, objc.Sel("attributedStringFromContact:defaultAttributes:"), objc.Ptr(contact), attributes)
	return rv
}

// Returns the delimiter to use between name components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactformatter/1397744-delimiterforcontact?language=objc
func (cc _ContactFormatterClass) DelimiterForContact(contact IContact) string {
	rv := objc.Call[string](cc, objc.Sel("delimiterForContact:"), objc.Ptr(contact))
	return rv
}

// Returns the delimiter to use between name components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactformatter/1397744-delimiterforcontact?language=objc
func ContactFormatter_DelimiterForContact(contact IContact) string {
	return ContactFormatterClass.DelimiterForContact(contact)
}

// The formatting style for the contact name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactformatter/1397758-style?language=objc
func (c_ ContactFormatter) Style() ContactFormatterStyle {
	rv := objc.Call[ContactFormatterStyle](c_, objc.Sel("style"))
	return rv
}

// The formatting style for the contact name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactformatter/1397758-style?language=objc
func (c_ ContactFormatter) SetStyle(value ContactFormatterStyle) {
	objc.Call[objc.Void](c_, objc.Sel("setStyle:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactformatter/3601127-descriptorforrequiredkeysforname?language=objc
func (cc _ContactFormatterClass) DescriptorForRequiredKeysForNameOrder() objc.Object {
	rv := objc.Call[objc.Object](cc, objc.Sel("descriptorForRequiredKeysForNameOrder"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactformatter/3601127-descriptorforrequiredkeysforname?language=objc
func ContactFormatter_DescriptorForRequiredKeysForNameOrder() objc.Object {
	return ContactFormatterClass.DescriptorForRequiredKeysForNameOrder()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactformatter/3601126-descriptorforrequiredkeysfordeli?language=objc
func (cc _ContactFormatterClass) DescriptorForRequiredKeysForDelimiter() objc.Object {
	rv := objc.Call[objc.Object](cc, objc.Sel("descriptorForRequiredKeysForDelimiter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactformatter/3601126-descriptorforrequiredkeysfordeli?language=objc
func ContactFormatter_DescriptorForRequiredKeysForDelimiter() objc.Object {
	return ContactFormatterClass.DescriptorForRequiredKeysForDelimiter()
}
