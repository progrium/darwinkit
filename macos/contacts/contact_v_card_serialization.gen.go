// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ContactVCardSerialization] class.
var ContactVCardSerializationClass = _ContactVCardSerializationClass{objc.GetClass("CNContactVCardSerialization")}

type _ContactVCardSerializationClass struct {
	objc.Class
}

// An interface definition for the [ContactVCardSerialization] class.
type IContactVCardSerialization interface {
	objc.IObject
}

// An object you use to convert to and from a vCard representation of the user's contacts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactvcardserialization?language=objc
type ContactVCardSerialization struct {
	objc.Object
}

func ContactVCardSerializationFrom(ptr unsafe.Pointer) ContactVCardSerialization {
	return ContactVCardSerialization{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ContactVCardSerializationClass) Alloc() ContactVCardSerialization {
	rv := objc.Call[ContactVCardSerialization](cc, objc.Sel("alloc"))
	return rv
}

func ContactVCardSerialization_Alloc() ContactVCardSerialization {
	return ContactVCardSerializationClass.Alloc()
}

func (cc _ContactVCardSerializationClass) New() ContactVCardSerialization {
	rv := objc.Call[ContactVCardSerialization](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewContactVCardSerialization() ContactVCardSerialization {
	return ContactVCardSerializationClass.New()
}

func (c_ ContactVCardSerialization) Init() ContactVCardSerialization {
	rv := objc.Call[ContactVCardSerialization](c_, objc.Sel("init"))
	return rv
}

// Use to fetch all contact keys required to create vCard data from a contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactvcardserialization/1403155-descriptorforrequiredkeys?language=objc
func (cc _ContactVCardSerializationClass) DescriptorForRequiredKeys() objc.Object {
	rv := objc.Call[objc.Object](cc, objc.Sel("descriptorForRequiredKeys"))
	return rv
}

// Use to fetch all contact keys required to create vCard data from a contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactvcardserialization/1403155-descriptorforrequiredkeys?language=objc
func ContactVCardSerialization_DescriptorForRequiredKeys() objc.Object {
	return ContactVCardSerializationClass.DescriptorForRequiredKeys()
}

// Returns the vCard representation of the specified contacts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactvcardserialization/1403357-datawithcontacts?language=objc
func (cc _ContactVCardSerializationClass) DataWithContactsError(contacts []IContact, error foundation.IError) []byte {
	rv := objc.Call[[]byte](cc, objc.Sel("dataWithContacts:error:"), contacts, objc.Ptr(error))
	return rv
}

// Returns the vCard representation of the specified contacts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactvcardserialization/1403357-datawithcontacts?language=objc
func ContactVCardSerialization_DataWithContactsError(contacts []IContact, error foundation.IError) []byte {
	return ContactVCardSerializationClass.DataWithContactsError(contacts, error)
}

// Returns the contacts from the vCard data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactvcardserialization/1403090-contactswithdata?language=objc
func (cc _ContactVCardSerializationClass) ContactsWithDataError(data []byte, error foundation.IError) []Contact {
	rv := objc.Call[[]Contact](cc, objc.Sel("contactsWithData:error:"), data, objc.Ptr(error))
	return rv
}

// Returns the contacts from the vCard data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactvcardserialization/1403090-contactswithdata?language=objc
func ContactVCardSerialization_ContactsWithDataError(data []byte, error foundation.IError) []Contact {
	return ContactVCardSerializationClass.ContactsWithDataError(data, error)
}
