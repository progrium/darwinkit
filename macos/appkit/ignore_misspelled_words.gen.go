// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that enables the Ignore button in the Spelling panel to function properly. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsignoremisspelledwords?language=objc
type PIgnoreMisspelledWords interface {
	// optional
	IgnoreSpelling(sender objc.Object)
	HasIgnoreSpelling() bool
}

// A concrete type wrapper for the [PIgnoreMisspelledWords] protocol.
type IgnoreMisspelledWordsWrapper struct {
	objc.Object
}

func (i_ IgnoreMisspelledWordsWrapper) HasIgnoreSpelling() bool {
	return i_.RespondsToSelector(objc.Sel("ignoreSpelling:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsignoremisspelledwords/1533150-ignorespelling?language=objc
func (i_ IgnoreMisspelledWordsWrapper) IgnoreSpelling(sender objc.IObject) {
	objc.Call[objc.Void](i_, objc.Sel("ignoreSpelling:"), sender)
}
