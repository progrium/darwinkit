// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a page-curl-with-shadow transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition?language=objc
type PPageCurlWithShadowTransition interface {
	// optional
	SetShadowExtent(value coregraphics.Rect)
	HasSetShadowExtent() bool

	// optional
	ShadowExtent() coregraphics.Rect
	HasShadowExtent() bool

	// optional
	SetShadowAmount(value float64)
	HasSetShadowAmount() bool

	// optional
	ShadowAmount() float64
	HasShadowAmount() bool

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
	SetShadowSize(value float64)
	HasSetShadowSize() bool

	// optional
	ShadowSize() float64
	HasShadowSize() bool

	// optional
	SetBacksideImage(value Image)
	HasSetBacksideImage() bool

	// optional
	BacksideImage() IImage
	HasBacksideImage() bool
}

// A concrete type wrapper for the [PPageCurlWithShadowTransition] protocol.
type PageCurlWithShadowTransitionWrapper struct {
	objc.Object
}

func (p_ PageCurlWithShadowTransitionWrapper) HasSetShadowExtent() bool {
	return p_.RespondsToSelector(objc.Sel("setShadowExtent:"))
}

// The rectagular portion of input image that casts a shadow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228629-shadowextent?language=objc
func (p_ PageCurlWithShadowTransitionWrapper) SetShadowExtent(value coregraphics.Rect) {
	objc.Call[objc.Void](p_, objc.Sel("setShadowExtent:"), value)
}

func (p_ PageCurlWithShadowTransitionWrapper) HasShadowExtent() bool {
	return p_.RespondsToSelector(objc.Sel("shadowExtent"))
}

// The rectagular portion of input image that casts a shadow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228629-shadowextent?language=objc
func (p_ PageCurlWithShadowTransitionWrapper) ShadowExtent() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](p_, objc.Sel("shadowExtent"))
	return rv
}

func (p_ PageCurlWithShadowTransitionWrapper) HasSetShadowAmount() bool {
	return p_.RespondsToSelector(objc.Sel("setShadowAmount:"))
}

// The strength of the shadow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228628-shadowamount?language=objc
func (p_ PageCurlWithShadowTransitionWrapper) SetShadowAmount(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setShadowAmount:"), value)
}

func (p_ PageCurlWithShadowTransitionWrapper) HasShadowAmount() bool {
	return p_.RespondsToSelector(objc.Sel("shadowAmount"))
}

// The strength of the shadow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228628-shadowamount?language=objc
func (p_ PageCurlWithShadowTransitionWrapper) ShadowAmount() float64 {
	rv := objc.Call[float64](p_, objc.Sel("shadowAmount"))
	return rv
}

func (p_ PageCurlWithShadowTransitionWrapper) HasSetRadius() bool {
	return p_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The radius of the curl. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228627-radius?language=objc
func (p_ PageCurlWithShadowTransitionWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setRadius:"), value)
}

func (p_ PageCurlWithShadowTransitionWrapper) HasRadius() bool {
	return p_.RespondsToSelector(objc.Sel("radius"))
}

// The radius of the curl. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228627-radius?language=objc
func (p_ PageCurlWithShadowTransitionWrapper) Radius() float64 {
	rv := objc.Call[float64](p_, objc.Sel("radius"))
	return rv
}

func (p_ PageCurlWithShadowTransitionWrapper) HasSetExtent() bool {
	return p_.RespondsToSelector(objc.Sel("setExtent:"))
}

// The extent of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228626-extent?language=objc
func (p_ PageCurlWithShadowTransitionWrapper) SetExtent(value coregraphics.Rect) {
	objc.Call[objc.Void](p_, objc.Sel("setExtent:"), value)
}

func (p_ PageCurlWithShadowTransitionWrapper) HasExtent() bool {
	return p_.RespondsToSelector(objc.Sel("extent"))
}

// The extent of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228626-extent?language=objc
func (p_ PageCurlWithShadowTransitionWrapper) Extent() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](p_, objc.Sel("extent"))
	return rv
}

func (p_ PageCurlWithShadowTransitionWrapper) HasSetAngle() bool {
	return p_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The angle of the curling page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228624-angle?language=objc
func (p_ PageCurlWithShadowTransitionWrapper) SetAngle(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setAngle:"), value)
}

func (p_ PageCurlWithShadowTransitionWrapper) HasAngle() bool {
	return p_.RespondsToSelector(objc.Sel("angle"))
}

// The angle of the curling page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228624-angle?language=objc
func (p_ PageCurlWithShadowTransitionWrapper) Angle() float64 {
	rv := objc.Call[float64](p_, objc.Sel("angle"))
	return rv
}

func (p_ PageCurlWithShadowTransitionWrapper) HasSetShadowSize() bool {
	return p_.RespondsToSelector(objc.Sel("setShadowSize:"))
}

// The maximum size, in pixels, of the shadow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228630-shadowsize?language=objc
func (p_ PageCurlWithShadowTransitionWrapper) SetShadowSize(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setShadowSize:"), value)
}

func (p_ PageCurlWithShadowTransitionWrapper) HasShadowSize() bool {
	return p_.RespondsToSelector(objc.Sel("shadowSize"))
}

// The maximum size, in pixels, of the shadow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228630-shadowsize?language=objc
func (p_ PageCurlWithShadowTransitionWrapper) ShadowSize() float64 {
	rv := objc.Call[float64](p_, objc.Sel("shadowSize"))
	return rv
}

func (p_ PageCurlWithShadowTransitionWrapper) HasSetBacksideImage() bool {
	return p_.RespondsToSelector(objc.Sel("setBacksideImage:"))
}

// The image that appears on the back of the source image as the page curls to reveal the target image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228625-backsideimage?language=objc
func (p_ PageCurlWithShadowTransitionWrapper) SetBacksideImage(value IImage) {
	objc.Call[objc.Void](p_, objc.Sel("setBacksideImage:"), objc.Ptr(value))
}

func (p_ PageCurlWithShadowTransitionWrapper) HasBacksideImage() bool {
	return p_.RespondsToSelector(objc.Sel("backsideImage"))
}

// The image that appears on the back of the source image as the page curls to reveal the target image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228625-backsideimage?language=objc
func (p_ PageCurlWithShadowTransitionWrapper) BacksideImage() Image {
	rv := objc.Call[Image](p_, objc.Sel("backsideImage"))
	return rv
}
