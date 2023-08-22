// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// Support for enumerating the file provider’s contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderenumerating?language=objc
type PFileProviderEnumerating interface {
	// optional
	EnumeratorForContainerItemIdentifierRequestError(containerItemIdentifier FileProviderItemIdentifier, request FileProviderRequest, error foundation.Error) PFileProviderEnumerator
	HasEnumeratorForContainerItemIdentifierRequestError() bool
}

// A concrete type wrapper for the [PFileProviderEnumerating] protocol.
type FileProviderEnumeratingWrapper struct {
	objc.Object
}

func (f_ FileProviderEnumeratingWrapper) HasEnumeratorForContainerItemIdentifierRequestError() bool {
	return f_.RespondsToSelector(objc.Sel("enumeratorForContainerItemIdentifier:request:error:"))
}

// Tells the file provider to return an enumerator for the provided directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderenumerating/3553295-enumeratorforcontaineritemidenti?language=objc
func (f_ FileProviderEnumeratingWrapper) EnumeratorForContainerItemIdentifierRequestError(containerItemIdentifier FileProviderItemIdentifier, request IFileProviderRequest, error foundation.IError) FileProviderEnumeratorWrapper {
	rv := objc.Call[FileProviderEnumeratorWrapper](f_, objc.Sel("enumeratorForContainerItemIdentifier:request:error:"), containerItemIdentifier, objc.Ptr(request), objc.Ptr(error))
	return rv
}
