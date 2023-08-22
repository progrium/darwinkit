// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"github.com/progrium/macdriver/objc"
)

// An operation that alerts the system to either local or remote storage changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingingestion?language=objc
type PFileProviderTestingIngestion interface {
	// optional
	Item() objc.IObject
	HasItem() bool

	// optional
	ItemIdentifier() FileProviderItemIdentifier
	HasItemIdentifier() bool

	// optional
	Side() FileProviderTestingOperationSide
	HasSide() bool
}

// A concrete type wrapper for the [PFileProviderTestingIngestion] protocol.
type FileProviderTestingIngestionWrapper struct {
	objc.Object
}

func (f_ FileProviderTestingIngestionWrapper) HasItem() bool {
	return f_.RespondsToSelector(objc.Sel("item"))
}

// A description of the item that changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingingestion/3727835-item?language=objc
func (f_ FileProviderTestingIngestionWrapper) Item() objc.Object {
	rv := objc.Call[objc.Object](f_, objc.Sel("item"))
	return rv
}

func (f_ FileProviderTestingIngestionWrapper) HasItemIdentifier() bool {
	return f_.RespondsToSelector(objc.Sel("itemIdentifier"))
}

// The unique identifier for the item that changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingingestion/3727836-itemidentifier?language=objc
func (f_ FileProviderTestingIngestionWrapper) ItemIdentifier() FileProviderItemIdentifier {
	rv := objc.Call[FileProviderItemIdentifier](f_, objc.Sel("itemIdentifier"))
	return rv
}

func (f_ FileProviderTestingIngestionWrapper) HasSide() bool {
	return f_.RespondsToSelector(objc.Sel("side"))
}

// The location where the change occurred. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingingestion/3727837-side?language=objc
func (f_ FileProviderTestingIngestionWrapper) Side() FileProviderTestingOperationSide {
	rv := objc.Call[FileProviderTestingOperationSide](f_, objc.Sel("side"))
	return rv
}
