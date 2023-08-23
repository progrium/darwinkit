// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [GeneratePersonSegmentationRequest] class.
var GeneratePersonSegmentationRequestClass = _GeneratePersonSegmentationRequestClass{objc.GetClass("VNGeneratePersonSegmentationRequest")}

type _GeneratePersonSegmentationRequestClass struct {
	objc.Class
}

// An interface definition for the [GeneratePersonSegmentationRequest] class.
type IGeneratePersonSegmentationRequest interface {
	IStatefulRequest
	QualityLevel() GeneratePersonSegmentationRequestQualityLevel
	SetQualityLevel(value GeneratePersonSegmentationRequestQualityLevel)
	OutputPixelFormat() uint
	SetOutputPixelFormat(value uint)
}

// An object that produces a matte image for a person it finds in the input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeneratepersonsegmentationrequest?language=objc
type GeneratePersonSegmentationRequest struct {
	StatefulRequest
}

func GeneratePersonSegmentationRequestFrom(ptr unsafe.Pointer) GeneratePersonSegmentationRequest {
	return GeneratePersonSegmentationRequest{
		StatefulRequest: StatefulRequestFrom(ptr),
	}
}

func (gc _GeneratePersonSegmentationRequestClass) New() GeneratePersonSegmentationRequest {
	rv := objc.Call[GeneratePersonSegmentationRequest](gc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewGeneratePersonSegmentationRequest() GeneratePersonSegmentationRequest {
	return GeneratePersonSegmentationRequestClass.New()
}

func (g_ GeneratePersonSegmentationRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) GeneratePersonSegmentationRequest {
	rv := objc.Call[GeneratePersonSegmentationRequest](g_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a generate person segmentation request with a completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeneratepersonsegmentationrequest/3783572-initwithcompletionhandler?language=objc
func NewGeneratePersonSegmentationRequestWithCompletionHandler(completionHandler RequestCompletionHandler) GeneratePersonSegmentationRequest {
	instance := GeneratePersonSegmentationRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}

func (g_ GeneratePersonSegmentationRequest) Init() GeneratePersonSegmentationRequest {
	rv := objc.Call[GeneratePersonSegmentationRequest](g_, objc.Sel("init"))
	return rv
}

func (gc _GeneratePersonSegmentationRequestClass) Alloc() GeneratePersonSegmentationRequest {
	rv := objc.Call[GeneratePersonSegmentationRequest](gc, objc.Sel("alloc"))
	return rv
}

func GeneratePersonSegmentationRequest_Alloc() GeneratePersonSegmentationRequest {
	return GeneratePersonSegmentationRequestClass.Alloc()
}

func (g_ GeneratePersonSegmentationRequest) InitWithFrameAnalysisSpacingCompletionHandler(frameAnalysisSpacing coremedia.Time, completionHandler RequestCompletionHandler) GeneratePersonSegmentationRequest {
	rv := objc.Call[GeneratePersonSegmentationRequest](g_, objc.Sel("initWithFrameAnalysisSpacing:completionHandler:"), frameAnalysisSpacing, completionHandler)
	return rv
}

// Initializes a video-based request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnstatefulrequest/3564828-initwithframeanalysisspacing?language=objc
func NewGeneratePersonSegmentationRequestWithFrameAnalysisSpacingCompletionHandler(frameAnalysisSpacing coremedia.Time, completionHandler RequestCompletionHandler) GeneratePersonSegmentationRequest {
	instance := GeneratePersonSegmentationRequestClass.Alloc().InitWithFrameAnalysisSpacingCompletionHandler(frameAnalysisSpacing, completionHandler)
	instance.Autorelease()
	return instance
}

// A value that indicates how the request balances accuracy and performance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeneratepersonsegmentationrequest/3750989-qualitylevel?language=objc
func (g_ GeneratePersonSegmentationRequest) QualityLevel() GeneratePersonSegmentationRequestQualityLevel {
	rv := objc.Call[GeneratePersonSegmentationRequestQualityLevel](g_, objc.Sel("qualityLevel"))
	return rv
}

// A value that indicates how the request balances accuracy and performance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeneratepersonsegmentationrequest/3750989-qualitylevel?language=objc
func (g_ GeneratePersonSegmentationRequest) SetQualityLevel(value GeneratePersonSegmentationRequestQualityLevel) {
	objc.Call[objc.Void](g_, objc.Sel("setQualityLevel:"), value)
}

// The pixel format of the output image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeneratepersonsegmentationrequest/3750988-outputpixelformat?language=objc
func (g_ GeneratePersonSegmentationRequest) OutputPixelFormat() uint {
	rv := objc.Call[uint](g_, objc.Sel("outputPixelFormat"))
	return rv
}

// The pixel format of the output image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeneratepersonsegmentationrequest/3750988-outputpixelformat?language=objc
func (g_ GeneratePersonSegmentationRequest) SetOutputPixelFormat(value uint) {
	objc.Call[objc.Void](g_, objc.Sel("setOutputPixelFormat:"), value)
}
