// Code generated by DarwinKit. DO NOT EDIT.

package fileprovider

import (
	"github.com/progrium/darwinkit/objc"
)

// An operation that resolves a collision by renaming the new item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingcollisionresolution?language=objc
type PFileProviderTestingCollisionResolution interface {
	// optional
	Side() FileProviderTestingOperationSide
	HasSide() bool

	// optional
	RenamedItem() objc.Object
	HasRenamedItem() bool
}

// ensure impl type implements protocol interface
var _ PFileProviderTestingCollisionResolution = (*FileProviderTestingCollisionResolutionObject)(nil)

// A concrete type for the [PFileProviderTestingCollisionResolution] protocol.
type FileProviderTestingCollisionResolutionObject struct {
	objc.Object
}

func (f_ FileProviderTestingCollisionResolutionObject) HasSide() bool {
	return f_.RespondsToSelector(objc.Sel("side"))
}

// The item’s location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingcollisionresolution/3736229-side?language=objc
func (f_ FileProviderTestingCollisionResolutionObject) Side() FileProviderTestingOperationSide {
	rv := objc.Call[FileProviderTestingOperationSide](f_, objc.Sel("side"))
	return rv
}

func (f_ FileProviderTestingCollisionResolutionObject) HasRenamedItem() bool {
	return f_.RespondsToSelector(objc.Sel("renamedItem"))
}

// A description of the renamed item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingcollisionresolution/3736228-renameditem?language=objc
func (f_ FileProviderTestingCollisionResolutionObject) RenamedItem() objc.Object {
	rv := objc.Call[objc.Object](f_, objc.Sel("renamedItem"))
	return rv
}
