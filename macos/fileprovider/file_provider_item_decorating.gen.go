// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"github.com/progrium/macdriver/objc"
)

// Support for decorating items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovideritemdecorating?language=objc
type PFileProviderItemDecorating interface {
	// optional
	Decorations() []FileProviderItemDecorationIdentifier
	HasDecorations() bool
}

// A concrete type wrapper for the [PFileProviderItemDecorating] protocol.
type FileProviderItemDecoratingWrapper struct {
	objc.Object
}

func (f_ FileProviderItemDecoratingWrapper) HasDecorations() bool {
	return f_.RespondsToSelector(objc.Sel("decorations"))
}

// Asks the item for an array of decorations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovideritemdecorating/3172722-decorations?language=objc
func (f_ FileProviderItemDecoratingWrapper) Decorations() []FileProviderItemDecorationIdentifier {
	rv := objc.Call[[]FileProviderItemDecorationIdentifier](f_, objc.Sel("decorations"))
	return rv
}
