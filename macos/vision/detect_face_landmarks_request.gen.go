// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DetectFaceLandmarksRequest] class.
var DetectFaceLandmarksRequestClass = _DetectFaceLandmarksRequestClass{objc.GetClass("VNDetectFaceLandmarksRequest")}

type _DetectFaceLandmarksRequestClass struct {
	objc.Class
}

// An interface definition for the [DetectFaceLandmarksRequest] class.
type IDetectFaceLandmarksRequest interface {
	IImageBasedRequest
	Constellation() RequestFaceLandmarksConstellation
	SetConstellation(value RequestFaceLandmarksConstellation)
}

// An image analysis request that finds facial features like eyes and mouth in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectfacelandmarksrequest?language=objc
type DetectFaceLandmarksRequest struct {
	ImageBasedRequest
}

func DetectFaceLandmarksRequestFrom(ptr unsafe.Pointer) DetectFaceLandmarksRequest {
	return DetectFaceLandmarksRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

func (dc _DetectFaceLandmarksRequestClass) Alloc() DetectFaceLandmarksRequest {
	rv := objc.Call[DetectFaceLandmarksRequest](dc, objc.Sel("alloc"))
	return rv
}

func DetectFaceLandmarksRequest_Alloc() DetectFaceLandmarksRequest {
	return DetectFaceLandmarksRequestClass.Alloc()
}

func (dc _DetectFaceLandmarksRequestClass) New() DetectFaceLandmarksRequest {
	rv := objc.Call[DetectFaceLandmarksRequest](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDetectFaceLandmarksRequest() DetectFaceLandmarksRequest {
	return DetectFaceLandmarksRequestClass.New()
}

func (d_ DetectFaceLandmarksRequest) Init() DetectFaceLandmarksRequest {
	rv := objc.Call[DetectFaceLandmarksRequest](d_, objc.Sel("init"))
	return rv
}

func (d_ DetectFaceLandmarksRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) DetectFaceLandmarksRequest {
	rv := objc.Call[DetectFaceLandmarksRequest](d_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewDetectFaceLandmarksRequestWithCompletionHandler(completionHandler RequestCompletionHandler) DetectFaceLandmarksRequest {
	instance := DetectFaceLandmarksRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}

// Returns a Boolean value that indicates whether a revision supports a constellation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectfacelandmarksrequest/3143665-revision?language=objc
func (dc _DetectFaceLandmarksRequestClass) RevisionSupportsConstellation(requestRevision uint, constellation RequestFaceLandmarksConstellation) bool {
	rv := objc.Call[bool](dc, objc.Sel("revision:supportsConstellation:"), requestRevision, constellation)
	return rv
}

// Returns a Boolean value that indicates whether a revision supports a constellation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectfacelandmarksrequest/3143665-revision?language=objc
func DetectFaceLandmarksRequest_RevisionSupportsConstellation(requestRevision uint, constellation RequestFaceLandmarksConstellation) bool {
	return DetectFaceLandmarksRequestClass.RevisionSupportsConstellation(requestRevision, constellation)
}

// A variable that describes how a face landmarks request orders or enumerates the resulting features. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectfacelandmarksrequest/3143664-constellation?language=objc
func (d_ DetectFaceLandmarksRequest) Constellation() RequestFaceLandmarksConstellation {
	rv := objc.Call[RequestFaceLandmarksConstellation](d_, objc.Sel("constellation"))
	return rv
}

// A variable that describes how a face landmarks request orders or enumerates the resulting features. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectfacelandmarksrequest/3143664-constellation?language=objc
func (d_ DetectFaceLandmarksRequest) SetConstellation(value RequestFaceLandmarksConstellation) {
	objc.Call[objc.Void](d_, objc.Sel("setConstellation:"), value)
}
