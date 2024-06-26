// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [Task] class.
var TaskClass = _TaskClass{objc.GetClass("NSTask")}

type _TaskClass struct {
	objc.Class
}

// An interface definition for the [Task] class.
type ITask interface {
	objc.IObject
	Suspend() bool
	Interrupt()
	Resume() bool
	Terminate()
	WaitUntilExit()
	LaunchAndReturnError(error unsafe.Pointer) bool
	StandardInput() objc.Object
	SetStandardInput(value objc.IObject)
	TerminationReason() TaskTerminationReason
	StandardOutput() objc.Object
	SetStandardOutput(value objc.IObject)
	TerminationStatus() int
	StandardError() objc.Object
	SetStandardError(value objc.IObject)
	ProcessIdentifier() int
	IsRunning() bool
	Environment() map[string]string
	SetEnvironment(value map[string]string)
	Arguments() []string
	SetArguments(value []string)
	ExecutableURL() URL
	SetExecutableURL(value IURL)
	TerminationHandler() func(arg0 Task)
	SetTerminationHandler(value func(arg0 Task))
	CurrentDirectoryURL() URL
	SetCurrentDirectoryURL(value IURL)
	QualityOfService() QualityOfService
	SetQualityOfService(value QualityOfService)
}

// An object that represents a subprocess of the current process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask?language=objc
type Task struct {
	objc.Object
}

func TaskFrom(ptr unsafe.Pointer) Task {
	return Task{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ Task) Init() Task {
	rv := objc.Call[Task](t_, objc.Sel("init"))
	return rv
}

func (tc _TaskClass) Alloc() Task {
	rv := objc.Call[Task](tc, objc.Sel("alloc"))
	return rv
}

func (tc _TaskClass) New() Task {
	rv := objc.Call[Task](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTask() Task {
	return TaskClass.New()
}

// Creates and runs a task with a specified executable and arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/2890108-launchedtaskwithexecutableurl?language=objc
func (tc _TaskClass) LaunchedTaskWithExecutableURLArgumentsErrorTerminationHandler(url IURL, arguments []string, error unsafe.Pointer, terminationHandler func(arg0 Task)) Task {
	rv := objc.Call[Task](tc, objc.Sel("launchedTaskWithExecutableURL:arguments:error:terminationHandler:"), url, arguments, error, terminationHandler)
	return rv
}

// Creates and runs a task with a specified executable and arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/2890108-launchedtaskwithexecutableurl?language=objc
func Task_LaunchedTaskWithExecutableURLArgumentsErrorTerminationHandler(url IURL, arguments []string, error unsafe.Pointer, terminationHandler func(arg0 Task)) Task {
	return TaskClass.LaunchedTaskWithExecutableURLArgumentsErrorTerminationHandler(url, arguments, error, terminationHandler)
}

// Suspends execution of the receiver task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/1411590-suspend?language=objc
func (t_ Task) Suspend() bool {
	rv := objc.Call[bool](t_, objc.Sel("suspend"))
	return rv
}

// Sends an interrupt signal to the receiver and all of its subtasks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/1410874-interrupt?language=objc
func (t_ Task) Interrupt() {
	objc.Call[objc.Void](t_, objc.Sel("interrupt"))
}

// Resumes execution of a suspended task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/1407819-resume?language=objc
func (t_ Task) Resume() bool {
	rv := objc.Call[bool](t_, objc.Sel("resume"))
	return rv
}

// Sends a terminate signal to the receiver and all of its subtasks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/1409620-terminate?language=objc
func (t_ Task) Terminate() {
	objc.Call[objc.Void](t_, objc.Sel("terminate"))
}

// Blocks the process until the receiver is finished. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/1415808-waituntilexit?language=objc
func (t_ Task) WaitUntilExit() {
	objc.Call[objc.Void](t_, objc.Sel("waitUntilExit"))
}

// Runs the process with the current environment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/2890105-launchandreturnerror?language=objc
func (t_ Task) LaunchAndReturnError(error unsafe.Pointer) bool {
	rv := objc.Call[bool](t_, objc.Sel("launchAndReturnError:"), error)
	return rv
}

// The standard input for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/1411576-standardinput?language=objc
func (t_ Task) StandardInput() objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("standardInput"))
	return rv
}

// The standard input for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/1411576-standardinput?language=objc
func (t_ Task) SetStandardInput(value objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setStandardInput:"), value)
}

// The reason the system terminated the task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/1415605-terminationreason?language=objc
func (t_ Task) TerminationReason() TaskTerminationReason {
	rv := objc.Call[TaskTerminationReason](t_, objc.Sel("terminationReason"))
	return rv
}

// The standard output for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/1407627-standardoutput?language=objc
func (t_ Task) StandardOutput() objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("standardOutput"))
	return rv
}

// The standard output for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/1407627-standardoutput?language=objc
func (t_ Task) SetStandardOutput(value objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setStandardOutput:"), value)
}

// The exit status the receiver’s executable returns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/1415801-terminationstatus?language=objc
func (t_ Task) TerminationStatus() int {
	rv := objc.Call[int](t_, objc.Sel("terminationStatus"))
	return rv
}

// The standard error for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/1414916-standarderror?language=objc
func (t_ Task) StandardError() objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("standardError"))
	return rv
}

// The standard error for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/1414916-standarderror?language=objc
func (t_ Task) SetStandardError(value objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setStandardError:"), value)
}

// The receiver’s process identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/1412022-processidentifier?language=objc
func (t_ Task) ProcessIdentifier() int {
	rv := objc.Call[int](t_, objc.Sel("processIdentifier"))
	return rv
}

// A status that indicates whether the receiver is still running. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/1415788-running?language=objc
func (t_ Task) IsRunning() bool {
	rv := objc.Call[bool](t_, objc.Sel("isRunning"))
	return rv
}

// The environment for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/1409412-environment?language=objc
func (t_ Task) Environment() map[string]string {
	rv := objc.Call[map[string]string](t_, objc.Sel("environment"))
	return rv
}

// The environment for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/1409412-environment?language=objc
func (t_ Task) SetEnvironment(value map[string]string) {
	objc.Call[objc.Void](t_, objc.Sel("setEnvironment:"), value)
}

// The command arguments that the system uses to launch the executable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/1408983-arguments?language=objc
func (t_ Task) Arguments() []string {
	rv := objc.Call[[]string](t_, objc.Sel("arguments"))
	return rv
}

// The command arguments that the system uses to launch the executable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/1408983-arguments?language=objc
func (t_ Task) SetArguments(value []string) {
	objc.Call[objc.Void](t_, objc.Sel("setArguments:"), value)
}

// The receiver’s executable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/2890106-executableurl?language=objc
func (t_ Task) ExecutableURL() URL {
	rv := objc.Call[URL](t_, objc.Sel("executableURL"))
	return rv
}

// The receiver’s executable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/2890106-executableurl?language=objc
func (t_ Task) SetExecutableURL(value IURL) {
	objc.Call[objc.Void](t_, objc.Sel("setExecutableURL:"), value)
}

// A completion block the system invokes when the task completes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/1408746-terminationhandler?language=objc
func (t_ Task) TerminationHandler() func(arg0 Task) {
	rv := objc.Call[func(arg0 Task)](t_, objc.Sel("terminationHandler"))
	return rv
}

// A completion block the system invokes when the task completes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/1408746-terminationhandler?language=objc
func (t_ Task) SetTerminationHandler(value func(arg0 Task)) {
	objc.Call[objc.Void](t_, objc.Sel("setTerminationHandler:"), value)
}

// The current directory for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/2890107-currentdirectoryurl?language=objc
func (t_ Task) CurrentDirectoryURL() URL {
	rv := objc.Call[URL](t_, objc.Sel("currentDirectoryURL"))
	return rv
}

// The current directory for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/2890107-currentdirectoryurl?language=objc
func (t_ Task) SetCurrentDirectoryURL(value IURL) {
	objc.Call[objc.Void](t_, objc.Sel("setCurrentDirectoryURL:"), value)
}

// The default quality of service level the system applies to operations the task executes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/1415794-qualityofservice?language=objc
func (t_ Task) QualityOfService() QualityOfService {
	rv := objc.Call[QualityOfService](t_, objc.Sel("qualityOfService"))
	return rv
}

// The default quality of service level the system applies to operations the task executes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstask/1415794-qualityofservice?language=objc
func (t_ Task) SetQualityOfService(value QualityOfService) {
	objc.Call[objc.Void](t_, objc.Sel("setQualityOfService:"), value)
}
