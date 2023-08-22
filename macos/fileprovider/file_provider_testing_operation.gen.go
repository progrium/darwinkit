// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"github.com/progrium/macdriver/objc"
)

// An operation that the system can schedule. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingoperation?language=objc
type PFileProviderTestingOperation interface {
	// optional
	AsModification() PFileProviderTestingModification
	HasAsModification() bool

	// optional
	AsCreation() PFileProviderTestingCreation
	HasAsCreation() bool

	// optional
	AsDeletion() PFileProviderTestingDeletion
	HasAsDeletion() bool

	// optional
	AsLookup() PFileProviderTestingLookup
	HasAsLookup() bool

	// optional
	AsContentFetch() PFileProviderTestingContentFetch
	HasAsContentFetch() bool

	// optional
	AsIngestion() PFileProviderTestingIngestion
	HasAsIngestion() bool

	// optional
	AsCollisionResolution() PFileProviderTestingCollisionResolution
	HasAsCollisionResolution() bool

	// optional
	AsChildrenEnumeration() PFileProviderTestingChildrenEnumeration
	HasAsChildrenEnumeration() bool

	// optional
	Type() FileProviderTestingOperationType
	HasType() bool
}

// A concrete type wrapper for the [PFileProviderTestingOperation] protocol.
type FileProviderTestingOperationWrapper struct {
	objc.Object
}

func (f_ FileProviderTestingOperationWrapper) HasAsModification() bool {
	return f_.RespondsToSelector(objc.Sel("asModification"))
}

// Returns the operation if it propagates a change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingoperation/3736254-asmodification?language=objc
func (f_ FileProviderTestingOperationWrapper) AsModification() FileProviderTestingModificationWrapper {
	rv := objc.Call[FileProviderTestingModificationWrapper](f_, objc.Sel("asModification"))
	return rv
}

func (f_ FileProviderTestingOperationWrapper) HasAsCreation() bool {
	return f_.RespondsToSelector(objc.Sel("asCreation"))
}

// Returns the operation if it propagates the creation of an item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingoperation/3736250-ascreation?language=objc
func (f_ FileProviderTestingOperationWrapper) AsCreation() FileProviderTestingCreationWrapper {
	rv := objc.Call[FileProviderTestingCreationWrapper](f_, objc.Sel("asCreation"))
	return rv
}

func (f_ FileProviderTestingOperationWrapper) HasAsDeletion() bool {
	return f_.RespondsToSelector(objc.Sel("asDeletion"))
}

// Returns the operation if it propagates the deletion of an item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingoperation/3736251-asdeletion?language=objc
func (f_ FileProviderTestingOperationWrapper) AsDeletion() FileProviderTestingDeletionWrapper {
	rv := objc.Call[FileProviderTestingDeletionWrapper](f_, objc.Sel("asDeletion"))
	return rv
}

func (f_ FileProviderTestingOperationWrapper) HasAsLookup() bool {
	return f_.RespondsToSelector(objc.Sel("asLookup"))
}

// Returns the operation if it looks up an item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingoperation/3736253-aslookup?language=objc
func (f_ FileProviderTestingOperationWrapper) AsLookup() FileProviderTestingLookupWrapper {
	rv := objc.Call[FileProviderTestingLookupWrapper](f_, objc.Sel("asLookup"))
	return rv
}

func (f_ FileProviderTestingOperationWrapper) HasAsContentFetch() bool {
	return f_.RespondsToSelector(objc.Sel("asContentFetch"))
}

// Returns the operation if it fetches an item’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingoperation/3736249-ascontentfetch?language=objc
func (f_ FileProviderTestingOperationWrapper) AsContentFetch() FileProviderTestingContentFetchWrapper {
	rv := objc.Call[FileProviderTestingContentFetchWrapper](f_, objc.Sel("asContentFetch"))
	return rv
}

func (f_ FileProviderTestingOperationWrapper) HasAsIngestion() bool {
	return f_.RespondsToSelector(objc.Sel("asIngestion"))
}

// Returns the operation if it alerts the system to changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingoperation/3736252-asingestion?language=objc
func (f_ FileProviderTestingOperationWrapper) AsIngestion() FileProviderTestingIngestionWrapper {
	rv := objc.Call[FileProviderTestingIngestionWrapper](f_, objc.Sel("asIngestion"))
	return rv
}

func (f_ FileProviderTestingOperationWrapper) HasAsCollisionResolution() bool {
	return f_.RespondsToSelector(objc.Sel("asCollisionResolution"))
}

// Returns the operation if it resolves a collision by renaming the new item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingoperation/3736248-ascollisionresolution?language=objc
func (f_ FileProviderTestingOperationWrapper) AsCollisionResolution() FileProviderTestingCollisionResolutionWrapper {
	rv := objc.Call[FileProviderTestingCollisionResolutionWrapper](f_, objc.Sel("asCollisionResolution"))
	return rv
}

func (f_ FileProviderTestingOperationWrapper) HasAsChildrenEnumeration() bool {
	return f_.RespondsToSelector(objc.Sel("asChildrenEnumeration"))
}

// Returns the operation if it enumerates contained items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingoperation/3736247-aschildrenenumeration?language=objc
func (f_ FileProviderTestingOperationWrapper) AsChildrenEnumeration() FileProviderTestingChildrenEnumerationWrapper {
	rv := objc.Call[FileProviderTestingChildrenEnumerationWrapper](f_, objc.Sel("asChildrenEnumeration"))
	return rv
}

func (f_ FileProviderTestingOperationWrapper) HasType() bool {
	return f_.RespondsToSelector(objc.Sel("type"))
}

// The operation’s type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingoperation/3736255-type?language=objc
func (f_ FileProviderTestingOperationWrapper) Type() FileProviderTestingOperationType {
	rv := objc.Call[FileProviderTestingOperationType](f_, objc.Sel("type"))
	return rv
}
