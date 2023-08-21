// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UserScriptTask] class.
var UserScriptTaskClass = _UserScriptTaskClass{objc.GetClass("NSUserScriptTask")}

type _UserScriptTaskClass struct {
	objc.Class
}

// An interface definition for the [UserScriptTask] class.
type IUserScriptTask interface {
	objc.IObject
	ExecuteWithCompletionHandler(handler UserScriptTaskCompletionHandler)
	ScriptURL() URL
}

// An object that executes scripts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserscripttask?language=objc
type UserScriptTask struct {
	objc.Object
}

func UserScriptTaskFrom(ptr unsafe.Pointer) UserScriptTask {
	return UserScriptTask{
		Object: objc.ObjectFrom(ptr),
	}
}

func (u_ UserScriptTask) InitWithURLError(url IURL, error IError) UserScriptTask {
	rv := objc.Call[UserScriptTask](u_, objc.Sel("initWithURL:error:"), objc.Ptr(url), objc.Ptr(error))
	return rv
}

// Return a user script task instance given a URL for a script file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserscripttask/1409998-initwithurl?language=objc
func NewUserScriptTaskWithURLError(url IURL, error IError) UserScriptTask {
	instance := UserScriptTaskClass.Alloc().InitWithURLError(url, error)
	instance.Autorelease()
	return instance
}

func (uc _UserScriptTaskClass) Alloc() UserScriptTask {
	rv := objc.Call[UserScriptTask](uc, objc.Sel("alloc"))
	return rv
}

func UserScriptTask_Alloc() UserScriptTask {
	return UserScriptTaskClass.Alloc()
}

func (uc _UserScriptTaskClass) New() UserScriptTask {
	rv := objc.Call[UserScriptTask](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUserScriptTask() UserScriptTask {
	return UserScriptTaskClass.New()
}

func (u_ UserScriptTask) Init() UserScriptTask {
	rv := objc.Call[UserScriptTask](u_, objc.Sel("init"))
	return rv
}

// Executes the script with no input and ignoring any result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserscripttask/1410967-executewithcompletionhandler?language=objc
func (u_ UserScriptTask) ExecuteWithCompletionHandler(handler UserScriptTaskCompletionHandler) {
	objc.Call[objc.Void](u_, objc.Sel("executeWithCompletionHandler:"), handler)
}

// The URL of the script file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserscripttask/1408618-scripturl?language=objc
func (u_ UserScriptTask) ScriptURL() URL {
	rv := objc.Call[URL](u_, objc.Sel("scriptURL"))
	return rv
}
