// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptureDeviceDiscoverySession] class.
var CaptureDeviceDiscoverySessionClass = _CaptureDeviceDiscoverySessionClass{objc.GetClass("AVCaptureDeviceDiscoverySession")}

type _CaptureDeviceDiscoverySessionClass struct {
	objc.Class
}

// An interface definition for the [CaptureDeviceDiscoverySession] class.
type ICaptureDeviceDiscoverySession interface {
	objc.IObject
	Devices() []CaptureDevice
}

// An object that finds capture devices that match specific search criteria. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevicediscoverysession?language=objc
type CaptureDeviceDiscoverySession struct {
	objc.Object
}

func CaptureDeviceDiscoverySessionFrom(ptr unsafe.Pointer) CaptureDeviceDiscoverySession {
	return CaptureDeviceDiscoverySession{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CaptureDeviceDiscoverySessionClass) DiscoverySessionWithDeviceTypesMediaTypePosition(deviceTypes []CaptureDeviceType, mediaType MediaType, position CaptureDevicePosition) CaptureDeviceDiscoverySession {
	rv := objc.Call[CaptureDeviceDiscoverySession](cc, objc.Sel("discoverySessionWithDeviceTypes:mediaType:position:"), deviceTypes, mediaType, position)
	return rv
}

// Creates a discovery session that finds devices that match the specified criteria. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevicediscoverysession/2361539-discoverysessionwithdevicetypes?language=objc
func CaptureDeviceDiscoverySession_DiscoverySessionWithDeviceTypesMediaTypePosition(deviceTypes []CaptureDeviceType, mediaType MediaType, position CaptureDevicePosition) CaptureDeviceDiscoverySession {
	return CaptureDeviceDiscoverySessionClass.DiscoverySessionWithDeviceTypesMediaTypePosition(deviceTypes, mediaType, position)
}

func (cc _CaptureDeviceDiscoverySessionClass) Alloc() CaptureDeviceDiscoverySession {
	rv := objc.Call[CaptureDeviceDiscoverySession](cc, objc.Sel("alloc"))
	return rv
}

func CaptureDeviceDiscoverySession_Alloc() CaptureDeviceDiscoverySession {
	return CaptureDeviceDiscoverySessionClass.Alloc()
}

func (cc _CaptureDeviceDiscoverySessionClass) New() CaptureDeviceDiscoverySession {
	rv := objc.Call[CaptureDeviceDiscoverySession](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptureDeviceDiscoverySession() CaptureDeviceDiscoverySession {
	return CaptureDeviceDiscoverySessionClass.New()
}

func (c_ CaptureDeviceDiscoverySession) Init() CaptureDeviceDiscoverySession {
	rv := objc.Call[CaptureDeviceDiscoverySession](c_, objc.Sel("init"))
	return rv
}

// A list of devices that match the search criteria of the discovery session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevicediscoverysession/2361002-devices?language=objc
func (c_ CaptureDeviceDiscoverySession) Devices() []CaptureDevice {
	rv := objc.Call[[]CaptureDevice](c_, objc.Sel("devices"))
	return rv
}
