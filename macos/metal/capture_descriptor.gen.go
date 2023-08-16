// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptureDescriptor] class.
var CaptureDescriptorClass = _CaptureDescriptorClass{objc.GetClass("MTLCaptureDescriptor")}

type _CaptureDescriptorClass struct {
	objc.Class
}

// An interface definition for the [CaptureDescriptor] class.
type ICaptureDescriptor interface {
	objc.IObject
	CaptureObject() objc.Object
	SetCaptureObject(value objc.IObject)
	Destination() CaptureDestination
	SetDestination(value CaptureDestination)
	OutputURL() foundation.URL
	SetOutputURL(value foundation.IURL)
}

// A configuration for a Metal capture session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturedescriptor?language=objc
type CaptureDescriptor struct {
	objc.Object
}

func CaptureDescriptorFrom(ptr unsafe.Pointer) CaptureDescriptor {
	return CaptureDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CaptureDescriptorClass) Alloc() CaptureDescriptor {
	rv := objc.Call[CaptureDescriptor](cc, objc.Sel("alloc"))
	return rv
}

func CaptureDescriptor_Alloc() CaptureDescriptor {
	return CaptureDescriptorClass.Alloc()
}

func (cc _CaptureDescriptorClass) New() CaptureDescriptor {
	rv := objc.Call[CaptureDescriptor](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptureDescriptor() CaptureDescriptor {
	return CaptureDescriptorClass.New()
}

func (c_ CaptureDescriptor) Init() CaptureDescriptor {
	rv := objc.Call[CaptureDescriptor](c_, objc.Sel("init"))
	return rv
}

// The object whose contents should be captured. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturedescriptor/3237248-captureobject?language=objc
func (c_ CaptureDescriptor) CaptureObject() objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("captureObject"))
	return rv
}

// The object whose contents should be captured. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturedescriptor/3237248-captureobject?language=objc
func (c_ CaptureDescriptor) SetCaptureObject(value objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setCaptureObject:"), value)
}

// The destination for any captured command data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturedescriptor/3237249-destination?language=objc
func (c_ CaptureDescriptor) Destination() CaptureDestination {
	rv := objc.Call[CaptureDestination](c_, objc.Sel("destination"))
	return rv
}

// The destination for any captured command data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturedescriptor/3237249-destination?language=objc
func (c_ CaptureDescriptor) SetDestination(value CaptureDestination) {
	objc.Call[objc.Void](c_, objc.Sel("setDestination:"), value)
}

// A URL for a file to write the capture data into. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturedescriptor/3237250-outputurl?language=objc
func (c_ CaptureDescriptor) OutputURL() foundation.URL {
	rv := objc.Call[foundation.URL](c_, objc.Sel("outputURL"))
	return rv
}

// A URL for a file to write the capture data into. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturedescriptor/3237250-outputurl?language=objc
func (c_ CaptureDescriptor) SetOutputURL(value foundation.IURL) {
	objc.Call[objc.Void](c_, objc.Sel("setOutputURL:"), objc.Ptr(value))
}
