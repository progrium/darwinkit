// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"github.com/progrium/macdriver/objc"
)

// An operation that looks up an item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestinglookup?language=objc
type PFileProviderTestingLookup interface {
	// optional
	ItemIdentifier() FileProviderItemIdentifier
	HasItemIdentifier() bool

	// optional
	Side() FileProviderTestingOperationSide
	HasSide() bool
}

// A concrete type wrapper for the [PFileProviderTestingLookup] protocol.
type FileProviderTestingLookupWrapper struct {
	objc.Object
}

func (f_ FileProviderTestingLookupWrapper) HasItemIdentifier() bool {
	return f_.RespondsToSelector(objc.Sel("itemIdentifier"))
}

// The unique identifier for the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestinglookup/3727839-itemidentifier?language=objc
func (f_ FileProviderTestingLookupWrapper) ItemIdentifier() FileProviderItemIdentifier {
	rv := objc.Call[FileProviderItemIdentifier](f_, objc.Sel("itemIdentifier"))
	return rv
}

func (f_ FileProviderTestingLookupWrapper) HasSide() bool {
	return f_.RespondsToSelector(objc.Sel("side"))
}

// The location where the lookup occurs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestinglookup/3727840-side?language=objc
func (f_ FileProviderTestingLookupWrapper) Side() FileProviderTestingOperationSide {
	rv := objc.Call[FileProviderTestingOperationSide](f_, objc.Sel("side"))
	return rv
}
