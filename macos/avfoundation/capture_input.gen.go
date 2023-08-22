// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptureInput] class.
var CaptureInputClass = _CaptureInputClass{objc.GetClass("AVCaptureInput")}

type _CaptureInputClass struct {
	objc.Class
}

// An interface definition for the [CaptureInput] class.
type ICaptureInput interface {
	objc.IObject
	Ports() []CaptureInputPort
}

// An abstract superclass for objects that provide input data to a capture session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureinput?language=objc
type CaptureInput struct {
	objc.Object
}

func CaptureInputFrom(ptr unsafe.Pointer) CaptureInput {
	return CaptureInput{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CaptureInputClass) Alloc() CaptureInput {
	rv := objc.Call[CaptureInput](cc, objc.Sel("alloc"))
	return rv
}

func CaptureInput_Alloc() CaptureInput {
	return CaptureInputClass.Alloc()
}

func (cc _CaptureInputClass) New() CaptureInput {
	rv := objc.Call[CaptureInput](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptureInput() CaptureInput {
	return CaptureInputClass.New()
}

func (c_ CaptureInput) Init() CaptureInput {
	rv := objc.Call[CaptureInput](c_, objc.Sel("init"))
	return rv
}

// The ports available on a capture input. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureinput/1386635-ports?language=objc
func (c_ CaptureInput) Ports() []CaptureInputPort {
	rv := objc.Call[[]CaptureInputPort](c_, objc.Sel("ports"))
	return rv
}
