// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Thread] class.
var ThreadClass = _ThreadClass{objc.GetClass("NSThread")}

type _ThreadClass struct {
	objc.Class
}

// An interface definition for the [Thread] class.
type IThread interface {
	objc.IObject
	Main()
	Start()
	Cancel()
	IsCancelled() bool
	IsExecuting() bool
	Name() string
	SetName(value string)
	QualityOfService() QualityOfService
	SetQualityOfService(value QualityOfService)
	IsMainThread() bool
	ThreadPriority() float64
	StackSize() uint
	SetStackSize(value uint)
	IsFinished() bool
	ThreadDictionary() MutableDictionary
}

// A thread of execution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread?language=objc
type Thread struct {
	objc.Object
}

func ThreadFrom(ptr unsafe.Pointer) Thread {
	return Thread{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ Thread) InitWithBlock(block func()) Thread {
	rv := objc.Call[Thread](t_, objc.Sel("initWithBlock:"), block)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/2088561-initwithblock?language=objc
func NewThreadWithBlock(block func()) Thread {
	instance := ThreadClass.Alloc().InitWithBlock(block)
	instance.Autorelease()
	return instance
}

func (t_ Thread) Init() Thread {
	rv := objc.Call[Thread](t_, objc.Sel("init"))
	return rv
}

func (t_ Thread) InitWithTargetSelectorObject(target objc.IObject, selector objc.Selector, argument objc.IObject) Thread {
	rv := objc.Call[Thread](t_, objc.Sel("initWithTarget:selector:object:"), target, selector, argument)
	return rv
}

// Returns an NSThread object initialized with the given arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1414773-initwithtarget?language=objc
func NewThreadWithTargetSelectorObject(target objc.IObject, selector objc.Selector, argument objc.IObject) Thread {
	instance := ThreadClass.Alloc().InitWithTargetSelectorObject(target, selector, argument)
	instance.Autorelease()
	return instance
}

func (tc _ThreadClass) Alloc() Thread {
	rv := objc.Call[Thread](tc, objc.Sel("alloc"))
	return rv
}

func Thread_Alloc() Thread {
	return ThreadClass.Alloc()
}

func (tc _ThreadClass) New() Thread {
	rv := objc.Call[Thread](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewThread() Thread {
	return ThreadClass.New()
}

// Terminates the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1409404-exit?language=objc
func (tc _ThreadClass) Exit() {
	objc.Call[objc.Void](tc, objc.Sel("exit"))
}

// Terminates the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1409404-exit?language=objc
func Thread_Exit() {
	ThreadClass.Exit()
}

// Blocks the current thread until the time specified. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1415959-sleepuntildate?language=objc
func (tc _ThreadClass) SleepUntilDate(date IDate) {
	objc.Call[objc.Void](tc, objc.Sel("sleepUntilDate:"), objc.Ptr(date))
}

// Blocks the current thread until the time specified. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1415959-sleepuntildate?language=objc
func Thread_SleepUntilDate(date IDate) {
	ThreadClass.SleepUntilDate(date)
}

// The main entry point routine for the thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1418421-main?language=objc
func (t_ Thread) Main() {
	objc.Call[objc.Void](t_, objc.Sel("main"))
}

// Sleeps the thread for a given time interval. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1413673-sleepfortimeinterval?language=objc
func (tc _ThreadClass) SleepForTimeInterval(ti TimeInterval) {
	objc.Call[objc.Void](tc, objc.Sel("sleepForTimeInterval:"), ti)
}

// Sleeps the thread for a given time interval. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1413673-sleepfortimeinterval?language=objc
func Thread_SleepForTimeInterval(ti TimeInterval) {
	ThreadClass.SleepForTimeInterval(ti)
}

// Starts the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1418166-start?language=objc
func (t_ Thread) Start() {
	objc.Call[objc.Void](t_, objc.Sel("start"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/2088563-detachnewthreadwithblock?language=objc
func (tc _ThreadClass) DetachNewThreadWithBlock(block func()) {
	objc.Call[objc.Void](tc, objc.Sel("detachNewThreadWithBlock:"), block)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/2088563-detachnewthreadwithblock?language=objc
func Thread_DetachNewThreadWithBlock(block func()) {
	ThreadClass.DetachNewThreadWithBlock(block)
}

// Returns whether the application is multithreaded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1410702-ismultithreaded?language=objc
func (tc _ThreadClass) IsMultiThreaded() bool {
	rv := objc.Call[bool](tc, objc.Sel("isMultiThreaded"))
	return rv
}

// Returns whether the application is multithreaded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1410702-ismultithreaded?language=objc
func Thread_IsMultiThreaded() bool {
	return ThreadClass.IsMultiThreaded()
}

// Detaches a new thread and uses the specified selector as the thread entry point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1415633-detachnewthreadselector?language=objc
func (tc _ThreadClass) DetachNewThreadSelectorToTargetWithObject(selector objc.Selector, target objc.IObject, argument objc.IObject) {
	objc.Call[objc.Void](tc, objc.Sel("detachNewThreadSelector:toTarget:withObject:"), selector, target, argument)
}

// Detaches a new thread and uses the specified selector as the thread entry point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1415633-detachnewthreadselector?language=objc
func Thread_DetachNewThreadSelectorToTargetWithObject(selector objc.Selector, target objc.IObject, argument objc.IObject) {
	ThreadClass.DetachNewThreadSelectorToTargetWithObject(selector, target, argument)
}

// Sets the current thread’s priority. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1407523-setthreadpriority?language=objc
func (tc _ThreadClass) SetThreadPriority(p float64) bool {
	rv := objc.Call[bool](tc, objc.Sel("setThreadPriority:"), p)
	return rv
}

// Sets the current thread’s priority. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1407523-setthreadpriority?language=objc
func Thread_SetThreadPriority(p float64) bool {
	return ThreadClass.SetThreadPriority(p)
}

// Changes the cancelled state of the receiver to indicate that it should exit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1411303-cancel?language=objc
func (t_ Thread) Cancel() {
	objc.Call[objc.Void](t_, objc.Sel("cancel"))
}

// A Boolean value that indicates whether the receiver is cancelled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1417366-cancelled?language=objc
func (t_ Thread) IsCancelled() bool {
	rv := objc.Call[bool](t_, objc.Sel("isCancelled"))
	return rv
}

// A Boolean value that indicates whether the receiver is executing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1411240-executing?language=objc
func (t_ Thread) IsExecuting() bool {
	rv := objc.Call[bool](t_, objc.Sel("isExecuting"))
	return rv
}

// Returns an array containing the call stack return addresses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1409565-callstackreturnaddresses?language=objc
func (tc _ThreadClass) CallStackReturnAddresses() []Number {
	rv := objc.Call[[]Number](tc, objc.Sel("callStackReturnAddresses"))
	return rv
}

// Returns an array containing the call stack return addresses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1409565-callstackreturnaddresses?language=objc
func Thread_CallStackReturnAddresses() []Number {
	return ThreadClass.CallStackReturnAddresses()
}

// The name of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1414122-name?language=objc
func (t_ Thread) Name() string {
	rv := objc.Call[string](t_, objc.Sel("name"))
	return rv
}

// The name of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1414122-name?language=objc
func (t_ Thread) SetName(value string) {
	objc.Call[objc.Void](t_, objc.Sel("setName:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1409426-qualityofservice?language=objc
func (t_ Thread) QualityOfService() QualityOfService {
	rv := objc.Call[QualityOfService](t_, objc.Sel("qualityOfService"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1409426-qualityofservice?language=objc
func (t_ Thread) SetQualityOfService(value QualityOfService) {
	objc.Call[objc.Void](t_, objc.Sel("setQualityOfService:"), value)
}

// A Boolean value that indicates whether the receiver is the main thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1408455-ismainthread?language=objc
func (t_ Thread) IsMainThread() bool {
	rv := objc.Call[bool](t_, objc.Sel("isMainThread"))
	return rv
}

// The receiver’s priority [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1411927-threadpriority?language=objc
func (t_ Thread) ThreadPriority() float64 {
	rv := objc.Call[float64](t_, objc.Sel("threadPriority"))
	return rv
}

// Returns an array containing the call stack symbols. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1414836-callstacksymbols?language=objc
func (tc _ThreadClass) CallStackSymbols() []string {
	rv := objc.Call[[]string](tc, objc.Sel("callStackSymbols"))
	return rv
}

// Returns an array containing the call stack symbols. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1414836-callstacksymbols?language=objc
func Thread_CallStackSymbols() []string {
	return ThreadClass.CallStackSymbols()
}

// The stack size of the receiver, in bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1415190-stacksize?language=objc
func (t_ Thread) StackSize() uint {
	rv := objc.Call[uint](t_, objc.Sel("stackSize"))
	return rv
}

// The stack size of the receiver, in bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1415190-stacksize?language=objc
func (t_ Thread) SetStackSize(value uint) {
	objc.Call[objc.Void](t_, objc.Sel("setStackSize:"), value)
}

// Returns the NSThread object representing the main thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1414782-mainthread?language=objc
func (tc _ThreadClass) MainThread() Thread {
	rv := objc.Call[Thread](tc, objc.Sel("mainThread"))
	return rv
}

// Returns the NSThread object representing the main thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1414782-mainthread?language=objc
func Thread_MainThread() Thread {
	return ThreadClass.MainThread()
}

// A Boolean value that indicates whether the receiver has finished execution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1409297-finished?language=objc
func (t_ Thread) IsFinished() bool {
	rv := objc.Call[bool](t_, objc.Sel("isFinished"))
	return rv
}

// Returns the thread object representing the current thread of execution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1410679-currentthread?language=objc
func (tc _ThreadClass) CurrentThread() Thread {
	rv := objc.Call[Thread](tc, objc.Sel("currentThread"))
	return rv
}

// Returns the thread object representing the current thread of execution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1410679-currentthread?language=objc
func Thread_CurrentThread() Thread {
	return ThreadClass.CurrentThread()
}

// The thread object's dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsthread/1411433-threaddictionary?language=objc
func (t_ Thread) ThreadDictionary() MutableDictionary {
	rv := objc.Call[MutableDictionary](t_, objc.Sel("threadDictionary"))
	return rv
}
