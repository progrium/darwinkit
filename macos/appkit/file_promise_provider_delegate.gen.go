// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of methods that provides the name of the promised file and writes the file to the destination directory when the file promise is fulfilled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseproviderdelegate?language=objc
type PFilePromiseProviderDelegate interface {
	// optional
	OperationQueueForFilePromiseProvider(filePromiseProvider FilePromiseProvider) foundation.IOperationQueue
	HasOperationQueueForFilePromiseProvider() bool

	// optional
	FilePromiseProviderFileNameForType(filePromiseProvider FilePromiseProvider, fileType string) string
	HasFilePromiseProviderFileNameForType() bool
}

// A delegate implementation builder for the [PFilePromiseProviderDelegate] protocol.
type FilePromiseProviderDelegate struct {
	_OperationQueueForFilePromiseProvider func(filePromiseProvider FilePromiseProvider) foundation.IOperationQueue
	_FilePromiseProviderFileNameForType   func(filePromiseProvider FilePromiseProvider, fileType string) string
}

func (di *FilePromiseProviderDelegate) HasOperationQueueForFilePromiseProvider() bool {
	return di._OperationQueueForFilePromiseProvider != nil
}

// Returns the operation queue from which to issue the write request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseproviderdelegate/2369279-operationqueueforfilepromiseprov?language=objc
func (di *FilePromiseProviderDelegate) SetOperationQueueForFilePromiseProvider(f func(filePromiseProvider FilePromiseProvider) foundation.IOperationQueue) {
	di._OperationQueueForFilePromiseProvider = f
}

// Returns the operation queue from which to issue the write request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseproviderdelegate/2369279-operationqueueforfilepromiseprov?language=objc
func (di *FilePromiseProviderDelegate) OperationQueueForFilePromiseProvider(filePromiseProvider FilePromiseProvider) foundation.IOperationQueue {
	return di._OperationQueueForFilePromiseProvider(filePromiseProvider)
}
func (di *FilePromiseProviderDelegate) HasFilePromiseProviderFileNameForType() bool {
	return di._FilePromiseProviderFileNameForType != nil
}

// Provides the drag destination file's name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseproviderdelegate/2369278-filepromiseprovider?language=objc
func (di *FilePromiseProviderDelegate) SetFilePromiseProviderFileNameForType(f func(filePromiseProvider FilePromiseProvider, fileType string) string) {
	di._FilePromiseProviderFileNameForType = f
}

// Provides the drag destination file's name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseproviderdelegate/2369278-filepromiseprovider?language=objc
func (di *FilePromiseProviderDelegate) FilePromiseProviderFileNameForType(filePromiseProvider FilePromiseProvider, fileType string) string {
	return di._FilePromiseProviderFileNameForType(filePromiseProvider, fileType)
}

// A concrete type wrapper for the [PFilePromiseProviderDelegate] protocol.
type FilePromiseProviderDelegateWrapper struct {
	objc.Object
}

func (f_ FilePromiseProviderDelegateWrapper) HasOperationQueueForFilePromiseProvider() bool {
	return f_.RespondsToSelector(objc.Sel("operationQueueForFilePromiseProvider:"))
}

// Returns the operation queue from which to issue the write request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseproviderdelegate/2369279-operationqueueforfilepromiseprov?language=objc
func (f_ FilePromiseProviderDelegateWrapper) OperationQueueForFilePromiseProvider(filePromiseProvider IFilePromiseProvider) foundation.OperationQueue {
	rv := objc.Call[foundation.OperationQueue](f_, objc.Sel("operationQueueForFilePromiseProvider:"), objc.Ptr(filePromiseProvider))
	return rv
}

func (f_ FilePromiseProviderDelegateWrapper) HasFilePromiseProviderFileNameForType() bool {
	return f_.RespondsToSelector(objc.Sel("filePromiseProvider:fileNameForType:"))
}

// Provides the drag destination file's name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromiseproviderdelegate/2369278-filepromiseprovider?language=objc
func (f_ FilePromiseProviderDelegateWrapper) FilePromiseProviderFileNameForType(filePromiseProvider IFilePromiseProvider, fileType string) string {
	rv := objc.Call[string](f_, objc.Sel("filePromiseProvider:fileNameForType:"), objc.Ptr(filePromiseProvider), fileType)
	return rv
}
