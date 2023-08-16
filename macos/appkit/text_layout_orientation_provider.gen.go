// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods that define the orientation of text for an object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutorientationprovider?language=objc
type PTextLayoutOrientationProvider interface {
	// optional
	LayoutOrientation() TextLayoutOrientation
	HasLayoutOrientation() bool
}

// A concrete type wrapper for the [PTextLayoutOrientationProvider] protocol.
type TextLayoutOrientationProviderWrapper struct {
	objc.Object
}

func (t_ TextLayoutOrientationProviderWrapper) HasLayoutOrientation() bool {
	return t_.RespondsToSelector(objc.Sel("layoutOrientation"))
}

// The default layout orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutorientationprovider/1402990-layoutorientation?language=objc
func (t_ TextLayoutOrientationProviderWrapper) LayoutOrientation() TextLayoutOrientation {
	rv := objc.Call[TextLayoutOrientation](t_, objc.Sel("layoutOrientation"))
	return rv
}
