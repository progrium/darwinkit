// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/corefoundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RunLoop] class.
var RunLoopClass = _RunLoopClass{objc.GetClass("NSRunLoop")}

type _RunLoopClass struct {
	objc.Class
}

// An interface definition for the [RunLoop] class.
type IRunLoop interface {
	objc.IObject
	AcceptInputForModeBeforeDate(mode RunLoopMode, limitDate IDate)
	CancelPerformSelectorTargetArgument(aSelector objc.Selector, target objc.IObject, arg objc.IObject)
	GetCFRunLoop() corefoundation.RunLoopRef
	RunModeBeforeDate(mode RunLoopMode, limitDate IDate) bool
	CancelPerformSelectorsWithTarget(target objc.IObject)
	AddTimerForMode(timer ITimer, mode RunLoopMode)
	PerformInModesBlock(modes []RunLoopMode, block func())
	AddPortForMode(aPort IPort, mode RunLoopMode)
	PerformBlock(block func())
	LimitDateForMode(mode RunLoopMode) Date
	PerformSelectorTargetArgumentOrderModes(aSelector objc.Selector, target objc.IObject, arg objc.IObject, order uint, modes []RunLoopMode)
	RunUntilDate(limitDate IDate)
	Run()
	RemovePortForMode(aPort IPort, mode RunLoopMode)
	CurrentMode() RunLoopMode
}

// The programmatic interface to objects that manage input sources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrunloop?language=objc
type RunLoop struct {
	objc.Object
}

func RunLoopFrom(ptr unsafe.Pointer) RunLoop {
	return RunLoop{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RunLoopClass) Alloc() RunLoop {
	rv := objc.Call[RunLoop](rc, objc.Sel("alloc"))
	return rv
}

func RunLoop_Alloc() RunLoop {
	return RunLoopClass.Alloc()
}

func (rc _RunLoopClass) New() RunLoop {
	rv := objc.Call[RunLoop](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRunLoop() RunLoop {
	return RunLoopClass.New()
}

func (r_ RunLoop) Init() RunLoop {
	rv := objc.Call[RunLoop](r_, objc.Sel("init"))
	return rv
}

// Runs the loop once or until the specified date, accepting input only for the specified mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrunloop/1417143-acceptinputformode?language=objc
func (r_ RunLoop) AcceptInputForModeBeforeDate(mode RunLoopMode, limitDate IDate) {
	objc.Call[objc.Void](r_, objc.Sel("acceptInputForMode:beforeDate:"), mode, objc.Ptr(limitDate))
}

// Cancels the sending of a previously scheduled message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrunloop/1418077-cancelperformselector?language=objc
func (r_ RunLoop) CancelPerformSelectorTargetArgument(aSelector objc.Selector, target objc.IObject, arg objc.IObject) {
	objc.Call[objc.Void](r_, objc.Sel("cancelPerformSelector:target:argument:"), aSelector, target, arg)
}

// Returns the receiver's underlying [corefoundation/cfrunloop] object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrunloop/1410140-getcfrunloop?language=objc
func (r_ RunLoop) GetCFRunLoop() corefoundation.RunLoopRef {
	rv := objc.Call[corefoundation.RunLoopRef](r_, objc.Sel("getCFRunLoop"))
	return rv
}

// Runs the loop once, blocking for input in the specified mode until a given date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrunloop/1411525-runmode?language=objc
func (r_ RunLoop) RunModeBeforeDate(mode RunLoopMode, limitDate IDate) bool {
	rv := objc.Call[bool](r_, objc.Sel("runMode:beforeDate:"), mode, objc.Ptr(limitDate))
	return rv
}

// Cancels all outstanding ordered performs scheduled with a given target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrunloop/1414208-cancelperformselectorswithtarget?language=objc
func (r_ RunLoop) CancelPerformSelectorsWithTarget(target objc.IObject) {
	objc.Call[objc.Void](r_, objc.Sel("cancelPerformSelectorsWithTarget:"), target)
}

// Registers a given timer with a given input mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrunloop/1418468-addtimer?language=objc
func (r_ RunLoop) AddTimerForMode(timer ITimer, mode RunLoopMode) {
	objc.Call[objc.Void](r_, objc.Sel("addTimer:forMode:"), objc.Ptr(timer), mode)
}

// Schedules a block that the run loop invokes when it’s running in any of the specified modes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrunloop/2091880-performinmodes?language=objc
func (r_ RunLoop) PerformInModesBlock(modes []RunLoopMode, block func()) {
	objc.Call[objc.Void](r_, objc.Sel("performInModes:block:"), modes, block)
}

// Adds a port as an input source to the specified mode of the run loop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrunloop/1417637-addport?language=objc
func (r_ RunLoop) AddPortForMode(aPort IPort, mode RunLoopMode) {
	objc.Call[objc.Void](r_, objc.Sel("addPort:forMode:"), objc.Ptr(aPort), mode)
}

// Schedules a block that the run loop invokes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrunloop/2091881-performblock?language=objc
func (r_ RunLoop) PerformBlock(block func()) {
	objc.Call[objc.Void](r_, objc.Sel("performBlock:"), block)
}

// Performs one pass through the run loop in the specified mode and returns the date at which the next timer is scheduled to fire. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrunloop/1412837-limitdateformode?language=objc
func (r_ RunLoop) LimitDateForMode(mode RunLoopMode) Date {
	rv := objc.Call[Date](r_, objc.Sel("limitDateForMode:"), mode)
	return rv
}

// Schedules the sending of a message on the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrunloop/1409310-performselector?language=objc
func (r_ RunLoop) PerformSelectorTargetArgumentOrderModes(aSelector objc.Selector, target objc.IObject, arg objc.IObject, order uint, modes []RunLoopMode) {
	objc.Call[objc.Void](r_, objc.Sel("performSelector:target:argument:order:modes:"), aSelector, target, arg, order, modes)
}

// Runs the loop until the specified date, during which time it processes data from all attached input sources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrunloop/1415778-rununtildate?language=objc
func (r_ RunLoop) RunUntilDate(limitDate IDate) {
	objc.Call[objc.Void](r_, objc.Sel("runUntilDate:"), objc.Ptr(limitDate))
}

// Puts the receiver into a permanent loop, during which time it processes data from all attached input sources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrunloop/1412430-run?language=objc
func (r_ RunLoop) Run() {
	objc.Call[objc.Void](r_, objc.Sel("run"))
}

// Removes a port from the specified input mode of the run loop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrunloop/1414332-removeport?language=objc
func (r_ RunLoop) RemovePortForMode(aPort IPort, mode RunLoopMode) {
	objc.Call[objc.Void](r_, objc.Sel("removePort:forMode:"), objc.Ptr(aPort), mode)
}

// Returns the run loop for the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrunloop/1412291-currentrunloop?language=objc
func (rc _RunLoopClass) CurrentRunLoop() RunLoop {
	rv := objc.Call[RunLoop](rc, objc.Sel("currentRunLoop"))
	return rv
}

// Returns the run loop for the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrunloop/1412291-currentrunloop?language=objc
func RunLoop_CurrentRunLoop() RunLoop {
	return RunLoopClass.CurrentRunLoop()
}

// Returns the run loop of the main thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrunloop/1418388-mainrunloop?language=objc
func (rc _RunLoopClass) MainRunLoop() RunLoop {
	rv := objc.Call[RunLoop](rc, objc.Sel("mainRunLoop"))
	return rv
}

// Returns the run loop of the main thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrunloop/1418388-mainrunloop?language=objc
func RunLoop_MainRunLoop() RunLoop {
	return RunLoopClass.MainRunLoop()
}

// The receiver's current input mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrunloop/1412652-currentmode?language=objc
func (r_ RunLoop) CurrentMode() RunLoopMode {
	rv := objc.Call[RunLoopMode](r_, objc.Sel("currentMode"))
	return rv
}
