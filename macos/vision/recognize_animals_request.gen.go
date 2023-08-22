// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RecognizeAnimalsRequest] class.
var RecognizeAnimalsRequestClass = _RecognizeAnimalsRequestClass{objc.GetClass("VNRecognizeAnimalsRequest")}

type _RecognizeAnimalsRequestClass struct {
	objc.Class
}

// An interface definition for the [RecognizeAnimalsRequest] class.
type IRecognizeAnimalsRequest interface {
	IImageBasedRequest
	SupportedIdentifiersAndReturnError(error foundation.IError) []AnimalIdentifier
}

// A request that recognizes animals in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizeanimalsrequest?language=objc
type RecognizeAnimalsRequest struct {
	ImageBasedRequest
}

func RecognizeAnimalsRequestFrom(ptr unsafe.Pointer) RecognizeAnimalsRequest {
	return RecognizeAnimalsRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

func (rc _RecognizeAnimalsRequestClass) Alloc() RecognizeAnimalsRequest {
	rv := objc.Call[RecognizeAnimalsRequest](rc, objc.Sel("alloc"))
	return rv
}

func RecognizeAnimalsRequest_Alloc() RecognizeAnimalsRequest {
	return RecognizeAnimalsRequestClass.Alloc()
}

func (rc _RecognizeAnimalsRequestClass) New() RecognizeAnimalsRequest {
	rv := objc.Call[RecognizeAnimalsRequest](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRecognizeAnimalsRequest() RecognizeAnimalsRequest {
	return RecognizeAnimalsRequestClass.New()
}

func (r_ RecognizeAnimalsRequest) Init() RecognizeAnimalsRequest {
	rv := objc.Call[RecognizeAnimalsRequest](r_, objc.Sel("init"))
	return rv
}

func (r_ RecognizeAnimalsRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) RecognizeAnimalsRequest {
	rv := objc.Call[RecognizeAnimalsRequest](r_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewRecognizeAnimalsRequestWithCompletionHandler(completionHandler RequestCompletionHandler) RecognizeAnimalsRequest {
	instance := RecognizeAnimalsRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}

// Returns the identifiers of the animals that the request detects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizeanimalsrequest/3751003-supportedidentifiersandreturnerr?language=objc
func (r_ RecognizeAnimalsRequest) SupportedIdentifiersAndReturnError(error foundation.IError) []AnimalIdentifier {
	rv := objc.Call[[]AnimalIdentifier](r_, objc.Sel("supportedIdentifiersAndReturnError:"), objc.Ptr(error))
	return rv
}
