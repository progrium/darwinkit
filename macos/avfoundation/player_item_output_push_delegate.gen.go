// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that defines the methods to implement to respond to changes in the media data sequence. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemoutputpushdelegate?language=objc
type PPlayerItemOutputPushDelegate interface {
	// optional
	OutputSequenceWasFlushed(output PlayerItemOutput)
	HasOutputSequenceWasFlushed() bool
}

// A delegate implementation builder for the [PPlayerItemOutputPushDelegate] protocol.
type PlayerItemOutputPushDelegate struct {
	_OutputSequenceWasFlushed func(output PlayerItemOutput)
}

func (di *PlayerItemOutputPushDelegate) HasOutputSequenceWasFlushed() bool {
	return di._OutputSequenceWasFlushed != nil
}

// Tells the delegate that the output is starting a new sequence of media data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemoutputpushdelegate/1390224-outputsequencewasflushed?language=objc
func (di *PlayerItemOutputPushDelegate) SetOutputSequenceWasFlushed(f func(output PlayerItemOutput)) {
	di._OutputSequenceWasFlushed = f
}

// Tells the delegate that the output is starting a new sequence of media data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemoutputpushdelegate/1390224-outputsequencewasflushed?language=objc
func (di *PlayerItemOutputPushDelegate) OutputSequenceWasFlushed(output PlayerItemOutput) {
	di._OutputSequenceWasFlushed(output)
}

// A concrete type wrapper for the [PPlayerItemOutputPushDelegate] protocol.
type PlayerItemOutputPushDelegateWrapper struct {
	objc.Object
}

func (p_ PlayerItemOutputPushDelegateWrapper) HasOutputSequenceWasFlushed() bool {
	return p_.RespondsToSelector(objc.Sel("outputSequenceWasFlushed:"))
}

// Tells the delegate that the output is starting a new sequence of media data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemoutputpushdelegate/1390224-outputsequencewasflushed?language=objc
func (p_ PlayerItemOutputPushDelegateWrapper) OutputSequenceWasFlushed(output IPlayerItemOutput) {
	objc.Call[objc.Void](p_, objc.Sel("outputSequenceWasFlushed:"), objc.Ptr(output))
}
