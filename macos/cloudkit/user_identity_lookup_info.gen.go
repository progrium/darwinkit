// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UserIdentityLookupInfo] class.
var UserIdentityLookupInfoClass = _UserIdentityLookupInfoClass{objc.GetClass("CKUserIdentityLookupInfo")}

type _UserIdentityLookupInfoClass struct {
	objc.Class
}

// An interface definition for the [UserIdentityLookupInfo] class.
type IUserIdentityLookupInfo interface {
	objc.IObject
	UserRecordID() RecordID
	PhoneNumber() string
	EmailAddress() string
}

// The criteria to use when searching for discoverable iCloud users. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentitylookupinfo?language=objc
type UserIdentityLookupInfo struct {
	objc.Object
}

func UserIdentityLookupInfoFrom(ptr unsafe.Pointer) UserIdentityLookupInfo {
	return UserIdentityLookupInfo{
		Object: objc.ObjectFrom(ptr),
	}
}

func (u_ UserIdentityLookupInfo) InitWithEmailAddress(emailAddress string) UserIdentityLookupInfo {
	rv := objc.Call[UserIdentityLookupInfo](u_, objc.Sel("initWithEmailAddress:"), emailAddress)
	return rv
}

// Creates a lookup info for the specified email address. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentitylookupinfo/1640484-initwithemailaddress?language=objc
func UserIdentityLookupInfo_InitWithEmailAddress(emailAddress string) UserIdentityLookupInfo {
	return UserIdentityLookupInfoClass.Alloc().InitWithEmailAddress(emailAddress)
}

func (u_ UserIdentityLookupInfo) InitWithUserRecordID(userRecordID IRecordID) UserIdentityLookupInfo {
	rv := objc.Call[UserIdentityLookupInfo](u_, objc.Sel("initWithUserRecordID:"), objc.Ptr(userRecordID))
	return rv
}

// Creates a lookup info for the specified user record ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentitylookupinfo/1640419-initwithuserrecordid?language=objc
func UserIdentityLookupInfo_InitWithUserRecordID(userRecordID IRecordID) UserIdentityLookupInfo {
	return UserIdentityLookupInfoClass.Alloc().InitWithUserRecordID(userRecordID)
}

func (u_ UserIdentityLookupInfo) InitWithPhoneNumber(phoneNumber string) UserIdentityLookupInfo {
	rv := objc.Call[UserIdentityLookupInfo](u_, objc.Sel("initWithPhoneNumber:"), phoneNumber)
	return rv
}

// Creates a lookup info for the specified phone number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentitylookupinfo/1640402-initwithphonenumber?language=objc
func UserIdentityLookupInfo_InitWithPhoneNumber(phoneNumber string) UserIdentityLookupInfo {
	return UserIdentityLookupInfoClass.Alloc().InitWithPhoneNumber(phoneNumber)
}

func (uc _UserIdentityLookupInfoClass) Alloc() UserIdentityLookupInfo {
	rv := objc.Call[UserIdentityLookupInfo](uc, objc.Sel("alloc"))
	return rv
}

func UserIdentityLookupInfo_Alloc() UserIdentityLookupInfo {
	return UserIdentityLookupInfoClass.Alloc()
}

func (uc _UserIdentityLookupInfoClass) New() UserIdentityLookupInfo {
	rv := objc.Call[UserIdentityLookupInfo](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUserIdentityLookupInfo() UserIdentityLookupInfo {
	return UserIdentityLookupInfoClass.New()
}

func (u_ UserIdentityLookupInfo) Init() UserIdentityLookupInfo {
	rv := objc.Call[UserIdentityLookupInfo](u_, objc.Sel("init"))
	return rv
}

// Returns an array of lookup infos for the specifed user record IDs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentitylookupinfo/1640407-lookupinfoswithrecordids?language=objc
func (uc _UserIdentityLookupInfoClass) LookupInfosWithRecordIDs(recordIDs []IRecordID) []UserIdentityLookupInfo {
	rv := objc.Call[[]UserIdentityLookupInfo](uc, objc.Sel("lookupInfosWithRecordIDs:"), recordIDs)
	return rv
}

// Returns an array of lookup infos for the specifed user record IDs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentitylookupinfo/1640407-lookupinfoswithrecordids?language=objc
func UserIdentityLookupInfo_LookupInfosWithRecordIDs(recordIDs []IRecordID) []UserIdentityLookupInfo {
	return UserIdentityLookupInfoClass.LookupInfosWithRecordIDs(recordIDs)
}

// Returns an array of lookup infos for the specifed email addresses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentitylookupinfo/1640439-lookupinfoswithemails?language=objc
func (uc _UserIdentityLookupInfoClass) LookupInfosWithEmails(emails []string) []UserIdentityLookupInfo {
	rv := objc.Call[[]UserIdentityLookupInfo](uc, objc.Sel("lookupInfosWithEmails:"), emails)
	return rv
}

// Returns an array of lookup infos for the specifed email addresses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentitylookupinfo/1640439-lookupinfoswithemails?language=objc
func UserIdentityLookupInfo_LookupInfosWithEmails(emails []string) []UserIdentityLookupInfo {
	return UserIdentityLookupInfoClass.LookupInfosWithEmails(emails)
}

// Returns an array of lookup infos for the specifed phone numbers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentitylookupinfo/1640429-lookupinfoswithphonenumbers?language=objc
func (uc _UserIdentityLookupInfoClass) LookupInfosWithPhoneNumbers(phoneNumbers []string) []UserIdentityLookupInfo {
	rv := objc.Call[[]UserIdentityLookupInfo](uc, objc.Sel("lookupInfosWithPhoneNumbers:"), phoneNumbers)
	return rv
}

// Returns an array of lookup infos for the specifed phone numbers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentitylookupinfo/1640429-lookupinfoswithphonenumbers?language=objc
func UserIdentityLookupInfo_LookupInfosWithPhoneNumbers(phoneNumbers []string) []UserIdentityLookupInfo {
	return UserIdentityLookupInfoClass.LookupInfosWithPhoneNumbers(phoneNumbers)
}

// The ID of the user record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentitylookupinfo/1640405-userrecordid?language=objc
func (u_ UserIdentityLookupInfo) UserRecordID() RecordID {
	rv := objc.Call[RecordID](u_, objc.Sel("userRecordID"))
	return rv
}

// The user’s phone number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentitylookupinfo/1640482-phonenumber?language=objc
func (u_ UserIdentityLookupInfo) PhoneNumber() string {
	rv := objc.Call[string](u_, objc.Sel("phoneNumber"))
	return rv
}

// The user’s email address. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentitylookupinfo/1640462-emailaddress?language=objc
func (u_ UserIdentityLookupInfo) EmailAddress() string {
	rv := objc.Call[string](u_, objc.Sel("emailAddress"))
	return rv
}
