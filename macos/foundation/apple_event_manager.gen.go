// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AppleEventManager] class.
var AppleEventManagerClass = _AppleEventManagerClass{objc.GetClass("NSAppleEventManager")}

type _AppleEventManagerClass struct {
	objc.Class
}

// An interface definition for the [AppleEventManager] class.
type IAppleEventManager interface {
	objc.IObject
	ReplyAppleEventForSuspensionID(suspensionID AppleEventManagerSuspensionID) AppleEventDescriptor
	SuspendCurrentAppleEvent() AppleEventManagerSuspensionID
	SetCurrentAppleEventAndReplyEventWithSuspensionID(suspensionID AppleEventManagerSuspensionID)
	AppleEventForSuspensionID(suspensionID AppleEventManagerSuspensionID) AppleEventDescriptor
	ResumeWithSuspensionID(suspensionID AppleEventManagerSuspensionID)
	CurrentAppleEvent() AppleEventDescriptor
	CurrentReplyAppleEvent() AppleEventDescriptor
}

// A mechanism for registering handler routines for specific types of Apple events and dispatching events to those handlers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventmanager?language=objc
type AppleEventManager struct {
	objc.Object
}

func AppleEventManagerFrom(ptr unsafe.Pointer) AppleEventManager {
	return AppleEventManager{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AppleEventManagerClass) Alloc() AppleEventManager {
	rv := objc.Call[AppleEventManager](ac, objc.Sel("alloc"))
	return rv
}

func AppleEventManager_Alloc() AppleEventManager {
	return AppleEventManagerClass.Alloc()
}

func (ac _AppleEventManagerClass) New() AppleEventManager {
	rv := objc.Call[AppleEventManager](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAppleEventManager() AppleEventManager {
	return AppleEventManagerClass.New()
}

func (a_ AppleEventManager) Init() AppleEventManager {
	rv := objc.Call[AppleEventManager](a_, objc.Sel("init"))
	return rv
}

// Given a nonzero suspensionID returned by an invocation of suspendCurrentAppleEvent, returns the corresponding reply event descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventmanager/1415286-replyappleeventforsuspensionid?language=objc
func (a_ AppleEventManager) ReplyAppleEventForSuspensionID(suspensionID AppleEventManagerSuspensionID) AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](a_, objc.Sel("replyAppleEventForSuspensionID:"), suspensionID)
	return rv
}

// Suspends the handling of the current event and returns an ID that must be used to resume the handling of the event if an Apple event is being handled on the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventmanager/1413501-suspendcurrentappleevent?language=objc
func (a_ AppleEventManager) SuspendCurrentAppleEvent() AppleEventManagerSuspensionID {
	rv := objc.Call[AppleEventManagerSuspensionID](a_, objc.Sel("suspendCurrentAppleEvent"))
	return rv
}

// Returns the single instance of NSAppleEventManager, creating it first if it doesn’t exist. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventmanager/1412900-sharedappleeventmanager?language=objc
func (ac _AppleEventManagerClass) SharedAppleEventManager() AppleEventManager {
	rv := objc.Call[AppleEventManager](ac, objc.Sel("sharedAppleEventManager"))
	return rv
}

// Returns the single instance of NSAppleEventManager, creating it first if it doesn’t exist. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventmanager/1412900-sharedappleeventmanager?language=objc
func AppleEventManager_SharedAppleEventManager() AppleEventManager {
	return AppleEventManagerClass.SharedAppleEventManager()
}

// Given a nonzero suspensionID returned by an invocation of suspendCurrentAppleEvent, sets the values that will be returned by subsequent invocations of currentAppleEvent and currentReplyAppleEvent to be the event whose handling was suspended and its corresponding reply event, respectively. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventmanager/1415997-setcurrentappleeventandreplyeven?language=objc
func (a_ AppleEventManager) SetCurrentAppleEventAndReplyEventWithSuspensionID(suspensionID AppleEventManagerSuspensionID) {
	objc.Call[objc.Void](a_, objc.Sel("setCurrentAppleEventAndReplyEventWithSuspensionID:"), suspensionID)
}

// Given a nonzero suspensionID returned by an invocation of suspendCurrentAppleEvent, returns the descriptor for the event whose handling was suspended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventmanager/1410477-appleeventforsuspensionid?language=objc
func (a_ AppleEventManager) AppleEventForSuspensionID(suspensionID AppleEventManagerSuspensionID) AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](a_, objc.Sel("appleEventForSuspensionID:"), suspensionID)
	return rv
}

// Given a nonzero suspensionID returned by an invocation of suspendCurrentAppleEvent, signal that handling of the suspended event may now continue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventmanager/1412315-resumewithsuspensionid?language=objc
func (a_ AppleEventManager) ResumeWithSuspensionID(suspensionID AppleEventManagerSuspensionID) {
	objc.Call[objc.Void](a_, objc.Sel("resumeWithSuspensionID:"), suspensionID)
}

// Returns the descriptor for currentAppleEvent if an Apple event is being handled on the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventmanager/1414690-currentappleevent?language=objc
func (a_ AppleEventManager) CurrentAppleEvent() AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](a_, objc.Sel("currentAppleEvent"))
	return rv
}

// Returns the corresponding reply event descriptor if an Apple event is being handled on the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventmanager/1413207-currentreplyappleevent?language=objc
func (a_ AppleEventManager) CurrentReplyAppleEvent() AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](a_, objc.Sel("currentReplyAppleEvent"))
	return rv
}
