// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptureInputPort] class.
var CaptureInputPortClass = _CaptureInputPortClass{objc.GetClass("AVCaptureInputPort")}

type _CaptureInputPortClass struct {
	objc.Class
}

// An interface definition for the [CaptureInputPort] class.
type ICaptureInputPort interface {
	objc.IObject
	Clock() coremedia.ClockRef
	Input() CaptureInput
	MediaType() MediaType
	IsEnabled() bool
	SetEnabled(value bool)
	FormatDescription() coremedia.FormatDescriptionRef
}

// An object that represents a stream of data that a capture input provides. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureinputport?language=objc
type CaptureInputPort struct {
	objc.Object
}

func CaptureInputPortFrom(ptr unsafe.Pointer) CaptureInputPort {
	return CaptureInputPort{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CaptureInputPortClass) Alloc() CaptureInputPort {
	rv := objc.Call[CaptureInputPort](cc, objc.Sel("alloc"))
	return rv
}

func CaptureInputPort_Alloc() CaptureInputPort {
	return CaptureInputPortClass.Alloc()
}

func (cc _CaptureInputPortClass) New() CaptureInputPort {
	rv := objc.Call[CaptureInputPort](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptureInputPort() CaptureInputPort {
	return CaptureInputPortClass.New()
}

func (c_ CaptureInputPort) Init() CaptureInputPort {
	rv := objc.Call[CaptureInputPort](c_, objc.Sel("init"))
	return rv
}

// An object that represents the capture device’s clock. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureinputport/1385908-clock?language=objc
func (c_ CaptureInputPort) Clock() coremedia.ClockRef {
	rv := objc.Call[coremedia.ClockRef](c_, objc.Sel("clock"))
	return rv
}

// The input object that owns the port. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureinputport/1387702-input?language=objc
func (c_ CaptureInputPort) Input() CaptureInput {
	rv := objc.Call[CaptureInput](c_, objc.Sel("input"))
	return rv
}

// The media type of the port. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureinputport/1387120-mediatype?language=objc
func (c_ CaptureInputPort) MediaType() MediaType {
	rv := objc.Call[MediaType](c_, objc.Sel("mediaType"))
	return rv
}

// A Boolean value that indicates whether the port is in an enabled state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureinputport/1386833-enabled?language=objc
func (c_ CaptureInputPort) IsEnabled() bool {
	rv := objc.Call[bool](c_, objc.Sel("isEnabled"))
	return rv
}

// A Boolean value that indicates whether the port is in an enabled state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureinputport/1386833-enabled?language=objc
func (c_ CaptureInputPort) SetEnabled(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setEnabled:"), value)
}

// A description of the port format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureinputport/1385890-formatdescription?language=objc
func (c_ CaptureInputPort) FormatDescription() coremedia.FormatDescriptionRef {
	rv := objc.Call[coremedia.FormatDescriptionRef](c_, objc.Sel("formatDescription"))
	return rv
}
