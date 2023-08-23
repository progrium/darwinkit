// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutableVideoComposition] class.
var MutableVideoCompositionClass = _MutableVideoCompositionClass{objc.GetClass("AVMutableVideoComposition")}

type _MutableVideoCompositionClass struct {
	objc.Class
}

// An interface definition for the [MutableVideoComposition] class.
type IMutableVideoComposition interface {
	IVideoComposition
	SetColorPrimaries(value string)
	SetColorYCbCrMatrix(value string)
	SetCustomVideoCompositorClass(value objc.IClass)
	SetFrameDuration(value coremedia.Time)
	SetInstructions(value []objc.IObject)
	SetSourceSampleDataTrackIDs(value []foundation.INumber)
	SetAnimationTool(value IVideoCompositionCoreAnimationTool)
	SetColorTransferFunction(value string)
	SetSourceTrackIDForFrameTiming(value objc.IObject)
	SetRenderSize(value coregraphics.Size)
	SetRenderScale(value float64)
}

// A mutable video composition subclass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocomposition?language=objc
type MutableVideoComposition struct {
	VideoComposition
}

func MutableVideoCompositionFrom(ptr unsafe.Pointer) MutableVideoComposition {
	return MutableVideoComposition{
		VideoComposition: VideoCompositionFrom(ptr),
	}
}

func (mc _MutableVideoCompositionClass) Alloc() MutableVideoComposition {
	rv := objc.Call[MutableVideoComposition](mc, objc.Sel("alloc"))
	return rv
}

func MutableVideoComposition_Alloc() MutableVideoComposition {
	return MutableVideoCompositionClass.Alloc()
}

func (mc _MutableVideoCompositionClass) New() MutableVideoComposition {
	rv := objc.Call[MutableVideoComposition](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableVideoComposition() MutableVideoComposition {
	return MutableVideoCompositionClass.New()
}

func (m_ MutableVideoComposition) Init() MutableVideoComposition {
	rv := objc.Call[MutableVideoComposition](m_, objc.Sel("init"))
	return rv
}

// Creates a new mutable video composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocomposition/1519720-videocomposition?language=objc
func (mc _MutableVideoCompositionClass) VideoComposition() MutableVideoComposition {
	rv := objc.Call[MutableVideoComposition](mc, objc.Sel("videoComposition"))
	return rv
}

// Creates a new mutable video composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocomposition/1519720-videocomposition?language=objc
func MutableVideoComposition_VideoComposition() MutableVideoComposition {
	return MutableVideoCompositionClass.VideoComposition()
}

// The color primaries used for video composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocomposition/1643234-colorprimaries?language=objc
func (m_ MutableVideoComposition) SetColorPrimaries(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setColorPrimaries:"), value)
}

// The YCbCr matrix used for video composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocomposition/1643231-colorycbcrmatrix?language=objc
func (m_ MutableVideoComposition) SetColorYCbCrMatrix(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setColorYCbCrMatrix:"), value)
}

// The custom compositor class to use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocomposition/1390649-customvideocompositorclass?language=objc
func (m_ MutableVideoComposition) SetCustomVideoCompositorClass(value objc.IClass) {
	objc.Call[objc.Void](m_, objc.Sel("setCustomVideoCompositorClass:"), objc.Ptr(value))
}

// A time interval for which the video composition should render composed video frames. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocomposition/1390059-frameduration?language=objc
func (m_ MutableVideoComposition) SetFrameDuration(value coremedia.Time) {
	objc.Call[objc.Void](m_, objc.Sel("setFrameDuration:"), value)
}

// The video composition instructions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocomposition/1385815-instructions?language=objc
func (m_ MutableVideoComposition) SetInstructions(value []objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("setInstructions:"), value)
}

// The identifiers of source sample data tracks in the composition that the compositor requires to compose frames. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocomposition/3750316-sourcesampledatatrackids?language=objc
func (m_ MutableVideoComposition) SetSourceSampleDataTrackIDs(value []foundation.INumber) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceSampleDataTrackIDs:"), value)
}

// A video composition tool to use with Core Animation in offline rendering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocomposition/1390395-animationtool?language=objc
func (m_ MutableVideoComposition) SetAnimationTool(value IVideoCompositionCoreAnimationTool) {
	objc.Call[objc.Void](m_, objc.Sel("setAnimationTool:"), objc.Ptr(value))
}

// The transfer function used for video composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocomposition/1643237-colortransferfunction?language=objc
func (m_ MutableVideoComposition) SetColorTransferFunction(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setColorTransferFunction:"), value)
}

// An identifier of the source track from which the video composition derives frame timing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocomposition/2873799-sourcetrackidforframetiming?language=objc
func (m_ MutableVideoComposition) SetSourceTrackIDForFrameTiming(value objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceTrackIDForFrameTiming:"), objc.Ptr(value))
}

// The size at which the video composition should render. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocomposition/1386365-rendersize?language=objc
func (m_ MutableVideoComposition) SetRenderSize(value coregraphics.Size) {
	objc.Call[objc.Void](m_, objc.Sel("setRenderSize:"), value)
}

// The scale at which the video composition should render. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocomposition/1615787-renderscale?language=objc
func (m_ MutableVideoComposition) SetRenderScale(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setRenderScale:"), value)
}
