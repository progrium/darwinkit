// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"github.com/progrium/macdriver/objc"
)

// Methods you can implement to provide additional metadata. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemmetadataoutputpushdelegate?language=objc
type PPlayerItemMetadataOutputPushDelegate interface {
	// optional
	MetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack(output PlayerItemMetadataOutput, groups []TimedMetadataGroup, track PlayerItemTrack)
	HasMetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack() bool
}

// A delegate implementation builder for the [PPlayerItemMetadataOutputPushDelegate] protocol.
type PlayerItemMetadataOutputPushDelegate struct {
	_MetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack func(output PlayerItemMetadataOutput, groups []TimedMetadataGroup, track PlayerItemTrack)
}

func (di *PlayerItemMetadataOutputPushDelegate) HasMetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack() bool {
	return di._MetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack != nil
}

// Tells the delegate a new collection of metadata items is available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemmetadataoutputpushdelegate/1388071-metadataoutput?language=objc
func (di *PlayerItemMetadataOutputPushDelegate) SetMetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack(f func(output PlayerItemMetadataOutput, groups []TimedMetadataGroup, track PlayerItemTrack)) {
	di._MetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack = f
}

// Tells the delegate a new collection of metadata items is available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemmetadataoutputpushdelegate/1388071-metadataoutput?language=objc
func (di *PlayerItemMetadataOutputPushDelegate) MetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack(output PlayerItemMetadataOutput, groups []TimedMetadataGroup, track PlayerItemTrack) {
	di._MetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack(output, groups, track)
}

// A concrete type wrapper for the [PPlayerItemMetadataOutputPushDelegate] protocol.
type PlayerItemMetadataOutputPushDelegateWrapper struct {
	objc.Object
}

func (p_ PlayerItemMetadataOutputPushDelegateWrapper) HasMetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack() bool {
	return p_.RespondsToSelector(objc.Sel("metadataOutput:didOutputTimedMetadataGroups:fromPlayerItemTrack:"))
}

// Tells the delegate a new collection of metadata items is available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemmetadataoutputpushdelegate/1388071-metadataoutput?language=objc
func (p_ PlayerItemMetadataOutputPushDelegateWrapper) MetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack(output IPlayerItemMetadataOutput, groups []ITimedMetadataGroup, track IPlayerItemTrack) {
	objc.Call[objc.Void](p_, objc.Sel("metadataOutput:didOutputTimedMetadataGroups:fromPlayerItemTrack:"), objc.Ptr(output), groups, objc.Ptr(track))
}
