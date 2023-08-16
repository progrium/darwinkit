// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssertionHandler] class.
var AssertionHandlerClass = _AssertionHandlerClass{objc.GetClass("NSAssertionHandler")}

type _AssertionHandlerClass struct {
	objc.Class
}

// An interface definition for the [AssertionHandler] class.
type IAssertionHandler interface {
	objc.IObject
	HandleFailureInMethodObjectFileLineNumberDescription(selector objc.Selector, object objc.IObject, fileName string, line int, format string)
	HandleFailureInFunctionFileLineNumberDescription(functionName string, fileName string, line int, format string)
}

// An object that logs an assertion to the console. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsassertionhandler?language=objc
type AssertionHandler struct {
	objc.Object
}

func AssertionHandlerFrom(ptr unsafe.Pointer) AssertionHandler {
	return AssertionHandler{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssertionHandlerClass) Alloc() AssertionHandler {
	rv := objc.Call[AssertionHandler](ac, objc.Sel("alloc"))
	return rv
}

func AssertionHandler_Alloc() AssertionHandler {
	return AssertionHandlerClass.Alloc()
}

func (ac _AssertionHandlerClass) New() AssertionHandler {
	rv := objc.Call[AssertionHandler](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssertionHandler() AssertionHandler {
	return AssertionHandlerClass.New()
}

func (a_ AssertionHandler) Init() AssertionHandler {
	rv := objc.Call[AssertionHandler](a_, objc.Sel("init"))
	return rv
}

// Logs (using NSLog) an error message that includes the name of the method that failed, the class name of the object, the name of the source file, and the line number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsassertionhandler/1569513-handlefailureinmethod?language=objc
func (a_ AssertionHandler) HandleFailureInMethodObjectFileLineNumberDescription(selector objc.Selector, object objc.IObject, fileName string, line int, format string) {
	objc.Call[objc.Void](a_, objc.Sel("handleFailureInMethod:object:file:lineNumber:description:"), selector, object, fileName, line, format)
}

// Logs (using NSLog) an error message that includes the name of the function, the name of the file, and the line number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsassertionhandler/1569532-handlefailureinfunction?language=objc
func (a_ AssertionHandler) HandleFailureInFunctionFileLineNumberDescription(functionName string, fileName string, line int, format string) {
	objc.Call[objc.Void](a_, objc.Sel("handleFailureInFunction:file:lineNumber:description:"), functionName, fileName, line, format)
}

// Returns the NSAssertionHandler object associated with the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsassertionhandler/1417391-currenthandler?language=objc
func (ac _AssertionHandlerClass) CurrentHandler() AssertionHandler {
	rv := objc.Call[AssertionHandler](ac, objc.Sel("currentHandler"))
	return rv
}

// Returns the NSAssertionHandler object associated with the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsassertionhandler/1417391-currenthandler?language=objc
func AssertionHandler_CurrentHandler() AssertionHandler {
	return AssertionHandlerClass.CurrentHandler()
}
