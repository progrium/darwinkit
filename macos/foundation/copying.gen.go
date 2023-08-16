// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// A protocol that objects adopt to provide functional copies of themselves. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscopying?language=objc
type PCopying interface {
	// optional
	CopyWithZone(zone unsafe.Pointer) objc.IObject
	HasCopyWithZone() bool
}

// A concrete type wrapper for the [PCopying] protocol.
type CopyingWrapper struct {
	objc.Object
}

func (c_ CopyingWrapper) HasCopyWithZone() bool {
	return c_.RespondsToSelector(objc.Sel("copyWithZone:"))
}

// Returns a new instance that’s a copy of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscopying/1410311-copywithzone?language=objc
func (c_ CopyingWrapper) CopyWithZone(zone unsafe.Pointer) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("copyWithZone:"), zone)
	rv.Autorelease()
	return rv
}
