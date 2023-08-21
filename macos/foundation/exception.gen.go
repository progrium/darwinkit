// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Exception] class.
var ExceptionClass = _ExceptionClass{objc.GetClass("NSException")}

type _ExceptionClass struct {
	objc.Class
}

// An interface definition for the [Exception] class.
type IException interface {
	objc.IObject
	Raise()
	CallStackReturnAddresses() []Number
	Name() ExceptionName
	Reason() string
	UserInfo() Dictionary
	CallStackSymbols() []string
}

// An object that represents a special condition that interrupts the normal flow of program execution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception?language=objc
type Exception struct {
	objc.Object
}

func ExceptionFrom(ptr unsafe.Pointer) Exception {
	return Exception{
		Object: objc.ObjectFrom(ptr),
	}
}

func (e_ Exception) InitWithNameReasonUserInfo(aName ExceptionName, aReason string, aUserInfo Dictionary) Exception {
	rv := objc.Call[Exception](e_, objc.Sel("initWithName:reason:userInfo:"), aName, aReason, aUserInfo)
	return rv
}

// Initializes and returns a newly allocated exception object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/1414506-initwithname?language=objc
func NewExceptionWithNameReasonUserInfo(aName ExceptionName, aReason string, aUserInfo Dictionary) Exception {
	instance := ExceptionClass.Alloc().InitWithNameReasonUserInfo(aName, aReason, aUserInfo)
	instance.Autorelease()
	return instance
}

func (ec _ExceptionClass) Alloc() Exception {
	rv := objc.Call[Exception](ec, objc.Sel("alloc"))
	return rv
}

func Exception_Alloc() Exception {
	return ExceptionClass.Alloc()
}

func (ec _ExceptionClass) New() Exception {
	rv := objc.Call[Exception](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewException() Exception {
	return ExceptionClass.New()
}

func (e_ Exception) Init() Exception {
	rv := objc.Call[Exception](e_, objc.Sel("init"))
	return rv
}

// Raises the receiver, causing program flow to jump to the local exception handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/1416135-raise?language=objc
func (e_ Exception) Raise() {
	objc.Call[objc.Void](e_, objc.Sel("raise"))
}

// Creates and returns an exception object . [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/1569530-exceptionwithname?language=objc
func (ec _ExceptionClass) ExceptionWithNameReasonUserInfo(name ExceptionName, reason string, userInfo Dictionary) Exception {
	rv := objc.Call[Exception](ec, objc.Sel("exceptionWithName:reason:userInfo:"), name, reason, userInfo)
	return rv
}

// Creates and returns an exception object . [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/1569530-exceptionwithname?language=objc
func Exception_ExceptionWithNameReasonUserInfo(name ExceptionName, reason string, userInfo Dictionary) Exception {
	return ExceptionClass.ExceptionWithNameReasonUserInfo(name, reason, userInfo)
}

// The call return addresses related to a raised exception. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/1412165-callstackreturnaddresses?language=objc
func (e_ Exception) CallStackReturnAddresses() []Number {
	rv := objc.Call[[]Number](e_, objc.Sel("callStackReturnAddresses"))
	return rv
}

// A string used to uniquely identify the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/1410925-name?language=objc
func (e_ Exception) Name() ExceptionName {
	rv := objc.Call[ExceptionName](e_, objc.Sel("name"))
	return rv
}

// A string containing a “human-readable” reason for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/1415537-reason?language=objc
func (e_ Exception) Reason() string {
	rv := objc.Call[string](e_, objc.Sel("reason"))
	return rv
}

// A dictionary containing application-specific data pertaining to the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/1418149-userinfo?language=objc
func (e_ Exception) UserInfo() Dictionary {
	rv := objc.Call[Dictionary](e_, objc.Sel("userInfo"))
	return rv
}

// An array containing the current call stack symbols. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/1416845-callstacksymbols?language=objc
func (e_ Exception) CallStackSymbols() []string {
	rv := objc.Call[[]string](e_, objc.Sel("callStackSymbols"))
	return rv
}
