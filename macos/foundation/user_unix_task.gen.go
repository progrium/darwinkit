// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UserUnixTask] class.
var UserUnixTaskClass = _UserUnixTaskClass{objc.GetClass("NSUserUnixTask")}

type _UserUnixTaskClass struct {
	objc.Class
}

// An interface definition for the [UserUnixTask] class.
type IUserUnixTask interface {
	IUserScriptTask
	ExecuteWithArgumentsCompletionHandler(arguments []string, handler UserUnixTaskCompletionHandler)
	StandardOutput() FileHandle
	SetStandardOutput(value IFileHandle)
	StandardError() FileHandle
	SetStandardError(value IFileHandle)
	StandardInput() FileHandle
	SetStandardInput(value IFileHandle)
}

// An object that executes unix applications. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserunixtask?language=objc
type UserUnixTask struct {
	UserScriptTask
}

func UserUnixTaskFrom(ptr unsafe.Pointer) UserUnixTask {
	return UserUnixTask{
		UserScriptTask: UserScriptTaskFrom(ptr),
	}
}

func (uc _UserUnixTaskClass) Alloc() UserUnixTask {
	rv := objc.Call[UserUnixTask](uc, objc.Sel("alloc"))
	return rv
}

func UserUnixTask_Alloc() UserUnixTask {
	return UserUnixTaskClass.Alloc()
}

func (uc _UserUnixTaskClass) New() UserUnixTask {
	rv := objc.Call[UserUnixTask](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUserUnixTask() UserUnixTask {
	return UserUnixTaskClass.New()
}

func (u_ UserUnixTask) Init() UserUnixTask {
	rv := objc.Call[UserUnixTask](u_, objc.Sel("init"))
	return rv
}

func (u_ UserUnixTask) InitWithURLError(url IURL, error IError) UserUnixTask {
	rv := objc.Call[UserUnixTask](u_, objc.Sel("initWithURL:error:"), objc.Ptr(url), objc.Ptr(error))
	return rv
}

// Return a user script task instance given a URL for a script file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserscripttask/1409998-initwithurl?language=objc
func UserUnixTask_InitWithURLError(url IURL, error IError) UserUnixTask {
	return UserUnixTaskClass.Alloc().InitWithURLError(url, error)
}

// Execute the unix script with the specified arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserunixtask/1412077-executewitharguments?language=objc
func (u_ UserUnixTask) ExecuteWithArgumentsCompletionHandler(arguments []string, handler UserUnixTaskCompletionHandler) {
	objc.Call[objc.Void](u_, objc.Sel("executeWithArguments:completionHandler:"), arguments, handler)
}

// The standard output stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserunixtask/1418151-standardoutput?language=objc
func (u_ UserUnixTask) StandardOutput() FileHandle {
	rv := objc.Call[FileHandle](u_, objc.Sel("standardOutput"))
	return rv
}

// The standard output stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserunixtask/1418151-standardoutput?language=objc
func (u_ UserUnixTask) SetStandardOutput(value IFileHandle) {
	objc.Call[objc.Void](u_, objc.Sel("setStandardOutput:"), objc.Ptr(value))
}

// The standard error stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserunixtask/1411522-standarderror?language=objc
func (u_ UserUnixTask) StandardError() FileHandle {
	rv := objc.Call[FileHandle](u_, objc.Sel("standardError"))
	return rv
}

// The standard error stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserunixtask/1411522-standarderror?language=objc
func (u_ UserUnixTask) SetStandardError(value IFileHandle) {
	objc.Call[objc.Void](u_, objc.Sel("setStandardError:"), objc.Ptr(value))
}

// The standard input stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserunixtask/1407821-standardinput?language=objc
func (u_ UserUnixTask) StandardInput() FileHandle {
	rv := objc.Call[FileHandle](u_, objc.Sel("standardInput"))
	return rv
}

// The standard input stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserunixtask/1407821-standardinput?language=objc
func (u_ UserUnixTask) SetStandardInput(value IFileHandle) {
	objc.Call[objc.Void](u_, objc.Sel("setStandardInput:"), objc.Ptr(value))
}
