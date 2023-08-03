// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var RunLoopClass = _RunLoopClass{objc.GetClass("NSRunLoop")}

type _RunLoopClass struct {
	objc.Class
}

type IRunLoop interface {
	objc.IObject
	LimitDateForMode(mode RunLoopMode) Date
	AddTimerForMode(timer ITimer, mode RunLoopMode)
	AddPortForMode(aPort IPort, mode RunLoopMode)
	RemovePortForMode(aPort IPort, mode RunLoopMode)
	Run()
	RunModeBeforeDate(mode RunLoopMode, limitDate IDate) bool
	RunUntilDate(limitDate IDate)
	AcceptInputForModeBeforeDate(mode RunLoopMode, limitDate IDate)
	PerformBlock(block func())
	PerformInModesBlock(modes []RunLoopMode, block func())
	PerformSelectorTargetArgumentOrderModes(aSelector objc.Selector, target objc.IObject, arg objc.IObject, order uint, modes []RunLoopMode)
	CancelPerformSelectorTargetArgument(aSelector objc.Selector, target objc.IObject, arg objc.IObject)
	CancelPerformSelectorsWithTarget(target objc.IObject)
	CurrentMode() RunLoopMode
}

type RunLoop struct {
	objc.Object
}

func MakeRunLoop(ptr unsafe.Pointer) RunLoop {
	return RunLoop{
		Object: objc.MakeObject(ptr),
	}
}

func (rc _RunLoopClass) Alloc() RunLoop {
	rv := objc.CallMethod[RunLoop](rc, objc.GetSelector("alloc"))
	return rv
}

func RunLoop_Alloc() RunLoop {
	return RunLoopClass.Alloc()
}

func (rc _RunLoopClass) New() RunLoop {
	rv := objc.CallMethod[RunLoop](rc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewRunLoop() RunLoop {
	return RunLoopClass.New()
}

func RunLoop_New() RunLoop {
	return RunLoopClass.New()
}

func (r_ RunLoop) Init() RunLoop {
	rv := objc.CallMethod[RunLoop](r_, objc.GetSelector("init"))
	return rv
}

func RunLoop_Init() RunLoop {
	return RunLoopClass.Alloc().Init()
}

func (r_ RunLoop) LimitDateForMode(mode RunLoopMode) Date {
	rv := objc.CallMethod[Date](r_, objc.GetSelector("limitDateForMode:"), mode)
	return rv
}

func (r_ RunLoop) AddTimerForMode(timer ITimer, mode RunLoopMode) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("addTimer:forMode:"), objc.ExtractPtr(timer), mode)
}

func (r_ RunLoop) AddPortForMode(aPort IPort, mode RunLoopMode) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("addPort:forMode:"), objc.ExtractPtr(aPort), mode)
}

func (r_ RunLoop) RemovePortForMode(aPort IPort, mode RunLoopMode) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("removePort:forMode:"), objc.ExtractPtr(aPort), mode)
}

func (r_ RunLoop) Run() {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("run"))
}

func (r_ RunLoop) RunModeBeforeDate(mode RunLoopMode, limitDate IDate) bool {
	rv := objc.CallMethod[bool](r_, objc.GetSelector("runMode:beforeDate:"), mode, objc.ExtractPtr(limitDate))
	return rv
}

func (r_ RunLoop) RunUntilDate(limitDate IDate) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("runUntilDate:"), objc.ExtractPtr(limitDate))
}

func (r_ RunLoop) AcceptInputForModeBeforeDate(mode RunLoopMode, limitDate IDate) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("acceptInputForMode:beforeDate:"), mode, objc.ExtractPtr(limitDate))
}

func (r_ RunLoop) PerformBlock(block func()) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("performBlock:"), block)
}

func (r_ RunLoop) PerformInModesBlock(modes []RunLoopMode, block func()) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("performInModes:block:"), modes, block)
}

func (r_ RunLoop) PerformSelectorTargetArgumentOrderModes(aSelector objc.Selector, target objc.IObject, arg objc.IObject, order uint, modes []RunLoopMode) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("performSelector:target:argument:order:modes:"), aSelector, objc.ExtractPtr(target), objc.ExtractPtr(arg), order, modes)
}

func (r_ RunLoop) CancelPerformSelectorTargetArgument(aSelector objc.Selector, target objc.IObject, arg objc.IObject) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("cancelPerformSelector:target:argument:"), aSelector, objc.ExtractPtr(target), objc.ExtractPtr(arg))
}

func (r_ RunLoop) CancelPerformSelectorsWithTarget(target objc.IObject) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("cancelPerformSelectorsWithTarget:"), objc.ExtractPtr(target))
}

func (rc _RunLoopClass) CurrentRunLoop() RunLoop {
	rv := objc.CallMethod[RunLoop](rc, objc.GetSelector("currentRunLoop"))
	return rv
}

func RunLoop_CurrentRunLoop() RunLoop {
	return RunLoopClass.CurrentRunLoop()
}

func (r_ RunLoop) CurrentMode() RunLoopMode {
	rv := objc.CallMethod[RunLoopMode](r_, objc.GetSelector("currentMode"))
	return rv
}

func (rc _RunLoopClass) MainRunLoop() RunLoop {
	rv := objc.CallMethod[RunLoop](rc, objc.GetSelector("mainRunLoop"))
	return rv
}

func RunLoop_MainRunLoop() RunLoop {
	return RunLoopClass.MainRunLoop()
}
