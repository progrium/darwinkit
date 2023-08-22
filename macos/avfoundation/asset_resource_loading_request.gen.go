// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetResourceLoadingRequest] class.
var AssetResourceLoadingRequestClass = _AssetResourceLoadingRequestClass{objc.GetClass("AVAssetResourceLoadingRequest")}

type _AssetResourceLoadingRequestClass struct {
	objc.Class
}

// An interface definition for the [AssetResourceLoadingRequest] class.
type IAssetResourceLoadingRequest interface {
	objc.IObject
	FinishLoading()
	FinishLoadingWithError(error foundation.IError)
	IsCancelled() bool
	Request() foundation.URLRequest
	Requestor() AssetResourceLoadingRequestor
	ContentInformationRequest() AssetResourceLoadingContentInformationRequest
	DataRequest() AssetResourceLoadingDataRequest
	Response() foundation.URLResponse
	SetResponse(value foundation.IURLResponse)
	IsFinished() bool
	Redirect() foundation.URLRequest
	SetRedirect(value foundation.IURLRequest)
}

// An object that encapsulates information about a resource request from a resource loader object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequest?language=objc
type AssetResourceLoadingRequest struct {
	objc.Object
}

func AssetResourceLoadingRequestFrom(ptr unsafe.Pointer) AssetResourceLoadingRequest {
	return AssetResourceLoadingRequest{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetResourceLoadingRequestClass) Alloc() AssetResourceLoadingRequest {
	rv := objc.Call[AssetResourceLoadingRequest](ac, objc.Sel("alloc"))
	return rv
}

func AssetResourceLoadingRequest_Alloc() AssetResourceLoadingRequest {
	return AssetResourceLoadingRequestClass.Alloc()
}

func (ac _AssetResourceLoadingRequestClass) New() AssetResourceLoadingRequest {
	rv := objc.Call[AssetResourceLoadingRequest](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetResourceLoadingRequest() AssetResourceLoadingRequest {
	return AssetResourceLoadingRequestClass.New()
}

func (a_ AssetResourceLoadingRequest) Init() AssetResourceLoadingRequest {
	rv := objc.Call[AssetResourceLoadingRequest](a_, objc.Sel("init"))
	return rv
}

// Causes the receiver to treat the processing of the request as complete. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequest/1388359-finishloading?language=objc
func (a_ AssetResourceLoadingRequest) FinishLoading() {
	objc.Call[objc.Void](a_, objc.Sel("finishLoading"))
}

// Causes the receiver to handle the failure to load a resource for which a resource loader’s delegate took responsibility. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequest/1390491-finishloadingwitherror?language=objc
func (a_ AssetResourceLoadingRequest) FinishLoadingWithError(error foundation.IError) {
	objc.Call[objc.Void](a_, objc.Sel("finishLoadingWithError:"), objc.Ptr(error))
}

// A Boolean value that indicates whether the request has been cancelled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequest/1389518-cancelled?language=objc
func (a_ AssetResourceLoadingRequest) IsCancelled() bool {
	rv := objc.Call[bool](a_, objc.Sel("isCancelled"))
	return rv
}

// The URL request object for the resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequest/1386220-request?language=objc
func (a_ AssetResourceLoadingRequest) Request() foundation.URLRequest {
	rv := objc.Call[foundation.URLRequest](a_, objc.Sel("request"))
	return rv
}

// The asset resource requestor that made the request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequest/2966509-requestor?language=objc
func (a_ AssetResourceLoadingRequest) Requestor() AssetResourceLoadingRequestor {
	rv := objc.Call[AssetResourceLoadingRequestor](a_, objc.Sel("requestor"))
	return rv
}

// The information for a requested resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequest/1390340-contentinformationrequest?language=objc
func (a_ AssetResourceLoadingRequest) ContentInformationRequest() AssetResourceLoadingContentInformationRequest {
	rv := objc.Call[AssetResourceLoadingContentInformationRequest](a_, objc.Sel("contentInformationRequest"))
	return rv
}

// The range of requested resource data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequest/1388779-datarequest?language=objc
func (a_ AssetResourceLoadingRequest) DataRequest() AssetResourceLoadingDataRequest {
	rv := objc.Call[AssetResourceLoadingDataRequest](a_, objc.Sel("dataRequest"))
	return rv
}

// The URL response for the loading request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequest/1389034-response?language=objc
func (a_ AssetResourceLoadingRequest) Response() foundation.URLResponse {
	rv := objc.Call[foundation.URLResponse](a_, objc.Sel("response"))
	return rv
}

// The URL response for the loading request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequest/1389034-response?language=objc
func (a_ AssetResourceLoadingRequest) SetResponse(value foundation.IURLResponse) {
	objc.Call[objc.Void](a_, objc.Sel("setResponse:"), objc.Ptr(value))
}

// A Boolean value that indicates whether loading of the resource has finished. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequest/1389270-finished?language=objc
func (a_ AssetResourceLoadingRequest) IsFinished() bool {
	rv := objc.Call[bool](a_, objc.Sel("isFinished"))
	return rv
}

// An URL request instance if the loading request was redirected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequest/1390854-redirect?language=objc
func (a_ AssetResourceLoadingRequest) Redirect() foundation.URLRequest {
	rv := objc.Call[foundation.URLRequest](a_, objc.Sel("redirect"))
	return rv
}

// An URL request instance if the loading request was redirected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequest/1390854-redirect?language=objc
func (a_ AssetResourceLoadingRequest) SetRedirect(value foundation.IURLRequest) {
	objc.Call[objc.Void](a_, objc.Sel("setRedirect:"), objc.Ptr(value))
}
