// AUTO-GENERATED CODE, DO NOT MODIFY

package avkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A protocol to adopt to respond to Picture in Picture events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontrollerdelegate?language=objc
type PPictureInPictureControllerDelegate interface {
	// optional
	PictureInPictureControllerDidStartPictureInPicture(pictureInPictureController PictureInPictureController)
	HasPictureInPictureControllerDidStartPictureInPicture() bool

	// optional
	PictureInPictureControllerDidStopPictureInPicture(pictureInPictureController PictureInPictureController)
	HasPictureInPictureControllerDidStopPictureInPicture() bool

	// optional
	PictureInPictureControllerWillStopPictureInPicture(pictureInPictureController PictureInPictureController)
	HasPictureInPictureControllerWillStopPictureInPicture() bool

	// optional
	PictureInPictureControllerFailedToStartPictureInPictureWithError(pictureInPictureController PictureInPictureController, error foundation.Error)
	HasPictureInPictureControllerFailedToStartPictureInPictureWithError() bool
}

// A delegate implementation builder for the [PPictureInPictureControllerDelegate] protocol.
type PictureInPictureControllerDelegate struct {
	_PictureInPictureControllerDidStartPictureInPicture               func(pictureInPictureController PictureInPictureController)
	_PictureInPictureControllerDidStopPictureInPicture                func(pictureInPictureController PictureInPictureController)
	_PictureInPictureControllerWillStopPictureInPicture               func(pictureInPictureController PictureInPictureController)
	_PictureInPictureControllerFailedToStartPictureInPictureWithError func(pictureInPictureController PictureInPictureController, error foundation.Error)
}

func (di *PictureInPictureControllerDelegate) HasPictureInPictureControllerDidStartPictureInPicture() bool {
	return di._PictureInPictureControllerDidStartPictureInPicture != nil
}

// Tells the delegate that Picture in Picture started. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontrollerdelegate/1614695-pictureinpicturecontrollerdidsta?language=objc
func (di *PictureInPictureControllerDelegate) SetPictureInPictureControllerDidStartPictureInPicture(f func(pictureInPictureController PictureInPictureController)) {
	di._PictureInPictureControllerDidStartPictureInPicture = f
}

// Tells the delegate that Picture in Picture started. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontrollerdelegate/1614695-pictureinpicturecontrollerdidsta?language=objc
func (di *PictureInPictureControllerDelegate) PictureInPictureControllerDidStartPictureInPicture(pictureInPictureController PictureInPictureController) {
	di._PictureInPictureControllerDidStartPictureInPicture(pictureInPictureController)
}
func (di *PictureInPictureControllerDelegate) HasPictureInPictureControllerDidStopPictureInPicture() bool {
	return di._PictureInPictureControllerDidStopPictureInPicture != nil
}

// Tells the delegate that Picture in Picture stopped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontrollerdelegate/1614717-pictureinpicturecontrollerdidsto?language=objc
func (di *PictureInPictureControllerDelegate) SetPictureInPictureControllerDidStopPictureInPicture(f func(pictureInPictureController PictureInPictureController)) {
	di._PictureInPictureControllerDidStopPictureInPicture = f
}

// Tells the delegate that Picture in Picture stopped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontrollerdelegate/1614717-pictureinpicturecontrollerdidsto?language=objc
func (di *PictureInPictureControllerDelegate) PictureInPictureControllerDidStopPictureInPicture(pictureInPictureController PictureInPictureController) {
	di._PictureInPictureControllerDidStopPictureInPicture(pictureInPictureController)
}
func (di *PictureInPictureControllerDelegate) HasPictureInPictureControllerWillStopPictureInPicture() bool {
	return di._PictureInPictureControllerWillStopPictureInPicture != nil
}

// Tells the delegate that Picture in Picture is about to stop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontrollerdelegate/1614719-pictureinpicturecontrollerwillst?language=objc
func (di *PictureInPictureControllerDelegate) SetPictureInPictureControllerWillStopPictureInPicture(f func(pictureInPictureController PictureInPictureController)) {
	di._PictureInPictureControllerWillStopPictureInPicture = f
}

// Tells the delegate that Picture in Picture is about to stop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontrollerdelegate/1614719-pictureinpicturecontrollerwillst?language=objc
func (di *PictureInPictureControllerDelegate) PictureInPictureControllerWillStopPictureInPicture(pictureInPictureController PictureInPictureController) {
	di._PictureInPictureControllerWillStopPictureInPicture(pictureInPictureController)
}
func (di *PictureInPictureControllerDelegate) HasPictureInPictureControllerFailedToStartPictureInPictureWithError() bool {
	return di._PictureInPictureControllerFailedToStartPictureInPictureWithError != nil
}

// Tells the delegate that Picture in Picture failed to start. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontrollerdelegate/1614697-pictureinpicturecontroller?language=objc
func (di *PictureInPictureControllerDelegate) SetPictureInPictureControllerFailedToStartPictureInPictureWithError(f func(pictureInPictureController PictureInPictureController, error foundation.Error)) {
	di._PictureInPictureControllerFailedToStartPictureInPictureWithError = f
}

// Tells the delegate that Picture in Picture failed to start. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontrollerdelegate/1614697-pictureinpicturecontroller?language=objc
func (di *PictureInPictureControllerDelegate) PictureInPictureControllerFailedToStartPictureInPictureWithError(pictureInPictureController PictureInPictureController, error foundation.Error) {
	di._PictureInPictureControllerFailedToStartPictureInPictureWithError(pictureInPictureController, error)
}

// A concrete type wrapper for the [PPictureInPictureControllerDelegate] protocol.
type PictureInPictureControllerDelegateWrapper struct {
	objc.Object
}

func (p_ PictureInPictureControllerDelegateWrapper) HasPictureInPictureControllerDidStartPictureInPicture() bool {
	return p_.RespondsToSelector(objc.Sel("pictureInPictureControllerDidStartPictureInPicture:"))
}

// Tells the delegate that Picture in Picture started. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontrollerdelegate/1614695-pictureinpicturecontrollerdidsta?language=objc
func (p_ PictureInPictureControllerDelegateWrapper) PictureInPictureControllerDidStartPictureInPicture(pictureInPictureController IPictureInPictureController) {
	objc.Call[objc.Void](p_, objc.Sel("pictureInPictureControllerDidStartPictureInPicture:"), objc.Ptr(pictureInPictureController))
}

func (p_ PictureInPictureControllerDelegateWrapper) HasPictureInPictureControllerDidStopPictureInPicture() bool {
	return p_.RespondsToSelector(objc.Sel("pictureInPictureControllerDidStopPictureInPicture:"))
}

// Tells the delegate that Picture in Picture stopped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontrollerdelegate/1614717-pictureinpicturecontrollerdidsto?language=objc
func (p_ PictureInPictureControllerDelegateWrapper) PictureInPictureControllerDidStopPictureInPicture(pictureInPictureController IPictureInPictureController) {
	objc.Call[objc.Void](p_, objc.Sel("pictureInPictureControllerDidStopPictureInPicture:"), objc.Ptr(pictureInPictureController))
}

func (p_ PictureInPictureControllerDelegateWrapper) HasPictureInPictureControllerWillStopPictureInPicture() bool {
	return p_.RespondsToSelector(objc.Sel("pictureInPictureControllerWillStopPictureInPicture:"))
}

// Tells the delegate that Picture in Picture is about to stop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontrollerdelegate/1614719-pictureinpicturecontrollerwillst?language=objc
func (p_ PictureInPictureControllerDelegateWrapper) PictureInPictureControllerWillStopPictureInPicture(pictureInPictureController IPictureInPictureController) {
	objc.Call[objc.Void](p_, objc.Sel("pictureInPictureControllerWillStopPictureInPicture:"), objc.Ptr(pictureInPictureController))
}

func (p_ PictureInPictureControllerDelegateWrapper) HasPictureInPictureControllerFailedToStartPictureInPictureWithError() bool {
	return p_.RespondsToSelector(objc.Sel("pictureInPictureController:failedToStartPictureInPictureWithError:"))
}

// Tells the delegate that Picture in Picture failed to start. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontrollerdelegate/1614697-pictureinpicturecontroller?language=objc
func (p_ PictureInPictureControllerDelegateWrapper) PictureInPictureControllerFailedToStartPictureInPictureWithError(pictureInPictureController IPictureInPictureController, error foundation.IError) {
	objc.Call[objc.Void](p_, objc.Sel("pictureInPictureController:failedToStartPictureInPictureWithError:"), objc.Ptr(pictureInPictureController), objc.Ptr(error))
}
