// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a perspective rotate filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectiverotate?language=objc
type PPerspectiveRotate interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetRoll(value float64)
	HasSetRoll() bool

	// optional
	Roll() float64
	HasRoll() bool

	// optional
	SetYaw(value float64)
	HasSetYaw() bool

	// optional
	Yaw() float64
	HasYaw() bool

	// optional
	SetFocalLength(value float64)
	HasSetFocalLength() bool

	// optional
	FocalLength() float64
	HasFocalLength() bool

	// optional
	SetPitch(value float64)
	HasSetPitch() bool

	// optional
	Pitch() float64
	HasPitch() bool
}

// A concrete type wrapper for the [PPerspectiveRotate] protocol.
type PerspectiveRotateWrapper struct {
	objc.Object
}

func (p_ PerspectiveRotateWrapper) HasSetInputImage() bool {
	return p_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectiverotate/3325538-inputimage?language=objc
func (p_ PerspectiveRotateWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](p_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (p_ PerspectiveRotateWrapper) HasInputImage() bool {
	return p_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectiverotate/3325538-inputimage?language=objc
func (p_ PerspectiveRotateWrapper) InputImage() Image {
	rv := objc.Call[Image](p_, objc.Sel("inputImage"))
	return rv
}

func (p_ PerspectiveRotateWrapper) HasSetRoll() bool {
	return p_.RespondsToSelector(objc.Sel("setRoll:"))
}

// The roll angle, in radians. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectiverotate/3325540-roll?language=objc
func (p_ PerspectiveRotateWrapper) SetRoll(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setRoll:"), value)
}

func (p_ PerspectiveRotateWrapper) HasRoll() bool {
	return p_.RespondsToSelector(objc.Sel("roll"))
}

// The roll angle, in radians. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectiverotate/3325540-roll?language=objc
func (p_ PerspectiveRotateWrapper) Roll() float64 {
	rv := objc.Call[float64](p_, objc.Sel("roll"))
	return rv
}

func (p_ PerspectiveRotateWrapper) HasSetYaw() bool {
	return p_.RespondsToSelector(objc.Sel("setYaw:"))
}

// The yaw angle, in radians. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectiverotate/3325541-yaw?language=objc
func (p_ PerspectiveRotateWrapper) SetYaw(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setYaw:"), value)
}

func (p_ PerspectiveRotateWrapper) HasYaw() bool {
	return p_.RespondsToSelector(objc.Sel("yaw"))
}

// The yaw angle, in radians. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectiverotate/3325541-yaw?language=objc
func (p_ PerspectiveRotateWrapper) Yaw() float64 {
	rv := objc.Call[float64](p_, objc.Sel("yaw"))
	return rv
}

func (p_ PerspectiveRotateWrapper) HasSetFocalLength() bool {
	return p_.RespondsToSelector(objc.Sel("setFocalLength:"))
}

// The 35mm equivalent focal length of the input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectiverotate/3325537-focallength?language=objc
func (p_ PerspectiveRotateWrapper) SetFocalLength(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setFocalLength:"), value)
}

func (p_ PerspectiveRotateWrapper) HasFocalLength() bool {
	return p_.RespondsToSelector(objc.Sel("focalLength"))
}

// The 35mm equivalent focal length of the input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectiverotate/3325537-focallength?language=objc
func (p_ PerspectiveRotateWrapper) FocalLength() float64 {
	rv := objc.Call[float64](p_, objc.Sel("focalLength"))
	return rv
}

func (p_ PerspectiveRotateWrapper) HasSetPitch() bool {
	return p_.RespondsToSelector(objc.Sel("setPitch:"))
}

// The pitch angle, in radians. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectiverotate/3325539-pitch?language=objc
func (p_ PerspectiveRotateWrapper) SetPitch(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setPitch:"), value)
}

func (p_ PerspectiveRotateWrapper) HasPitch() bool {
	return p_.RespondsToSelector(objc.Sel("pitch"))
}

// The pitch angle, in radians. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectiverotate/3325539-pitch?language=objc
func (p_ PerspectiveRotateWrapper) Pitch() float64 {
	rv := objc.Call[float64](p_, objc.Sel("pitch"))
	return rv
}
