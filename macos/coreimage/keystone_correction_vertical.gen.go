// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a keystone correction vertical filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikeystonecorrectionvertical?language=objc
type PKeystoneCorrectionVertical interface {
	// optional
	SetFocalLength(value float64)
	HasSetFocalLength() bool

	// optional
	FocalLength() float64
	HasFocalLength() bool
}

// A concrete type wrapper for the [PKeystoneCorrectionVertical] protocol.
type KeystoneCorrectionVerticalWrapper struct {
	objc.Object
}

func (k_ KeystoneCorrectionVerticalWrapper) HasSetFocalLength() bool {
	return k_.RespondsToSelector(objc.Sel("setFocalLength:"))
}

// The 35mm equivalent focal length of the input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikeystonecorrectionvertical/3325532-focallength?language=objc
func (k_ KeystoneCorrectionVerticalWrapper) SetFocalLength(value float64) {
	objc.Call[objc.Void](k_, objc.Sel("setFocalLength:"), value)
}

func (k_ KeystoneCorrectionVerticalWrapper) HasFocalLength() bool {
	return k_.RespondsToSelector(objc.Sel("focalLength"))
}

// The 35mm equivalent focal length of the input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikeystonecorrectionvertical/3325532-focallength?language=objc
func (k_ KeystoneCorrectionVerticalWrapper) FocalLength() float64 {
	rv := objc.Call[float64](k_, objc.Sel("focalLength"))
	return rv
}
