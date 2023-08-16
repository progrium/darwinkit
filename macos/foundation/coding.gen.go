// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that enables an object to be encoded and decoded for archiving and distribution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoding?language=objc
type PCoding interface {
	// optional
	EncodeWithCoder(coder Coder)
	HasEncodeWithCoder() bool

	// optional
	InitWithCoder(coder Coder) objc.IObject
	HasInitWithCoder() bool
}

// A concrete type wrapper for the [PCoding] protocol.
type CodingWrapper struct {
	objc.Object
}

func (c_ CodingWrapper) HasEncodeWithCoder() bool {
	return c_.RespondsToSelector(objc.Sel("encodeWithCoder:"))
}

// Encodes the receiver using a given archiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoding/1413933-encodewithcoder?language=objc
func (c_ CodingWrapper) EncodeWithCoder(coder ICoder) {
	objc.Call[objc.Void](c_, objc.Sel("encodeWithCoder:"), objc.Ptr(coder))
}

func (c_ CodingWrapper) HasInitWithCoder() bool {
	return c_.RespondsToSelector(objc.Sel("initWithCoder:"))
}

// Returns an object initialized from data in a given unarchiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoding/1416145-initwithcoder?language=objc
func (c_ CodingWrapper) InitWithCoder(coder ICoder) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("initWithCoder:"), objc.Ptr(coder))
	return rv
}
