// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DetectBarcodesRequest] class.
var DetectBarcodesRequestClass = _DetectBarcodesRequestClass{objc.GetClass("VNDetectBarcodesRequest")}

type _DetectBarcodesRequestClass struct {
	objc.Class
}

// An interface definition for the [DetectBarcodesRequest] class.
type IDetectBarcodesRequest interface {
	IImageBasedRequest
	SupportedSymbologiesAndReturnError(error foundation.IError) []BarcodeSymbology
	Symbologies() []BarcodeSymbology
	SetSymbologies(value []BarcodeSymbology)
}

// A request that detects barcodes in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectbarcodesrequest?language=objc
type DetectBarcodesRequest struct {
	ImageBasedRequest
}

func DetectBarcodesRequestFrom(ptr unsafe.Pointer) DetectBarcodesRequest {
	return DetectBarcodesRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

func (dc _DetectBarcodesRequestClass) Alloc() DetectBarcodesRequest {
	rv := objc.Call[DetectBarcodesRequest](dc, objc.Sel("alloc"))
	return rv
}

func DetectBarcodesRequest_Alloc() DetectBarcodesRequest {
	return DetectBarcodesRequestClass.Alloc()
}

func (dc _DetectBarcodesRequestClass) New() DetectBarcodesRequest {
	rv := objc.Call[DetectBarcodesRequest](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDetectBarcodesRequest() DetectBarcodesRequest {
	return DetectBarcodesRequestClass.New()
}

func (d_ DetectBarcodesRequest) Init() DetectBarcodesRequest {
	rv := objc.Call[DetectBarcodesRequest](d_, objc.Sel("init"))
	return rv
}

func (d_ DetectBarcodesRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) DetectBarcodesRequest {
	rv := objc.Call[DetectBarcodesRequest](d_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewDetectBarcodesRequestWithCompletionHandler(completionHandler RequestCompletionHandler) DetectBarcodesRequest {
	instance := DetectBarcodesRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}

// Returns the barcode symbologies that the request supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectbarcodesrequest/3750959-supportedsymbologiesandreturnerr?language=objc
func (d_ DetectBarcodesRequest) SupportedSymbologiesAndReturnError(error foundation.IError) []BarcodeSymbology {
	rv := objc.Call[[]BarcodeSymbology](d_, objc.Sel("supportedSymbologiesAndReturnError:"), objc.Ptr(error))
	return rv
}

// The barcode symbologies that the request detects in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectbarcodesrequest/2875397-symbologies?language=objc
func (d_ DetectBarcodesRequest) Symbologies() []BarcodeSymbology {
	rv := objc.Call[[]BarcodeSymbology](d_, objc.Sel("symbologies"))
	return rv
}

// The barcode symbologies that the request detects in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectbarcodesrequest/2875397-symbologies?language=objc
func (d_ DetectBarcodesRequest) SetSymbologies(value []BarcodeSymbology) {
	objc.Call[objc.Void](d_, objc.Sel("setSymbologies:"), value)
}
