// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"github.com/progrium/macdriver/objc"
)

// An operation that resolves a collision by renaming the new item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingcollisionresolution?language=objc
type PFileProviderTestingCollisionResolution interface {
	// optional
	RenamedItem() objc.IObject
	HasRenamedItem() bool

	// optional
	Side() FileProviderTestingOperationSide
	HasSide() bool
}

// A concrete type wrapper for the [PFileProviderTestingCollisionResolution] protocol.
type FileProviderTestingCollisionResolutionWrapper struct {
	objc.Object
}

func (f_ FileProviderTestingCollisionResolutionWrapper) HasRenamedItem() bool {
	return f_.RespondsToSelector(objc.Sel("renamedItem"))
}

// A description of the renamed item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingcollisionresolution/3736228-renameditem?language=objc
func (f_ FileProviderTestingCollisionResolutionWrapper) RenamedItem() objc.Object {
	rv := objc.Call[objc.Object](f_, objc.Sel("renamedItem"))
	return rv
}

func (f_ FileProviderTestingCollisionResolutionWrapper) HasSide() bool {
	return f_.RespondsToSelector(objc.Sel("side"))
}

// The item’s location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingcollisionresolution/3736229-side?language=objc
func (f_ FileProviderTestingCollisionResolutionWrapper) Side() FileProviderTestingOperationSide {
	rv := objc.Call[FileProviderTestingOperationSide](f_, objc.Sel("side"))
	return rv
}
