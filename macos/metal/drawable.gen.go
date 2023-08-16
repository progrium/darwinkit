// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/macos/corefoundation"
	"github.com/progrium/macdriver/objc"
)

// A displayable resource that can be rendered or written to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldrawable?language=objc
type PDrawable interface {
	// optional
	AddPresentedHandler(block DrawablePresentedHandler)
	HasAddPresentedHandler() bool

	// optional
	Present()
	HasPresent() bool

	// optional
	PresentAfterMinimumDuration(duration corefoundation.TimeInterval)
	HasPresentAfterMinimumDuration() bool

	// optional
	PresentAtTime(presentationTime corefoundation.TimeInterval)
	HasPresentAtTime() bool

	// optional
	DrawableID() uint
	HasDrawableID() bool

	// optional
	PresentedTime() corefoundation.TimeInterval
	HasPresentedTime() bool
}

// A concrete type wrapper for the [PDrawable] protocol.
type DrawableWrapper struct {
	objc.Object
}

func (d_ DrawableWrapper) HasAddPresentedHandler() bool {
	return d_.RespondsToSelector(objc.Sel("addPresentedHandler:"))
}

// Registers a block of code to be called immediately after the drawable is presented. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldrawable/2806858-addpresentedhandler?language=objc
func (d_ DrawableWrapper) AddPresentedHandler(block DrawablePresentedHandler) {
	objc.Call[objc.Void](d_, objc.Sel("addPresentedHandler:"), block)
}

func (d_ DrawableWrapper) HasPresent() bool {
	return d_.RespondsToSelector(objc.Sel("present"))
}

// Presents the drawable onscreen as soon as possible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldrawable/1470284-present?language=objc
func (d_ DrawableWrapper) Present() {
	objc.Call[objc.Void](d_, objc.Sel("present"))
}

func (d_ DrawableWrapper) HasPresentAfterMinimumDuration() bool {
	return d_.RespondsToSelector(objc.Sel("presentAfterMinimumDuration:"))
}

// Presents the drawable onscreen as soon as possible after a previous drawable is visible for the specified duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldrawable/2806859-presentafterminimumduration?language=objc
func (d_ DrawableWrapper) PresentAfterMinimumDuration(duration corefoundation.TimeInterval) {
	objc.Call[objc.Void](d_, objc.Sel("presentAfterMinimumDuration:"), duration)
}

func (d_ DrawableWrapper) HasPresentAtTime() bool {
	return d_.RespondsToSelector(objc.Sel("presentAtTime:"))
}

// Presents the drawable onscreen at a specific host time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldrawable/1470282-presentattime?language=objc
func (d_ DrawableWrapper) PresentAtTime(presentationTime corefoundation.TimeInterval) {
	objc.Call[objc.Void](d_, objc.Sel("presentAtTime:"), presentationTime)
}

func (d_ DrawableWrapper) HasDrawableID() bool {
	return d_.RespondsToSelector(objc.Sel("drawableID"))
}

// A positive integer that identifies the drawable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldrawable/2806860-drawableid?language=objc
func (d_ DrawableWrapper) DrawableID() uint {
	rv := objc.Call[uint](d_, objc.Sel("drawableID"))
	return rv
}

func (d_ DrawableWrapper) HasPresentedTime() bool {
	return d_.RespondsToSelector(objc.Sel("presentedTime"))
}

// The host time, in seconds, when the drawable was displayed onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldrawable/2806855-presentedtime?language=objc
func (d_ DrawableWrapper) PresentedTime() corefoundation.TimeInterval {
	rv := objc.Call[corefoundation.TimeInterval](d_, objc.Sel("presentedTime"))
	return rv
}
