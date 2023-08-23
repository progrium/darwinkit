// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DetectContoursRequest] class.
var DetectContoursRequestClass = _DetectContoursRequestClass{objc.GetClass("VNDetectContoursRequest")}

type _DetectContoursRequestClass struct {
	objc.Class
}

// An interface definition for the [DetectContoursRequest] class.
type IDetectContoursRequest interface {
	IImageBasedRequest
	ContrastPivot() foundation.Number
	SetContrastPivot(value foundation.INumber)
	DetectsDarkOnLight() bool
	SetDetectsDarkOnLight(value bool)
	ContrastAdjustment() float64
	SetContrastAdjustment(value float64)
	MaximumImageDimension() uint
	SetMaximumImageDimension(value uint)
}

// A request that detects the contours of the edges of an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectcontoursrequest?language=objc
type DetectContoursRequest struct {
	ImageBasedRequest
}

func DetectContoursRequestFrom(ptr unsafe.Pointer) DetectContoursRequest {
	return DetectContoursRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

func (dc _DetectContoursRequestClass) Alloc() DetectContoursRequest {
	rv := objc.Call[DetectContoursRequest](dc, objc.Sel("alloc"))
	return rv
}

func DetectContoursRequest_Alloc() DetectContoursRequest {
	return DetectContoursRequestClass.Alloc()
}

func (dc _DetectContoursRequestClass) New() DetectContoursRequest {
	rv := objc.Call[DetectContoursRequest](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDetectContoursRequest() DetectContoursRequest {
	return DetectContoursRequestClass.New()
}

func (d_ DetectContoursRequest) Init() DetectContoursRequest {
	rv := objc.Call[DetectContoursRequest](d_, objc.Sel("init"))
	return rv
}

func (d_ DetectContoursRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) DetectContoursRequest {
	rv := objc.Call[DetectContoursRequest](d_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewDetectContoursRequestWithCompletionHandler(completionHandler RequestCompletionHandler) DetectContoursRequest {
	instance := DetectContoursRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}

// The pixel value to use as a pivot for the contrast. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectcontoursrequest/3750961-contrastpivot?language=objc
func (d_ DetectContoursRequest) ContrastPivot() foundation.Number {
	rv := objc.Call[foundation.Number](d_, objc.Sel("contrastPivot"))
	return rv
}

// The pixel value to use as a pivot for the contrast. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectcontoursrequest/3750961-contrastpivot?language=objc
func (d_ DetectContoursRequest) SetContrastPivot(value foundation.INumber) {
	objc.Call[objc.Void](d_, objc.Sel("setContrastPivot:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the request detects a dark object on a light background to aid in detection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectcontoursrequest/3675596-detectsdarkonlight?language=objc
func (d_ DetectContoursRequest) DetectsDarkOnLight() bool {
	rv := objc.Call[bool](d_, objc.Sel("detectsDarkOnLight"))
	return rv
}

// A Boolean value that indicates whether the request detects a dark object on a light background to aid in detection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectcontoursrequest/3675596-detectsdarkonlight?language=objc
func (d_ DetectContoursRequest) SetDetectsDarkOnLight(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setDetectsDarkOnLight:"), value)
}

// The amount by which to adjust the image contrast. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectcontoursrequest/3548236-contrastadjustment?language=objc
func (d_ DetectContoursRequest) ContrastAdjustment() float64 {
	rv := objc.Call[float64](d_, objc.Sel("contrastAdjustment"))
	return rv
}

// The amount by which to adjust the image contrast. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectcontoursrequest/3548236-contrastadjustment?language=objc
func (d_ DetectContoursRequest) SetContrastAdjustment(value float64) {
	objc.Call[objc.Void](d_, objc.Sel("setContrastAdjustment:"), value)
}

// The maximum image dimension to use for contour detection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectcontoursrequest/3548238-maximumimagedimension?language=objc
func (d_ DetectContoursRequest) MaximumImageDimension() uint {
	rv := objc.Call[uint](d_, objc.Sel("maximumImageDimension"))
	return rv
}

// The maximum image dimension to use for contour detection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectcontoursrequest/3548238-maximumimagedimension?language=objc
func (d_ DetectContoursRequest) SetMaximumImageDimension(value uint) {
	objc.Call[objc.Void](d_, objc.Sel("setMaximumImageDimension:"), value)
}
