// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetReaderVideoCompositionOutput] class.
var AssetReaderVideoCompositionOutputClass = _AssetReaderVideoCompositionOutputClass{objc.GetClass("AVAssetReaderVideoCompositionOutput")}

type _AssetReaderVideoCompositionOutputClass struct {
	objc.Class
}

// An interface definition for the [AssetReaderVideoCompositionOutput] class.
type IAssetReaderVideoCompositionOutput interface {
	IAssetReaderOutput
	VideoComposition() VideoComposition
	SetVideoComposition(value IVideoComposition)
	VideoSettings() map[string]objc.Object
	CustomVideoCompositor() VideoCompositingWrapper
	VideoTracks() []AssetTrack
}

// An object that reads composited video frames from one or more tracks of an asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadervideocompositionoutput?language=objc
type AssetReaderVideoCompositionOutput struct {
	AssetReaderOutput
}

func AssetReaderVideoCompositionOutputFrom(ptr unsafe.Pointer) AssetReaderVideoCompositionOutput {
	return AssetReaderVideoCompositionOutput{
		AssetReaderOutput: AssetReaderOutputFrom(ptr),
	}
}

func (a_ AssetReaderVideoCompositionOutput) InitWithVideoTracksVideoSettings(videoTracks []IAssetTrack, videoSettings map[string]objc.IObject) AssetReaderVideoCompositionOutput {
	rv := objc.Call[AssetReaderVideoCompositionOutput](a_, objc.Sel("initWithVideoTracks:videoSettings:"), videoTracks, videoSettings)
	return rv
}

// Creates an object that reads composited video frames from the specified video tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadervideocompositionoutput/1386676-initwithvideotracks?language=objc
func NewAssetReaderVideoCompositionOutputWithVideoTracksVideoSettings(videoTracks []IAssetTrack, videoSettings map[string]objc.IObject) AssetReaderVideoCompositionOutput {
	instance := AssetReaderVideoCompositionOutputClass.Alloc().InitWithVideoTracksVideoSettings(videoTracks, videoSettings)
	instance.Autorelease()
	return instance
}

func (ac _AssetReaderVideoCompositionOutputClass) AssetReaderVideoCompositionOutputWithVideoTracksVideoSettings(videoTracks []IAssetTrack, videoSettings map[string]objc.IObject) AssetReaderVideoCompositionOutput {
	rv := objc.Call[AssetReaderVideoCompositionOutput](ac, objc.Sel("assetReaderVideoCompositionOutputWithVideoTracks:videoSettings:"), videoTracks, videoSettings)
	return rv
}

// Returns a new object that reads composited video from the specified video tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadervideocompositionoutput/1490331-assetreadervideocompositionoutpu?language=objc
func AssetReaderVideoCompositionOutput_AssetReaderVideoCompositionOutputWithVideoTracksVideoSettings(videoTracks []IAssetTrack, videoSettings map[string]objc.IObject) AssetReaderVideoCompositionOutput {
	return AssetReaderVideoCompositionOutputClass.AssetReaderVideoCompositionOutputWithVideoTracksVideoSettings(videoTracks, videoSettings)
}

func (ac _AssetReaderVideoCompositionOutputClass) Alloc() AssetReaderVideoCompositionOutput {
	rv := objc.Call[AssetReaderVideoCompositionOutput](ac, objc.Sel("alloc"))
	return rv
}

func AssetReaderVideoCompositionOutput_Alloc() AssetReaderVideoCompositionOutput {
	return AssetReaderVideoCompositionOutputClass.Alloc()
}

func (ac _AssetReaderVideoCompositionOutputClass) New() AssetReaderVideoCompositionOutput {
	rv := objc.Call[AssetReaderVideoCompositionOutput](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetReaderVideoCompositionOutput() AssetReaderVideoCompositionOutput {
	return AssetReaderVideoCompositionOutputClass.New()
}

func (a_ AssetReaderVideoCompositionOutput) Init() AssetReaderVideoCompositionOutput {
	rv := objc.Call[AssetReaderVideoCompositionOutput](a_, objc.Sel("init"))
	return rv
}

// The video composition to use for the output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadervideocompositionoutput/1388927-videocomposition?language=objc
func (a_ AssetReaderVideoCompositionOutput) VideoComposition() VideoComposition {
	rv := objc.Call[VideoComposition](a_, objc.Sel("videoComposition"))
	return rv
}

// The video composition to use for the output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadervideocompositionoutput/1388927-videocomposition?language=objc
func (a_ AssetReaderVideoCompositionOutput) SetVideoComposition(value IVideoComposition) {
	objc.Call[objc.Void](a_, objc.Sel("setVideoComposition:"), objc.Ptr(value))
}

// The video settings that the output uses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadervideocompositionoutput/1389247-videosettings?language=objc
func (a_ AssetReaderVideoCompositionOutput) VideoSettings() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](a_, objc.Sel("videoSettings"))
	return rv
}

// A custom video compositor for the output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadervideocompositionoutput/1388310-customvideocompositor?language=objc
func (a_ AssetReaderVideoCompositionOutput) CustomVideoCompositor() VideoCompositingWrapper {
	rv := objc.Call[VideoCompositingWrapper](a_, objc.Sel("customVideoCompositor"))
	return rv
}

// The tracks from which the output reads the composited video. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadervideocompositionoutput/1389000-videotracks?language=objc
func (a_ AssetReaderVideoCompositionOutput) VideoTracks() []AssetTrack {
	rv := objc.Call[[]AssetTrack](a_, objc.Sel("videoTracks"))
	return rv
}
