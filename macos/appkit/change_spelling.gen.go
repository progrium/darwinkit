// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that responder objects can implement to correct a misspelled word. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nschangespelling?language=objc
type PChangeSpelling interface {
	// optional
	ChangeSpelling(sender objc.Object)
	HasChangeSpelling() bool
}

// A concrete type wrapper for the [PChangeSpelling] protocol.
type ChangeSpellingWrapper struct {
	objc.Object
}

func (c_ ChangeSpellingWrapper) HasChangeSpelling() bool {
	return c_.RespondsToSelector(objc.Sel("changeSpelling:"))
}

// Replaces the selected word in the receiver with a corrected version from the Spelling panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nschangespelling/1526512-changespelling?language=objc
func (c_ ChangeSpellingWrapper) ChangeSpelling(sender objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("changeSpelling:"), sender)
}
