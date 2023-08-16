// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PhoneNumber] class.
var PhoneNumberClass = _PhoneNumberClass{objc.GetClass("CNPhoneNumber")}

type _PhoneNumberClass struct {
	objc.Class
}

// An interface definition for the [PhoneNumber] class.
type IPhoneNumber interface {
	objc.IObject
	StringValue() string
}

// An immutable object representing a phone number for a contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnphonenumber?language=objc
type PhoneNumber struct {
	objc.Object
}

func PhoneNumberFrom(ptr unsafe.Pointer) PhoneNumber {
	return PhoneNumber{
		Object: objc.ObjectFrom(ptr),
	}
}

func (p_ PhoneNumber) InitWithStringValue(string_ string) PhoneNumber {
	rv := objc.Call[PhoneNumber](p_, objc.Sel("initWithStringValue:"), string_)
	return rv
}

// Returns a new phone number object initialized with the specified phone number string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnphonenumber/1401355-initwithstringvalue?language=objc
func PhoneNumber_InitWithStringValue(string_ string) PhoneNumber {
	return PhoneNumberClass.Alloc().InitWithStringValue(string_)
}

func (pc _PhoneNumberClass) PhoneNumberWithStringValue(stringValue string) PhoneNumber {
	rv := objc.Call[PhoneNumber](pc, objc.Sel("phoneNumberWithStringValue:"), stringValue)
	return rv
}

// Returns a new phone number object initialized with the specified phone number string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnphonenumber/1401363-phonenumberwithstringvalue?language=objc
func PhoneNumber_PhoneNumberWithStringValue(stringValue string) PhoneNumber {
	return PhoneNumberClass.PhoneNumberWithStringValue(stringValue)
}

func (pc _PhoneNumberClass) Alloc() PhoneNumber {
	rv := objc.Call[PhoneNumber](pc, objc.Sel("alloc"))
	return rv
}

func PhoneNumber_Alloc() PhoneNumber {
	return PhoneNumberClass.Alloc()
}

func (pc _PhoneNumberClass) New() PhoneNumber {
	rv := objc.Call[PhoneNumber](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPhoneNumber() PhoneNumber {
	return PhoneNumberClass.New()
}

func (p_ PhoneNumber) Init() PhoneNumber {
	rv := objc.Call[PhoneNumber](p_, objc.Sel("init"))
	return rv
}

// The string value of the phone number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnphonenumber/1401367-stringvalue?language=objc
func (p_ PhoneNumber) StringValue() string {
	rv := objc.Call[string](p_, objc.Sel("stringValue"))
	return rv
}
