// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var PortClass = _PortClass{objc.GetClass("NSPort")}

type _PortClass struct {
	objc.Class
}

type IPort interface {
	objc.IObject
	Invalidate()
	SetDelegate(anObject IPortDelegate)
	SetDelegate0(anObject objc.IObject)
	Delegate() PortDelegateWrapper
	RemoveFromRunLoopForMode(runLoop IRunLoop, mode RunLoopMode)
	ScheduleInRunLoopForMode(runLoop IRunLoop, mode RunLoopMode)
	IsValid() bool
	ReservedSpaceLength() uint
}

type Port struct {
	objc.Object
}

func MakePort(ptr unsafe.Pointer) Port {
	return Port{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PortClass) Alloc() Port {
	rv := objc.CallMethod[Port](pc, objc.GetSelector("alloc"))
	return rv
}

func Port_Alloc() Port {
	return PortClass.Alloc()
}

func (pc _PortClass) New() Port {
	rv := objc.CallMethod[Port](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewPort() Port {
	return PortClass.New()
}

func Port_New() Port {
	return PortClass.New()
}

func (p_ Port) Init() Port {
	rv := objc.CallMethod[Port](p_, objc.GetSelector("init"))
	return rv
}

func Port_Init() Port {
	return PortClass.Alloc().Init()
}

func (pc _PortClass) Port() Port {
	rv := objc.CallMethod[Port](pc, objc.GetSelector("port"))
	return rv
}

func Port_Port() Port {
	return PortClass.Port()
}

func (p_ Port) Invalidate() {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("invalidate"))
}

func (p_ Port) SetDelegate(anObject IPortDelegate) {
	po := objc.WrapAsProtocol("NSPortDelegate", anObject)
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setDelegate:"), po)
}

func (p_ Port) SetDelegate0(anObject objc.IObject) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(anObject))
}

func (p_ Port) Delegate() PortDelegateWrapper {
	rv := objc.CallMethod[PortDelegateWrapper](p_, objc.GetSelector("delegate"))
	return rv
}

func (p_ Port) RemoveFromRunLoopForMode(runLoop IRunLoop, mode RunLoopMode) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("removeFromRunLoop:forMode:"), objc.ExtractPtr(runLoop), mode)
}

func (p_ Port) ScheduleInRunLoopForMode(runLoop IRunLoop, mode RunLoopMode) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("scheduleInRunLoop:forMode:"), objc.ExtractPtr(runLoop), mode)
}

func (p_ Port) IsValid() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isValid"))
	return rv
}

func (p_ Port) ReservedSpaceLength() uint {
	rv := objc.CallMethod[uint](p_, objc.GetSelector("reservedSpaceLength"))
	return rv
}
