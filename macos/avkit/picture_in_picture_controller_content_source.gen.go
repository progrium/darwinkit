// AUTO-GENERATED CODE, DO NOT MODIFY

package avkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/avfoundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PictureInPictureControllerContentSource] class.
var PictureInPictureControllerContentSourceClass = _PictureInPictureControllerContentSourceClass{objc.GetClass("AVPictureInPictureControllerContentSource")}

type _PictureInPictureControllerContentSourceClass struct {
	objc.Class
}

// An interface definition for the [PictureInPictureControllerContentSource] class.
type IPictureInPictureControllerContentSource interface {
	objc.IObject
	SampleBufferDisplayLayer() avfoundation.SampleBufferDisplayLayer
	PlayerLayer() avfoundation.PlayerLayer
	SampleBufferPlaybackDelegate() PictureInPictureSampleBufferPlaybackDelegateWrapper
}

// An object that represents the source of the content to present in Picture in Picture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontrollercontentsource?language=objc
type PictureInPictureControllerContentSource struct {
	objc.Object
}

func PictureInPictureControllerContentSourceFrom(ptr unsafe.Pointer) PictureInPictureControllerContentSource {
	return PictureInPictureControllerContentSource{
		Object: objc.ObjectFrom(ptr),
	}
}

func (p_ PictureInPictureControllerContentSource) InitWithPlayerLayer(playerLayer avfoundation.IPlayerLayer) PictureInPictureControllerContentSource {
	rv := objc.Call[PictureInPictureControllerContentSource](p_, objc.Sel("initWithPlayerLayer:"), objc.Ptr(playerLayer))
	return rv
}

// Creates a content source with a player layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontrollercontentsource/3750326-initwithplayerlayer?language=objc
func NewPictureInPictureControllerContentSourceWithPlayerLayer(playerLayer avfoundation.IPlayerLayer) PictureInPictureControllerContentSource {
	instance := PictureInPictureControllerContentSourceClass.Alloc().InitWithPlayerLayer(playerLayer)
	instance.Autorelease()
	return instance
}

func (p_ PictureInPictureControllerContentSource) InitWithSampleBufferDisplayLayerPlaybackDelegate(sampleBufferDisplayLayer avfoundation.ISampleBufferDisplayLayer, playbackDelegate PPictureInPictureSampleBufferPlaybackDelegate) PictureInPictureControllerContentSource {
	po1 := objc.WrapAsProtocol("AVPictureInPictureSampleBufferPlaybackDelegate", playbackDelegate)
	rv := objc.Call[PictureInPictureControllerContentSource](p_, objc.Sel("initWithSampleBufferDisplayLayer:playbackDelegate:"), objc.Ptr(sampleBufferDisplayLayer), po1)
	return rv
}

// Creates a content source with a sample buffer display layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontrollercontentsource/3750329-initwithsamplebufferdisplaylayer?language=objc
func NewPictureInPictureControllerContentSourceWithSampleBufferDisplayLayerPlaybackDelegate(sampleBufferDisplayLayer avfoundation.ISampleBufferDisplayLayer, playbackDelegate PPictureInPictureSampleBufferPlaybackDelegate) PictureInPictureControllerContentSource {
	instance := PictureInPictureControllerContentSourceClass.Alloc().InitWithSampleBufferDisplayLayerPlaybackDelegate(sampleBufferDisplayLayer, playbackDelegate)
	instance.Autorelease()
	return instance
}

func (pc _PictureInPictureControllerContentSourceClass) Alloc() PictureInPictureControllerContentSource {
	rv := objc.Call[PictureInPictureControllerContentSource](pc, objc.Sel("alloc"))
	return rv
}

func PictureInPictureControllerContentSource_Alloc() PictureInPictureControllerContentSource {
	return PictureInPictureControllerContentSourceClass.Alloc()
}

func (pc _PictureInPictureControllerContentSourceClass) New() PictureInPictureControllerContentSource {
	rv := objc.Call[PictureInPictureControllerContentSource](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPictureInPictureControllerContentSource() PictureInPictureControllerContentSource {
	return PictureInPictureControllerContentSourceClass.New()
}

func (p_ PictureInPictureControllerContentSource) Init() PictureInPictureControllerContentSource {
	rv := objc.Call[PictureInPictureControllerContentSource](p_, objc.Sel("init"))
	return rv
}

// The presenting sample buffer display layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontrollercontentsource/3750330-samplebufferdisplaylayer?language=objc
func (p_ PictureInPictureControllerContentSource) SampleBufferDisplayLayer() avfoundation.SampleBufferDisplayLayer {
	rv := objc.Call[avfoundation.SampleBufferDisplayLayer](p_, objc.Sel("sampleBufferDisplayLayer"))
	return rv
}

// The presenting player layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontrollercontentsource/3750327-playerlayer?language=objc
func (p_ PictureInPictureControllerContentSource) PlayerLayer() avfoundation.PlayerLayer {
	rv := objc.Call[avfoundation.PlayerLayer](p_, objc.Sel("playerLayer"))
	return rv
}

// A delegate object that responds to sample buffer playback events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontrollercontentsource/3750331-samplebufferplaybackdelegate?language=objc
func (p_ PictureInPictureControllerContentSource) SampleBufferPlaybackDelegate() PictureInPictureSampleBufferPlaybackDelegateWrapper {
	rv := objc.Call[PictureInPictureSampleBufferPlaybackDelegateWrapper](p_, objc.Sel("sampleBufferPlaybackDelegate"))
	return rv
}
