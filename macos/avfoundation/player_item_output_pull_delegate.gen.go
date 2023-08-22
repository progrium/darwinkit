// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"github.com/progrium/macdriver/objc"
)

// Methods you can implement to respond to pixel buffer changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemoutputpulldelegate?language=objc
type PPlayerItemOutputPullDelegate interface {
	// optional
	OutputSequenceWasFlushed(output PlayerItemOutput)
	HasOutputSequenceWasFlushed() bool

	// optional
	OutputMediaDataWillChange(sender PlayerItemOutput)
	HasOutputMediaDataWillChange() bool
}

// A delegate implementation builder for the [PPlayerItemOutputPullDelegate] protocol.
type PlayerItemOutputPullDelegate struct {
	_OutputSequenceWasFlushed  func(output PlayerItemOutput)
	_OutputMediaDataWillChange func(sender PlayerItemOutput)
}

func (di *PlayerItemOutputPullDelegate) HasOutputSequenceWasFlushed() bool {
	return di._OutputSequenceWasFlushed != nil
}

// Tells the delegate that a new sample sequence is commencing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemoutputpulldelegate/1387279-outputsequencewasflushed?language=objc
func (di *PlayerItemOutputPullDelegate) SetOutputSequenceWasFlushed(f func(output PlayerItemOutput)) {
	di._OutputSequenceWasFlushed = f
}

// Tells the delegate that a new sample sequence is commencing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemoutputpulldelegate/1387279-outputsequencewasflushed?language=objc
func (di *PlayerItemOutputPullDelegate) OutputSequenceWasFlushed(output PlayerItemOutput) {
	di._OutputSequenceWasFlushed(output)
}
func (di *PlayerItemOutputPullDelegate) HasOutputMediaDataWillChange() bool {
	return di._OutputMediaDataWillChange != nil
}

// Tells the delegate that new samples are about to arrive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemoutputpulldelegate/1387498-outputmediadatawillchange?language=objc
func (di *PlayerItemOutputPullDelegate) SetOutputMediaDataWillChange(f func(sender PlayerItemOutput)) {
	di._OutputMediaDataWillChange = f
}

// Tells the delegate that new samples are about to arrive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemoutputpulldelegate/1387498-outputmediadatawillchange?language=objc
func (di *PlayerItemOutputPullDelegate) OutputMediaDataWillChange(sender PlayerItemOutput) {
	di._OutputMediaDataWillChange(sender)
}

// A concrete type wrapper for the [PPlayerItemOutputPullDelegate] protocol.
type PlayerItemOutputPullDelegateWrapper struct {
	objc.Object
}

func (p_ PlayerItemOutputPullDelegateWrapper) HasOutputSequenceWasFlushed() bool {
	return p_.RespondsToSelector(objc.Sel("outputSequenceWasFlushed:"))
}

// Tells the delegate that a new sample sequence is commencing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemoutputpulldelegate/1387279-outputsequencewasflushed?language=objc
func (p_ PlayerItemOutputPullDelegateWrapper) OutputSequenceWasFlushed(output IPlayerItemOutput) {
	objc.Call[objc.Void](p_, objc.Sel("outputSequenceWasFlushed:"), objc.Ptr(output))
}

func (p_ PlayerItemOutputPullDelegateWrapper) HasOutputMediaDataWillChange() bool {
	return p_.RespondsToSelector(objc.Sel("outputMediaDataWillChange:"))
}

// Tells the delegate that new samples are about to arrive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemoutputpulldelegate/1387498-outputmediadatawillchange?language=objc
func (p_ PlayerItemOutputPullDelegateWrapper) OutputMediaDataWillChange(sender IPlayerItemOutput) {
	objc.Call[objc.Void](p_, objc.Sel("outputMediaDataWillChange:"), objc.Ptr(sender))
}
