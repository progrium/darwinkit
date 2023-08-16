// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciareareductionfilter?language=objc
type PAreaReductionFilter interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetExtent(value coregraphics.Rect)
	HasSetExtent() bool

	// optional
	Extent() coregraphics.Rect
	HasExtent() bool
}

// A concrete type wrapper for the [PAreaReductionFilter] protocol.
type AreaReductionFilterWrapper struct {
	objc.Object
}

func (a_ AreaReductionFilterWrapper) HasSetInputImage() bool {
	return a_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciareareductionfilter/3547102-inputimage?language=objc
func (a_ AreaReductionFilterWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](a_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (a_ AreaReductionFilterWrapper) HasInputImage() bool {
	return a_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciareareductionfilter/3547102-inputimage?language=objc
func (a_ AreaReductionFilterWrapper) InputImage() Image {
	rv := objc.Call[Image](a_, objc.Sel("inputImage"))
	return rv
}

func (a_ AreaReductionFilterWrapper) HasSetExtent() bool {
	return a_.RespondsToSelector(objc.Sel("setExtent:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciareareductionfilter/3547101-extent?language=objc
func (a_ AreaReductionFilterWrapper) SetExtent(value coregraphics.Rect) {
	objc.Call[objc.Void](a_, objc.Sel("setExtent:"), value)
}

func (a_ AreaReductionFilterWrapper) HasExtent() bool {
	return a_.RespondsToSelector(objc.Sel("extent"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciareareductionfilter/3547101-extent?language=objc
func (a_ AreaReductionFilterWrapper) Extent() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](a_, objc.Sel("extent"))
	return rv
}
