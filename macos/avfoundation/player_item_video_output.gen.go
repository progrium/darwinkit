// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/dispatch"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/corevideo"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PlayerItemVideoOutput] class.
var PlayerItemVideoOutputClass = _PlayerItemVideoOutputClass{objc.GetClass("AVPlayerItemVideoOutput")}

type _PlayerItemVideoOutputClass struct {
	objc.Class
}

// An interface definition for the [PlayerItemVideoOutput] class.
type IPlayerItemVideoOutput interface {
	IPlayerItemOutput
	SetDelegateQueue(delegate PPlayerItemOutputPullDelegate, delegateQueue dispatch.Queue)
	SetDelegateObjectQueue(delegateObject objc.IObject, delegateQueue dispatch.Queue)
	CopyPixelBufferForItemTimeItemTimeForDisplay(itemTime coremedia.Time, outItemTimeForDisplay *coremedia.Time) corevideo.PixelBufferRef
	RequestNotificationOfMediaDataChangeWithAdvanceInterval(interval foundation.TimeInterval)
	HasNewPixelBufferForItemTime(itemTime coremedia.Time) bool
	Delegate() PlayerItemOutputPullDelegateWrapper
	DelegateQueue() dispatch.Queue
}

// An object that outputs video frames from a player item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemvideooutput?language=objc
type PlayerItemVideoOutput struct {
	PlayerItemOutput
}

func PlayerItemVideoOutputFrom(ptr unsafe.Pointer) PlayerItemVideoOutput {
	return PlayerItemVideoOutput{
		PlayerItemOutput: PlayerItemOutputFrom(ptr),
	}
}

func (p_ PlayerItemVideoOutput) InitWithOutputSettings(outputSettings map[string]objc.IObject) PlayerItemVideoOutput {
	rv := objc.Call[PlayerItemVideoOutput](p_, objc.Sel("initWithOutputSettings:"), outputSettings)
	return rv
}

// Creates a video output object initialized with the specified output settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemvideooutput/1643270-initwithoutputsettings?language=objc
func NewPlayerItemVideoOutputWithOutputSettings(outputSettings map[string]objc.IObject) PlayerItemVideoOutput {
	instance := PlayerItemVideoOutputClass.Alloc().InitWithOutputSettings(outputSettings)
	instance.Autorelease()
	return instance
}

func (p_ PlayerItemVideoOutput) InitWithPixelBufferAttributes(pixelBufferAttributes map[string]objc.IObject) PlayerItemVideoOutput {
	rv := objc.Call[PlayerItemVideoOutput](p_, objc.Sel("initWithPixelBufferAttributes:"), pixelBufferAttributes)
	return rv
}

// Creates a video output object using the specified pixel buffer attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemvideooutput/1389231-initwithpixelbufferattributes?language=objc
func NewPlayerItemVideoOutputWithPixelBufferAttributes(pixelBufferAttributes map[string]objc.IObject) PlayerItemVideoOutput {
	instance := PlayerItemVideoOutputClass.Alloc().InitWithPixelBufferAttributes(pixelBufferAttributes)
	instance.Autorelease()
	return instance
}

func (pc _PlayerItemVideoOutputClass) Alloc() PlayerItemVideoOutput {
	rv := objc.Call[PlayerItemVideoOutput](pc, objc.Sel("alloc"))
	return rv
}

func PlayerItemVideoOutput_Alloc() PlayerItemVideoOutput {
	return PlayerItemVideoOutputClass.Alloc()
}

func (pc _PlayerItemVideoOutputClass) New() PlayerItemVideoOutput {
	rv := objc.Call[PlayerItemVideoOutput](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPlayerItemVideoOutput() PlayerItemVideoOutput {
	return PlayerItemVideoOutputClass.New()
}

func (p_ PlayerItemVideoOutput) Init() PlayerItemVideoOutput {
	rv := objc.Call[PlayerItemVideoOutput](p_, objc.Sel("init"))
	return rv
}

// Sets the delegate and dispatch queue for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemvideooutput/1386824-setdelegate?language=objc
func (p_ PlayerItemVideoOutput) SetDelegateQueue(delegate PPlayerItemOutputPullDelegate, delegateQueue dispatch.Queue) {
	po0 := objc.WrapAsProtocol("AVPlayerItemOutputPullDelegate", delegate)
	objc.Call[objc.Void](p_, objc.Sel("setDelegate:queue:"), po0, delegateQueue)
}

// Sets the delegate and dispatch queue for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemvideooutput/1386824-setdelegate?language=objc
func (p_ PlayerItemVideoOutput) SetDelegateObjectQueue(delegateObject objc.IObject, delegateQueue dispatch.Queue) {
	objc.Call[objc.Void](p_, objc.Sel("setDelegate:queue:"), objc.Ptr(delegateObject), delegateQueue)
}

// Retrieves an image that is appropriate for display at the specified item time, and marks the image as acquired. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemvideooutput/1386148-copypixelbufferforitemtime?language=objc
func (p_ PlayerItemVideoOutput) CopyPixelBufferForItemTimeItemTimeForDisplay(itemTime coremedia.Time, outItemTimeForDisplay *coremedia.Time) corevideo.PixelBufferRef {
	rv := objc.Call[corevideo.PixelBufferRef](p_, objc.Sel("copyPixelBufferForItemTime:itemTimeForDisplay:"), itemTime, outItemTimeForDisplay)
	return rv
}

// Tells the receiver that the video out put client is entering a quiescent state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemvideooutput/1386046-requestnotificationofmediadatach?language=objc
func (p_ PlayerItemVideoOutput) RequestNotificationOfMediaDataChangeWithAdvanceInterval(interval foundation.TimeInterval) {
	objc.Call[objc.Void](p_, objc.Sel("requestNotificationOfMediaDataChangeWithAdvanceInterval:"), interval)
}

// Returns a Boolean value that indicates whether video output is available for the specified item time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemvideooutput/1386444-hasnewpixelbufferforitemtime?language=objc
func (p_ PlayerItemVideoOutput) HasNewPixelBufferForItemTime(itemTime coremedia.Time) bool {
	rv := objc.Call[bool](p_, objc.Sel("hasNewPixelBufferForItemTime:"), itemTime)
	return rv
}

// The delegate for the video output object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemvideooutput/1385827-delegate?language=objc
func (p_ PlayerItemVideoOutput) Delegate() PlayerItemOutputPullDelegateWrapper {
	rv := objc.Call[PlayerItemOutputPullDelegateWrapper](p_, objc.Sel("delegate"))
	return rv
}

// The dispatch queue on which to call delegate methods. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemvideooutput/1388108-delegatequeue?language=objc
func (p_ PlayerItemVideoOutput) DelegateQueue() dispatch.Queue {
	rv := objc.Call[dispatch.Queue](p_, objc.Sel("delegateQueue"))
	return rv
}
