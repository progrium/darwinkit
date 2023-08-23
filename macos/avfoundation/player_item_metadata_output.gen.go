// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/dispatch"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PlayerItemMetadataOutput] class.
var PlayerItemMetadataOutputClass = _PlayerItemMetadataOutputClass{objc.GetClass("AVPlayerItemMetadataOutput")}

type _PlayerItemMetadataOutputClass struct {
	objc.Class
}

// An interface definition for the [PlayerItemMetadataOutput] class.
type IPlayerItemMetadataOutput interface {
	IPlayerItemOutput
	SetDelegateQueue(delegate PPlayerItemMetadataOutputPushDelegate, delegateQueue dispatch.Queue)
	SetDelegateObjectQueue(delegateObject objc.IObject, delegateQueue dispatch.Queue)
	AdvanceIntervalForDelegateInvocation() foundation.TimeInterval
	SetAdvanceIntervalForDelegateInvocation(value foundation.TimeInterval)
	Delegate() PlayerItemMetadataOutputPushDelegateWrapper
	DelegateQueue() dispatch.Queue
}

// An object that vends collections of metadata items that a player item’s tracks carry. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemmetadataoutput?language=objc
type PlayerItemMetadataOutput struct {
	PlayerItemOutput
}

func PlayerItemMetadataOutputFrom(ptr unsafe.Pointer) PlayerItemMetadataOutput {
	return PlayerItemMetadataOutput{
		PlayerItemOutput: PlayerItemOutputFrom(ptr),
	}
}

func (p_ PlayerItemMetadataOutput) InitWithIdentifiers(identifiers []string) PlayerItemMetadataOutput {
	rv := objc.Call[PlayerItemMetadataOutput](p_, objc.Sel("initWithIdentifiers:"), identifiers)
	return rv
}

// Creates an instance of AVPlayerItemMetadataOutput. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemmetadataoutput/1390205-initwithidentifiers?language=objc
func NewPlayerItemMetadataOutputWithIdentifiers(identifiers []string) PlayerItemMetadataOutput {
	instance := PlayerItemMetadataOutputClass.Alloc().InitWithIdentifiers(identifiers)
	instance.Autorelease()
	return instance
}

func (pc _PlayerItemMetadataOutputClass) Alloc() PlayerItemMetadataOutput {
	rv := objc.Call[PlayerItemMetadataOutput](pc, objc.Sel("alloc"))
	return rv
}

func PlayerItemMetadataOutput_Alloc() PlayerItemMetadataOutput {
	return PlayerItemMetadataOutputClass.Alloc()
}

func (pc _PlayerItemMetadataOutputClass) New() PlayerItemMetadataOutput {
	rv := objc.Call[PlayerItemMetadataOutput](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPlayerItemMetadataOutput() PlayerItemMetadataOutput {
	return PlayerItemMetadataOutputClass.New()
}

func (p_ PlayerItemMetadataOutput) Init() PlayerItemMetadataOutput {
	rv := objc.Call[PlayerItemMetadataOutput](p_, objc.Sel("init"))
	return rv
}

// Sets the delegate and a dispatch queue on which the delegate is called. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemmetadataoutput/1385728-setdelegate?language=objc
func (p_ PlayerItemMetadataOutput) SetDelegateQueue(delegate PPlayerItemMetadataOutputPushDelegate, delegateQueue dispatch.Queue) {
	po0 := objc.WrapAsProtocol("AVPlayerItemMetadataOutputPushDelegate", delegate)
	objc.Call[objc.Void](p_, objc.Sel("setDelegate:queue:"), po0, delegateQueue)
}

// Sets the delegate and a dispatch queue on which the delegate is called. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemmetadataoutput/1385728-setdelegate?language=objc
func (p_ PlayerItemMetadataOutput) SetDelegateObjectQueue(delegateObject objc.IObject, delegateQueue dispatch.Queue) {
	objc.Call[objc.Void](p_, objc.Sel("setDelegate:queue:"), objc.Ptr(delegateObject), delegateQueue)
}

// The time interval, in seconds, the player item metadata output object messages its delegate earlier than normal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemmetadataoutput/1387372-advanceintervalfordelegateinvoca?language=objc
func (p_ PlayerItemMetadataOutput) AdvanceIntervalForDelegateInvocation() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](p_, objc.Sel("advanceIntervalForDelegateInvocation"))
	return rv
}

// The time interval, in seconds, the player item metadata output object messages its delegate earlier than normal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemmetadataoutput/1387372-advanceintervalfordelegateinvoca?language=objc
func (p_ PlayerItemMetadataOutput) SetAdvanceIntervalForDelegateInvocation(value foundation.TimeInterval) {
	objc.Call[objc.Void](p_, objc.Sel("setAdvanceIntervalForDelegateInvocation:"), value)
}

// The delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemmetadataoutput/1387200-delegate?language=objc
func (p_ PlayerItemMetadataOutput) Delegate() PlayerItemMetadataOutputPushDelegateWrapper {
	rv := objc.Call[PlayerItemMetadataOutputPushDelegateWrapper](p_, objc.Sel("delegate"))
	return rv
}

// The dispatch queue on which messages are sent to the delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemmetadataoutput/1387265-delegatequeue?language=objc
func (p_ PlayerItemMetadataOutput) DelegateQueue() dispatch.Queue {
	rv := objc.Call[dispatch.Queue](p_, objc.Sel("delegateQueue"))
	return rv
}
