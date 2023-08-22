// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SequenceRequestHandler] class.
var SequenceRequestHandlerClass = _SequenceRequestHandlerClass{objc.GetClass("VNSequenceRequestHandler")}

type _SequenceRequestHandlerClass struct {
	objc.Class
}

// An interface definition for the [SequenceRequestHandler] class.
type ISequenceRequestHandler interface {
	objc.IObject
	PerformRequestsOnCMSampleBufferError(requests []IRequest, sampleBuffer coremedia.SampleBufferRef, error foundation.IError) bool
}

// An object that processes image analysis requests for each frame in a sequence. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnsequencerequesthandler?language=objc
type SequenceRequestHandler struct {
	objc.Object
}

func SequenceRequestHandlerFrom(ptr unsafe.Pointer) SequenceRequestHandler {
	return SequenceRequestHandler{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ SequenceRequestHandler) Init() SequenceRequestHandler {
	rv := objc.Call[SequenceRequestHandler](s_, objc.Sel("init"))
	return rv
}

func (sc _SequenceRequestHandlerClass) Alloc() SequenceRequestHandler {
	rv := objc.Call[SequenceRequestHandler](sc, objc.Sel("alloc"))
	return rv
}

func SequenceRequestHandler_Alloc() SequenceRequestHandler {
	return SequenceRequestHandlerClass.Alloc()
}

func (sc _SequenceRequestHandlerClass) New() SequenceRequestHandler {
	rv := objc.Call[SequenceRequestHandler](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSequenceRequestHandler() SequenceRequestHandler {
	return SequenceRequestHandlerClass.New()
}

// Performs one or more requests on an image contained within a sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnsequencerequesthandler/3571272-performrequests?language=objc
func (s_ SequenceRequestHandler) PerformRequestsOnCMSampleBufferError(requests []IRequest, sampleBuffer coremedia.SampleBufferRef, error foundation.IError) bool {
	rv := objc.Call[bool](s_, objc.Sel("performRequests:onCMSampleBuffer:error:"), requests, sampleBuffer, objc.Ptr(error))
	return rv
}
