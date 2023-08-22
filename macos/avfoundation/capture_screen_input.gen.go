// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptureScreenInput] class.
var CaptureScreenInputClass = _CaptureScreenInputClass{objc.GetClass("AVCaptureScreenInput")}

type _CaptureScreenInputClass struct {
	objc.Class
}

// An interface definition for the [CaptureScreenInput] class.
type ICaptureScreenInput interface {
	ICaptureInput
	CapturesCursor() bool
	SetCapturesCursor(value bool)
	ScaleFactor() float64
	SetScaleFactor(value float64)
	CapturesMouseClicks() bool
	SetCapturesMouseClicks(value bool)
	MinFrameDuration() coremedia.Time
	SetMinFrameDuration(value coremedia.Time)
	CropRect() coregraphics.Rect
	SetCropRect(value coregraphics.Rect)
}

// A capture input for recording from a screen in macOS. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturescreeninput?language=objc
type CaptureScreenInput struct {
	CaptureInput
}

func CaptureScreenInputFrom(ptr unsafe.Pointer) CaptureScreenInput {
	return CaptureScreenInput{
		CaptureInput: CaptureInputFrom(ptr),
	}
}

func (cc _CaptureScreenInputClass) New() CaptureScreenInput {
	rv := objc.Call[CaptureScreenInput](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptureScreenInput() CaptureScreenInput {
	return CaptureScreenInputClass.New()
}

func (c_ CaptureScreenInput) InitWithDisplayID(displayID coregraphics.DirectDisplayID) CaptureScreenInput {
	rv := objc.Call[CaptureScreenInput](c_, objc.Sel("initWithDisplayID:"), displayID)
	return rv
}

// Initializes a capture screen input that provides media data from the specified display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturescreeninput/1386383-initwithdisplayid?language=objc
func NewCaptureScreenInputWithDisplayID(displayID coregraphics.DirectDisplayID) CaptureScreenInput {
	instance := CaptureScreenInputClass.Alloc().InitWithDisplayID(displayID)
	instance.Autorelease()
	return instance
}

func (c_ CaptureScreenInput) Init() CaptureScreenInput {
	rv := objc.Call[CaptureScreenInput](c_, objc.Sel("init"))
	return rv
}

func (cc _CaptureScreenInputClass) Alloc() CaptureScreenInput {
	rv := objc.Call[CaptureScreenInput](cc, objc.Sel("alloc"))
	return rv
}

func CaptureScreenInput_Alloc() CaptureScreenInput {
	return CaptureScreenInputClass.Alloc()
}

// A Boolean value that specifies whether the mouse cursor appears in the captured output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturescreeninput/1385601-capturescursor?language=objc
func (c_ CaptureScreenInput) CapturesCursor() bool {
	rv := objc.Call[bool](c_, objc.Sel("capturesCursor"))
	return rv
}

// A Boolean value that specifies whether the mouse cursor appears in the captured output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturescreeninput/1385601-capturescursor?language=objc
func (c_ CaptureScreenInput) SetCapturesCursor(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setCapturesCursor:"), value)
}

// Indicates the factor by which video buffers captured from the screen are to be scaled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturescreeninput/1390311-scalefactor?language=objc
func (c_ CaptureScreenInput) ScaleFactor() float64 {
	rv := objc.Call[float64](c_, objc.Sel("scaleFactor"))
	return rv
}

// Indicates the factor by which video buffers captured from the screen are to be scaled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturescreeninput/1390311-scalefactor?language=objc
func (c_ CaptureScreenInput) SetScaleFactor(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setScaleFactor:"), value)
}

// A Boolean value that specifies whether mouse clicks appear highlighted in the captured output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturescreeninput/1385722-capturesmouseclicks?language=objc
func (c_ CaptureScreenInput) CapturesMouseClicks() bool {
	rv := objc.Call[bool](c_, objc.Sel("capturesMouseClicks"))
	return rv
}

// A Boolean value that specifies whether mouse clicks appear highlighted in the captured output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturescreeninput/1385722-capturesmouseclicks?language=objc
func (c_ CaptureScreenInput) SetCapturesMouseClicks(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setCapturesMouseClicks:"), value)
}

// The screen input's minimum frame duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturescreeninput/1387216-minframeduration?language=objc
func (c_ CaptureScreenInput) MinFrameDuration() coremedia.Time {
	rv := objc.Call[coremedia.Time](c_, objc.Sel("minFrameDuration"))
	return rv
}

// The screen input's minimum frame duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturescreeninput/1387216-minframeduration?language=objc
func (c_ CaptureScreenInput) SetMinFrameDuration(value coremedia.Time) {
	objc.Call[objc.Void](c_, objc.Sel("setMinFrameDuration:"), value)
}

// Indicates the bounding rectangle of the screen area to be captured, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturescreeninput/1390518-croprect?language=objc
func (c_ CaptureScreenInput) CropRect() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](c_, objc.Sel("cropRect"))
	return rv
}

// Indicates the bounding rectangle of the screen area to be captured, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturescreeninput/1390518-croprect?language=objc
func (c_ CaptureScreenInput) SetCropRect(value coregraphics.Rect) {
	objc.Call[objc.Void](c_, objc.Sel("setCropRect:"), value)
}
