// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"github.com/progrium/macdriver/objc"
)

// An operation that syncs the deletion of the source item to the target location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingdeletion?language=objc
type PFileProviderTestingDeletion interface {
	// optional
	TargetSide() FileProviderTestingOperationSide
	HasTargetSide() bool

	// optional
	TargetItemIdentifier() FileProviderItemIdentifier
	HasTargetItemIdentifier() bool

	// optional
	DomainVersion() IFileProviderDomainVersion
	HasDomainVersion() bool

	// optional
	TargetItemBaseVersion() IFileProviderItemVersion
	HasTargetItemBaseVersion() bool

	// optional
	SourceItemIdentifier() FileProviderItemIdentifier
	HasSourceItemIdentifier() bool
}

// A concrete type wrapper for the [PFileProviderTestingDeletion] protocol.
type FileProviderTestingDeletionWrapper struct {
	objc.Object
}

func (f_ FileProviderTestingDeletionWrapper) HasTargetSide() bool {
	return f_.RespondsToSelector(objc.Sel("targetSide"))
}

// The target location for the delete operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingdeletion/3736239-targetside?language=objc
func (f_ FileProviderTestingDeletionWrapper) TargetSide() FileProviderTestingOperationSide {
	rv := objc.Call[FileProviderTestingOperationSide](f_, objc.Sel("targetSide"))
	return rv
}

func (f_ FileProviderTestingDeletionWrapper) HasTargetItemIdentifier() bool {
	return f_.RespondsToSelector(objc.Sel("targetItemIdentifier"))
}

// The unique identifier for the target item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingdeletion/3736238-targetitemidentifier?language=objc
func (f_ FileProviderTestingDeletionWrapper) TargetItemIdentifier() FileProviderItemIdentifier {
	rv := objc.Call[FileProviderItemIdentifier](f_, objc.Sel("targetItemIdentifier"))
	return rv
}

func (f_ FileProviderTestingDeletionWrapper) HasDomainVersion() bool {
	return f_.RespondsToSelector(objc.Sel("domainVersion"))
}

// The domain’s version when the source location deleted the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingdeletion/3736235-domainversion?language=objc
func (f_ FileProviderTestingDeletionWrapper) DomainVersion() FileProviderDomainVersion {
	rv := objc.Call[FileProviderDomainVersion](f_, objc.Sel("domainVersion"))
	return rv
}

func (f_ FileProviderTestingDeletionWrapper) HasTargetItemBaseVersion() bool {
	return f_.RespondsToSelector(objc.Sel("targetItemBaseVersion"))
}

// The version of the deleted item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingdeletion/3736237-targetitembaseversion?language=objc
func (f_ FileProviderTestingDeletionWrapper) TargetItemBaseVersion() FileProviderItemVersion {
	rv := objc.Call[FileProviderItemVersion](f_, objc.Sel("targetItemBaseVersion"))
	return rv
}

func (f_ FileProviderTestingDeletionWrapper) HasSourceItemIdentifier() bool {
	return f_.RespondsToSelector(objc.Sel("sourceItemIdentifier"))
}

// The unique identifier for the source item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingdeletion/3736236-sourceitemidentifier?language=objc
func (f_ FileProviderTestingDeletionWrapper) SourceItemIdentifier() FileProviderItemIdentifier {
	rv := objc.Call[FileProviderItemIdentifier](f_, objc.Sel("sourceItemIdentifier"))
	return rv
}
