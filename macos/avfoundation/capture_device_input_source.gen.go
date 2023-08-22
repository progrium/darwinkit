// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptureDeviceInputSource] class.
var CaptureDeviceInputSourceClass = _CaptureDeviceInputSourceClass{objc.GetClass("AVCaptureDeviceInputSource")}

type _CaptureDeviceInputSourceClass struct {
	objc.Class
}

// An interface definition for the [CaptureDeviceInputSource] class.
type ICaptureDeviceInputSource interface {
	objc.IObject
	InputSourceID() string
	LocalizedName() string
}

// A distinct input source on a capture device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinputsource?language=objc
type CaptureDeviceInputSource struct {
	objc.Object
}

func CaptureDeviceInputSourceFrom(ptr unsafe.Pointer) CaptureDeviceInputSource {
	return CaptureDeviceInputSource{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CaptureDeviceInputSourceClass) Alloc() CaptureDeviceInputSource {
	rv := objc.Call[CaptureDeviceInputSource](cc, objc.Sel("alloc"))
	return rv
}

func CaptureDeviceInputSource_Alloc() CaptureDeviceInputSource {
	return CaptureDeviceInputSourceClass.Alloc()
}

func (cc _CaptureDeviceInputSourceClass) New() CaptureDeviceInputSource {
	rv := objc.Call[CaptureDeviceInputSource](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptureDeviceInputSource() CaptureDeviceInputSource {
	return CaptureDeviceInputSourceClass.New()
}

func (c_ CaptureDeviceInputSource) Init() CaptureDeviceInputSource {
	rv := objc.Call[CaptureDeviceInputSource](c_, objc.Sel("init"))
	return rv
}

// An identifier for an input source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinputsource/1387788-inputsourceid?language=objc
func (c_ CaptureDeviceInputSource) InputSourceID() string {
	rv := objc.Call[string](c_, objc.Sel("inputSourceID"))
	return rv
}

// A localized, human-readable name for the input source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinputsource/1390422-localizedname?language=objc
func (c_ CaptureDeviceInputSource) LocalizedName() string {
	rv := objc.Call[string](c_, objc.Sel("localizedName"))
	return rv
}
