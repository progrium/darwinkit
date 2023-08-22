// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coreimage"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BarcodeObservation] class.
var BarcodeObservationClass = _BarcodeObservationClass{objc.GetClass("VNBarcodeObservation")}

type _BarcodeObservationClass struct {
	objc.Class
}

// An interface definition for the [BarcodeObservation] class.
type IBarcodeObservation interface {
	IRectangleObservation
	PayloadStringValue() string
	BarcodeDescriptor() coreimage.BarcodeDescriptor
	Symbology() BarcodeSymbology
}

// An object that represents barcode information that an image analysis request detects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnbarcodeobservation?language=objc
type BarcodeObservation struct {
	RectangleObservation
}

func BarcodeObservationFrom(ptr unsafe.Pointer) BarcodeObservation {
	return BarcodeObservation{
		RectangleObservation: RectangleObservationFrom(ptr),
	}
}

func (bc _BarcodeObservationClass) Alloc() BarcodeObservation {
	rv := objc.Call[BarcodeObservation](bc, objc.Sel("alloc"))
	return rv
}

func BarcodeObservation_Alloc() BarcodeObservation {
	return BarcodeObservationClass.Alloc()
}

func (bc _BarcodeObservationClass) New() BarcodeObservation {
	rv := objc.Call[BarcodeObservation](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBarcodeObservation() BarcodeObservation {
	return BarcodeObservationClass.New()
}

func (b_ BarcodeObservation) Init() BarcodeObservation {
	rv := objc.Call[BarcodeObservation](b_, objc.Sel("init"))
	return rv
}

func (bc _BarcodeObservationClass) ObservationWithBoundingBox(boundingBox coregraphics.Rect) BarcodeObservation {
	rv := objc.Call[BarcodeObservation](bc, objc.Sel("observationWithBoundingBox:"), boundingBox)
	return rv
}

// Creates an observation with a bounding box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectedobjectobservation/2879297-observationwithboundingbox?language=objc
func BarcodeObservation_ObservationWithBoundingBox(boundingBox coregraphics.Rect) BarcodeObservation {
	return BarcodeObservationClass.ObservationWithBoundingBox(boundingBox)
}

func (bc _BarcodeObservationClass) ObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox coregraphics.Rect) BarcodeObservation {
	rv := objc.Call[BarcodeObservation](bc, objc.Sel("observationWithRequestRevision:boundingBox:"), requestRevision, boundingBox)
	return rv
}

// Creates an observation with a revision number and bounding box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectedobjectobservation/2967105-observationwithrequestrevision?language=objc
func BarcodeObservation_ObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox coregraphics.Rect) BarcodeObservation {
	return BarcodeObservationClass.ObservationWithRequestRevisionBoundingBox(requestRevision, boundingBox)
}

// A string value that represents the barcode payload. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnbarcodeobservation/2923485-payloadstringvalue?language=objc
func (b_ BarcodeObservation) PayloadStringValue() string {
	rv := objc.Call[string](b_, objc.Sel("payloadStringValue"))
	return rv
}

// An object that describes the low-level details about the barcode and its data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnbarcodeobservation/2880296-barcodedescriptor?language=objc
func (b_ BarcodeObservation) BarcodeDescriptor() coreimage.BarcodeDescriptor {
	rv := objc.Call[coreimage.BarcodeDescriptor](b_, objc.Sel("barcodeDescriptor"))
	return rv
}

// The symbology of the observed barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnbarcodeobservation/2869611-symbology?language=objc
func (b_ BarcodeObservation) Symbology() BarcodeSymbology {
	rv := objc.Call[BarcodeSymbology](b_, objc.Sel("symbology"))
	return rv
}
