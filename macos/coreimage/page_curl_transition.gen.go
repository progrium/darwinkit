// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a page curl transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurltransition?language=objc
type PPageCurlTransition interface {
	// optional
	SetRadius(value float64)
	HasSetRadius() bool

	// optional
	Radius() float64
	HasRadius() bool

	// optional
	SetExtent(value coregraphics.Rect)
	HasSetExtent() bool

	// optional
	Extent() coregraphics.Rect
	HasExtent() bool

	// optional
	SetAngle(value float64)
	HasSetAngle() bool

	// optional
	Angle() float64
	HasAngle() bool

	// optional
	SetBacksideImage(value Image)
	HasSetBacksideImage() bool

	// optional
	BacksideImage() IImage
	HasBacksideImage() bool

	// optional
	SetShadingImage(value Image)
	HasSetShadingImage() bool

	// optional
	ShadingImage() IImage
	HasShadingImage() bool
}

// A concrete type wrapper for the [PPageCurlTransition] protocol.
type PageCurlTransitionWrapper struct {
	objc.Object
}

func (p_ PageCurlTransitionWrapper) HasSetRadius() bool {
	return p_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The radius of the curl. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurltransition/3228621-radius?language=objc
func (p_ PageCurlTransitionWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setRadius:"), value)
}

func (p_ PageCurlTransitionWrapper) HasRadius() bool {
	return p_.RespondsToSelector(objc.Sel("radius"))
}

// The radius of the curl. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurltransition/3228621-radius?language=objc
func (p_ PageCurlTransitionWrapper) Radius() float64 {
	rv := objc.Call[float64](p_, objc.Sel("radius"))
	return rv
}

func (p_ PageCurlTransitionWrapper) HasSetExtent() bool {
	return p_.RespondsToSelector(objc.Sel("setExtent:"))
}

// The extent of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurltransition/3228620-extent?language=objc
func (p_ PageCurlTransitionWrapper) SetExtent(value coregraphics.Rect) {
	objc.Call[objc.Void](p_, objc.Sel("setExtent:"), value)
}

func (p_ PageCurlTransitionWrapper) HasExtent() bool {
	return p_.RespondsToSelector(objc.Sel("extent"))
}

// The extent of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurltransition/3228620-extent?language=objc
func (p_ PageCurlTransitionWrapper) Extent() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](p_, objc.Sel("extent"))
	return rv
}

func (p_ PageCurlTransitionWrapper) HasSetAngle() bool {
	return p_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The angle of the curling page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurltransition/3228618-angle?language=objc
func (p_ PageCurlTransitionWrapper) SetAngle(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setAngle:"), value)
}

func (p_ PageCurlTransitionWrapper) HasAngle() bool {
	return p_.RespondsToSelector(objc.Sel("angle"))
}

// The angle of the curling page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurltransition/3228618-angle?language=objc
func (p_ PageCurlTransitionWrapper) Angle() float64 {
	rv := objc.Call[float64](p_, objc.Sel("angle"))
	return rv
}

func (p_ PageCurlTransitionWrapper) HasSetBacksideImage() bool {
	return p_.RespondsToSelector(objc.Sel("setBacksideImage:"))
}

// The image that appears on the back of the source image as the page curls to reveal the target image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurltransition/3228619-backsideimage?language=objc
func (p_ PageCurlTransitionWrapper) SetBacksideImage(value IImage) {
	objc.Call[objc.Void](p_, objc.Sel("setBacksideImage:"), objc.Ptr(value))
}

func (p_ PageCurlTransitionWrapper) HasBacksideImage() bool {
	return p_.RespondsToSelector(objc.Sel("backsideImage"))
}

// The image that appears on the back of the source image as the page curls to reveal the target image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurltransition/3228619-backsideimage?language=objc
func (p_ PageCurlTransitionWrapper) BacksideImage() Image {
	rv := objc.Call[Image](p_, objc.Sel("backsideImage"))
	return rv
}

func (p_ PageCurlTransitionWrapper) HasSetShadingImage() bool {
	return p_.RespondsToSelector(objc.Sel("setShadingImage:"))
}

// An image that looks like a shaded sphere enclosed in a square. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurltransition/3228622-shadingimage?language=objc
func (p_ PageCurlTransitionWrapper) SetShadingImage(value IImage) {
	objc.Call[objc.Void](p_, objc.Sel("setShadingImage:"), objc.Ptr(value))
}

func (p_ PageCurlTransitionWrapper) HasShadingImage() bool {
	return p_.RespondsToSelector(objc.Sel("shadingImage"))
}

// An image that looks like a shaded sphere enclosed in a square. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurltransition/3228622-shadingimage?language=objc
func (p_ PageCurlTransitionWrapper) ShadingImage() Image {
	rv := objc.Call[Image](p_, objc.Sel("shadingImage"))
	return rv
}
