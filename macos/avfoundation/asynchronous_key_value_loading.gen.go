// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A protocol that defines the interface to load media data asynchronously. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasynchronouskeyvalueloading?language=objc
type PAsynchronousKeyValueLoading interface {
	// optional
	StatusOfValueForKeyError(key string, outError foundation.Error) KeyValueStatus
	HasStatusOfValueForKeyError() bool

	// optional
	LoadValuesAsynchronouslyForKeysCompletionHandler(keys []string, handler func())
	HasLoadValuesAsynchronouslyForKeysCompletionHandler() bool
}

// A concrete type wrapper for the [PAsynchronousKeyValueLoading] protocol.
type AsynchronousKeyValueLoadingWrapper struct {
	objc.Object
}

func (a_ AsynchronousKeyValueLoadingWrapper) HasStatusOfValueForKeyError() bool {
	return a_.RespondsToSelector(objc.Sel("statusOfValueForKey:error:"))
}

// Returns a status that indicates whether a property value is immediately available without blocking the calling thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasynchronouskeyvalueloading/1386816-statusofvalueforkey?language=objc
func (a_ AsynchronousKeyValueLoadingWrapper) StatusOfValueForKeyError(key string, outError foundation.IError) KeyValueStatus {
	rv := objc.Call[KeyValueStatus](a_, objc.Sel("statusOfValueForKey:error:"), key, objc.Ptr(outError))
	return rv
}

func (a_ AsynchronousKeyValueLoadingWrapper) HasLoadValuesAsynchronouslyForKeysCompletionHandler() bool {
	return a_.RespondsToSelector(objc.Sel("loadValuesAsynchronouslyForKeys:completionHandler:"))
}

// Tells the asset to load the values of all of the specified keys that aren’t already loaded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasynchronouskeyvalueloading/1387321-loadvaluesasynchronouslyforkeys?language=objc
func (a_ AsynchronousKeyValueLoadingWrapper) LoadValuesAsynchronouslyForKeysCompletionHandler(keys []string, handler func()) {
	objc.Call[objc.Void](a_, objc.Sel("loadValuesAsynchronouslyForKeys:completionHandler:"), keys, handler)
}
