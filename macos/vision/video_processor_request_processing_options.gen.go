// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [VideoProcessorRequestProcessingOptions] class.
var VideoProcessorRequestProcessingOptionsClass = _VideoProcessorRequestProcessingOptionsClass{objc.GetClass("VNVideoProcessorRequestProcessingOptions")}

type _VideoProcessorRequestProcessingOptionsClass struct {
	objc.Class
}

// An interface definition for the [VideoProcessorRequestProcessingOptions] class.
type IVideoProcessorRequestProcessingOptions interface {
	objc.IObject
	Cadence() VideoProcessorCadence
	SetCadence(value IVideoProcessorCadence)
}

// An object that defines a video processor’s configuration options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvideoprocessorrequestprocessingoptions?language=objc
type VideoProcessorRequestProcessingOptions struct {
	objc.Object
}

func VideoProcessorRequestProcessingOptionsFrom(ptr unsafe.Pointer) VideoProcessorRequestProcessingOptions {
	return VideoProcessorRequestProcessingOptions{
		Object: objc.ObjectFrom(ptr),
	}
}

func (vc _VideoProcessorRequestProcessingOptionsClass) Alloc() VideoProcessorRequestProcessingOptions {
	rv := objc.Call[VideoProcessorRequestProcessingOptions](vc, objc.Sel("alloc"))
	return rv
}

func VideoProcessorRequestProcessingOptions_Alloc() VideoProcessorRequestProcessingOptions {
	return VideoProcessorRequestProcessingOptionsClass.Alloc()
}

func (vc _VideoProcessorRequestProcessingOptionsClass) New() VideoProcessorRequestProcessingOptions {
	rv := objc.Call[VideoProcessorRequestProcessingOptions](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewVideoProcessorRequestProcessingOptions() VideoProcessorRequestProcessingOptions {
	return VideoProcessorRequestProcessingOptionsClass.New()
}

func (v_ VideoProcessorRequestProcessingOptions) Init() VideoProcessorRequestProcessingOptions {
	rv := objc.Call[VideoProcessorRequestProcessingOptions](v_, objc.Sel("init"))
	return rv
}

// The cadence the video processor maintains to process the request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvideoprocessorrequestprocessingoptions/3675684-cadence?language=objc
func (v_ VideoProcessorRequestProcessingOptions) Cadence() VideoProcessorCadence {
	rv := objc.Call[VideoProcessorCadence](v_, objc.Sel("cadence"))
	return rv
}

// The cadence the video processor maintains to process the request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvideoprocessorrequestprocessingoptions/3675684-cadence?language=objc
func (v_ VideoProcessorRequestProcessingOptions) SetCadence(value IVideoProcessorCadence) {
	objc.Call[objc.Void](v_, objc.Sel("setCadence:"), objc.Ptr(value))
}
