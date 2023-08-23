// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// An observer that receives batches of items during enumeration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderenumerationobserver?language=objc
type PFileProviderEnumerationObserver interface {
	// optional
	FinishEnumeratingWithError(error foundation.Error)
	HasFinishEnumeratingWithError() bool

	// optional
	DidEnumerateItems(updatedItems []objc.Object)
	HasDidEnumerateItems() bool

	// optional
	FinishEnumeratingUpToPage(nextPage FileProviderPage)
	HasFinishEnumeratingUpToPage() bool

	// optional
	SuggestedPageSize() int
	HasSuggestedPageSize() bool
}

// A concrete type wrapper for the [PFileProviderEnumerationObserver] protocol.
type FileProviderEnumerationObserverWrapper struct {
	objc.Object
}

func (f_ FileProviderEnumerationObserverWrapper) HasFinishEnumeratingWithError() bool {
	return f_.RespondsToSelector(objc.Sel("finishEnumeratingWithError:"))
}

// Tells the observer that an error occurred during item enumeration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderenumerationobserver/2879612-finishenumeratingwitherror?language=objc
func (f_ FileProviderEnumerationObserverWrapper) FinishEnumeratingWithError(error foundation.IError) {
	objc.Call[objc.Void](f_, objc.Sel("finishEnumeratingWithError:"), objc.Ptr(error))
}

func (f_ FileProviderEnumerationObserverWrapper) HasDidEnumerateItems() bool {
	return f_.RespondsToSelector(objc.Sel("didEnumerateItems:"))
}

// Provides a batch of enumerated items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderenumerationobserver/2879615-didenumerateitems?language=objc
func (f_ FileProviderEnumerationObserverWrapper) DidEnumerateItems(updatedItems []objc.IObject) {
	objc.Call[objc.Void](f_, objc.Sel("didEnumerateItems:"), updatedItems)
}

func (f_ FileProviderEnumerationObserverWrapper) HasFinishEnumeratingUpToPage() bool {
	return f_.RespondsToSelector(objc.Sel("finishEnumeratingUpToPage:"))
}

// Tells the observer that all of the items have been enumerated up to the specified page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderenumerationobserver/2879602-finishenumeratinguptopage?language=objc
func (f_ FileProviderEnumerationObserverWrapper) FinishEnumeratingUpToPage(nextPage FileProviderPage) {
	objc.Call[objc.Void](f_, objc.Sel("finishEnumeratingUpToPage:"), nextPage)
}

func (f_ FileProviderEnumerationObserverWrapper) HasSuggestedPageSize() bool {
	return f_.RespondsToSelector(objc.Sel("suggestedPageSize"))
}

// The page size that the system recommends. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderenumerationobserver/3656526-suggestedpagesize?language=objc
func (f_ FileProviderEnumerationObserverWrapper) SuggestedPageSize() int {
	rv := objc.Call[int](f_, objc.Sel("suggestedPageSize"))
	return rv
}
