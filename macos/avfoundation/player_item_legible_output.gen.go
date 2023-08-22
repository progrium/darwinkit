// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/dispatch"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PlayerItemLegibleOutput] class.
var PlayerItemLegibleOutputClass = _PlayerItemLegibleOutputClass{objc.GetClass("AVPlayerItemLegibleOutput")}

type _PlayerItemLegibleOutputClass struct {
	objc.Class
}

// An interface definition for the [PlayerItemLegibleOutput] class.
type IPlayerItemLegibleOutput interface {
	IPlayerItemOutput
	SetDelegateQueue(delegate PPlayerItemLegibleOutputPushDelegate, delegateQueue dispatch.Queue)
	SetDelegateObjectQueue(delegateObject objc.IObject, delegateQueue dispatch.Queue)
	AdvanceIntervalForDelegateInvocation() foundation.TimeInterval
	SetAdvanceIntervalForDelegateInvocation(value foundation.TimeInterval)
	Delegate() PlayerItemLegibleOutputPushDelegateWrapper
	TextStylingResolution() PlayerItemLegibleOutputTextStylingResolution
	SetTextStylingResolution(value PlayerItemLegibleOutputTextStylingResolution)
	DelegateQueue() dispatch.Queue
}

// An object that vends attributed strings for media with a legible characteristic. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemlegibleoutput?language=objc
type PlayerItemLegibleOutput struct {
	PlayerItemOutput
}

func PlayerItemLegibleOutputFrom(ptr unsafe.Pointer) PlayerItemLegibleOutput {
	return PlayerItemLegibleOutput{
		PlayerItemOutput: PlayerItemOutputFrom(ptr),
	}
}

func (p_ PlayerItemLegibleOutput) InitWithMediaSubtypesForNativeRepresentation(subtypes []foundation.INumber) PlayerItemLegibleOutput {
	rv := objc.Call[PlayerItemLegibleOutput](p_, objc.Sel("initWithMediaSubtypesForNativeRepresentation:"), subtypes)
	return rv
}

// Creates an initialized legible-output object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemlegibleoutput/1390500-initwithmediasubtypesfornativere?language=objc
func NewPlayerItemLegibleOutputWithMediaSubtypesForNativeRepresentation(subtypes []foundation.INumber) PlayerItemLegibleOutput {
	instance := PlayerItemLegibleOutputClass.Alloc().InitWithMediaSubtypesForNativeRepresentation(subtypes)
	instance.Autorelease()
	return instance
}

func (pc _PlayerItemLegibleOutputClass) Alloc() PlayerItemLegibleOutput {
	rv := objc.Call[PlayerItemLegibleOutput](pc, objc.Sel("alloc"))
	return rv
}

func PlayerItemLegibleOutput_Alloc() PlayerItemLegibleOutput {
	return PlayerItemLegibleOutputClass.Alloc()
}

func (pc _PlayerItemLegibleOutputClass) New() PlayerItemLegibleOutput {
	rv := objc.Call[PlayerItemLegibleOutput](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPlayerItemLegibleOutput() PlayerItemLegibleOutput {
	return PlayerItemLegibleOutputClass.New()
}

func (p_ PlayerItemLegibleOutput) Init() PlayerItemLegibleOutput {
	rv := objc.Call[PlayerItemLegibleOutput](p_, objc.Sel("init"))
	return rv
}

// Sets the receiver's delegate and a dispatch queue on which the delegate is called. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemlegibleoutput/1386204-setdelegate?language=objc
func (p_ PlayerItemLegibleOutput) SetDelegateQueue(delegate PPlayerItemLegibleOutputPushDelegate, delegateQueue dispatch.Queue) {
	po0 := objc.WrapAsProtocol("AVPlayerItemLegibleOutputPushDelegate", delegate)
	objc.Call[objc.Void](p_, objc.Sel("setDelegate:queue:"), po0, delegateQueue)
}

// Sets the receiver's delegate and a dispatch queue on which the delegate is called. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemlegibleoutput/1386204-setdelegate?language=objc
func (p_ PlayerItemLegibleOutput) SetDelegateObjectQueue(delegateObject objc.IObject, delegateQueue dispatch.Queue) {
	objc.Call[objc.Void](p_, objc.Sel("setDelegate:queue:"), objc.Ptr(delegateObject), delegateQueue)
}

// The time interval, in seconds, that a player item legible output object messages its delegate earlier than normal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemlegibleoutput/1388098-advanceintervalfordelegateinvoca?language=objc
func (p_ PlayerItemLegibleOutput) AdvanceIntervalForDelegateInvocation() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](p_, objc.Sel("advanceIntervalForDelegateInvocation"))
	return rv
}

// The time interval, in seconds, that a player item legible output object messages its delegate earlier than normal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemlegibleoutput/1388098-advanceintervalfordelegateinvoca?language=objc
func (p_ PlayerItemLegibleOutput) SetAdvanceIntervalForDelegateInvocation(value foundation.TimeInterval) {
	objc.Call[objc.Void](p_, objc.Sel("setAdvanceIntervalForDelegateInvocation:"), value)
}

// The delegate of the output class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemlegibleoutput/1387877-delegate?language=objc
func (p_ PlayerItemLegibleOutput) Delegate() PlayerItemLegibleOutputPushDelegateWrapper {
	rv := objc.Call[PlayerItemLegibleOutputPushDelegateWrapper](p_, objc.Sel("delegate"))
	return rv
}

// A string identifier indicating the degree of text styling to be applied to attributed strings vended by the  object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemlegibleoutput/1385803-textstylingresolution?language=objc
func (p_ PlayerItemLegibleOutput) TextStylingResolution() PlayerItemLegibleOutputTextStylingResolution {
	rv := objc.Call[PlayerItemLegibleOutputTextStylingResolution](p_, objc.Sel("textStylingResolution"))
	return rv
}

// A string identifier indicating the degree of text styling to be applied to attributed strings vended by the  object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemlegibleoutput/1385803-textstylingresolution?language=objc
func (p_ PlayerItemLegibleOutput) SetTextStylingResolution(value PlayerItemLegibleOutputTextStylingResolution) {
	objc.Call[objc.Void](p_, objc.Sel("setTextStylingResolution:"), value)
}

// The dispatch queue on which the delegate is called. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemlegibleoutput/1386275-delegatequeue?language=objc
func (p_ PlayerItemLegibleOutput) DelegateQueue() dispatch.Queue {
	rv := objc.Call[dispatch.Queue](p_, objc.Sel("delegateQueue"))
	return rv
}
