// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Port] class.
var PortClass = _PortClass{objc.GetClass("NSPort")}

type _PortClass struct {
	objc.Class
}

// An interface definition for the [Port] class.
type IPort interface {
	objc.IObject
	ScheduleInRunLoopForMode(runLoop IRunLoop, mode RunLoopMode)
	RemoveFromRunLoopForMode(runLoop IRunLoop, mode RunLoopMode)
	SetDelegate(anObject PPortDelegate)
	SetDelegateObject(anObjectObject objc.IObject)
	SendBeforeDateComponentsFromReserved(limitDate IDate, components IMutableArray, receivePort IPort, headerSpaceReserved uint) bool
	Delegate() PortDelegateWrapper
	Invalidate()
	IsValid() bool
	ReservedSpaceLength() uint
}

// An abstract class that represents a communication channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsport?language=objc
type Port struct {
	objc.Object
}

func PortFrom(ptr unsafe.Pointer) Port {
	return Port{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PortClass) Alloc() Port {
	rv := objc.Call[Port](pc, objc.Sel("alloc"))
	return rv
}

func Port_Alloc() Port {
	return PortClass.Alloc()
}

func (pc _PortClass) New() Port {
	rv := objc.Call[Port](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPort() Port {
	return PortClass.New()
}

func (p_ Port) Init() Port {
	rv := objc.Call[Port](p_, objc.Sel("init"))
	return rv
}

// This method should be implemented by a subclass to set up monitoring of a port when added to a given run loop in a given input mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsport/1399517-scheduleinrunloop?language=objc
func (p_ Port) ScheduleInRunLoopForMode(runLoop IRunLoop, mode RunLoopMode) {
	objc.Call[objc.Void](p_, objc.Sel("scheduleInRunLoop:forMode:"), objc.Ptr(runLoop), mode)
}

// This method should be implemented by a subclass to stop monitoring of a port when removed from a give run loop in a given input mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsport/1399525-removefromrunloop?language=objc
func (p_ Port) RemoveFromRunLoopForMode(runLoop IRunLoop, mode RunLoopMode) {
	objc.Call[objc.Void](p_, objc.Sel("removeFromRunLoop:forMode:"), objc.Ptr(runLoop), mode)
}

// Sets the receiver’s delegate to a given object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsport/1399527-setdelegate?language=objc
func (p_ Port) SetDelegate(anObject PPortDelegate) {
	po0 := objc.WrapAsProtocol("NSPortDelegate", anObject)
	objc.Call[objc.Void](p_, objc.Sel("setDelegate:"), po0)
}

// Sets the receiver’s delegate to a given object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsport/1399527-setdelegate?language=objc
func (p_ Port) SetDelegateObject(anObjectObject objc.IObject) {
	objc.Call[objc.Void](p_, objc.Sel("setDelegate:"), objc.Ptr(anObjectObject))
}

// This method is provided for subclasses that have custom types of NSPort. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsport/1399537-sendbeforedate?language=objc
func (p_ Port) SendBeforeDateComponentsFromReserved(limitDate IDate, components IMutableArray, receivePort IPort, headerSpaceReserved uint) bool {
	rv := objc.Call[bool](p_, objc.Sel("sendBeforeDate:components:from:reserved:"), objc.Ptr(limitDate), objc.Ptr(components), objc.Ptr(receivePort), headerSpaceReserved)
	return rv
}

// Returns the receiver’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsport/1399545-delegate?language=objc
func (p_ Port) Delegate() PortDelegateWrapper {
	rv := objc.Call[PortDelegateWrapper](p_, objc.Sel("delegate"))
	return rv
}

// Creates and returns a new NSPort object capable of both sending and receiving messages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsport/1399515-port?language=objc
func (pc _PortClass) Port() Port {
	rv := objc.Call[Port](pc, objc.Sel("port"))
	return rv
}

// Creates and returns a new NSPort object capable of both sending and receiving messages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsport/1399515-port?language=objc
func Port_Port() Port {
	return PortClass.Port()
}

// Marks the receiver as invalid and posts an NSPortDidBecomeInvalidNotification to the default notification center. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsport/1399521-invalidate?language=objc
func (p_ Port) Invalidate() {
	objc.Call[objc.Void](p_, objc.Sel("invalidate"))
}

// A Boolean value that indicates whether the receiver is valid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsport/1399503-valid?language=objc
func (p_ Port) IsValid() bool {
	rv := objc.Call[bool](p_, objc.Sel("isValid"))
	return rv
}

// The number of bytes of space reserved by the receiver for sending data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsport/1399529-reservedspacelength?language=objc
func (p_ Port) ReservedSpaceLength() uint {
	rv := objc.Call[uint](p_, objc.Sel("reservedSpaceLength"))
	return rv
}
