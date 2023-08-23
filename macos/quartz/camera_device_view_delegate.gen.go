// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The IKCameraDeviceViewDelegate protocol is adopted by the delegate of the IKCameraDeviceView class. It allows downloading of camera content, handling selection changes, and handling errors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceviewdelegate?language=objc
type PCameraDeviceViewDelegate interface {
	// optional
	CameraDeviceViewDidEncounterError(cameraDeviceView CameraDeviceView, error foundation.Error)
	HasCameraDeviceViewDidEncounterError() bool

	// optional
	CameraDeviceViewSelectionDidChange(cameraDeviceView CameraDeviceView)
	HasCameraDeviceViewSelectionDidChange() bool
}

// A delegate implementation builder for the [PCameraDeviceViewDelegate] protocol.
type CameraDeviceViewDelegate struct {
	_CameraDeviceViewDidEncounterError  func(cameraDeviceView CameraDeviceView, error foundation.Error)
	_CameraDeviceViewSelectionDidChange func(cameraDeviceView CameraDeviceView)
}

func (di *CameraDeviceViewDelegate) HasCameraDeviceViewDidEncounterError() bool {
	return di._CameraDeviceViewDidEncounterError != nil
}

// Invoked when the camera encounters an error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceviewdelegate/1505239-cameradeviceview?language=objc
func (di *CameraDeviceViewDelegate) SetCameraDeviceViewDidEncounterError(f func(cameraDeviceView CameraDeviceView, error foundation.Error)) {
	di._CameraDeviceViewDidEncounterError = f
}

// Invoked when the camera encounters an error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceviewdelegate/1505239-cameradeviceview?language=objc
func (di *CameraDeviceViewDelegate) CameraDeviceViewDidEncounterError(cameraDeviceView CameraDeviceView, error foundation.Error) {
	di._CameraDeviceViewDidEncounterError(cameraDeviceView, error)
}
func (di *CameraDeviceViewDelegate) HasCameraDeviceViewSelectionDidChange() bool {
	return di._CameraDeviceViewSelectionDidChange != nil
}

// Invoked when the selection changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceviewdelegate/1505308-cameradeviceviewselectiondidchan?language=objc
func (di *CameraDeviceViewDelegate) SetCameraDeviceViewSelectionDidChange(f func(cameraDeviceView CameraDeviceView)) {
	di._CameraDeviceViewSelectionDidChange = f
}

// Invoked when the selection changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceviewdelegate/1505308-cameradeviceviewselectiondidchan?language=objc
func (di *CameraDeviceViewDelegate) CameraDeviceViewSelectionDidChange(cameraDeviceView CameraDeviceView) {
	di._CameraDeviceViewSelectionDidChange(cameraDeviceView)
}

// A concrete type wrapper for the [PCameraDeviceViewDelegate] protocol.
type CameraDeviceViewDelegateWrapper struct {
	objc.Object
}

func (c_ CameraDeviceViewDelegateWrapper) HasCameraDeviceViewDidEncounterError() bool {
	return c_.RespondsToSelector(objc.Sel("cameraDeviceView:didEncounterError:"))
}

// Invoked when the camera encounters an error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceviewdelegate/1505239-cameradeviceview?language=objc
func (c_ CameraDeviceViewDelegateWrapper) CameraDeviceViewDidEncounterError(cameraDeviceView ICameraDeviceView, error foundation.IError) {
	objc.Call[objc.Void](c_, objc.Sel("cameraDeviceView:didEncounterError:"), objc.Ptr(cameraDeviceView), objc.Ptr(error))
}

func (c_ CameraDeviceViewDelegateWrapper) HasCameraDeviceViewSelectionDidChange() bool {
	return c_.RespondsToSelector(objc.Sel("cameraDeviceViewSelectionDidChange:"))
}

// Invoked when the selection changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceviewdelegate/1505308-cameradeviceviewselectiondidchan?language=objc
func (c_ CameraDeviceViewDelegateWrapper) CameraDeviceViewSelectionDidChange(cameraDeviceView ICameraDeviceView) {
	objc.Call[objc.Void](c_, objc.Sel("cameraDeviceViewSelectionDidChange:"), objc.Ptr(cameraDeviceView))
}
