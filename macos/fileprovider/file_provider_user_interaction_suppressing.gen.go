// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"github.com/progrium/macdriver/objc"
)

// Support for suppressing user-interaction alerts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovideruserinteractionsuppressing?language=objc
type PFileProviderUserInteractionSuppressing interface {
	// optional
	IsInteractionSuppressedForIdentifier(suppressionIdentifier string) bool
	HasIsInteractionSuppressedForIdentifier() bool

	// optional
	SetInteractionSuppressedForIdentifier(suppression bool, suppressionIdentifier string)
	HasSetInteractionSuppressedForIdentifier() bool
}

// A concrete type wrapper for the [PFileProviderUserInteractionSuppressing] protocol.
type FileProviderUserInteractionSuppressingWrapper struct {
	objc.Object
}

func (f_ FileProviderUserInteractionSuppressingWrapper) HasIsInteractionSuppressedForIdentifier() bool {
	return f_.RespondsToSelector(objc.Sel("isInteractionSuppressedForIdentifier:"))
}

// Asks the File Provider extension if the user suppressed the specified interaction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovideruserinteractionsuppressing/3762902-isinteractionsuppressedforidenti?language=objc
func (f_ FileProviderUserInteractionSuppressingWrapper) IsInteractionSuppressedForIdentifier(suppressionIdentifier string) bool {
	rv := objc.Call[bool](f_, objc.Sel("isInteractionSuppressedForIdentifier:"), suppressionIdentifier)
	return rv
}

func (f_ FileProviderUserInteractionSuppressingWrapper) HasSetInteractionSuppressedForIdentifier() bool {
	return f_.RespondsToSelector(objc.Sel("setInteractionSuppressed:forIdentifier:"))
}

// Tells the File Provider extension that the user wants to suppress the user interaction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovideruserinteractionsuppressing/3762903-setinteractionsuppressed?language=objc
func (f_ FileProviderUserInteractionSuppressingWrapper) SetInteractionSuppressedForIdentifier(suppression bool, suppressionIdentifier string) {
	objc.Call[objc.Void](f_, objc.Sel("setInteractionSuppressed:forIdentifier:"), suppression, suppressionIdentifier)
}
