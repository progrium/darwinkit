// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetResourceLoadingContentInformationRequest] class.
var AssetResourceLoadingContentInformationRequestClass = _AssetResourceLoadingContentInformationRequestClass{objc.GetClass("AVAssetResourceLoadingContentInformationRequest")}

type _AssetResourceLoadingContentInformationRequestClass struct {
	objc.Class
}

// An interface definition for the [AssetResourceLoadingContentInformationRequest] class.
type IAssetResourceLoadingContentInformationRequest interface {
	objc.IObject
	RenewalDate() foundation.Date
	SetRenewalDate(value foundation.IDate)
	ContentLength() int64
	SetContentLength(value int64)
	IsByteRangeAccessSupported() bool
	SetByteRangeAccessSupported(value bool)
	ContentType() string
	SetContentType(value string)
	AllowedContentTypes() []string
}

// A query for retrieving essential information about a resource that an asset resource-loading request references. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingcontentinformationrequest?language=objc
type AssetResourceLoadingContentInformationRequest struct {
	objc.Object
}

func AssetResourceLoadingContentInformationRequestFrom(ptr unsafe.Pointer) AssetResourceLoadingContentInformationRequest {
	return AssetResourceLoadingContentInformationRequest{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetResourceLoadingContentInformationRequestClass) Alloc() AssetResourceLoadingContentInformationRequest {
	rv := objc.Call[AssetResourceLoadingContentInformationRequest](ac, objc.Sel("alloc"))
	return rv
}

func AssetResourceLoadingContentInformationRequest_Alloc() AssetResourceLoadingContentInformationRequest {
	return AssetResourceLoadingContentInformationRequestClass.Alloc()
}

func (ac _AssetResourceLoadingContentInformationRequestClass) New() AssetResourceLoadingContentInformationRequest {
	rv := objc.Call[AssetResourceLoadingContentInformationRequest](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetResourceLoadingContentInformationRequest() AssetResourceLoadingContentInformationRequest {
	return AssetResourceLoadingContentInformationRequestClass.New()
}

func (a_ AssetResourceLoadingContentInformationRequest) Init() AssetResourceLoadingContentInformationRequest {
	rv := objc.Call[AssetResourceLoadingContentInformationRequest](a_, objc.Sel("init"))
	return rv
}

// The date at which a new resource loading request will be issued for resources that expire, if the media system still requires it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingcontentinformationrequest/1390683-renewaldate?language=objc
func (a_ AssetResourceLoadingContentInformationRequest) RenewalDate() foundation.Date {
	rv := objc.Call[foundation.Date](a_, objc.Sel("renewalDate"))
	return rv
}

// The date at which a new resource loading request will be issued for resources that expire, if the media system still requires it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingcontentinformationrequest/1390683-renewaldate?language=objc
func (a_ AssetResourceLoadingContentInformationRequest) SetRenewalDate(value foundation.IDate) {
	objc.Call[objc.Void](a_, objc.Sel("setRenewalDate:"), objc.Ptr(value))
}

// The length, in bytes, of the requested resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingcontentinformationrequest/1389390-contentlength?language=objc
func (a_ AssetResourceLoadingContentInformationRequest) ContentLength() int64 {
	rv := objc.Call[int64](a_, objc.Sel("contentLength"))
	return rv
}

// The length, in bytes, of the requested resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingcontentinformationrequest/1389390-contentlength?language=objc
func (a_ AssetResourceLoadingContentInformationRequest) SetContentLength(value int64) {
	objc.Call[objc.Void](a_, objc.Sel("setContentLength:"), value)
}

// A Boolean value that indicates whether random access to arbitrary ranges of bytes of the resource is supported. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingcontentinformationrequest/1386054-byterangeaccesssupported?language=objc
func (a_ AssetResourceLoadingContentInformationRequest) IsByteRangeAccessSupported() bool {
	rv := objc.Call[bool](a_, objc.Sel("isByteRangeAccessSupported"))
	return rv
}

// A Boolean value that indicates whether random access to arbitrary ranges of bytes of the resource is supported. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingcontentinformationrequest/1386054-byterangeaccesssupported?language=objc
func (a_ AssetResourceLoadingContentInformationRequest) SetByteRangeAccessSupported(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setByteRangeAccessSupported:"), value)
}

// The UTI that specifies the type of data contained by the requested resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingcontentinformationrequest/1388529-contenttype?language=objc
func (a_ AssetResourceLoadingContentInformationRequest) ContentType() string {
	rv := objc.Call[string](a_, objc.Sel("contentType"))
	return rv
}

// The UTI that specifies the type of data contained by the requested resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingcontentinformationrequest/1388529-contenttype?language=objc
func (a_ AssetResourceLoadingContentInformationRequest) SetContentType(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setContentType:"), value)
}

// The types of data that are accepted as a valid response for the requested resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingcontentinformationrequest/2936886-allowedcontenttypes?language=objc
func (a_ AssetResourceLoadingContentInformationRequest) AllowedContentTypes() []string {
	rv := objc.Call[[]string](a_, objc.Sel("allowedContentTypes"))
	return rv
}
