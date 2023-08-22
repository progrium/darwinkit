// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptureDeviceInput] class.
var CaptureDeviceInputClass = _CaptureDeviceInputClass{objc.GetClass("AVCaptureDeviceInput")}

type _CaptureDeviceInputClass struct {
	objc.Class
}

// An interface definition for the [CaptureDeviceInput] class.
type ICaptureDeviceInput interface {
	ICaptureInput
	Device() CaptureDevice
}

// An object that provides media input from a capture device to a capture session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput?language=objc
type CaptureDeviceInput struct {
	CaptureInput
}

func CaptureDeviceInputFrom(ptr unsafe.Pointer) CaptureDeviceInput {
	return CaptureDeviceInput{
		CaptureInput: CaptureInputFrom(ptr),
	}
}

func (c_ CaptureDeviceInput) InitWithDeviceError(device ICaptureDevice, outError foundation.IError) CaptureDeviceInput {
	rv := objc.Call[CaptureDeviceInput](c_, objc.Sel("initWithDevice:error:"), objc.Ptr(device), objc.Ptr(outError))
	return rv
}

// Creates an input for the specified capture device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/1387609-initwithdevice?language=objc
func NewCaptureDeviceInputWithDeviceError(device ICaptureDevice, outError foundation.IError) CaptureDeviceInput {
	instance := CaptureDeviceInputClass.Alloc().InitWithDeviceError(device, outError)
	instance.Autorelease()
	return instance
}

func (cc _CaptureDeviceInputClass) DeviceInputWithDeviceError(device ICaptureDevice, outError foundation.IError) CaptureDeviceInput {
	rv := objc.Call[CaptureDeviceInput](cc, objc.Sel("deviceInputWithDevice:error:"), objc.Ptr(device), objc.Ptr(outError))
	return rv
}

// Returns a new input for the specified capture device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/1450880-deviceinputwithdevice?language=objc
func CaptureDeviceInput_DeviceInputWithDeviceError(device ICaptureDevice, outError foundation.IError) CaptureDeviceInput {
	return CaptureDeviceInputClass.DeviceInputWithDeviceError(device, outError)
}

func (cc _CaptureDeviceInputClass) Alloc() CaptureDeviceInput {
	rv := objc.Call[CaptureDeviceInput](cc, objc.Sel("alloc"))
	return rv
}

func CaptureDeviceInput_Alloc() CaptureDeviceInput {
	return CaptureDeviceInputClass.Alloc()
}

func (cc _CaptureDeviceInputClass) New() CaptureDeviceInput {
	rv := objc.Call[CaptureDeviceInput](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptureDeviceInput() CaptureDeviceInput {
	return CaptureDeviceInputClass.New()
}

func (c_ CaptureDeviceInput) Init() CaptureDeviceInput {
	rv := objc.Call[CaptureDeviceInput](c_, objc.Sel("init"))
	return rv
}

// A capture device associated with this input. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/1387915-device?language=objc
func (c_ CaptureDeviceInput) Device() CaptureDevice {
	rv := objc.Call[CaptureDevice](c_, objc.Sel("device"))
	return rv
}
