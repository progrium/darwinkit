// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UserAppleScriptTask] class.
var UserAppleScriptTaskClass = _UserAppleScriptTaskClass{objc.GetClass("NSUserAppleScriptTask")}

type _UserAppleScriptTaskClass struct {
	objc.Class
}

// An interface definition for the [UserAppleScriptTask] class.
type IUserAppleScriptTask interface {
	IUserScriptTask
	ExecuteWithAppleEventCompletionHandler(event IAppleEventDescriptor, handler UserAppleScriptTaskCompletionHandler)
}

// An object that executes AppleScript scripts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserapplescripttask?language=objc
type UserAppleScriptTask struct {
	UserScriptTask
}

func UserAppleScriptTaskFrom(ptr unsafe.Pointer) UserAppleScriptTask {
	return UserAppleScriptTask{
		UserScriptTask: UserScriptTaskFrom(ptr),
	}
}

func (uc _UserAppleScriptTaskClass) Alloc() UserAppleScriptTask {
	rv := objc.Call[UserAppleScriptTask](uc, objc.Sel("alloc"))
	return rv
}

func UserAppleScriptTask_Alloc() UserAppleScriptTask {
	return UserAppleScriptTaskClass.Alloc()
}

func (uc _UserAppleScriptTaskClass) New() UserAppleScriptTask {
	rv := objc.Call[UserAppleScriptTask](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUserAppleScriptTask() UserAppleScriptTask {
	return UserAppleScriptTaskClass.New()
}

func (u_ UserAppleScriptTask) Init() UserAppleScriptTask {
	rv := objc.Call[UserAppleScriptTask](u_, objc.Sel("init"))
	return rv
}

func (u_ UserAppleScriptTask) InitWithURLError(url IURL, error IError) UserAppleScriptTask {
	rv := objc.Call[UserAppleScriptTask](u_, objc.Sel("initWithURL:error:"), objc.Ptr(url), objc.Ptr(error))
	return rv
}

// Return a user script task instance given a URL for a script file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserscripttask/1409998-initwithurl?language=objc
func NewUserAppleScriptTaskWithURLError(url IURL, error IError) UserAppleScriptTask {
	instance := UserAppleScriptTaskClass.Alloc().InitWithURLError(url, error)
	instance.Autorelease()
	return instance
}

// Execute the AppleScript script by sending it the specified Apple event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserapplescripttask/1416515-executewithappleevent?language=objc
func (u_ UserAppleScriptTask) ExecuteWithAppleEventCompletionHandler(event IAppleEventDescriptor, handler UserAppleScriptTaskCompletionHandler) {
	objc.Call[objc.Void](u_, objc.Sel("executeWithAppleEvent:completionHandler:"), objc.Ptr(event), handler)
}
