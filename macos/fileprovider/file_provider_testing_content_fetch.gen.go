// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"github.com/progrium/macdriver/objc"
)

// An operation that fetches an item’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingcontentfetch?language=objc
type PFileProviderTestingContentFetch interface {
	// optional
	ItemIdentifier() FileProviderItemIdentifier
	HasItemIdentifier() bool

	// optional
	Side() FileProviderTestingOperationSide
	HasSide() bool
}

// A concrete type wrapper for the [PFileProviderTestingContentFetch] protocol.
type FileProviderTestingContentFetchWrapper struct {
	objc.Object
}

func (f_ FileProviderTestingContentFetchWrapper) HasItemIdentifier() bool {
	return f_.RespondsToSelector(objc.Sel("itemIdentifier"))
}

// The containing item’s unique identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingcontentfetch/3727832-itemidentifier?language=objc
func (f_ FileProviderTestingContentFetchWrapper) ItemIdentifier() FileProviderItemIdentifier {
	rv := objc.Call[FileProviderItemIdentifier](f_, objc.Sel("itemIdentifier"))
	return rv
}

func (f_ FileProviderTestingContentFetchWrapper) HasSide() bool {
	return f_.RespondsToSelector(objc.Sel("side"))
}

// The item’s location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingcontentfetch/3727833-side?language=objc
func (f_ FileProviderTestingContentFetchWrapper) Side() FileProviderTestingOperationSide {
	rv := objc.Call[FileProviderTestingOperationSide](f_, objc.Sel("side"))
	return rv
}
