// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"github.com/progrium/macdriver/objc"
)

// An operation that syncs the creation of the source item to the target location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingcreation?language=objc
type PFileProviderTestingCreation interface {
	// optional
	TargetSide() FileProviderTestingOperationSide
	HasTargetSide() bool

	// optional
	DomainVersion() IFileProviderDomainVersion
	HasDomainVersion() bool

	// optional
	SourceItem() objc.IObject
	HasSourceItem() bool
}

// A concrete type wrapper for the [PFileProviderTestingCreation] protocol.
type FileProviderTestingCreationWrapper struct {
	objc.Object
}

func (f_ FileProviderTestingCreationWrapper) HasTargetSide() bool {
	return f_.RespondsToSelector(objc.Sel("targetSide"))
}

// The target location for the new item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingcreation/3736233-targetside?language=objc
func (f_ FileProviderTestingCreationWrapper) TargetSide() FileProviderTestingOperationSide {
	rv := objc.Call[FileProviderTestingOperationSide](f_, objc.Sel("targetSide"))
	return rv
}

func (f_ FileProviderTestingCreationWrapper) HasDomainVersion() bool {
	return f_.RespondsToSelector(objc.Sel("domainVersion"))
}

// The domain’s version when the system discovered the item at the source location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingcreation/3736231-domainversion?language=objc
func (f_ FileProviderTestingCreationWrapper) DomainVersion() FileProviderDomainVersion {
	rv := objc.Call[FileProviderDomainVersion](f_, objc.Sel("domainVersion"))
	return rv
}

func (f_ FileProviderTestingCreationWrapper) HasSourceItem() bool {
	return f_.RespondsToSelector(objc.Sel("sourceItem"))
}

// A description of the item stored at the source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingcreation/3736232-sourceitem?language=objc
func (f_ FileProviderTestingCreationWrapper) SourceItem() objc.Object {
	rv := objc.Call[objc.Object](f_, objc.Sel("sourceItem"))
	return rv
}
