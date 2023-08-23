// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol for objects that contain information about an image size elsewhere in the graph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagesizeencodingstate?language=objc
type PImageSizeEncodingState interface {
	// optional
	SourceWidth() uint
	HasSourceWidth() bool

	// optional
	SourceHeight() uint
	HasSourceHeight() bool
}

// A concrete type wrapper for the [PImageSizeEncodingState] protocol.
type ImageSizeEncodingStateWrapper struct {
	objc.Object
}

func (i_ ImageSizeEncodingStateWrapper) HasSourceWidth() bool {
	return i_.RespondsToSelector(objc.Sel("sourceWidth"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagesizeencodingstate/2873364-sourcewidth?language=objc
func (i_ ImageSizeEncodingStateWrapper) SourceWidth() uint {
	rv := objc.Call[uint](i_, objc.Sel("sourceWidth"))
	return rv
}

func (i_ ImageSizeEncodingStateWrapper) HasSourceHeight() bool {
	return i_.RespondsToSelector(objc.Sel("sourceHeight"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagesizeencodingstate/2873329-sourceheight?language=objc
func (i_ ImageSizeEncodingStateWrapper) SourceHeight() uint {
	rv := objc.Call[uint](i_, objc.Sel("sourceHeight"))
	return rv
}
