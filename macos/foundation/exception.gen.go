// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var ExceptionClass = _ExceptionClass{objc.GetClass("NSException")}

type _ExceptionClass struct {
	objc.Class
}

type IException interface {
	objc.IObject
	Raise()
	Name() ExceptionName
	Reason() string
	CallStackReturnAddresses() []Number
	CallStackSymbols() []string
}

type Exception struct {
	objc.Object
}

func MakeException(ptr unsafe.Pointer) Exception {
	return Exception{
		Object: objc.MakeObject(ptr),
	}
}

func (ec _ExceptionClass) Alloc() Exception {
	rv := objc.CallMethod[Exception](ec, objc.GetSelector("alloc"))
	return rv
}

func Exception_Alloc() Exception {
	return ExceptionClass.Alloc()
}

func (ec _ExceptionClass) New() Exception {
	rv := objc.CallMethod[Exception](ec, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewException() Exception {
	return ExceptionClass.New()
}

func Exception_New() Exception {
	return ExceptionClass.New()
}

func (e_ Exception) Init() Exception {
	rv := objc.CallMethod[Exception](e_, objc.GetSelector("init"))
	return rv
}

func Exception_Init() Exception {
	return ExceptionClass.Alloc().Init()
}

func (e_ Exception) Raise() {
	objc.CallMethod[objc.Void](e_, objc.GetSelector("raise"))
}

func (e_ Exception) Name() ExceptionName {
	rv := objc.CallMethod[ExceptionName](e_, objc.GetSelector("name"))
	return rv
}

func (e_ Exception) Reason() string {
	rv := objc.CallMethod[string](e_, objc.GetSelector("reason"))
	return rv
}

func (e_ Exception) CallStackReturnAddresses() []Number {
	rv := objc.CallMethod[[]Number](e_, objc.GetSelector("callStackReturnAddresses"))
	return rv
}

func (e_ Exception) CallStackSymbols() []string {
	rv := objc.CallMethod[[]string](e_, objc.GetSelector("callStackSymbols"))
	return rv
}
