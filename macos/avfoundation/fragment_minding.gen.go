// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that defines whether an asset supports fragment minding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentminding?language=objc
type PFragmentMinding interface {
	// optional
	IsAssociatedWithFragmentMinder() bool
	HasIsAssociatedWithFragmentMinder() bool
}

// A concrete type wrapper for the [PFragmentMinding] protocol.
type FragmentMindingWrapper struct {
	objc.Object
}

func (f_ FragmentMindingWrapper) HasIsAssociatedWithFragmentMinder() bool {
	return f_.RespondsToSelector(objc.Sel("isAssociatedWithFragmentMinder"))
}

// A Boolean value that indicates whether an asset that supports fragment minding is currently associated with a fragment minder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentminding/1390175-associatedwithfragmentminder?language=objc
func (f_ FragmentMindingWrapper) IsAssociatedWithFragmentMinder() bool {
	rv := objc.Call[bool](f_, objc.Sel("isAssociatedWithFragmentMinder"))
	return rv
}
