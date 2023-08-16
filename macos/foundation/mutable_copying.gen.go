// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// A protocol that mutable objects adopt to provide functional copies of themselves. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablecopying?language=objc
type PMutableCopying interface {
	// optional
	MutableCopyWithZone(zone unsafe.Pointer) objc.IObject
	HasMutableCopyWithZone() bool
}

// A concrete type wrapper for the [PMutableCopying] protocol.
type MutableCopyingWrapper struct {
	objc.Object
}

func (m_ MutableCopyingWrapper) HasMutableCopyWithZone() bool {
	return m_.RespondsToSelector(objc.Sel("mutableCopyWithZone:"))
}

// Returns a new instance that’s a mutable copy of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutablecopying/1414175-mutablecopywithzone?language=objc
func (m_ MutableCopyingWrapper) MutableCopyWithZone(zone unsafe.Pointer) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("mutableCopyWithZone:"), zone)
	rv.Autorelease()
	return rv
}
