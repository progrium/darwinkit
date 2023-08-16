// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// You implement this protocol when a class’s objects have subcomponents that can be discarded when not being used, thereby giving an application a smaller memory footprint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdiscardablecontent?language=objc
type PDiscardableContent interface {
	// optional
	IsContentDiscarded() bool
	HasIsContentDiscarded() bool

	// optional
	DiscardContentIfPossible()
	HasDiscardContentIfPossible() bool

	// optional
	EndContentAccess()
	HasEndContentAccess() bool

	// optional
	BeginContentAccess() bool
	HasBeginContentAccess() bool
}

// A concrete type wrapper for the [PDiscardableContent] protocol.
type DiscardableContentWrapper struct {
	objc.Object
}

func (d_ DiscardableContentWrapper) HasIsContentDiscarded() bool {
	return d_.RespondsToSelector(objc.Sel("isContentDiscarded"))
}

// Returns a Boolean value indicating whether the content has been discarded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdiscardablecontent/1417470-iscontentdiscarded?language=objc
func (d_ DiscardableContentWrapper) IsContentDiscarded() bool {
	rv := objc.Call[bool](d_, objc.Sel("isContentDiscarded"))
	return rv
}

func (d_ DiscardableContentWrapper) HasDiscardContentIfPossible() bool {
	return d_.RespondsToSelector(objc.Sel("discardContentIfPossible"))
}

// Called to discard the contents of the receiver if the value of the accessed counter is 0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdiscardablecontent/1408998-discardcontentifpossible?language=objc
func (d_ DiscardableContentWrapper) DiscardContentIfPossible() {
	objc.Call[objc.Void](d_, objc.Sel("discardContentIfPossible"))
}

func (d_ DiscardableContentWrapper) HasEndContentAccess() bool {
	return d_.RespondsToSelector(objc.Sel("endContentAccess"))
}

// Called if the discardable contents are no longer being accessed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdiscardablecontent/1407535-endcontentaccess?language=objc
func (d_ DiscardableContentWrapper) EndContentAccess() {
	objc.Call[objc.Void](d_, objc.Sel("endContentAccess"))
}

func (d_ DiscardableContentWrapper) HasBeginContentAccess() bool {
	return d_.RespondsToSelector(objc.Sel("beginContentAccess"))
}

// Returns a Boolean value indicating whether the discardable contents are still available and have been successfully accessed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdiscardablecontent/1412187-begincontentaccess?language=objc
func (d_ DiscardableContentWrapper) BeginContentAccess() bool {
	rv := objc.Call[bool](d_, objc.Sel("beginContentAccess"))
	return rv
}
