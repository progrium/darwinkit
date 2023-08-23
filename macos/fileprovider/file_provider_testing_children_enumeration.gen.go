// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"github.com/progrium/macdriver/objc"
)

// An operation that lists a directory’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingchildrenenumeration?language=objc
type PFileProviderTestingChildrenEnumeration interface {
	// optional
	ItemIdentifier() FileProviderItemIdentifier
	HasItemIdentifier() bool

	// optional
	Side() FileProviderTestingOperationSide
	HasSide() bool
}

// A concrete type wrapper for the [PFileProviderTestingChildrenEnumeration] protocol.
type FileProviderTestingChildrenEnumerationWrapper struct {
	objc.Object
}

func (f_ FileProviderTestingChildrenEnumerationWrapper) HasItemIdentifier() bool {
	return f_.RespondsToSelector(objc.Sel("itemIdentifier"))
}

// The containing identifier’s unique identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingchildrenenumeration/3727829-itemidentifier?language=objc
func (f_ FileProviderTestingChildrenEnumerationWrapper) ItemIdentifier() FileProviderItemIdentifier {
	rv := objc.Call[FileProviderItemIdentifier](f_, objc.Sel("itemIdentifier"))
	return rv
}

func (f_ FileProviderTestingChildrenEnumerationWrapper) HasSide() bool {
	return f_.RespondsToSelector(objc.Sel("side"))
}

// The item’s location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingchildrenenumeration/3727830-side?language=objc
func (f_ FileProviderTestingChildrenEnumerationWrapper) Side() FileProviderTestingOperationSide {
	rv := objc.Call[FileProviderTestingOperationSide](f_, objc.Sel("side"))
	return rv
}
