// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/macos/coremedia"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [MutableVideoComposition] class.
var MutableVideoCompositionClass = _MutableVideoCompositionClass{objc.GetClass("AVMutableVideoComposition")}

type _MutableVideoCompositionClass struct {
	objc.Class
}

// An interface definition for the [MutableVideoComposition] class.
type IMutableVideoComposition interface {
	IVideoComposition
	SetInstructions(value []objc.IObject)
	SetCustomVideoCompositorClass(value objc.IClass)
	SetAnimationTool(value IVideoCompositionCoreAnimationTool)
	SetColorPrimaries(value string)
	SetSourceTrackIDForFrameTiming(value objc.IObject)
	SetSourceSampleDataTrackIDs(value []foundation.INumber)
	SetColorTransferFunction(value string)
	SetColorYCbCrMatrix(value string)
	SetRenderSize(value coregraphics.Size)
	SetFrameDuration(value coremedia.Time)
	SetRenderScale(value float32)
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

// The video composition instructions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocomposition/1385815-instructions?language=objc
func (m_ MutableVideoComposition) SetInstructions(value []objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("setInstructions:"), value)
}

// The custom compositor class to use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocomposition/1390649-customvideocompositorclass?language=objc
func (m_ MutableVideoComposition) SetCustomVideoCompositorClass(value objc.IClass) {
	objc.Call[objc.Void](m_, objc.Sel("setCustomVideoCompositorClass:"), value)
}

// A video composition tool to use with Core Animation in offline rendering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocomposition/1390395-animationtool?language=objc
func (m_ MutableVideoComposition) SetAnimationTool(value IVideoCompositionCoreAnimationTool) {
	objc.Call[objc.Void](m_, objc.Sel("setAnimationTool:"), value)
}

// The color primaries used for video composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocomposition/1643234-colorprimaries?language=objc
func (m_ MutableVideoComposition) SetColorPrimaries(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setColorPrimaries:"), value)
}

// An identifier of the source track from which the video composition derives frame timing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocomposition/2873799-sourcetrackidforframetiming?language=objc
func (m_ MutableVideoComposition) SetSourceTrackIDForFrameTiming(value objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceTrackIDForFrameTiming:"), value)
}

// The identifiers of source sample data tracks in the composition that the compositor requires to compose frames. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocomposition/3750316-sourcesampledatatrackids?language=objc
func (m_ MutableVideoComposition) SetSourceSampleDataTrackIDs(value []foundation.INumber) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceSampleDataTrackIDs:"), value)
}

// The transfer function used for video composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocomposition/1643237-colortransferfunction?language=objc
func (m_ MutableVideoComposition) SetColorTransferFunction(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setColorTransferFunction:"), value)
}

// The YCbCr matrix used for video composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocomposition/1643231-colorycbcrmatrix?language=objc
func (m_ MutableVideoComposition) SetColorYCbCrMatrix(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setColorYCbCrMatrix:"), value)
}

// The size at which the video composition should render. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocomposition/1386365-rendersize?language=objc
func (m_ MutableVideoComposition) SetRenderSize(value coregraphics.Size) {
	objc.Call[objc.Void](m_, objc.Sel("setRenderSize:"), value)
}

// A time interval for which the video composition should render composed video frames. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocomposition/1390059-frameduration?language=objc
func (m_ MutableVideoComposition) SetFrameDuration(value coremedia.Time) {
	objc.Call[objc.Void](m_, objc.Sel("setFrameDuration:"), value)
}

// The scale at which the video composition should render. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocomposition/1615787-renderscale?language=objc
func (m_ MutableVideoComposition) SetRenderScale(value float32) {
	objc.Call[objc.Void](m_, objc.Sel("setRenderScale:"), value)
}
