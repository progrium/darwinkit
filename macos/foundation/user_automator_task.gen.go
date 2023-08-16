// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UserAutomatorTask] class.
var UserAutomatorTaskClass = _UserAutomatorTaskClass{objc.GetClass("NSUserAutomatorTask")}

type _UserAutomatorTaskClass struct {
	objc.Class
}

// An interface definition for the [UserAutomatorTask] class.
type IUserAutomatorTask interface {
	IUserScriptTask
	ExecuteWithInputCompletionHandler(input objc.IObject, handler UserAutomatorTaskCompletionHandler)
	Variables() map[string]objc.Object
	SetVariables(value map[string]objc.IObject)
}

// An object that executes Automator workflows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserautomatortask?language=objc
type UserAutomatorTask struct {
	UserScriptTask
}

func UserAutomatorTaskFrom(ptr unsafe.Pointer) UserAutomatorTask {
	return UserAutomatorTask{
		UserScriptTask: UserScriptTaskFrom(ptr),
	}
}

func (uc _UserAutomatorTaskClass) Alloc() UserAutomatorTask {
	rv := objc.Call[UserAutomatorTask](uc, objc.Sel("alloc"))
	return rv
}

func UserAutomatorTask_Alloc() UserAutomatorTask {
	return UserAutomatorTaskClass.Alloc()
}

func (uc _UserAutomatorTaskClass) New() UserAutomatorTask {
	rv := objc.Call[UserAutomatorTask](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUserAutomatorTask() UserAutomatorTask {
	return UserAutomatorTaskClass.New()
}

func (u_ UserAutomatorTask) Init() UserAutomatorTask {
	rv := objc.Call[UserAutomatorTask](u_, objc.Sel("init"))
	return rv
}

func (u_ UserAutomatorTask) InitWithURLError(url IURL, error IError) UserAutomatorTask {
	rv := objc.Call[UserAutomatorTask](u_, objc.Sel("initWithURL:error:"), objc.Ptr(url), objc.Ptr(error))
	return rv
}

// Return a user script task instance given a URL for a script file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserscripttask/1409998-initwithurl?language=objc
func UserAutomatorTask_InitWithURLError(url IURL, error IError) UserAutomatorTask {
	return UserAutomatorTaskClass.Alloc().InitWithURLError(url, error)
}

// Execute the Automator workflow by providing it as securely coded input. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserautomatortask/1418079-executewithinput?language=objc
func (u_ UserAutomatorTask) ExecuteWithInputCompletionHandler(input objc.IObject, handler UserAutomatorTaskCompletionHandler) {
	objc.Call[objc.Void](u_, objc.Sel("executeWithInput:completionHandler:"), input, handler)
}

// The variables required by the Automator workflow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserautomatortask/1418099-variables?language=objc
func (u_ UserAutomatorTask) Variables() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](u_, objc.Sel("variables"))
	return rv
}

// The variables required by the Automator workflow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserautomatortask/1418099-variables?language=objc
func (u_ UserAutomatorTask) SetVariables(value map[string]objc.IObject) {
	objc.Call[objc.Void](u_, objc.Sel("setVariables:"), value)
}
