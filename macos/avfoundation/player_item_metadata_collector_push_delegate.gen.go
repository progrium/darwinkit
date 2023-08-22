// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A protocol you implement to receive metadata callbacks from a player item metadata collector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemmetadatacollectorpushdelegate?language=objc
type PPlayerItemMetadataCollectorPushDelegate interface {
	// optional
	MetadataCollectorDidCollectDateRangeMetadataGroupsIndexesOfNewGroupsIndexesOfModifiedGroups(metadataCollector PlayerItemMetadataCollector, metadataGroups []DateRangeMetadataGroup, indexesOfNewGroups foundation.IndexSet, indexesOfModifiedGroups foundation.IndexSet)
	HasMetadataCollectorDidCollectDateRangeMetadataGroupsIndexesOfNewGroupsIndexesOfModifiedGroups() bool
}

// A delegate implementation builder for the [PPlayerItemMetadataCollectorPushDelegate] protocol.
type PlayerItemMetadataCollectorPushDelegate struct {
	_MetadataCollectorDidCollectDateRangeMetadataGroupsIndexesOfNewGroupsIndexesOfModifiedGroups func(metadataCollector PlayerItemMetadataCollector, metadataGroups []DateRangeMetadataGroup, indexesOfNewGroups foundation.IndexSet, indexesOfModifiedGroups foundation.IndexSet)
}

func (di *PlayerItemMetadataCollectorPushDelegate) HasMetadataCollectorDidCollectDateRangeMetadataGroupsIndexesOfNewGroupsIndexesOfModifiedGroups() bool {
	return di._MetadataCollectorDidCollectDateRangeMetadataGroupsIndexesOfNewGroupsIndexesOfModifiedGroups != nil
}

// Tells the delegate the collected metadata group information has changed and needs to be updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemmetadatacollectorpushdelegate/1617190-metadatacollector?language=objc
func (di *PlayerItemMetadataCollectorPushDelegate) SetMetadataCollectorDidCollectDateRangeMetadataGroupsIndexesOfNewGroupsIndexesOfModifiedGroups(f func(metadataCollector PlayerItemMetadataCollector, metadataGroups []DateRangeMetadataGroup, indexesOfNewGroups foundation.IndexSet, indexesOfModifiedGroups foundation.IndexSet)) {
	di._MetadataCollectorDidCollectDateRangeMetadataGroupsIndexesOfNewGroupsIndexesOfModifiedGroups = f
}

// Tells the delegate the collected metadata group information has changed and needs to be updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemmetadatacollectorpushdelegate/1617190-metadatacollector?language=objc
func (di *PlayerItemMetadataCollectorPushDelegate) MetadataCollectorDidCollectDateRangeMetadataGroupsIndexesOfNewGroupsIndexesOfModifiedGroups(metadataCollector PlayerItemMetadataCollector, metadataGroups []DateRangeMetadataGroup, indexesOfNewGroups foundation.IndexSet, indexesOfModifiedGroups foundation.IndexSet) {
	di._MetadataCollectorDidCollectDateRangeMetadataGroupsIndexesOfNewGroupsIndexesOfModifiedGroups(metadataCollector, metadataGroups, indexesOfNewGroups, indexesOfModifiedGroups)
}

// A concrete type wrapper for the [PPlayerItemMetadataCollectorPushDelegate] protocol.
type PlayerItemMetadataCollectorPushDelegateWrapper struct {
	objc.Object
}

func (p_ PlayerItemMetadataCollectorPushDelegateWrapper) HasMetadataCollectorDidCollectDateRangeMetadataGroupsIndexesOfNewGroupsIndexesOfModifiedGroups() bool {
	return p_.RespondsToSelector(objc.Sel("metadataCollector:didCollectDateRangeMetadataGroups:indexesOfNewGroups:indexesOfModifiedGroups:"))
}

// Tells the delegate the collected metadata group information has changed and needs to be updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemmetadatacollectorpushdelegate/1617190-metadatacollector?language=objc
func (p_ PlayerItemMetadataCollectorPushDelegateWrapper) MetadataCollectorDidCollectDateRangeMetadataGroupsIndexesOfNewGroupsIndexesOfModifiedGroups(metadataCollector IPlayerItemMetadataCollector, metadataGroups []IDateRangeMetadataGroup, indexesOfNewGroups foundation.IIndexSet, indexesOfModifiedGroups foundation.IIndexSet) {
	objc.Call[objc.Void](p_, objc.Sel("metadataCollector:didCollectDateRangeMetadataGroups:indexesOfNewGroups:indexesOfModifiedGroups:"), objc.Ptr(metadataCollector), metadataGroups, objc.Ptr(indexesOfNewGroups), objc.Ptr(indexesOfModifiedGroups))
}
