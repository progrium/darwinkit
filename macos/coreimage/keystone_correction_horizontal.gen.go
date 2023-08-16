// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a keystone correction horizontal filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikeystonecorrectionhorizontal?language=objc
type PKeystoneCorrectionHorizontal interface {
	// optional
	SetFocalLength(value float64)
	HasSetFocalLength() bool

	// optional
	FocalLength() float64
	HasFocalLength() bool
}

// A concrete type wrapper for the [PKeystoneCorrectionHorizontal] protocol.
type KeystoneCorrectionHorizontalWrapper struct {
	objc.Object
}

func (k_ KeystoneCorrectionHorizontalWrapper) HasSetFocalLength() bool {
	return k_.RespondsToSelector(objc.Sel("setFocalLength:"))
}

// The 35mm equivalent focal length of the input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikeystonecorrectionhorizontal/3325525-focallength?language=objc
func (k_ KeystoneCorrectionHorizontalWrapper) SetFocalLength(value float64) {
	objc.Call[objc.Void](k_, objc.Sel("setFocalLength:"), value)
}

func (k_ KeystoneCorrectionHorizontalWrapper) HasFocalLength() bool {
	return k_.RespondsToSelector(objc.Sel("focalLength"))
}

// The 35mm equivalent focal length of the input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikeystonecorrectionhorizontal/3325525-focallength?language=objc
func (k_ KeystoneCorrectionHorizontalWrapper) FocalLength() float64 {
	rv := objc.Call[float64](k_, objc.Sel("focalLength"))
	return rv
}
