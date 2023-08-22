// AUTO-GENERATED CODE, DO NOT MODIFY

package avkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/avfoundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PictureInPictureController] class.
var PictureInPictureControllerClass = _PictureInPictureControllerClass{objc.GetClass("AVPictureInPictureController")}

type _PictureInPictureControllerClass struct {
	objc.Class
}

// An interface definition for the [PictureInPictureController] class.
type IPictureInPictureController interface {
	objc.IObject
	StopPictureInPicture()
	InvalidatePlaybackState()
	StartPictureInPicture()
	IsPictureInPictureSuspended() bool
	RequiresLinearPlayback() bool
	SetRequiresLinearPlayback(value bool)
	Delegate() PictureInPictureControllerDelegateWrapper
	SetDelegate(value PPictureInPictureControllerDelegate)
	SetDelegateObject(valueObject objc.IObject)
	IsPictureInPictureActive() bool
	PlayerLayer() avfoundation.PlayerLayer
	IsPictureInPicturePossible() bool
	ContentSource() PictureInPictureControllerContentSource
	SetContentSource(value IPictureInPictureControllerContentSource)
}

// A controller that responds to user-initiated Picture in Picture playback of video in a floating, resizable window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller?language=objc
type PictureInPictureController struct {
	objc.Object
}

func PictureInPictureControllerFrom(ptr unsafe.Pointer) PictureInPictureController {
	return PictureInPictureController{
		Object: objc.ObjectFrom(ptr),
	}
}

func (p_ PictureInPictureController) InitWithPlayerLayer(playerLayer avfoundation.IPlayerLayer) PictureInPictureController {
	rv := objc.Call[PictureInPictureController](p_, objc.Sel("initWithPlayerLayer:"), objc.Ptr(playerLayer))
	return rv
}

// Creates a Picture in Picture controller with a player layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/1614707-initwithplayerlayer?language=objc
func NewPictureInPictureControllerWithPlayerLayer(playerLayer avfoundation.IPlayerLayer) PictureInPictureController {
	instance := PictureInPictureControllerClass.Alloc().InitWithPlayerLayer(playerLayer)
	instance.Autorelease()
	return instance
}

func (p_ PictureInPictureController) InitWithContentSource(contentSource IPictureInPictureControllerContentSource) PictureInPictureController {
	rv := objc.Call[PictureInPictureController](p_, objc.Sel("initWithContentSource:"), objc.Ptr(contentSource))
	return rv
}

// Creates a Picture in Picture controller with a content source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/3750324-initwithcontentsource?language=objc
func NewPictureInPictureControllerWithContentSource(contentSource IPictureInPictureControllerContentSource) PictureInPictureController {
	instance := PictureInPictureControllerClass.Alloc().InitWithContentSource(contentSource)
	instance.Autorelease()
	return instance
}

func (pc _PictureInPictureControllerClass) Alloc() PictureInPictureController {
	rv := objc.Call[PictureInPictureController](pc, objc.Sel("alloc"))
	return rv
}

func PictureInPictureController_Alloc() PictureInPictureController {
	return PictureInPictureControllerClass.Alloc()
}

func (pc _PictureInPictureControllerClass) New() PictureInPictureController {
	rv := objc.Call[PictureInPictureController](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPictureInPictureController() PictureInPictureController {
	return PictureInPictureControllerClass.New()
}

func (p_ PictureInPictureController) Init() PictureInPictureController {
	rv := objc.Call[PictureInPictureController](p_, objc.Sel("init"))
	return rv
}

// Returns a Boolean value that indicates whether the current device supports Picture in Picture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/1614693-ispictureinpicturesupported?language=objc
func (pc _PictureInPictureControllerClass) IsPictureInPictureSupported() bool {
	rv := objc.Call[bool](pc, objc.Sel("isPictureInPictureSupported"))
	return rv
}

// Returns a Boolean value that indicates whether the current device supports Picture in Picture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/1614693-ispictureinpicturesupported?language=objc
func PictureInPictureController_IsPictureInPictureSupported() bool {
	return PictureInPictureControllerClass.IsPictureInPictureSupported()
}

// Stops Picture in Picture, if active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/1614701-stoppictureinpicture?language=objc
func (p_ PictureInPictureController) StopPictureInPicture() {
	objc.Call[objc.Void](p_, objc.Sel("stopPictureInPicture"))
}

// Invalidates the controller’s current playback state and fetches the updated state from the sample buffer playback delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/3750328-invalidateplaybackstate?language=objc
func (p_ PictureInPictureController) InvalidatePlaybackState() {
	objc.Call[objc.Void](p_, objc.Sel("invalidatePlaybackState"))
}

// Starts Picture in Picture, if possible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/1614687-startpictureinpicture?language=objc
func (p_ PictureInPictureController) StartPictureInPicture() {
	objc.Call[objc.Void](p_, objc.Sel("startPictureInPicture"))
}

// A Boolean value that indicates whether the system suspends the controller’s Picture in Picture window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/1614689-pictureinpicturesuspended?language=objc
func (p_ PictureInPictureController) IsPictureInPictureSuspended() bool {
	rv := objc.Call[bool](p_, objc.Sel("isPictureInPictureSuspended"))
	return rv
}

// A Boolean value that determines whether the controller allows the user to skip media content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/3566335-requireslinearplayback?language=objc
func (p_ PictureInPictureController) RequiresLinearPlayback() bool {
	rv := objc.Call[bool](p_, objc.Sel("requiresLinearPlayback"))
	return rv
}

// A Boolean value that determines whether the controller allows the user to skip media content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/3566335-requireslinearplayback?language=objc
func (p_ PictureInPictureController) SetRequiresLinearPlayback(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setRequiresLinearPlayback:"), value)
}

// A delegate object for a Picture in Picture controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/1614709-delegate?language=objc
func (p_ PictureInPictureController) Delegate() PictureInPictureControllerDelegateWrapper {
	rv := objc.Call[PictureInPictureControllerDelegateWrapper](p_, objc.Sel("delegate"))
	return rv
}

// A delegate object for a Picture in Picture controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/1614709-delegate?language=objc
func (p_ PictureInPictureController) SetDelegate(value PPictureInPictureControllerDelegate) {
	po0 := objc.WrapAsProtocol("AVPictureInPictureControllerDelegate", value)
	objc.SetAssociatedObject(p_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](p_, objc.Sel("setDelegate:"), po0)
}

// A delegate object for a Picture in Picture controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/1614709-delegate?language=objc
func (p_ PictureInPictureController) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](p_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// A system-default template image for the button that starts Picture in Picture in your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/3172686-pictureinpicturebuttonstartimage?language=objc
func (pc _PictureInPictureControllerClass) PictureInPictureButtonStartImage() appkit.Image {
	rv := objc.Call[appkit.Image](pc, objc.Sel("pictureInPictureButtonStartImage"))
	return rv
}

// A system-default template image for the button that starts Picture in Picture in your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/3172686-pictureinpicturebuttonstartimage?language=objc
func PictureInPictureController_PictureInPictureButtonStartImage() appkit.Image {
	return PictureInPictureControllerClass.PictureInPictureButtonStartImage()
}

// A Boolean value that indicates whether the Picture in Picture window is onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/1614720-pictureinpictureactive?language=objc
func (p_ PictureInPictureController) IsPictureInPictureActive() bool {
	rv := objc.Call[bool](p_, objc.Sel("isPictureInPictureActive"))
	return rv
}

// The layer that displays the video content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/1614706-playerlayer?language=objc
func (p_ PictureInPictureController) PlayerLayer() avfoundation.PlayerLayer {
	rv := objc.Call[avfoundation.PlayerLayer](p_, objc.Sel("playerLayer"))
	return rv
}

// A system-default template image for the button that stops Picture in Picture in your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/3172687-pictureinpicturebuttonstopimage?language=objc
func (pc _PictureInPictureControllerClass) PictureInPictureButtonStopImage() appkit.Image {
	rv := objc.Call[appkit.Image](pc, objc.Sel("pictureInPictureButtonStopImage"))
	return rv
}

// A system-default template image for the button that stops Picture in Picture in your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/3172687-pictureinpicturebuttonstopimage?language=objc
func PictureInPictureController_PictureInPictureButtonStopImage() appkit.Image {
	return PictureInPictureControllerClass.PictureInPictureButtonStopImage()
}

// A Boolean value that indicates whether Picture in Picture playback is currently possible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/1614691-pictureinpicturepossible?language=objc
func (p_ PictureInPictureController) IsPictureInPicturePossible() bool {
	rv := objc.Call[bool](p_, objc.Sel("isPictureInPicturePossible"))
	return rv
}

// The source of the controller’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/3750323-contentsource?language=objc
func (p_ PictureInPictureController) ContentSource() PictureInPictureControllerContentSource {
	rv := objc.Call[PictureInPictureControllerContentSource](p_, objc.Sel("contentSource"))
	return rv
}

// The source of the controller’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/3750323-contentsource?language=objc
func (p_ PictureInPictureController) SetContentSource(value IPictureInPictureControllerContentSource) {
	objc.Call[objc.Void](p_, objc.Sel("setContentSource:"), objc.Ptr(value))
}
