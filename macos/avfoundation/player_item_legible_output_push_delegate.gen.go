// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// Methods you can implement to provide alternative attributed-string output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemlegibleoutputpushdelegate?language=objc
type PPlayerItemLegibleOutputPushDelegate interface {
	// optional
	LegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime(output PlayerItemLegibleOutput, strings []foundation.AttributedString, nativeSamples []objc.Object, itemTime coremedia.Time)
	HasLegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime() bool
}

// A delegate implementation builder for the [PPlayerItemLegibleOutputPushDelegate] protocol.
type PlayerItemLegibleOutputPushDelegate struct {
	_LegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime func(output PlayerItemLegibleOutput, strings []foundation.AttributedString, nativeSamples []objc.Object, itemTime coremedia.Time)
}

func (di *PlayerItemLegibleOutputPushDelegate) HasLegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime() bool {
	return di._LegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime != nil
}

// Asks the delegate to process the delivery of new textual samples. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemlegibleoutputpushdelegate/1386790-legibleoutput?language=objc
func (di *PlayerItemLegibleOutputPushDelegate) SetLegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime(f func(output PlayerItemLegibleOutput, strings []foundation.AttributedString, nativeSamples []objc.Object, itemTime coremedia.Time)) {
	di._LegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime = f
}

// Asks the delegate to process the delivery of new textual samples. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemlegibleoutputpushdelegate/1386790-legibleoutput?language=objc
func (di *PlayerItemLegibleOutputPushDelegate) LegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime(output PlayerItemLegibleOutput, strings []foundation.AttributedString, nativeSamples []objc.Object, itemTime coremedia.Time) {
	di._LegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime(output, strings, nativeSamples, itemTime)
}

// A concrete type wrapper for the [PPlayerItemLegibleOutputPushDelegate] protocol.
type PlayerItemLegibleOutputPushDelegateWrapper struct {
	objc.Object
}

func (p_ PlayerItemLegibleOutputPushDelegateWrapper) HasLegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime() bool {
	return p_.RespondsToSelector(objc.Sel("legibleOutput:didOutputAttributedStrings:nativeSampleBuffers:forItemTime:"))
}

// Asks the delegate to process the delivery of new textual samples. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemlegibleoutputpushdelegate/1386790-legibleoutput?language=objc
func (p_ PlayerItemLegibleOutputPushDelegateWrapper) LegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime(output IPlayerItemLegibleOutput, strings []foundation.IAttributedString, nativeSamples []objc.IObject, itemTime coremedia.Time) {
	objc.Call[objc.Void](p_, objc.Sel("legibleOutput:didOutputAttributedStrings:nativeSampleBuffers:forItemTime:"), objc.Ptr(output), strings, nativeSamples, itemTime)
}
