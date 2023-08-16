// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a perspective transform with extent filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectivetransformwithextent?language=objc
type PPerspectiveTransformWithExtent interface {
	// optional
	SetExtent(value coregraphics.Rect)
	HasSetExtent() bool

	// optional
	Extent() coregraphics.Rect
	HasExtent() bool
}

// A concrete type wrapper for the [PPerspectiveTransformWithExtent] protocol.
type PerspectiveTransformWithExtentWrapper struct {
	objc.Object
}

func (p_ PerspectiveTransformWithExtentWrapper) HasSetExtent() bool {
	return p_.RespondsToSelector(objc.Sel("setExtent:"))
}

// A rectangle that defines the extent of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectivetransformwithextent/3228667-extent?language=objc
func (p_ PerspectiveTransformWithExtentWrapper) SetExtent(value coregraphics.Rect) {
	objc.Call[objc.Void](p_, objc.Sel("setExtent:"), value)
}

func (p_ PerspectiveTransformWithExtentWrapper) HasExtent() bool {
	return p_.RespondsToSelector(objc.Sel("extent"))
}

// A rectangle that defines the extent of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectivetransformwithextent/3228667-extent?language=objc
func (p_ PerspectiveTransformWithExtentWrapper) Extent() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](p_, objc.Sel("extent"))
	return rv
}
