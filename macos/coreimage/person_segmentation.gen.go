// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipersonsegmentation?language=objc
type PPersonSegmentation interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetQualityLevel(value uint)
	HasSetQualityLevel() bool

	// optional
	QualityLevel() uint
	HasQualityLevel() bool
}

// A concrete type wrapper for the [PPersonSegmentation] protocol.
type PersonSegmentationWrapper struct {
	objc.Object
}

func (p_ PersonSegmentationWrapper) HasSetInputImage() bool {
	return p_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipersonsegmentation/3750392-inputimage?language=objc
func (p_ PersonSegmentationWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](p_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (p_ PersonSegmentationWrapper) HasInputImage() bool {
	return p_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipersonsegmentation/3750392-inputimage?language=objc
func (p_ PersonSegmentationWrapper) InputImage() Image {
	rv := objc.Call[Image](p_, objc.Sel("inputImage"))
	return rv
}

func (p_ PersonSegmentationWrapper) HasSetQualityLevel() bool {
	return p_.RespondsToSelector(objc.Sel("setQualityLevel:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipersonsegmentation/3784637-qualitylevel?language=objc
func (p_ PersonSegmentationWrapper) SetQualityLevel(value uint) {
	objc.Call[objc.Void](p_, objc.Sel("setQualityLevel:"), value)
}

func (p_ PersonSegmentationWrapper) HasQualityLevel() bool {
	return p_.RespondsToSelector(objc.Sel("qualityLevel"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipersonsegmentation/3784637-qualitylevel?language=objc
func (p_ PersonSegmentationWrapper) QualityLevel() uint {
	rv := objc.Call[uint](p_, objc.Sel("qualityLevel"))
	return rv
}
