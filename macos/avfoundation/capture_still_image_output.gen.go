// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptureStillImageOutput] class.
var CaptureStillImageOutputClass = _CaptureStillImageOutputClass{objc.GetClass("AVCaptureStillImageOutput")}

type _CaptureStillImageOutputClass struct {
	objc.Class
}

// An interface definition for the [CaptureStillImageOutput] class.
type ICaptureStillImageOutput interface {
	ICaptureOutput
}

// A capture output for capturing still photos. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput?language=objc
type CaptureStillImageOutput struct {
	CaptureOutput
}

func CaptureStillImageOutputFrom(ptr unsafe.Pointer) CaptureStillImageOutput {
	return CaptureStillImageOutput{
		CaptureOutput: CaptureOutputFrom(ptr),
	}
}

func (cc _CaptureStillImageOutputClass) Alloc() CaptureStillImageOutput {
	rv := objc.Call[CaptureStillImageOutput](cc, objc.Sel("alloc"))
	return rv
}

func CaptureStillImageOutput_Alloc() CaptureStillImageOutput {
	return CaptureStillImageOutputClass.Alloc()
}

func (cc _CaptureStillImageOutputClass) New() CaptureStillImageOutput {
	rv := objc.Call[CaptureStillImageOutput](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptureStillImageOutput() CaptureStillImageOutput {
	return CaptureStillImageOutputClass.New()
}

func (c_ CaptureStillImageOutput) Init() CaptureStillImageOutput {
	rv := objc.Call[CaptureStillImageOutput](c_, objc.Sel("init"))
	return rv
}
