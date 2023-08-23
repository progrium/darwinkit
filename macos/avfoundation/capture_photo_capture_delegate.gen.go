// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"github.com/progrium/macdriver/objc"
)

// Methods for monitoring progress and receiving results from a photo capture output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotocapturedelegate?language=objc
type PCapturePhotoCaptureDelegate interface {
	// optional
	CaptureOutputDidCapturePhotoForResolvedSettings(output CapturePhotoOutput, resolvedSettings CaptureResolvedPhotoSettings)
	HasCaptureOutputDidCapturePhotoForResolvedSettings() bool
}

// A delegate implementation builder for the [PCapturePhotoCaptureDelegate] protocol.
type CapturePhotoCaptureDelegate struct {
	_CaptureOutputDidCapturePhotoForResolvedSettings func(output CapturePhotoOutput, resolvedSettings CaptureResolvedPhotoSettings)
}

func (di *CapturePhotoCaptureDelegate) HasCaptureOutputDidCapturePhotoForResolvedSettings() bool {
	return di._CaptureOutputDidCapturePhotoForResolvedSettings != nil
}

// Notifies the delegate that the photo has been taken. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotocapturedelegate/1778632-captureoutput?language=objc
func (di *CapturePhotoCaptureDelegate) SetCaptureOutputDidCapturePhotoForResolvedSettings(f func(output CapturePhotoOutput, resolvedSettings CaptureResolvedPhotoSettings)) {
	di._CaptureOutputDidCapturePhotoForResolvedSettings = f
}

// Notifies the delegate that the photo has been taken. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotocapturedelegate/1778632-captureoutput?language=objc
func (di *CapturePhotoCaptureDelegate) CaptureOutputDidCapturePhotoForResolvedSettings(output CapturePhotoOutput, resolvedSettings CaptureResolvedPhotoSettings) {
	di._CaptureOutputDidCapturePhotoForResolvedSettings(output, resolvedSettings)
}

// A concrete type wrapper for the [PCapturePhotoCaptureDelegate] protocol.
type CapturePhotoCaptureDelegateWrapper struct {
	objc.Object
}

func (c_ CapturePhotoCaptureDelegateWrapper) HasCaptureOutputDidCapturePhotoForResolvedSettings() bool {
	return c_.RespondsToSelector(objc.Sel("captureOutput:didCapturePhotoForResolvedSettings:"))
}

// Notifies the delegate that the photo has been taken. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotocapturedelegate/1778632-captureoutput?language=objc
func (c_ CapturePhotoCaptureDelegateWrapper) CaptureOutputDidCapturePhotoForResolvedSettings(output ICapturePhotoOutput, resolvedSettings ICaptureResolvedPhotoSettings) {
	objc.Call[objc.Void](c_, objc.Sel("captureOutput:didCapturePhotoForResolvedSettings:"), objc.Ptr(output), objc.Ptr(resolvedSettings))
}
