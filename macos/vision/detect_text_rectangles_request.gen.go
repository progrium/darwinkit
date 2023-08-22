// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DetectTextRectanglesRequest] class.
var DetectTextRectanglesRequestClass = _DetectTextRectanglesRequestClass{objc.GetClass("VNDetectTextRectanglesRequest")}

type _DetectTextRectanglesRequestClass struct {
	objc.Class
}

// An interface definition for the [DetectTextRectanglesRequest] class.
type IDetectTextRectanglesRequest interface {
	IImageBasedRequest
	ReportCharacterBoxes() bool
	SetReportCharacterBoxes(value bool)
}

// An image analysis request that finds regions of visible text in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttextrectanglesrequest?language=objc
type DetectTextRectanglesRequest struct {
	ImageBasedRequest
}

func DetectTextRectanglesRequestFrom(ptr unsafe.Pointer) DetectTextRectanglesRequest {
	return DetectTextRectanglesRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

func (dc _DetectTextRectanglesRequestClass) Alloc() DetectTextRectanglesRequest {
	rv := objc.Call[DetectTextRectanglesRequest](dc, objc.Sel("alloc"))
	return rv
}

func DetectTextRectanglesRequest_Alloc() DetectTextRectanglesRequest {
	return DetectTextRectanglesRequestClass.Alloc()
}

func (dc _DetectTextRectanglesRequestClass) New() DetectTextRectanglesRequest {
	rv := objc.Call[DetectTextRectanglesRequest](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDetectTextRectanglesRequest() DetectTextRectanglesRequest {
	return DetectTextRectanglesRequestClass.New()
}

func (d_ DetectTextRectanglesRequest) Init() DetectTextRectanglesRequest {
	rv := objc.Call[DetectTextRectanglesRequest](d_, objc.Sel("init"))
	return rv
}

func (d_ DetectTextRectanglesRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) DetectTextRectanglesRequest {
	rv := objc.Call[DetectTextRectanglesRequest](d_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewDetectTextRectanglesRequestWithCompletionHandler(completionHandler RequestCompletionHandler) DetectTextRectanglesRequest {
	instance := DetectTextRectanglesRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}

// A Boolean value that indicates whether the request detects character bounding boxes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttextrectanglesrequest/2875410-reportcharacterboxes?language=objc
func (d_ DetectTextRectanglesRequest) ReportCharacterBoxes() bool {
	rv := objc.Call[bool](d_, objc.Sel("reportCharacterBoxes"))
	return rv
}

// A Boolean value that indicates whether the request detects character bounding boxes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttextrectanglesrequest/2875410-reportcharacterboxes?language=objc
func (d_ DetectTextRectanglesRequest) SetReportCharacterBoxes(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setReportCharacterBoxes:"), value)
}
