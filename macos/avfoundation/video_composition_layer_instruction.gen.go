// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [VideoCompositionLayerInstruction] class.
var VideoCompositionLayerInstructionClass = _VideoCompositionLayerInstructionClass{objc.GetClass("AVVideoCompositionLayerInstruction")}

type _VideoCompositionLayerInstructionClass struct {
	objc.Class
}

// An interface definition for the [VideoCompositionLayerInstruction] class.
type IVideoCompositionLayerInstruction interface {
	objc.IObject
	GetCropRectangleRampForTimeStartCropRectangleEndCropRectangleTimeRange(time coremedia.Time, startCropRectangle *coregraphics.Rect, endCropRectangle *coregraphics.Rect, timeRange *coremedia.TimeRange) bool
	GetOpacityRampForTimeStartOpacityEndOpacityTimeRange(time coremedia.Time, startOpacity *float64, endOpacity *float64, timeRange *coremedia.TimeRange) bool
	GetTransformRampForTimeStartTransformEndTransformTimeRange(time coremedia.Time, startTransform *coregraphics.AffineTransform, endTransform *coregraphics.AffineTransform, timeRange *coremedia.TimeRange) bool
	TrackID() objc.Object
}

// An object used to modify the transform, cropping, and opacity ramps applied to a given track in a composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositionlayerinstruction?language=objc
type VideoCompositionLayerInstruction struct {
	objc.Object
}

func VideoCompositionLayerInstructionFrom(ptr unsafe.Pointer) VideoCompositionLayerInstruction {
	return VideoCompositionLayerInstruction{
		Object: objc.ObjectFrom(ptr),
	}
}

func (vc _VideoCompositionLayerInstructionClass) Alloc() VideoCompositionLayerInstruction {
	rv := objc.Call[VideoCompositionLayerInstruction](vc, objc.Sel("alloc"))
	return rv
}

func VideoCompositionLayerInstruction_Alloc() VideoCompositionLayerInstruction {
	return VideoCompositionLayerInstructionClass.Alloc()
}

func (vc _VideoCompositionLayerInstructionClass) New() VideoCompositionLayerInstruction {
	rv := objc.Call[VideoCompositionLayerInstruction](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewVideoCompositionLayerInstruction() VideoCompositionLayerInstruction {
	return VideoCompositionLayerInstructionClass.New()
}

func (v_ VideoCompositionLayerInstruction) Init() VideoCompositionLayerInstruction {
	rv := objc.Call[VideoCompositionLayerInstruction](v_, objc.Sel("init"))
	return rv
}

// Obtains the crop rectangle ramp that includes the specified time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositionlayerinstruction/1387998-getcroprectanglerampfortime?language=objc
func (v_ VideoCompositionLayerInstruction) GetCropRectangleRampForTimeStartCropRectangleEndCropRectangleTimeRange(time coremedia.Time, startCropRectangle *coregraphics.Rect, endCropRectangle *coregraphics.Rect, timeRange *coremedia.TimeRange) bool {
	rv := objc.Call[bool](v_, objc.Sel("getCropRectangleRampForTime:startCropRectangle:endCropRectangle:timeRange:"), time, startCropRectangle, endCropRectangle, timeRange)
	return rv
}

// Obtains the opacity ramp that includes a specified time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositionlayerinstruction/1388471-getopacityrampfortime?language=objc
func (v_ VideoCompositionLayerInstruction) GetOpacityRampForTimeStartOpacityEndOpacityTimeRange(time coremedia.Time, startOpacity *float64, endOpacity *float64, timeRange *coremedia.TimeRange) bool {
	rv := objc.Call[bool](v_, objc.Sel("getOpacityRampForTime:startOpacity:endOpacity:timeRange:"), time, startOpacity, endOpacity, timeRange)
	return rv
}

// Obtains the transform ramp that includes a specified time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositionlayerinstruction/1387257-gettransformrampfortime?language=objc
func (v_ VideoCompositionLayerInstruction) GetTransformRampForTimeStartTransformEndTransformTimeRange(time coremedia.Time, startTransform *coregraphics.AffineTransform, endTransform *coregraphics.AffineTransform, timeRange *coremedia.TimeRange) bool {
	rv := objc.Call[bool](v_, objc.Sel("getTransformRampForTime:startTransform:endTransform:timeRange:"), time, startTransform, endTransform, timeRange)
	return rv
}

// The track identifier of the source track to which the compositor will apply the instruction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositionlayerinstruction/1390240-trackid?language=objc
func (v_ VideoCompositionLayerInstruction) TrackID() objc.Object {
	rv := objc.Call[objc.Object](v_, objc.Sel("trackID"))
	return rv
}
