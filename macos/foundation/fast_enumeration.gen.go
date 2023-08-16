// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that objects adopt to support fast enumeration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfastenumeration?language=objc
type PFastEnumeration interface {
	// optional
	CountByEnumeratingWithStateObjectsCount(state *FastEnumerationState, buffer objc.Object, len uint) uint
	HasCountByEnumeratingWithStateObjectsCount() bool
}

// A concrete type wrapper for the [PFastEnumeration] protocol.
type FastEnumerationWrapper struct {
	objc.Object
}

func (f_ FastEnumerationWrapper) HasCountByEnumeratingWithStateObjectsCount() bool {
	return f_.RespondsToSelector(objc.Sel("countByEnumeratingWithState:objects:count:"))
}

// Returns by reference a C array of objects over which the sender should iterate, and as the return value the number of objects in the array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfastenumeration/1412867-countbyenumeratingwithstate?language=objc
func (f_ FastEnumerationWrapper) CountByEnumeratingWithStateObjectsCount(state *FastEnumerationState, buffer objc.IObject, len uint) uint {
	rv := objc.Call[uint](f_, objc.Sel("countByEnumeratingWithState:objects:count:"), state, buffer, len)
	return rv
}
