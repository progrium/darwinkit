// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [VideoProcessorCadence] class.
var VideoProcessorCadenceClass = _VideoProcessorCadenceClass{objc.GetClass("VNVideoProcessorCadence")}

type _VideoProcessorCadenceClass struct {
	objc.Class
}

// An interface definition for the [VideoProcessorCadence] class.
type IVideoProcessorCadence interface {
	objc.IObject
}

// An object that defines the cadence at which to process video. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvideoprocessorcadence?language=objc
type VideoProcessorCadence struct {
	objc.Object
}

func VideoProcessorCadenceFrom(ptr unsafe.Pointer) VideoProcessorCadence {
	return VideoProcessorCadence{
		Object: objc.ObjectFrom(ptr),
	}
}

func (vc _VideoProcessorCadenceClass) Alloc() VideoProcessorCadence {
	rv := objc.Call[VideoProcessorCadence](vc, objc.Sel("alloc"))
	return rv
}

func VideoProcessorCadence_Alloc() VideoProcessorCadence {
	return VideoProcessorCadenceClass.Alloc()
}

func (vc _VideoProcessorCadenceClass) New() VideoProcessorCadence {
	rv := objc.Call[VideoProcessorCadence](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewVideoProcessorCadence() VideoProcessorCadence {
	return VideoProcessorCadenceClass.New()
}

func (v_ VideoProcessorCadence) Init() VideoProcessorCadence {
	rv := objc.Call[VideoProcessorCadence](v_, objc.Sel("init"))
	return rv
}
