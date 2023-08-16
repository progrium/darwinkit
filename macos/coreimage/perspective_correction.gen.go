// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a perspective correction filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectivecorrection?language=objc
type PPerspectiveCorrection interface {
	// optional
	SetCrop(value bool)
	HasSetCrop() bool

	// optional
	Crop() bool
	HasCrop() bool
}

// A concrete type wrapper for the [PPerspectiveCorrection] protocol.
type PerspectiveCorrectionWrapper struct {
	objc.Object
}

func (p_ PerspectiveCorrectionWrapper) HasSetCrop() bool {
	return p_.RespondsToSelector(objc.Sel("setCrop:"))
}

// A rectangle that specifies the extent of the corrected image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectivecorrection/3228648-crop?language=objc
func (p_ PerspectiveCorrectionWrapper) SetCrop(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setCrop:"), value)
}

func (p_ PerspectiveCorrectionWrapper) HasCrop() bool {
	return p_.RespondsToSelector(objc.Sel("crop"))
}

// A rectangle that specifies the extent of the corrected image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectivecorrection/3228648-crop?language=objc
func (p_ PerspectiveCorrectionWrapper) Crop() bool {
	rv := objc.Call[bool](p_, objc.Sel("crop"))
	return rv
}
