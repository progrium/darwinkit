// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DetectRectanglesRequest] class.
var DetectRectanglesRequestClass = _DetectRectanglesRequestClass{objc.GetClass("VNDetectRectanglesRequest")}

type _DetectRectanglesRequestClass struct {
	objc.Class
}

// An interface definition for the [DetectRectanglesRequest] class.
type IDetectRectanglesRequest interface {
	IImageBasedRequest
	MinimumAspectRatio() AspectRatio
	SetMinimumAspectRatio(value AspectRatio)
	MinimumSize() float64
	SetMinimumSize(value float64)
	MinimumConfidence() Confidence
	SetMinimumConfidence(value Confidence)
	MaximumObservations() uint
	SetMaximumObservations(value uint)
	MaximumAspectRatio() AspectRatio
	SetMaximumAspectRatio(value AspectRatio)
	QuadratureTolerance() Degrees
	SetQuadratureTolerance(value Degrees)
}

// An image analysis request that finds projected rectangular regions in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectrectanglesrequest?language=objc
type DetectRectanglesRequest struct {
	ImageBasedRequest
}

func DetectRectanglesRequestFrom(ptr unsafe.Pointer) DetectRectanglesRequest {
	return DetectRectanglesRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

func (dc _DetectRectanglesRequestClass) Alloc() DetectRectanglesRequest {
	rv := objc.Call[DetectRectanglesRequest](dc, objc.Sel("alloc"))
	return rv
}

func DetectRectanglesRequest_Alloc() DetectRectanglesRequest {
	return DetectRectanglesRequestClass.Alloc()
}

func (dc _DetectRectanglesRequestClass) New() DetectRectanglesRequest {
	rv := objc.Call[DetectRectanglesRequest](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDetectRectanglesRequest() DetectRectanglesRequest {
	return DetectRectanglesRequestClass.New()
}

func (d_ DetectRectanglesRequest) Init() DetectRectanglesRequest {
	rv := objc.Call[DetectRectanglesRequest](d_, objc.Sel("init"))
	return rv
}

func (d_ DetectRectanglesRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) DetectRectanglesRequest {
	rv := objc.Call[DetectRectanglesRequest](d_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewDetectRectanglesRequestWithCompletionHandler(completionHandler RequestCompletionHandler) DetectRectanglesRequest {
	instance := DetectRectanglesRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}

// A float specifying the minimum aspect ratio of the rectangle to detect, defined as the shorter dimension over the longer dimension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectrectanglesrequest/2875378-minimumaspectratio?language=objc
func (d_ DetectRectanglesRequest) MinimumAspectRatio() AspectRatio {
	rv := objc.Call[AspectRatio](d_, objc.Sel("minimumAspectRatio"))
	return rv
}

// A float specifying the minimum aspect ratio of the rectangle to detect, defined as the shorter dimension over the longer dimension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectrectanglesrequest/2875378-minimumaspectratio?language=objc
func (d_ DetectRectanglesRequest) SetMinimumAspectRatio(value AspectRatio) {
	objc.Call[objc.Void](d_, objc.Sel("setMinimumAspectRatio:"), value)
}

// The minimum size of a rectangle to detect, as a proportion of the smallest dimension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectrectanglesrequest/2875374-minimumsize?language=objc
func (d_ DetectRectanglesRequest) MinimumSize() float64 {
	rv := objc.Call[float64](d_, objc.Sel("minimumSize"))
	return rv
}

// The minimum size of a rectangle to detect, as a proportion of the smallest dimension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectrectanglesrequest/2875374-minimumsize?language=objc
func (d_ DetectRectanglesRequest) SetMinimumSize(value float64) {
	objc.Call[objc.Void](d_, objc.Sel("setMinimumSize:"), value)
}

// A value specifying the minimum acceptable confidence level. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectrectanglesrequest/2875375-minimumconfidence?language=objc
func (d_ DetectRectanglesRequest) MinimumConfidence() Confidence {
	rv := objc.Call[Confidence](d_, objc.Sel("minimumConfidence"))
	return rv
}

// A value specifying the minimum acceptable confidence level. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectrectanglesrequest/2875375-minimumconfidence?language=objc
func (d_ DetectRectanglesRequest) SetMinimumConfidence(value Confidence) {
	objc.Call[objc.Void](d_, objc.Sel("setMinimumConfidence:"), value)
}

// An integer specifying the maximum number of rectangles Vision returns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectrectanglesrequest/2875373-maximumobservations?language=objc
func (d_ DetectRectanglesRequest) MaximumObservations() uint {
	rv := objc.Call[uint](d_, objc.Sel("maximumObservations"))
	return rv
}

// An integer specifying the maximum number of rectangles Vision returns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectrectanglesrequest/2875373-maximumobservations?language=objc
func (d_ DetectRectanglesRequest) SetMaximumObservations(value uint) {
	objc.Call[objc.Void](d_, objc.Sel("setMaximumObservations:"), value)
}

// A float specifying the maximum aspect ratio of the rectangle to detect, defined as the shorter dimension over the longer dimension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectrectanglesrequest/2875376-maximumaspectratio?language=objc
func (d_ DetectRectanglesRequest) MaximumAspectRatio() AspectRatio {
	rv := objc.Call[AspectRatio](d_, objc.Sel("maximumAspectRatio"))
	return rv
}

// A float specifying the maximum aspect ratio of the rectangle to detect, defined as the shorter dimension over the longer dimension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectrectanglesrequest/2875376-maximumaspectratio?language=objc
func (d_ DetectRectanglesRequest) SetMaximumAspectRatio(value AspectRatio) {
	objc.Call[objc.Void](d_, objc.Sel("setMaximumAspectRatio:"), value)
}

// A float specifying the number of degrees a rectangle corner angle can deviate from 90°. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectrectanglesrequest/2875379-quadraturetolerance?language=objc
func (d_ DetectRectanglesRequest) QuadratureTolerance() Degrees {
	rv := objc.Call[Degrees](d_, objc.Sel("quadratureTolerance"))
	return rv
}

// A float specifying the number of degrees a rectangle corner angle can deviate from 90°. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectrectanglesrequest/2875379-quadraturetolerance?language=objc
func (d_ DetectRectanglesRequest) SetQuadratureTolerance(value Degrees) {
	objc.Call[objc.Void](d_, objc.Sel("setQuadratureTolerance:"), value)
}
