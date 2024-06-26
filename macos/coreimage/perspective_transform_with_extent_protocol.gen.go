// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/objc"
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

// ensure impl type implements protocol interface
var _ PPerspectiveTransformWithExtent = (*PerspectiveTransformWithExtentObject)(nil)

// A concrete type for the [PPerspectiveTransformWithExtent] protocol.
type PerspectiveTransformWithExtentObject struct {
	objc.Object
}

func (p_ PerspectiveTransformWithExtentObject) HasSetExtent() bool {
	return p_.RespondsToSelector(objc.Sel("setExtent:"))
}

// A rectangle that defines the extent of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectivetransformwithextent/3228667-extent?language=objc
func (p_ PerspectiveTransformWithExtentObject) SetExtent(value coregraphics.Rect) {
	objc.Call[objc.Void](p_, objc.Sel("setExtent:"), value)
}

func (p_ PerspectiveTransformWithExtentObject) HasExtent() bool {
	return p_.RespondsToSelector(objc.Sel("extent"))
}

// A rectangle that defines the extent of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectivetransformwithextent/3228667-extent?language=objc
func (p_ PerspectiveTransformWithExtentObject) Extent() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](p_, objc.Sel("extent"))
	return rv
}
