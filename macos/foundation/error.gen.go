// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Error] class.
var ErrorClass = _ErrorClass{objc.GetClass("NSError")}

type _ErrorClass struct {
	objc.Class
}

// An interface definition for the [Error] class.
type IError interface {
	objc.IObject
	Domain() ErrorDomain
	LocalizedDescription() string
	LocalizedFailureReason() string
	LocalizedRecoverySuggestion() string
	Code() int
	LocalizedRecoveryOptions() []string
	UserInfo() map[ErrorUserInfoKey]objc.Object
	UnderlyingErrors() []Error
	RecoveryAttempter() objc.Object
	HelpAnchor() string
}

// Information about an error condition including a domain, a domain-specific error code, and application-specific information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror?language=objc
type Error struct {
	objc.Object
}

func ErrorFrom(ptr unsafe.Pointer) Error {
	return Error{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ec _ErrorClass) FileProviderErrorForRejectedDeletionOfItem(updatedVersion objc.IObject) Error {
	rv := objc.Call[Error](ec, objc.Sel("fileProviderErrorForRejectedDeletionOfItem:"), objc.Ptr(updatedVersion))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/3603574-fileprovidererrorforrejecteddele?language=objc
func Error_FileProviderErrorForRejectedDeletionOfItem(updatedVersion objc.IObject) Error {
	return ErrorClass.FileProviderErrorForRejectedDeletionOfItem(updatedVersion)
}

func (ec _ErrorClass) FileProviderErrorForCollisionWithItem(existingItem objc.IObject) Error {
	rv := objc.Call[Error](ec, objc.Sel("fileProviderErrorForCollisionWithItem:"), objc.Ptr(existingItem))
	return rv
}

// Returns a properly formatted error object with a NSFileProviderItemCollisionError error code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/2882067-fileprovidererrorforcollisionwit?language=objc
func Error_FileProviderErrorForCollisionWithItem(existingItem objc.IObject) Error {
	return ErrorClass.FileProviderErrorForCollisionWithItem(existingItem)
}

func (ec _ErrorClass) FileProviderErrorForNonExistentItemWithIdentifier(itemIdentifier objc.IObject) Error {
	rv := objc.Call[Error](ec, objc.Sel("fileProviderErrorForNonExistentItemWithIdentifier:"), objc.Ptr(itemIdentifier))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/2915899-fileprovidererrorfornonexistenti?language=objc
func Error_FileProviderErrorForNonExistentItemWithIdentifier(itemIdentifier objc.IObject) Error {
	return ErrorClass.FileProviderErrorForNonExistentItemWithIdentifier(itemIdentifier)
}

func (ec _ErrorClass) ErrorWithDomainCodeUserInfo(domain ErrorDomain, code int, dict map[ErrorUserInfoKey]objc.IObject) Error {
	rv := objc.Call[Error](ec, objc.Sel("errorWithDomain:code:userInfo:"), domain, code, dict)
	return rv
}

// Creates and initializes an NSError object for a given domain and code with a given userInfo dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/1522782-errorwithdomain?language=objc
func Error_ErrorWithDomainCodeUserInfo(domain ErrorDomain, code int, dict map[ErrorUserInfoKey]objc.IObject) Error {
	return ErrorClass.ErrorWithDomainCodeUserInfo(domain, code, dict)
}

func (e_ Error) InitWithDomainCodeUserInfo(domain ErrorDomain, code int, dict map[ErrorUserInfoKey]objc.IObject) Error {
	rv := objc.Call[Error](e_, objc.Sel("initWithDomain:code:userInfo:"), domain, code, dict)
	return rv
}

// Returns an NSError object initialized for a given domain and code with a given userInfo dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/1417063-initwithdomain?language=objc
func NewErrorWithDomainCodeUserInfo(domain ErrorDomain, code int, dict map[ErrorUserInfoKey]objc.IObject) Error {
	instance := ErrorClass.Alloc().InitWithDomainCodeUserInfo(domain, code, dict)
	instance.Autorelease()
	return instance
}

func (ec _ErrorClass) Alloc() Error {
	rv := objc.Call[Error](ec, objc.Sel("alloc"))
	return rv
}

func Error_Alloc() Error {
	return ErrorClass.Alloc()
}

func (ec _ErrorClass) New() Error {
	rv := objc.Call[Error](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewError() Error {
	return ErrorClass.New()
}

func (e_ Error) Init() Error {
	rv := objc.Call[Error](e_, objc.Sel("init"))
	return rv
}

// Specifies a block to call when the corresponding property is not present in the user info dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/1408064-setuserinfovalueproviderfordomai?language=objc
func (ec _ErrorClass) SetUserInfoValueProviderForDomainProvider(errorDomain ErrorDomain, provider func(err Error, userInfoKey ErrorUserInfoKey) objc.Object) {
	objc.Call[objc.Void](ec, objc.Sel("setUserInfoValueProviderForDomain:provider:"), errorDomain, provider)
}

// Specifies a block to call when the corresponding property is not present in the user info dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/1408064-setuserinfovalueproviderfordomai?language=objc
func Error_SetUserInfoValueProviderForDomainProvider(errorDomain ErrorDomain, provider func(err Error, userInfoKey ErrorUserInfoKey) objc.Object) {
	ErrorClass.SetUserInfoValueProviderForDomainProvider(errorDomain, provider)
}

// Returns any user info provider specified for a given error domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/1413427-userinfovalueproviderfordomain?language=objc
func (ec _ErrorClass) UserInfoValueProviderForDomain(errorDomain ErrorDomain) func(arg0 Error, arg1 ErrorUserInfoKey) objc.Object {
	rv := objc.Call[func(arg0 Error, arg1 ErrorUserInfoKey) objc.Object](ec, objc.Sel("userInfoValueProviderForDomain:"), errorDomain)
	return rv
}

// Returns any user info provider specified for a given error domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/1413427-userinfovalueproviderfordomain?language=objc
func Error_UserInfoValueProviderForDomain(errorDomain ErrorDomain) func(arg0 Error, arg1 ErrorUserInfoKey) objc.Object {
	return ErrorClass.UserInfoValueProviderForDomain(errorDomain)
}

// A string containing the error domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/1413924-domain?language=objc
func (e_ Error) Domain() ErrorDomain {
	rv := objc.Call[ErrorDomain](e_, objc.Sel("domain"))
	return rv
}

// A string containing the localized description of the error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/1414418-localizeddescription?language=objc
func (e_ Error) LocalizedDescription() string {
	rv := objc.Call[string](e_, objc.Sel("localizedDescription"))
	return rv
}

// A string containing the localized explanation of the reason for the error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/1412752-localizedfailurereason?language=objc
func (e_ Error) LocalizedFailureReason() string {
	rv := objc.Call[string](e_, objc.Sel("localizedFailureReason"))
	return rv
}

// A string containing the localized recovery suggestion for the error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/1407500-localizedrecoverysuggestion?language=objc
func (e_ Error) LocalizedRecoverySuggestion() string {
	rv := objc.Call[string](e_, objc.Sel("localizedRecoverySuggestion"))
	return rv
}

// The error code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/1409165-code?language=objc
func (e_ Error) Code() int {
	rv := objc.Call[int](e_, objc.Sel("code"))
	return rv
}

// An array containing the localized titles of buttons appropriate for displaying in an alert panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/1415950-localizedrecoveryoptions?language=objc
func (e_ Error) LocalizedRecoveryOptions() []string {
	rv := objc.Call[[]string](e_, objc.Sel("localizedRecoveryOptions"))
	return rv
}

// The user info dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/1411580-userinfo?language=objc
func (e_ Error) UserInfo() map[ErrorUserInfoKey]objc.Object {
	rv := objc.Call[map[ErrorUserInfoKey]objc.Object](e_, objc.Sel("userInfo"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/3738169-underlyingerrors?language=objc
func (e_ Error) UnderlyingErrors() []Error {
	rv := objc.Call[[]Error](e_, objc.Sel("underlyingErrors"))
	return rv
}

// The object in the user info dictionary corresponding to the NSRecoveryAttempterErrorKey key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/1408864-recoveryattempter?language=objc
func (e_ Error) RecoveryAttempter() objc.Object {
	rv := objc.Call[objc.Object](e_, objc.Sel("recoveryAttempter"))
	return rv
}

// A string to display in response to an alert panel help anchor button being pressed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nserror/1414718-helpanchor?language=objc
func (e_ Error) HelpAnchor() string {
	rv := objc.Call[string](e_, objc.Sel("helpAnchor"))
	return rv
}
