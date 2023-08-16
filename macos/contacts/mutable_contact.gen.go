// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutableContact] class.
var MutableContactClass = _MutableContactClass{objc.GetClass("CNMutableContact")}

type _MutableContactClass struct {
	objc.Class
}

// An interface definition for the [MutableContact] class.
type IMutableContact interface {
	IContact
	SetDates(value []ILabeledValue)
	SetPhoneNumbers(value []ILabeledValue)
	SetPreviousFamilyName(value string)
	SetPhoneticOrganizationName(value string)
	SetContactType(value ContactType)
	SetPhoneticGivenName(value string)
	SetPhoneticMiddleName(value string)
	SetNamePrefix(value string)
	SetPostalAddresses(value []ILabeledValue)
	SetNonGregorianBirthday(value foundation.IDateComponents)
	SetContactRelations(value []ILabeledValue)
	SetBirthday(value foundation.IDateComponents)
	SetImageData(value []byte)
	SetNote(value string)
	SetMiddleName(value string)
	SetJobTitle(value string)
	SetGivenName(value string)
	SetEmailAddresses(value []ILabeledValue)
	SetInstantMessageAddresses(value []ILabeledValue)
	SetUrlAddresses(value []ILabeledValue)
	SetOrganizationName(value string)
	SetNameSuffix(value string)
	SetDepartmentName(value string)
	SetPhoneticFamilyName(value string)
	SetSocialProfiles(value []ILabeledValue)
	SetFamilyName(value string)
	SetNickname(value string)
}

// A mutable object that stores information about a single contact, such as the contact's first name, phone numbers, and addresses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact?language=objc
type MutableContact struct {
	Contact
}

func MutableContactFrom(ptr unsafe.Pointer) MutableContact {
	return MutableContact{
		Contact: ContactFrom(ptr),
	}
}

func (mc _MutableContactClass) Alloc() MutableContact {
	rv := objc.Call[MutableContact](mc, objc.Sel("alloc"))
	return rv
}

func MutableContact_Alloc() MutableContact {
	return MutableContactClass.Alloc()
}

func (mc _MutableContactClass) New() MutableContact {
	rv := objc.Call[MutableContact](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableContact() MutableContact {
	return MutableContactClass.New()
}

func (m_ MutableContact) Init() MutableContact {
	rv := objc.Call[MutableContact](m_, objc.Sel("init"))
	return rv
}

// An array containing labeled Gregorian dates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1403311-dates?language=objc
func (m_ MutableContact) SetDates(value []ILabeledValue) {
	objc.Call[objc.Void](m_, objc.Sel("setDates:"), value)
}

// An array of labeled phone numbers for a contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1403202-phonenumbers?language=objc
func (m_ MutableContact) SetPhoneNumbers(value []ILabeledValue) {
	objc.Call[objc.Void](m_, objc.Sel("setPhoneNumbers:"), value)
}

// The previous family name of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1403113-previousfamilyname?language=objc
func (m_ MutableContact) SetPreviousFamilyName(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setPreviousFamilyName:"), value)
}

// The phonetic name of the organization associated with the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/2138282-phoneticorganizationname?language=objc
func (m_ MutableContact) SetPhoneticOrganizationName(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setPhoneticOrganizationName:"), value)
}

// An enum identifying the contact type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1403242-contacttype?language=objc
func (m_ MutableContact) SetContactType(value ContactType) {
	objc.Call[objc.Void](m_, objc.Sel("setContactType:"), value)
}

// The phonetic given name of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1402897-phoneticgivenname?language=objc
func (m_ MutableContact) SetPhoneticGivenName(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setPhoneticGivenName:"), value)
}

// The phonetic middle name of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1403340-phoneticmiddlename?language=objc
func (m_ MutableContact) SetPhoneticMiddleName(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setPhoneticMiddleName:"), value)
}

// The name prefix of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1403184-nameprefix?language=objc
func (m_ MutableContact) SetNamePrefix(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setNamePrefix:"), value)
}

// An array of labeled postal addresses for a contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1402823-postaladdresses?language=objc
func (m_ MutableContact) SetPostalAddresses(value []ILabeledValue) {
	objc.Call[objc.Void](m_, objc.Sel("setPostalAddresses:"), value)
}

// A date component for the non-Gregorian birthday of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1402994-nongregorianbirthday?language=objc
func (m_ MutableContact) SetNonGregorianBirthday(value foundation.IDateComponents) {
	objc.Call[objc.Void](m_, objc.Sel("setNonGregorianBirthday:"), objc.Ptr(value))
}

// An array of labeled contact relations for the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1403206-contactrelations?language=objc
func (m_ MutableContact) SetContactRelations(value []ILabeledValue) {
	objc.Call[objc.Void](m_, objc.Sel("setContactRelations:"), value)
}

// A date component for the Gregorian birthday of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1403236-birthday?language=objc
func (m_ MutableContact) SetBirthday(value foundation.IDateComponents) {
	objc.Call[objc.Void](m_, objc.Sel("setBirthday:"), objc.Ptr(value))
}

// The profile picture of a contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1403400-imagedata?language=objc
func (m_ MutableContact) SetImageData(value []byte) {
	objc.Call[objc.Void](m_, objc.Sel("setImageData:"), value)
}

// A string containing notes for the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1403096-note?language=objc
func (m_ MutableContact) SetNote(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setNote:"), value)
}

// The middle name of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1403029-middlename?language=objc
func (m_ MutableContact) SetMiddleName(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setMiddleName:"), value)
}

// The contact’s job title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1403224-jobtitle?language=objc
func (m_ MutableContact) SetJobTitle(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setJobTitle:"), value)
}

// The given name of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1403228-givenname?language=objc
func (m_ MutableContact) SetGivenName(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setGivenName:"), value)
}

// An array of labeled email addresses for the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1402792-emailaddresses?language=objc
func (m_ MutableContact) SetEmailAddresses(value []ILabeledValue) {
	objc.Call[objc.Void](m_, objc.Sel("setEmailAddresses:"), value)
}

// An array of labeled IM addresses for the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1403362-instantmessageaddresses?language=objc
func (m_ MutableContact) SetInstantMessageAddresses(value []ILabeledValue) {
	objc.Call[objc.Void](m_, objc.Sel("setInstantMessageAddresses:"), value)
}

// An array of labeled URL addresses for a contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1403006-urladdresses?language=objc
func (m_ MutableContact) SetUrlAddresses(value []ILabeledValue) {
	objc.Call[objc.Void](m_, objc.Sel("setUrlAddresses:"), value)
}

// The name of the organization associated with the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1403381-organizationname?language=objc
func (m_ MutableContact) SetOrganizationName(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setOrganizationName:"), value)
}

// The name suffix of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1403162-namesuffix?language=objc
func (m_ MutableContact) SetNameSuffix(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setNameSuffix:"), value)
}

// The name of the department associated with the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1403398-departmentname?language=objc
func (m_ MutableContact) SetDepartmentName(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setDepartmentName:"), value)
}

// The phonetic family name of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1402985-phoneticfamilyname?language=objc
func (m_ MutableContact) SetPhoneticFamilyName(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setPhoneticFamilyName:"), value)
}

// An array of labeled social profiles for a contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1403038-socialprofiles?language=objc
func (m_ MutableContact) SetSocialProfiles(value []ILabeledValue) {
	objc.Call[objc.Void](m_, objc.Sel("setSocialProfiles:"), value)
}

// The family name of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1403056-familyname?language=objc
func (m_ MutableContact) SetFamilyName(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setFamilyName:"), value)
}

// The nickname of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/1403135-nickname?language=objc
func (m_ MutableContact) SetNickname(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setNickname:"), value)
}
