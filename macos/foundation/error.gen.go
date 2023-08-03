// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var ErrorClass = _ErrorClass{objc.GetClass("NSError")}

type _ErrorClass struct {
	objc.Class
}

type IError interface {
	objc.IObject
	AttemptRecoveryFromErrorOptionIndexDelegateDidRecoverSelectorContextInfo(error IError, recoveryOptionIndex uint, delegate objc.IObject, didRecoverSelector objc.Selector, contextInfo unsafe.Pointer)
	AttemptRecoveryFromErrorOptionIndex(error IError, recoveryOptionIndex uint) bool
	Code() int
	Domain() ErrorDomain
	UserInfo() map[ErrorUserInfoKey]objc.Object
	LocalizedDescription() string
	LocalizedRecoveryOptions() []string
	LocalizedRecoverySuggestion() string
	LocalizedFailureReason() string
	RecoveryAttempter() objc.Object
	HelpAnchor() string
	UnderlyingErrors() []Error
}

type Error struct {
	objc.Object
}

func MakeError(ptr unsafe.Pointer) Error {
	return Error{
		Object: objc.MakeObject(ptr),
	}
}

func (ec _ErrorClass) ErrorWithDomainCodeUserInfo(domain ErrorDomain, code int, dict map[ErrorUserInfoKey]objc.IObject) Error {
	rv := objc.CallMethod[Error](ec, objc.GetSelector("errorWithDomain:code:userInfo:"), domain, code, dict)
	return rv
}

func Error_ErrorWithDomainCodeUserInfo(domain ErrorDomain, code int, dict map[ErrorUserInfoKey]objc.IObject) Error {
	return ErrorClass.ErrorWithDomainCodeUserInfo(domain, code, dict)
}

func (e_ Error) InitWithDomainCodeUserInfo(domain ErrorDomain, code int, dict map[ErrorUserInfoKey]objc.IObject) Error {
	rv := objc.CallMethod[Error](e_, objc.GetSelector("initWithDomain:code:userInfo:"), domain, code, dict)
	return rv
}

func Error_InitWithDomainCodeUserInfo(domain ErrorDomain, code int, dict map[ErrorUserInfoKey]objc.IObject) Error {
	return ErrorClass.Alloc().InitWithDomainCodeUserInfo(domain, code, dict)
}

func (ec _ErrorClass) Alloc() Error {
	rv := objc.CallMethod[Error](ec, objc.GetSelector("alloc"))
	return rv
}

func Error_Alloc() Error {
	return ErrorClass.Alloc()
}

func (ec _ErrorClass) New() Error {
	rv := objc.CallMethod[Error](ec, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewError() Error {
	return ErrorClass.New()
}

func Error_New() Error {
	return ErrorClass.New()
}

func (e_ Error) Init() Error {
	rv := objc.CallMethod[Error](e_, objc.GetSelector("init"))
	return rv
}

func Error_Init() Error {
	return ErrorClass.Alloc().Init()
}

func (ec _ErrorClass) SetUserInfoValueProviderForDomainProvider(errorDomain ErrorDomain, provider func(err Error, userInfoKey ErrorUserInfoKey) objc.Object) {
	objc.CallMethod[objc.Void](ec, objc.GetSelector("setUserInfoValueProviderForDomain:provider:"), errorDomain, provider)
}

func Error_SetUserInfoValueProviderForDomainProvider(errorDomain ErrorDomain, provider func(err Error, userInfoKey ErrorUserInfoKey) objc.Object) {
	ErrorClass.SetUserInfoValueProviderForDomainProvider(errorDomain, provider)
}

func (ec _ErrorClass) UserInfoValueProviderForDomain(errorDomain ErrorDomain) func(param1 Error, param2 ErrorUserInfoKey) objc.Object {
	rv := objc.CallMethod[func(param1 Error, param2 ErrorUserInfoKey) objc.Object](ec, objc.GetSelector("userInfoValueProviderForDomain:"), errorDomain)
	return rv
}

func Error_UserInfoValueProviderForDomain(errorDomain ErrorDomain) func(param1 Error, param2 ErrorUserInfoKey) objc.Object {
	return ErrorClass.UserInfoValueProviderForDomain(errorDomain)
}

func (e_ Error) AttemptRecoveryFromErrorOptionIndexDelegateDidRecoverSelectorContextInfo(error IError, recoveryOptionIndex uint, delegate objc.IObject, didRecoverSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](e_, objc.GetSelector("attemptRecoveryFromError:optionIndex:delegate:didRecoverSelector:contextInfo:"), objc.ExtractPtr(error), recoveryOptionIndex, objc.ExtractPtr(delegate), didRecoverSelector, contextInfo)
}

func (e_ Error) AttemptRecoveryFromErrorOptionIndex(error IError, recoveryOptionIndex uint) bool {
	rv := objc.CallMethod[bool](e_, objc.GetSelector("attemptRecoveryFromError:optionIndex:"), objc.ExtractPtr(error), recoveryOptionIndex)
	return rv
}

func (e_ Error) Code() int {
	rv := objc.CallMethod[int](e_, objc.GetSelector("code"))
	return rv
}

func (e_ Error) Domain() ErrorDomain {
	rv := objc.CallMethod[ErrorDomain](e_, objc.GetSelector("domain"))
	return rv
}

func (e_ Error) UserInfo() map[ErrorUserInfoKey]objc.Object {
	rv := objc.CallMethod[map[ErrorUserInfoKey]objc.Object](e_, objc.GetSelector("userInfo"))
	return rv
}

func (e_ Error) LocalizedDescription() string {
	rv := objc.CallMethod[string](e_, objc.GetSelector("localizedDescription"))
	return rv
}

func (e_ Error) LocalizedRecoveryOptions() []string {
	rv := objc.CallMethod[[]string](e_, objc.GetSelector("localizedRecoveryOptions"))
	return rv
}

func (e_ Error) LocalizedRecoverySuggestion() string {
	rv := objc.CallMethod[string](e_, objc.GetSelector("localizedRecoverySuggestion"))
	return rv
}

func (e_ Error) LocalizedFailureReason() string {
	rv := objc.CallMethod[string](e_, objc.GetSelector("localizedFailureReason"))
	return rv
}

func (e_ Error) RecoveryAttempter() objc.Object {
	rv := objc.CallMethod[objc.Object](e_, objc.GetSelector("recoveryAttempter"))
	return rv
}

func (e_ Error) HelpAnchor() string {
	rv := objc.CallMethod[string](e_, objc.GetSelector("helpAnchor"))
	return rv
}

func (e_ Error) UnderlyingErrors() []Error {
	rv := objc.CallMethod[[]Error](e_, objc.GetSelector("underlyingErrors"))
	// TODO: convert slice items...
	return rv
}
