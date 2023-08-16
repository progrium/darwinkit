// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of methods that saves a machine learning type to the file system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlwritable?language=objc
type PWritable interface {
	// optional
	WriteToURLError(url foundation.URL, error foundation.Error) bool
	HasWriteToURLError() bool
}

// A concrete type wrapper for the [PWritable] protocol.
type WritableWrapper struct {
	objc.Object
}

func (w_ WritableWrapper) HasWriteToURLError() bool {
	return w_.RespondsToSelector(objc.Sel("writeToURL:error:"))
}

// Exports a machine learning file to the file system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlwritable/3197650-writetourl?language=objc
func (w_ WritableWrapper) WriteToURLError(url foundation.IURL, error foundation.IError) bool {
	rv := objc.Call[bool](w_, objc.Sel("writeToURL:error:"), objc.Ptr(url), objc.Ptr(error))
	return rv
}
