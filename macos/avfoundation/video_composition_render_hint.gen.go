// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [VideoCompositionRenderHint] class.
var VideoCompositionRenderHintClass = _VideoCompositionRenderHintClass{objc.GetClass("AVVideoCompositionRenderHint")}

type _VideoCompositionRenderHintClass struct {
	objc.Class
}

// An interface definition for the [VideoCompositionRenderHint] class.
type IVideoCompositionRenderHint interface {
	objc.IObject
	StartCompositionTime() coremedia.Time
	EndCompositionTime() coremedia.Time
}

// Information about upcoming composition requests, such as composition start time and end time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositionrenderhint?language=objc
type VideoCompositionRenderHint struct {
	objc.Object
}

func VideoCompositionRenderHintFrom(ptr unsafe.Pointer) VideoCompositionRenderHint {
	return VideoCompositionRenderHint{
		Object: objc.ObjectFrom(ptr),
	}
}

func (vc _VideoCompositionRenderHintClass) Alloc() VideoCompositionRenderHint {
	rv := objc.Call[VideoCompositionRenderHint](vc, objc.Sel("alloc"))
	return rv
}

func VideoCompositionRenderHint_Alloc() VideoCompositionRenderHint {
	return VideoCompositionRenderHintClass.Alloc()
}

func (vc _VideoCompositionRenderHintClass) New() VideoCompositionRenderHint {
	rv := objc.Call[VideoCompositionRenderHint](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewVideoCompositionRenderHint() VideoCompositionRenderHint {
	return VideoCompositionRenderHintClass.New()
}

func (v_ VideoCompositionRenderHint) Init() VideoCompositionRenderHint {
	rv := objc.Call[VideoCompositionRenderHint](v_, objc.Sel("init"))
	return rv
}

// The start time of the upcoming composition requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositionrenderhint/3227889-startcompositiontime?language=objc
func (v_ VideoCompositionRenderHint) StartCompositionTime() coremedia.Time {
	rv := objc.Call[coremedia.Time](v_, objc.Sel("startCompositionTime"))
	return rv
}

// The end time of the upcoming composition requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositionrenderhint/3227888-endcompositiontime?language=objc
func (v_ VideoCompositionRenderHint) EndCompositionTime() coremedia.Time {
	rv := objc.Call[coremedia.Time](v_, objc.Sel("endCompositionTime"))
	return rv
}
